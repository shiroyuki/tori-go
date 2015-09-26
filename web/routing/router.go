package routing

import web "../"

type Action func(h *web.Handler)

type Router struct {
    // the method-to-route-to-handler map
    HandlerMap map[string]map[string]Action

    // the route pattern-to-object map
    ReferenceMap map[string](*Route)
}

func NewRouter() *Router {
    hMap := make(map[string]map[string]Action)
    rMap := make(map[string]*Route)

    return &Router{
        HandlerMap:   hMap,
        ReferenceMap: rMap,
    }
}

func (self *Router) AddRoute(method string, pattern string, handler Action, reversible bool) {
    var route *Route

    route, ok := self.ReferenceMap[pattern]

    if !ok {
        route := NewRoute(pattern, reversible)

        self.ReferenceMap[pattern] = route
    }

    self.HandlerMap[method][route.Pattern] = handler
}

func (self *Router) AddSimpleRoute(method string, pattern string, handler Action) {
    self.AddRoute(method, pattern, handler, true)
}
