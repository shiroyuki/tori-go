// Package routing to manage request routes and deliver request to a corresponding handler.
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
package routing

import tori_re "../../re"

var toriWebRoutingSimplePattern = tori_re.Compile("<(?P<key>[^>]+)>")

type Route struct {
    Pattern    string
    Reversible bool
    RePattern  *tori_re.Expression
}

// Create a route.
func NewRoute(pattern string, reversible bool) *Route {
    var route Route

    // TODO add the assertion to check if the pattern has a prefix "/". Raise exceptions if necessary.

    route = Route{
        Pattern:    pattern,
        Reversible: reversible,
        RePattern:  nil, // This is lazy-loading.
    }

    return &route
}

func (self *Route) GetCompiledPattern() (*tori_re.Expression, error) {
    if self.RePattern != nil {
        return self.RePattern, nil
    }

    // Handle a non-reversible route.
    if (!self.Reversible) {
        return self.compileNonReversiblePattern(), nil
    }

    // Handle a reversible route.
    return self.compileReversiblePattern()
}

func (self *Route) compileReversiblePattern() (*tori_re.Expression, error) {
    var alternativePattern string

    matches := toriWebRoutingSimplePattern.SearchAll(self.Pattern)

    if matches.HasAny() {
        values := matches.Key("key")

        for _, key := range *values {
            spotCheckPattern := tori_re.Compile("<" + key + ">")
            spotCheckMatches := spotCheckPattern.SearchAll(self.Pattern)

            if spotCheckMatches.CountIndices() > 1 {
                return nil, RouteWithDuplicatedKeyError
            }
        }
    }

    alternativePattern = toriWebRoutingSimplePattern.ReplaceAll(self.Pattern, "(?P<${key}>[^/]+)")
    self.RePattern     = tori_re.Compile(alternativePattern)

    return self.RePattern, nil
}

func (self *Route) compileNonReversiblePattern() *tori_re.Expression {
    self.RePattern = tori_re.Compile(self.Pattern)

    return self.RePattern
}
