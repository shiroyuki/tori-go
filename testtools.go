package tori

import "testing"

var NumberOfProcessedAssertion uint32 = 0

type Assertion struct {
    T *testing.T
}

func NewAssertion(t *testing.T) Assertion {
    return Assertion{ t }
}

func (self *Assertion) IsTrue(result bool, reason string) bool {
    NumberOfProcessedAssertion += 1

    if !result {
        self.T.Log("[ Assertion", NumberOfProcessedAssertion, "]", reason)
        self.T.Fail()

        return false
    }

    self.T.Log("[Assertion", NumberOfProcessedAssertion, "] Passed")

    return true
}

func (self *Assertion) IsFalse(result bool, reason string) bool {
    return !self.IsTrue(!result, reason)
}
