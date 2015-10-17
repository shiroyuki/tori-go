package tori

import "github.com/shiroyuki/re"

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

func (self *Router) Handle(
    id         string,
    method     string,
    pattern    string,
    action     Action,
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
        Action:    &action,
        Cacheable: cacheable,
    })
}

func (self *Router) Find(method string, path string) (*Record, *re.MultipleResult) {
    r, p := self.PriorityList.Find(method, path)

    return r, p
}
