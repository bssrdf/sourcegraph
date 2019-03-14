import { ProxyValue, proxyValue, proxyValueSymbol } from '@sourcegraph/comlink'
import { from, of, ReplaySubject, Unsubscribable } from 'rxjs'
import { map } from 'rxjs/operators'
import { PanelView } from 'sourcegraph'
import { ContributableViewContainer, TextDocumentPositionParams } from '../../protocol'
import { modelToTextDocumentPositionParams } from '../model'
import { TextDocumentLocationProviderIDRegistry } from '../services/location'
import { ModelService } from '../services/modelService'
import { PanelViewWithComponent, ViewProviderRegistry } from '../services/view'

/** @internal */
export interface PanelViewData extends Pick<PanelView, 'title' | 'content' | 'priority' | 'component'> {}

export interface PanelUpdater extends Unsubscribable, ProxyValue {
    update(data: PanelViewData): void
}

/** @internal */
export interface ClientViewsAPI extends ProxyValue {
    $registerPanelViewProvider(provider: { id: string }): PanelUpdater
}

/** @internal */
export class ClientViews implements ClientViewsAPI {
    public readonly [proxyValueSymbol] = true

    constructor(
        private viewRegistry: ViewProviderRegistry,
        private textDocumentLocations: TextDocumentLocationProviderIDRegistry,
        private modelService: ModelService
    ) {}

    public $registerPanelViewProvider(provider: { id: string }): PanelUpdater {
        // TODO(sqs): This will probably hang forever if an extension neglects to set any of the fields on a
        // PanelView because this subject will never emit.
        const panelView = new ReplaySubject<PanelViewData>(1)
        const registryUnsubscribable = this.viewRegistry.registerProvider(
            { ...provider, container: ContributableViewContainer.Panel },
            panelView.pipe(
                map(({ title, content, priority, component }) => {
                    const panelView: PanelViewWithComponent = {
                        title,
                        content,
                        priority,
                        locationProvider: from(this.modelService.model).pipe(
                            map(model => {
                                if (!component) {
                                    return of(null)
                                }
                                const params: TextDocumentPositionParams | null = modelToTextDocumentPositionParams(
                                    model
                                )
                                if (!params) {
                                    return of(null)
                                }
                                return this.textDocumentLocations.getLocations(component.locationProvider, params)
                            })
                        ),
                    }
                    return panelView
                })
            )
        )
        return proxyValue({
            update: (data: PanelViewData) => {
                panelView.next(data)
            },
            unsubscribe: () => {
                registryUnsubscribable.unsubscribe()
            },
        })
    }
}
