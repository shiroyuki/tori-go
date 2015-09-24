// Package re provides a simplified Regular Expression class
package re // TODO rename this file to expression.go

//import "fmt" // for debugging
import "regexp"

type Expression struct {
    Internal *regexp.Regexp
}

func Compile(pattern string) *Expression {
    return &Expression{regexp.MustCompile(pattern)}
}

func (self *Expression) SearchOne(content string) SingleResult {
    matches := self.Internal.FindStringSubmatch(content)

    return self.makeSingleResult(matches)
}

func (self *Expression) SearchAll(content string) MultipleResult {
    var cursor     SingleResult
    var itemList   []string
    var dictionary map[string][]string

    matches := self.Internal.FindAllStringSubmatch(content, -1)

    dictionary = make(map[string][]string)

    for i := range matches {
        cursor = self.makeSingleResult(matches[i])

        if itemList == nil {
            itemList = cursor.ItemList
        } else {
            itemList = self.combineList(itemList, cursor.ItemList)
        }

        for k, v := range cursor.Dictionary {
            _, ok := dictionary[k]

            if !ok {
                dictionary[k] = make([]string, 0)
            }

            dictionary[k] = append(dictionary[k], v)
        }
    }

    return NewMultipleResult(itemList, dictionary)
}

// Replace all matches with the given replacement.
//
// This is a simple proxy to Regexp.ReplaceAllString. No test required.
func (self *Expression) ReplaceAll(content string, replacement string) string {
    return self.Internal.ReplaceAllString(content, replacement)
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
    var j int = 0
    var c = make([]string, len(a) + len(b))

    for i = range a {
        c[j] = a[i]
        j += 1
    }

    for i = range b {
        c[j] = b[i]
        j += 1
    }

    return c
}
