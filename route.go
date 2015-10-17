package tori

import "github.com/shiroyuki/re"

var toriWebRoutingSimplePattern = re.Compile("<(?P<key>[^>]+)>")

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
type Route struct {
    Id         string
    Method     string
    Pattern    string
    RePattern  *re.Expression
    Handler    Action
    Reversible bool
    Cacheable  bool
}

// Create a route.
func NewRoute(pattern string, reversible bool) *Route {
    // TODO add the assertion to check if the pattern has a prefix "/". Raise exceptions if necessary.

    return &Route{
        Pattern:    pattern,
        Reversible: reversible,
        RePattern:  nil,   // This is lazy-loading.
        Cacheable:  false, // By default, no response is cacheable.
    }
}

func (self *Route) GetCompiledPattern() (*re.Expression, error) {
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

func (self *Route) Match(method string, path string) (*re.MultipleResult) {
    pattern, err := self.GetCompiledPattern()

    if err != nil {
        return nil
    }

    result := pattern.SearchAll(path)

    if self.Method != method || !result.HasAny() {
        return nil
    }

    return &result
}

func (self *Route) compileReversiblePattern() (*re.Expression, error) {
    var alternativePattern string

    matches := toriWebRoutingSimplePattern.SearchAll(self.Pattern)

    if matches.HasAny() {
        values := matches.Key("key")

        for _, key := range *values {
            spotCheckPattern := re.Compile("<" + key + ">")
            spotCheckMatches := spotCheckPattern.SearchAll(self.Pattern)

            if spotCheckMatches.CountIndices() > 1 {
                return nil, RouteWithDuplicatedKeyError
            }
        }
    }

    alternativePattern = toriWebRoutingSimplePattern.ReplaceAll(self.Pattern, "(?P<${key}>[^/]+)")
    self.RePattern     = re.Compile(alternativePattern)

    return self.RePattern, nil
}

func (self *Route) compileNonReversiblePattern() *re.Expression {
    self.RePattern = re.Compile(self.Pattern)

    return self.RePattern
}
