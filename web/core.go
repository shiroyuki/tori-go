// Core for web framework
package web

import (
    "fmt"
    "net/http"
)

import tori  "../"
import tori_cache  "../cache"

type Core struct { // implements http.Handler
    Cache      tori_cache.Driver
    Enigma     tori.Enigma
    Compressed bool
}

func NewCore(
    cache      tori_cache.Driver,
    enigma     tori.Enigma,
    compressed bool,
) Core {
    return Core{
        Cache:      cache,
        Enigma:     enigma,
        Compressed: compressed,
    }
}

// Handle the request and delegate the request to a proper handler.
func (self *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    self.write(w, "text/plain", []byte("sample"))
}

func (self *Core) write(w http.ResponseWriter, kind string, content []byte) {
    w.Header().Set("Content-Type", kind)

    if !self.Compressed {
        w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
        w.Write(content)

        return
    }

    compressed := self.Enigma.Compress(content)

    w.Header().Set("Content-Encoding", "gzip")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", len(compressed)))
    w.Write(compressed)
}
