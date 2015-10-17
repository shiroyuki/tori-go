package tori

import "testing"
import tu "github.com/shiroyuki/tameshigiri"

func TestRecordListInitial(t *testing.T) {
    assertion  := tu.NewAssertion(t)
    recordList := RecordList{}

    assertion.IsFalse(recordList.HasAny(), "This list should have not had the head record.")
}

func TestRecordListAppendOneRecordAndFindWithOneResult(t *testing.T) {
    assertion  := tu.NewAssertion(t)
    recordList := RecordList{DebugMode: false}

    head := &Record{
        Method: METHOD_GET,
        Route:  NewRoute("/a/<k1>/<k2>", true),
    }

    assertion.IsFalse(recordList.HasAny(), "This list should not have the head record.")

    recordList.Append(head)

    assertion.IsTrue(recordList.HasAny(), "This list should have the head record.")

    r, m := recordList.Find(METHOD_GET, "/a/b/c")

    assertion.IsTrue(r != nil && m != nil, "There should be a match.")
}

func TestRecordListAppendOneRecordAndFindWithNoResult(t *testing.T) {
    assertion  := tu.NewAssertion(t)
    recordList := RecordList{DebugMode: true}

    head := &Record{
        Method: METHOD_GET,
        Route:  NewRoute("/a/<k1>/<k2>", true),
    }

    assertion.IsFalse(recordList.HasAny(), "This list should not have the head record.")

    recordList.Append(head)

    assertion.IsTrue(recordList.HasAny(), "This list should have the head record.")

    r, m := recordList.Find(METHOD_GET, "/a/b")

    assertion.IsTrue(r == nil && m == nil, "There should be no matches.")
}
