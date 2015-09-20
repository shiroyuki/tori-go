package web

import "net/http"

type RequestHandler interface {
    Initialize(w http.ResponseWriter, r *http.Request)
    Write(content []byte)
}
