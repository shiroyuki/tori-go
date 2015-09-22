// Package re provides a simplified Regular Expression class
package re // TODO rename this file to expression.go

//import "fmt" // for debugging
import "regexp"

type Expression struct {
    Internal *regexp.Regexp
}

func Compile(pattern string) Expression {
    return Expression{regexp.MustCompile(pattern)}
}

func (self *Expression) SearchOne(content string) SingleResult {
    matches := self.Internal.FindStringSubmatch(content)

    return self.makeSingleResult(matches)
}

func (self *Expression) SearchAll(content string) MultipleResult {
    var cursor     SingleResult
    var itemList   []string
    var dictionary map[string][]string
    var completeMatchIncluded bool

    completeMatchIncluded = false

    matches := self.Internal.FindAllStringSubmatch(content, -1)

    dictionary = make(map[string][]string)

    for i := range matches {
        cursor = self.makeSingleResult(matches[i])

        if !completeMatchIncluded {
            if itemList == nil {
                itemList = cursor.ItemList
            } else {
                itemList = self.combineList(itemList, cursor.ItemList)
            }
        } else {
            if itemList == nil {
                itemList = cursor.ItemList[1:]
            } else {
                itemList = self.combineList(itemList, cursor.ItemList[1:])
            }
        }

        for k, v := range cursor.Dictionary {
            _, ok := dictionary[k]

            if !ok {
                dictionary[k] = make([]string, 1)
            }

            dictionary[k] = append(dictionary[k], v)
        }
    }

    return NewMultipleResult(itemList, dictionary)
}

func (self *Expression) makeSingleResult(matches []string) SingleResult {
    var itemList   []string
    var nextIndex  int
    var dictionary map[string]string

    nextIndex  = 0
    dictionary = make(map[string]string)
    itemList   = make([]string, len(matches)) // allocate the memory to the maximum length first.

    //fmt.Println(content, matches)

    if len(matches) == 0 {
        return NewSingleResult(itemList, dictionary)
    }

    for i, name := range self.Internal.SubexpNames() {
        //fmt.Println("I", i, "K", name, "KL", len(name), "V", matches[i])

        value := matches[i]

        if len(name) == 0 {
            itemList[nextIndex] = value

            nextIndex += 1

            continue
        }

        dictionary[name] = matches[i]
    }

    //fmt.Println(itemList[:nextIndex], dictionary)

    return NewSingleResult(
        itemList[:nextIndex], // minimize the memory usage by trimming itemList.
        dictionary,
    )
}

func (self *Expression) combineList(a []string, b[]string) []string {
    var i int
    var c = make([]string, len(a) + len(b))

    for i = range a {
        c = append(c, a[i])
    }

    for i = range b {
        c = append(c, b[i])
    }

    return c
}
