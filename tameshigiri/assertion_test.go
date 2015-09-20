package tameshigiri

import "testing"

func ExampleAssertion() {
    var t         = &testing.T{}
    var assertion = Assertion{ T: t }
    var expected  = 123
    var actual    = 123

    assertion.Equals(expected, actual, "Something must be wrong.")
}
