package web

import "net/http"
import "../re"

type Route interface {
    GetCompiledPattern() (*re.Expression, error)
    Match(request *http.Request) (*re.MultipleResult)
}
