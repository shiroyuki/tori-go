package routing

type Record struct {
    Id        string
    Method    string
    Handler   Action
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
    for cursor.Next != nil {
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

    cursor.Next = record
}

func (self *RecordList) Search(requestPath string) *Record {
    return nil
}
