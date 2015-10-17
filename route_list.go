package tori

import "github.com/shiroyuki/re"

type Record struct {
    Id        string
    Method    string
    Route     *Route
    Action    *Action
    Cacheable bool
    Previous  *Record
    Next      *Record
}

type RecordList struct {
    head *Record
}

func (self *RecordList) Append(record *Record) {
    var cursor *Record

    if self.head == nil {
        self.head = record

        return
    }

    cursor = self.head

    // Iterate all the way to the tail of the list.
    for cursor != nil && cursor.Next != nil {
        var detachAndAppend = (cursor.Id == record.Id && cursor.Method == record.Method)

        // When the new record is referred to the same ID and method, the record
        // list detach the cursor from the list and the new record will become
        // the replacement but at the bottom of the order.
        if detachAndAppend {
            cursor.Previous.Next = cursor.Next     // The previous's next is the current's next.
            cursor.Next.Previous = cursor.Previous // The next's previous is the current's previous.
        }

        cursor = cursor.Next

        // Deference the links from the cursor
        if detachAndAppend {
            cursor.Next     = nil
            cursor.Previous = nil
        }
    }

    // Assign a new records.
    cursor.Next = record
}

func (self *RecordList) Find(method string, path string) (*Record, *re.MultipleResult) {
    var cursor *Record = self.head

    for cursor != nil && cursor.Next != nil {
        matches := cursor.Route.Match(method, path)

        if matches != nil {
            return cursor, matches
        }

        cursor = cursor.Next
    }

    return nil, nil
}
