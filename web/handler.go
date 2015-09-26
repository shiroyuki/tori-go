package web

import "net/http"
import "../re"

type Handler struct {
    Response   *http.ResponseWriter
    Request    *http.Request
    Parameters *re.MultipleResult
}

func NewHandler(w *http.ResponseWriter, r *http.Request, p *re.MultipleResult) *Handler {
    handler := &Handler{
        Response:   w,
        Request:    r,
        Parameters: p,
    }

    return handler
}
