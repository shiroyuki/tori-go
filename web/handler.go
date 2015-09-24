package web

import "net/http"

type Handler struct {
    //
}

func (self *Handler) Initialize(w http.ResponseWriter, r *http.Request) {}

func (self *Handler) Write(content []byte) {}
