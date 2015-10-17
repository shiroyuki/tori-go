package tori

import "testing"
import "github.com/shiroyuki/tameshigiri"

func localTestWebRoutingRouteNewRoute(t *testing.T, givenPattern string, reversible bool) {
    var assertion tameshigiri.Assertion
    var newRoute  *Route

    assertion = tameshigiri.NewAssertion(t)
    newRoute  = NewRoute(givenPattern, reversible)

    assertion.IsTrue(
        newRoute.Pattern == givenPattern,
        "The registered pattern must be exactly the same as the given one.",
    )

    assertion.IsTrue(
        newRoute.Reversible == reversible,
        "The registered pattern must be reversible.",
    )

    assertion.IsTrue(
        newRoute.RePattern == nil,
        "The compiled pattern should not be compiled during initialization.",
    )
}

func TestWebRoutingRouteNewRouteNormal(t *testing.T) {
    localTestWebRoutingRouteNewRoute(t, "/", true)
    localTestWebRoutingRouteNewRoute(t, "/", false)
}

func TestWebRoutingReversibleRouteGetCompiledPatternOkay(t *testing.T) {
    var assertion    tameshigiri.Assertion
    var givenPattern string
    var newRoute     *Route

    assertion = tameshigiri.NewAssertion(t)

    givenPattern = "/user/<alias>"
    newRoute     = NewRoute(givenPattern, true)

    compiled, err := newRoute.GetCompiledPattern()

    assertion.IsTrue(compiled != nil, "Unexpected compiled pattern")
    assertion.IsTrue(err      == nil, "No error raised")
}

func TestWebRoutingReversibleRouteGetCompiledPatternFailed(t *testing.T) {
    var assertion    tameshigiri.Assertion
    var givenPattern string
    var newRoute     *Route

    assertion = tameshigiri.NewAssertion(t)

    givenPattern = "/user/<alias>/<alias>"
    newRoute     = NewRoute(givenPattern, true)

    compiled, err := newRoute.GetCompiledPattern()

    assertion.IsTrue(compiled == nil, "The pattern should not be compiled.")
    assertion.IsTrue(err      != nil, "An error should be raised.")
}

func TestWebRoutingNonReversibleRouteGetCompiledPattern(t *testing.T) {
    var assertion    tameshigiri.Assertion
    var givenPattern string
    var newRoute     *Route

    assertion = tameshigiri.NewAssertion(t)

    givenPattern = "/user/(?P<alias>[^/]+)"
    newRoute     = NewRoute(givenPattern, false)

    compiled, err := newRoute.GetCompiledPattern()

    assertion.IsTrue(compiled != nil, "Unexpected compiled pattern")
    assertion.IsTrue(err      == nil, "No error raised")
}

func TestRouteMatchOk(t *testing.T) {
    var assertion    tameshigiri.Assertion
    var givenPattern string
    var newRoute     *Route

    assertion = tameshigiri.NewAssertion(t)

    givenPattern = "/user/<alias>"
    newRoute     = NewRoute(givenPattern, true)

    matches := newRoute.Match("/user/shiroyuki")

    assertion.IsTrue(matches != nil, "There should be a match.")

    actual  := (*matches.Key("alias"))[0]

    assertion.Equals("shiroyuki", actual, "This should be 'shiroyuki'.")
}

func TestRouteMatchFails(t *testing.T) {
    var assertion    tameshigiri.Assertion
    var givenPattern string
    var newRoute     *Route

    assertion = tameshigiri.NewAssertion(t)

    givenPattern = "/user/<alias>"
    newRoute     = NewRoute(givenPattern, true)

    matches := newRoute.Match("/user/shiroyuki/asdf")

    assertion.IsTrue(matches == nil, "There should not be a match.")
}
