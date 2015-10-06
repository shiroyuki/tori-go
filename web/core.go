// Core for web framework
package web

import "net/http"
import tori  "../"
import re    "../re"
import cache "../cache"

type Core struct { // implements http.Handler
    Router     *Router
    Cache      *cache.Driver
    Enigma     *tori.Enigma
    Compressed bool
}

// Create a core of the web framework with everything pre-configured for development.
func NewSimpleCore() *Core {
    var enigma = tori.Enigma{}
    var router = NewRouter()
    var actualCacheDriver = cache.NewInMemoryCacheDriver(&enigma, false)
    var castedCacheDriver = cache.Driver(actualCacheDriver)

    return NewCore(
        router,
        &castedCacheDriver,
        &enigma,
        false,
    )
}

func NewCore(
    router     *Router,
    cache      *cache.Driver,
    enigma     *tori.Enigma,
    compressed bool,
) *Core {
    return &Core{
        Router:     router,
        Cache:      cache,
        Enigma:     enigma,
        Compressed: compressed,
    }
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
