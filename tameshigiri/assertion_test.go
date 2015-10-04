package tameshigiri

import "testing"

//
func ExampleAssertion() {
    var t         = &testing.T{}
    var assertion = Assertion{ T: t }
    var expected  = 123
    var actual    = 123

    assertion.IsTrue(expected == actual, "Ye")
    assertion.IsFalse(expected != actual, "Ne")
    assertion.Equals(expected, actual, "Eh?")

    // This example shows all passed assertions.
}
