package tori

import "fmt"
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
    var firstLine string = fmt.Sprintf("%s %s", method, pattern)
    var logPrefix string = fmt.Sprintf("Route #%s (%s)", id, firstLine)
    var newRecord Record

    self.log(logPrefix + " Registering...")

    _, ok := self.IdToRouteMap[pattern]

    if !ok {
        self.log(logPrefix + " Adding a new route...")

        route := NewRoute(pattern, reversible)

        self.IdToRouteMap[pattern] = route

        self.log(logPrefix + " Added a new route.")
    }

    newRecord = Record{
        Id:        id,
        Method:    method,
        Action:    action,
        Cacheable: cacheable,
    }

    self.PriorityList.Append(&newRecord)

    self.log(logPrefix + " Registered.")
}

func (self *Router) OnGet(
    id         string,
    pattern    string,
    action     Action,
    cacheable  bool,
) {
    self.Handle(id, METHOD_GET, pattern, action, true, cacheable)
}

func (self *Router) OnPost(
    id         string,
    pattern    string,
    action     Action,
) {
    self.Handle(id, METHOD_POST, pattern, action, true, false)
}

func (self *Router) OnPut(
    id         string,
    pattern    string,
    action     Action,
) {
    self.Handle(id, METHOD_PUT, pattern, action, true, false)
}

func (self *Router) OnDelete(
    id         string,
    pattern    string,
    action     Action,
) {
    self.Handle(id, METHOD_DELETE, pattern, action, true, false)
}

func (self *Router) Find(method string, path string) (*Record, *re.MultipleResult) {
    var firstLine string = fmt.Sprintf("%s %s", method, path)

    self.log(fmt.Sprintf("Incomming request for %s", firstLine))

    r, p := self.PriorityList.Find(method, path)

    if r == nil {
        self.log(fmt.Sprintf("Unable to find the request handler for %s", firstLine))
    }

    return r, p
}

func (self *Router) log(message string) {
    if !self.DebugMode {
        return
    }

    log.Println(message)
}
