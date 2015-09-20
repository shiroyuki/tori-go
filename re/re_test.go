package re

import "testing"
import tori "../"

func TestReSearch(t *testing.T) {
    var assertion = tori.NewAssertion(t)
    var compiled  = Compile("shiroyuki")

    assertion.IsTrue(true, "Hehe")
}
