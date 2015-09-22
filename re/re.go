// Package re provides a simplified Regular Expression class
package re // TODO rename this file to expression.go

import "fmt" // for debugging
import "regexp"

type Expression struct {
    Internal *regexp.Regexp
}

func Compile(pattern string) Expression {
    return Expression{regexp.MustCompile(pattern)}
}

func (self *Expression) SearchOne(content string) SingleResult {
    var itemList   []string
    var nextIndex  int
    var dictionary map[string]string

    matches := self.Internal.FindStringSubmatch(content)

    nextIndex  = 0
    dictionary = make(map[string]string)
    itemList   = make([]string, len(matches)) // allocate the memory to the maximum length first.

    fmt.Println(content, matches)

    if len(matches) == 0 {
        return NewSingleResult(itemList, dictionary)
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

    return NewSingleResult(
        itemList[:nextIndex], // minimize the memory usage by trimming itemList.
        dictionary,
    )
}

func (self *Expression) SearchAll(content string) *MultipleResult {
    //matches := self.Internal.FindAllStringSubmatch(content, -1)

    return nil
}
