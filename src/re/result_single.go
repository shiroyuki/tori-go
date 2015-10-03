package re

// Result for single substing searches
//
// This class is designed to minimize memory usage. Hence, the class properties
// is not protected by design.
type SingleResult struct {
    Result     // implement the interface
    ItemList   []string
    Dictionary map[string]string
}

func NewSingleResult(items []string, dictionary map[string]string) SingleResult {
    return SingleResult{
        ItemList:   items,
        Dictionary: dictionary,
    }
}

func (self *SingleResult) Index(i int) *string {
    value := self.ItemList[i]

    return &value
}

func (self *SingleResult) Key(k string) *string {
    value := self.Dictionary[k]

    return &value
}

func (self *SingleResult) HasAny() bool {
    return self.Count() > 0
}

func (self *SingleResult) Count() int {
    return self.CountIndices() + self.CountKeys()
}

func (self *SingleResult) CountIndices() int {
    return len(self.ItemList)
}

func (self *SingleResult) CountKeys() int {
    return len(self.Dictionary)
}
