import { BehaviorSubject, Observable, TeardownLogic } from 'rxjs'
import { ExecuteCommandParams } from '../../protocol'

export type ExecuteCommandSignature = (params: ExecuteCommandParams) => Promise<any>

/** A registered command in the command registry. */
export interface CommandEntry {
    /** The command ID (conventionally, e.g., "myextension.mycommand"). */
    command: string

    /** The function called to run the command and return an async value. */
    run: (...args: any[]) => Promise<any>
}

/** Manages and executes commands from all extensions. */
export class CommandRegistry {
    private entries = new BehaviorSubject<CommandEntry[]>([])

    public registerCommand(entry: CommandEntry): TeardownLogic {
        // Enforce uniqueness of command IDs.
        for (const e of this.entries.value) {
            if (e.command === entry.command) {
                throw new Error(`command is already registered: ${JSON.stringify(entry.command)}`)
            }
        }

        this.entries.next([...this.entries.value, entry])
        return () => {
            this.entries.next(this.entries.value.filter(e => e !== entry))
        }
    }

    public executeCommand(params: ExecuteCommandParams): Promise<any> {
        return executeCommand(this.commandsSnapshot, params)
    }

    /** All commands, emitted whenever the set of registered commands changed. */
    public readonly commands: Observable<CommandEntry[]> = this.entries

    /**
     * The current set of commands. Used by callers that do not need to react to commands being registered or
     * unregistered.
     */
    public get commandsSnapshot(): CommandEntry[] {
        return this.entries.value
    }
}

/**
 * Returns an observable that emits all providers' hovers whenever any of the last-emitted set of providers emits
 * hovers.
 *
 * Most callers should use TextDocumentHoverProviderRegistry, which sources hovers from the current set of
 * registered providers (and then completes).
 */
export function executeCommand(commands: CommandEntry[], params: ExecuteCommandParams): Promise<any> {
    const command = commands.find(c => c.command === params.command)
    if (!command) {
        throw new Error(`command not found: ${JSON.stringify(params.command)}`)
    }
    return command.run(...(params.arguments || []))
}