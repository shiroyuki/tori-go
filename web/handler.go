package web

import "bytes"
import "net/http"
import "../re"

type Handler struct {
    Route         *Route
    Response      *http.ResponseWriter
    Request       *http.Request
    Parameters    *re.MultipleResult
    BufferEnabled bool
    Status        uint16
    buffer        *bytes.Buffer
}

// Construct a handler.
//
// The constructed handler will use a buffer by default. This is to allow a
// middleware to work on the response content and minimize the I/O interuption.
//
// Disabling buffering will automatically prevent a middleware from doing
// post-processing on the response content for the corresponding request.
func NewHandler(
    route      *Route,
    response   *http.ResponseWriter,
    request    *http.Request,
    parameters *re.MultipleResult,
) *Handler {
    handler := &Handler{
        Route:         route,
        Response:      response,
        Request:       request,
        Parameters:    parameters,
        BufferEnabled: true,
    }

    handler.SetStatus(http.StatusOK)

    return handler
}

// Set the HTTP status code for the response.
func (self *Handler) SetStatus(statusCode uint16) {
    self.Status = statusCode
}

// Get the request header.
func (self *Handler) GetHeader(key string) string {
    return self.Request.Header.Get(key)
}

// Set the response header.
func (self *Handler) SetHeader(key string, value string) {
    (*self.Response).Header().Add(key, value)
}

// Disable buffering.
func (self *Handler) DisableBuffering() {
    self.BufferEnabled = false
}

func (self *Handler) Write(content string) {
    self.WriteByte([]byte(content))
}

func (self *Handler) WriteByte(content []byte) {
    if self.BufferEnabled {
        self.buffer.Write(content)

        return
    }

    (*self.Response).Write(content)
}
