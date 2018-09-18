package gobuildserver

import (
	"fmt"

	"github.com/sourcegraph/ctxvfs"
)

// addSysZversionFile adds the zversion.go file, which is generated
// during the Go release process and does not exist in the VCS repo
// archive zips. We need to create it here, or else we'll see
// typechecker errors like "StackGuardMultiplier not declared by
// package sys" when any packages import from the Go stdlib.
func addSysZversionFile(fs ctxvfs.FileSystem) ctxvfs.FileSystem {
	return ctxvfs.SingleFileOverlay(fs,
		"/src/runtime/internal/sys/zversion.go",
		[]byte(fmt.Sprintf(`
package sys

const DefaultGoroot = %q
const TheVersion = %q
const Goexperiment=""
const StackGuardMultiplier=1`,
			goroot, RuntimeVersion)))
}

// go list std | awk '{ print "\"" $1 "\": struct{}{}," }'
var stdlibPackagePaths = map[string]struct{}{
	"archive/tar":                       struct{}{},
	"archive/zip":                       struct{}{},
	"bufio":                             struct{}{},
	"bytes":                             struct{}{},
	"compress/bzip2":                    struct{}{},
	"compress/flate":                    struct{}{},
	"compress/gzip":                     struct{}{},
	"compress/lzw":                      struct{}{},
	"compress/zlib":                     struct{}{},
	"container/heap":                    struct{}{},
	"container/list":                    struct{}{},
	"container/ring":                    struct{}{},
	"context":                           struct{}{},
	"crypto":                            struct{}{},
	"crypto/aes":                        struct{}{},
	"crypto/cipher":                     struct{}{},
	"crypto/des":                        struct{}{},
	"crypto/dsa":                        struct{}{},
	"crypto/ecdsa":                      struct{}{},
	"crypto/elliptic":                   struct{}{},
	"crypto/hmac":                       struct{}{},
	"crypto/internal/randutil":          struct{}{},
	"crypto/internal/subtle":            struct{}{},
	"crypto/md5":                        struct{}{},
	"crypto/rand":                       struct{}{},
	"crypto/rc4":                        struct{}{},
	"crypto/rsa":                        struct{}{},
	"crypto/sha1":                       struct{}{},
	"crypto/sha256":                     struct{}{},
	"crypto/sha512":                     struct{}{},
	"crypto/subtle":                     struct{}{},
	"crypto/tls":                        struct{}{},
	"crypto/x509":                       struct{}{},
	"crypto/x509/pkix":                  struct{}{},
	"database/sql":                      struct{}{},
	"database/sql/driver":               struct{}{},
	"debug/dwarf":                       struct{}{},
	"debug/elf":                         struct{}{},
	"debug/gosym":                       struct{}{},
	"debug/macho":                       struct{}{},
	"debug/pe":                          struct{}{},
	"debug/plan9obj":                    struct{}{},
	"encoding":                          struct{}{},
	"encoding/ascii85":                  struct{}{},
	"encoding/asn1":                     struct{}{},
	"encoding/base32":                   struct{}{},
	"encoding/base64":                   struct{}{},
	"encoding/binary":                   struct{}{},
	"encoding/csv":                      struct{}{},
	"encoding/gob":                      struct{}{},
	"encoding/hex":                      struct{}{},
	"encoding/json":                     struct{}{},
	"encoding/pem":                      struct{}{},
	"encoding/xml":                      struct{}{},
	"errors":                            struct{}{},
	"expvar":                            struct{}{},
	"flag":                              struct{}{},
	"fmt":                               struct{}{},
	"go/ast":                            struct{}{},
	"go/build":                          struct{}{},
	"go/constant":                       struct{}{},
	"go/doc":                            struct{}{},
	"go/format":                         struct{}{},
	"go/importer":                       struct{}{},
	"go/internal/gccgoimporter":         struct{}{},
	"go/internal/gcimporter":            struct{}{},
	"go/internal/srcimporter":           struct{}{},
	"go/parser":                         struct{}{},
	"go/printer":                        struct{}{},
	"go/scanner":                        struct{}{},
	"go/token":                          struct{}{},
	"go/types":                          struct{}{},
	"hash":                              struct{}{},
	"hash/adler32":                      struct{}{},
	"hash/crc32":                        struct{}{},
	"hash/crc64":                        struct{}{},
	"hash/fnv":                          struct{}{},
	"html":                              struct{}{},
	"html/template":                     struct{}{},
	"image":                             struct{}{},
	"image/color":                       struct{}{},
	"image/color/palette":               struct{}{},
	"image/draw":                        struct{}{},
	"image/gif":                         struct{}{},
	"image/internal/imageutil":          struct{}{},
	"image/jpeg":                        struct{}{},
	"image/png":                         struct{}{},
	"index/suffixarray":                 struct{}{},
	"internal/bytealg":                  struct{}{},
	"internal/cpu":                      struct{}{},
	"internal/nettrace":                 struct{}{},
	"internal/poll":                     struct{}{},
	"internal/race":                     struct{}{},
	"internal/singleflight":             struct{}{},
	"internal/syscall/unix":             struct{}{},
	"internal/syscall/windows":          struct{}{},
	"internal/syscall/windows/registry": struct{}{},
	"internal/syscall/windows/sysdll":   struct{}{},
	"internal/testenv":                  struct{}{},
	"internal/testlog":                  struct{}{},
	"internal/trace":                    struct{}{},
	"io":                                struct{}{},
	"io/ioutil":                         struct{}{},
	"log":                               struct{}{},
	"log/syslog":                        struct{}{},
	"math":                              struct{}{},
	"math/big":                          struct{}{},
	"math/bits":                         struct{}{},
	"math/cmplx":                        struct{}{},
	"math/rand":                         struct{}{},
	"mime":                              struct{}{},
	"mime/multipart":                    struct{}{},
	"mime/quotedprintable":              struct{}{},
	"net":                               struct{}{},
	"net/http":                          struct{}{},
	"net/http/cgi":                      struct{}{},
	"net/http/cookiejar":                struct{}{},
	"net/http/fcgi":                     struct{}{},
	"net/http/httptest":                 struct{}{},
	"net/http/httptrace":                struct{}{},
	"net/http/httputil":                 struct{}{},
	"net/http/internal":                 struct{}{},
	"net/http/pprof":                    struct{}{},
	"net/internal/socktest":             struct{}{},
	"net/mail":                          struct{}{},
	"net/rpc":                           struct{}{},
	"net/rpc/jsonrpc":                   struct{}{},
	"net/smtp":                          struct{}{},
	"net/textproto":                     struct{}{},
	"net/url":                           struct{}{},
	"os":                                struct{}{},
	"os/exec":                           struct{}{},
	"os/signal":                         struct{}{},
	"os/signal/internal/pty":            struct{}{},
	"os/user":                           struct{}{},
	"path":                              struct{}{},
	"path/filepath":                     struct{}{},
	"plugin":                            struct{}{},
	"reflect":                           struct{}{},
	"regexp":                            struct{}{},
	"regexp/syntax":                     struct{}{},
	"runtime":                           struct{}{},
	"runtime/cgo":                       struct{}{},
	"runtime/debug":                     struct{}{},
	"runtime/internal/atomic":           struct{}{},
	"runtime/internal/sys":              struct{}{},
	"runtime/pprof":                     struct{}{},
	"runtime/pprof/internal/profile":    struct{}{},
	"runtime/race":                      struct{}{},
	"runtime/trace":                     struct{}{},
	"sort":                              struct{}{},
	"strconv":                           struct{}{},
	"strings":                           struct{}{},
	"sync":                              struct{}{},
	"sync/atomic":                       struct{}{},
	"syscall":                           struct{}{},
	"testing":                           struct{}{},
	"testing/internal/testdeps":         struct{}{},
	"testing/iotest":                    struct{}{},
	"testing/quick":                     struct{}{},
	"text/scanner":                      struct{}{},
	"text/tabwriter":                    struct{}{},
	"text/template":                     struct{}{},
	"text/template/parse":               struct{}{},
	"time":                              struct{}{},
	"unicode":                           struct{}{},
	"unicode/utf16":                     struct{}{},
	"unicode/utf8":                      struct{}{},
	"unsafe":                            struct{}{},
	"vendor/golang_org/x/crypto/chacha20poly1305":  struct{}{},
	"vendor/golang_org/x/crypto/cryptobyte":        struct{}{},
	"vendor/golang_org/x/crypto/cryptobyte/asn1":   struct{}{},
	"vendor/golang_org/x/crypto/curve25519":        struct{}{},
	"vendor/golang_org/x/crypto/internal/chacha20": struct{}{},
	"vendor/golang_org/x/crypto/poly1305":          struct{}{},
	"vendor/golang_org/x/net/dns/dnsmessage":       struct{}{},
	"vendor/golang_org/x/net/http/httpguts":        struct{}{},
	"vendor/golang_org/x/net/http/httpproxy":       struct{}{},
	"vendor/golang_org/x/net/http2/hpack":          struct{}{},
	"vendor/golang_org/x/net/idna":                 struct{}{},
	"vendor/golang_org/x/net/internal/nettest":     struct{}{},
	"vendor/golang_org/x/net/nettest":              struct{}{},
	"vendor/golang_org/x/net/route":                struct{}{},
	"vendor/golang_org/x/text/secure":              struct{}{},
	"vendor/golang_org/x/text/secure/bidirule":     struct{}{},
	"vendor/golang_org/x/text/transform":           struct{}{},
	"vendor/golang_org/x/text/unicode":             struct{}{},
	"vendor/golang_org/x/text/unicode/bidi":        struct{}{},
	"vendor/golang_org/x/text/unicode/norm":        struct{}{},
}
