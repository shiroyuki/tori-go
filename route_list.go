package tori

import "log"
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
    head      *Record
    DebugMode bool
}

func (self *RecordList) Append(record *Record) {
    var cursor *Record

    if self.head == nil {
        self.log("Make " + (*record.Route).Pattern + " the head of the list.")

        self.head = record

        return
    }

    self.log("Appending " + (*record.Route).Pattern + " at the end of the list...")

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
    var cursor   *Record = self.head
    var firstLine string = method + " " + path

    self.log(firstLine + ": Looking for a record...")

    for cursor != nil {
        matches          := cursor.Route.Match(path)
        isMethodExpected := (*cursor).Method == method

        self.log(firstLine + ": Compare against " + cursor.Route.Pattern)

        if matches != nil && isMethodExpected {
            self.log(firstLine + ": Found a match.")

            return cursor, matches
        }

        if isMethodExpected {
            self.log(firstLine + ": Not a match due to unexpected method.")
        }

        if cursor.Next == nil {
            self.log(firstLine + ": End of iteration")

            break
        }

        cursor = cursor.Next
    }

    self.log(firstLine + ": Failed to find a match.")

    return nil, nil
}

func (self *RecordList) HasAny() bool {
    return self.head != nil
}

func (self *RecordList) log(message string) {
    if !self.DebugMode {
        return
    }

    log.Println(message)
}
