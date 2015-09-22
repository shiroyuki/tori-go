// Routing collection
//
// There are two types of routes in Tori Framework: reversible and non-reversible.
// The reversible one can be use to with other classes to re-create the route that
// refers to the given one. For example:
//
//      // Suppose we have a router object.
//      route  := NewRoute("/user/{alias}")
//      params := map[string]string{
//          "alias": "shiroyuki"
//      }
//
//      requestPath := route.For(params) // expected: /user/shiroyuki
package web

import tori_re "../re"

type Route struct {
    Pattern    string
    Reversible bool
    RePattern  *tori_re.Expression
}

// Create a route.
func NewRoute(pattern string, reversible bool) Route {
    var route Route

    // TODO add the assertion to check if the pattern has a prefix "/". Raise exceptions if necessary.

    route = Route{
        Pattern:    pattern,
        Reversible: reversible,
        RePattern:  nil, // This is lazy-loading.
    }

    return route
}

func (self *Route) GetCompiledPattern() *tori_re.Expression {
    var compiled             tori_re.Expression
    var simpleRoutingPattern tori_re.Expression

    if self.RePattern != nil {
        return self.RePattern
    }

    // Handle a non-reversible route.
    if (!self.Reversible) {
        compiled = tori_re.Compile(self.Pattern)
        self.RePattern = &compiled

        return self.RePattern
    }

    // Handle a reversible route.
    simpleRoutingPattern.SearchAll(self.Pattern)

    return self.RePattern
}
