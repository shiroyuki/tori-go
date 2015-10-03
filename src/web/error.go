package web

import "errors"

var RouteWithDuplicatedKeyError = errors.New("In a route, any keys must appear only once.")
