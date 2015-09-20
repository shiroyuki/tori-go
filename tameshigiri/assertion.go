package tameshigiri

import "runtime"
import "testing"

// Static reference number of the processed assertions
var NumberOfProcessedAssertion uint32 = 0

// Assertion class
//
// When the assertion fails, assuming that the given result is unexpected, the
// most recent call stacks (up to the size of 2KB) will be provided along with
// the human-readable description.
//
// Example:
//
//      package panda
//
//      import "testing"
//      import "github.com/shiroyuki/tori-go/tameshigiri"
//
//      func TestPanda(t *testing.T) {
//          var assertion = tameshigiri.NewAssertion(t)
//          var expected  = 123
//          var actual    = 123
//
//          t.Equals(expected, actual, "Something must be wrong.")
//      }
//
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
func (self *Assertion) IsFalse(result bool, description string) bool {
    return self.IsTrue(!result, description)
}

// Assert if the actual value is equal to the expected value
func (self *Assertion) Equals(expected interface{}, actual interface{}, description string) bool {
    var yes bool

    var stackDumpEnabled = self.stackDumpEnabled

    if stackDumpEnabled {
        self.DisableStackDump()
    }

    yes = self.IsTrue(expected == actual, "")

    if stackDumpEnabled && !yes {
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
