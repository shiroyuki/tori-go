// Core for web framework
package tori

import "log"
import "net/http"
import "github.com/shiroyuki/re"
import yotsuba "github.com/shiroyuki/yotsuba-go"

type Core struct { // implements http.Handler
    Router     *Router
    Cache      *yotsuba.CacheDriver
    Enigma     *yotsuba.Enigma
    Internal   *http.Server
    Compressed bool
}

// Create a core of the web framework with everything pre-configured for development.
func NewSimpleCore() *Core {
    var enigma = yotsuba.Enigma{}
    var router = NewRouter()
    var actualCacheDriver = yotsuba.NewInMemoryCacheDriver(&enigma, false)
    var castedCacheDriver = yotsuba.CacheDriver(actualCacheDriver)

    return NewCore(
        router,
        &castedCacheDriver,
        &enigma,
        false,
    )
}

func NewCore(
    router     *Router,
    cache      *yotsuba.CacheDriver,
    enigma     *yotsuba.Enigma,
    compressed bool,
) *Core {
    appCore := Core{
        Router:     router,
        Cache:      cache,
        Enigma:     enigma,
        Compressed: compressed,
    }

    internalServer := &http.Server{
        Addr:    "0.0.0.0:8000",
        Handler: &appCore,
    }

    appCore.Internal = internalServer

    return &appCore
}

// Handle the request and delegate the request to a proper handler.
func (self *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var routingRecord *Record
    var parameters    *re.MultipleResult

    var method string = r.Method
    var path   string = r.URL.Path

    routingRecord, parameters = self.Router.Find(method, path)

    if routingRecord == nil {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Not Found"))

        // TODO Event "web.core.error.404@default": allow flexible error handling for HTTP 404.

        return
    }

    handler := NewHandler(routingRecord.Route, &w, r, parameters)
    action  := routingRecord.Action

    // TODO Event "web.handler.pre.<route_id>": allow flexible interceptions before processing requests.

    (*action)(handler)

    // TODO Event "web.handler.post.<route_id>": allow flexible interceptions before processing requests.

    self.response(handler)
}

func (self *Core) response(handler *Handler) {
    var content []byte = handler.Content()

    if !self.Compressed {
        handler.SetContentLength(len(content))
        (*handler.Response).Write(content)

        return
    }

    compressed := self.Enigma.Compress(content)

    handler.SetContentEncoding("gzip")
    handler.SetContentLength(len(compressed))
    (*handler.Response).Write(compressed)
}

func (self *Core) listen(address *string) {
    if address != nil {
        self.Internal.Addr = *address
    }

    log.Println("Listening at:", self.Internal.Addr)
    log.Fatal("Terminated due to:", self.Internal.ListenAndServe())
}
