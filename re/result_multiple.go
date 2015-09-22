package re

// Result for single substing searches
//
// This class is designed to minimize memory usage. Hence, the class properties
// is not protected by design.
type MultipleResult struct {
    Result     // implement the interface
    ItemList   []string
    Dictionary map[string][]string
}

func NewMultipleResult(items []string, dictionary map[string][]string) MultipleResult {
    return MultipleResult{
        ItemList:   items,
        Dictionary: dictionary,
    }
}

func (self *MultipleResult) Index(i int) *string {
    value := self.ItemList[i]

    return &value
}

func (self *MultipleResult) Key(k string) *[]string {
    value := self.Dictionary[k]

    return &value
}

func (self *MultipleResult) HasAny() bool {
    return self.Count() > 0
}

func (self *MultipleResult) Count() int {
    return self.CountIndices() + self.CountKeys()
}

func (self *MultipleResult) CountIndices() int {
    return len(self.ItemList)
}

func (self *MultipleResult) CountKeys() int {
    return len(self.Dictionary)
}
