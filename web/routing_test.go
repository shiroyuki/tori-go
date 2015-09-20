package web

import "testing"
import "../tameshigiri"

func localTestWebRoutingRouteNewRoute(t *testing.T, givenPattern string, reversible bool) {
    var assertion tameshigiri.Assertion
    var newRoute  Route

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

func _TestWebRoutingReversibleRouteGetCompiledPattern(t *testing.T) {
    var assertion    tameshigiri.Assertion
    var givenPattern string
    var newRoute     Route

    assertion = tameshigiri.NewAssertion(t)

    givenPattern = "/user/{alias}"
    newRoute     = NewRoute(givenPattern, true)

    compiled := newRoute.GetCompiledPattern()

    assertion.IsTrue(compiled != nil, "Unexpected compiled pattern")
}

func TestWebRoutingNonReversibleRouteGetCompiledPattern(t *testing.T) {
    var assertion    tameshigiri.Assertion
    var givenPattern string
    var newRoute     Route

    assertion = tameshigiri.NewAssertion(t)

    givenPattern = "/user/(?P<alias>[^/]+)"
    newRoute     = NewRoute(givenPattern, false)

    assertion.IsTrue(newRoute.GetCompiledPattern() != nil, "Unexpected compiled pattern")
}
