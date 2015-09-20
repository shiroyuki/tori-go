// Package re provides a simplified Regular Expression class
package re

import "fmt" // for debugging
import "regexp"

type Result struct {
    ItemList   []string
    Dictionary map[string]string
}

func NewResult(items []string, dictionary map[string]string) Result {
    return Result{
        ItemList:   items,
        Dictionary: dictionary,
    }
}

func (self *Result) Index(i int) *string {
    value := self.ItemList[i]

    return &value
}

func (self *Result) Key(k string) *string {
    value := self.Dictionary[k]

    return &value
}

func (self *Result) HasAny() bool {
    return self.Count() > 0
}

func (self *Result) Count() int {
    return self.CountIndices() + self.CountKeys()
}

func (self *Result) CountIndices() int {
    return len(self.ItemList)
}

func (self *Result) CountKeys() int {
    return len(self.Dictionary)
}

type Expression struct {
    Internal *regexp.Regexp
}

func Compile(pattern string) Expression {
    return Expression{regexp.MustCompile(pattern)}
}

func (self *Expression) Search(content string) Result {
    var itemList   []string
    var nextIndex  int
    var dictionary map[string]string

    matches := self.Internal.FindStringSubmatch(content)

    nextIndex  = 0
    dictionary = make(map[string]string)
    itemList   = make([]string, len(matches)) // allocate the memory to the maximum length first.

    fmt.Println(content, matches)

    if len(matches) == 0 {
        return NewResult(itemList, dictionary)
    }

    for i, name := range self.Internal.SubexpNames() {
        fmt.Println("I", i, "K", name, "KL", len(name), "V", matches[i])

        value := matches[i]

        if len(name) == 0 {
            itemList[nextIndex] = value

            nextIndex += 1

            continue
        }

        dictionary[name] = matches[i]
    }

    fmt.Println(itemList[:nextIndex], dictionary)

    return NewResult(
        itemList[:nextIndex], // minimize the memory usage by trimming itemList.
        dictionary,
    )
}
