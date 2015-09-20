package tori

import "runtime"
import "testing"

// Static reference number of the processed assertions
var NumberOfProcessedAssertion uint32 = 0

type Assertion struct {
    T                *testing.T
    stackDumpEnabled bool
}

// Create a new assertion.
func NewAssertion(t *testing.T) Assertion {
    assertion := Assertion{ T: t }

    assertion.EnableStackDump()

    return assertion
}

// Check if the result is true.
//
// When the result is false, assuming that the given result is unexpected, the
// most recent call stacks (up to the size of 2KB) will be provided along with
// the human-readable description.
func (self *Assertion) IsTrue(result bool, description string) bool {
    NumberOfProcessedAssertion += 1

    if !result {
        self.T.Fail()

        if self.stackDumpEnabled {
            self.T.Logf("#%d FAILED\n", NumberOfProcessedAssertion)
            self.T.Log("#", NumberOfProcessedAssertion, description)
        }

        return false
    }

    return true
}

// Check if the result is false.
//
// This is the opposite of (*Assertion) IsTrue(...)
func (self *Assertion) IsFalse(result bool, description string) bool {
    return self.IsTrue(!result, description)
}

func (self *Assertion) Equals(expected interface{}, actual interface{}, description string) bool {
    var yes bool

    var stackDumpEnabled = self.stackDumpEnabled

    if stackDumpEnabled {
        self.DisableStackDump()
    }

    yes = self.IsTrue(expected == actual, "")

    if !yes {
        self.T.Logf("#%d FAILED\n", NumberOfProcessedAssertion)
        self.T.Log("#", NumberOfProcessedAssertion, description)
        self.T.Log("#", NumberOfProcessedAssertion, "Expected:", expected)
        self.T.Log("#", NumberOfProcessedAssertion, "Given:", actual)
    }

    if stackDumpEnabled {
        self.EnableStackDump()
    }

    return yes
}

// Enable stack dump
func (self *Assertion) EnableStackDump() {
    self.stackDumpEnabled = true
}

// Disable stack dump
func (self *Assertion) DisableStackDump() {
    self.stackDumpEnabled = false
}

// Dump call stacks
func (self *Assertion) dumpStack() {
    var buffer []byte

    if !self.stackDumpEnabled {
        return
    }

    buffer = make([]byte, 2048); // keep 2KB

    runtime.Stack(buffer, true)

    self.T.Logf("#%d Detail: %s\n", NumberOfProcessedAssertion, buffer)
}
