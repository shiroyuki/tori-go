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
    var compiled             tori_re.Expression
    var simpleRoutingPattern tori_re.Expression
    var alternativePattern   string

    simpleRoutingPattern = tori_re.Compile("<(?P<key>[^>]+)>")
    matches := simpleRoutingPattern.SearchAll(self.Pattern)

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

    alternativePattern = simpleRoutingPattern.ReplaceAll(self.Pattern, "(?P<${key}>[^/]+)")

    compiled       = tori_re.Compile(alternativePattern)
    self.RePattern = &compiled

    return self.RePattern, nil
}

func (self *Route) compileNonReversiblePattern() *tori_re.Expression {
    var compiled tori_re.Expression

    compiled       = tori_re.Compile(self.Pattern)
    self.RePattern = &compiled

    return self.RePattern
}
