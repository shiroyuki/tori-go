package routing

type Router struct {
    PriorityList *RecordList
    // the route id-to-route map
    IdToRouteMap map[string](*Route)
}

func NewRouter() *Router {
    priorityList := &RecordList{}
    idToRouteMap := make(map[string]*Route)

    return &Router{
        PriorityList: priorityList,
        IdToRouteMap: idToRouteMap,
    }
}

func (self *Router) AddRoute(
    id         string,
    method     string,
    pattern    string,
    handler    Action,
    reversible bool,
    cacheable  bool,
) {
    _, ok := self.IdToRouteMap[pattern]

    if !ok {
        route := NewRoute(pattern, reversible)

        self.IdToRouteMap[pattern] = route
    }

    self.PriorityList.Append(&Record{
        Id:        id,
        Method:    method,
        Handler:   handler,
        Cacheable: cacheable,
    })
}
