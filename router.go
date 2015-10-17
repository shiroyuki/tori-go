package tori

import "log"
import "github.com/shiroyuki/re"

type Router struct {
    PriorityList *RecordList
    // the route id-to-route map
    IdToRouteMap map[string](*Route)
    DebugMode    bool
}

func NewRouter() *Router {
    priorityList := &RecordList{}
    idToRouteMap := make(map[string]*Route)

    return &Router{
        PriorityList: priorityList,
        IdToRouteMap: idToRouteMap,
        DebugMode:    false,
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
    if self.DebugMode {
        log.Printf("Incomming request for %s %s\n", method, path)
    }

    r, p := self.PriorityList.Find(method, path)

    if self.DebugMode {
        if r == nil {
            log.Printf("Unable to find the request handler for %s %s\n", method, path)
        }
    }

    return r, p
}
