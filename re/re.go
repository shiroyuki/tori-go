// Package re provides a simplified Regular Expression class
package re

import "regexp"

type Result struct {
    ItemList   []string
    Dictionary map[string]string
}

func NewResult(items []string, dictionary map[string]string) {
    return Result{
        ItemList:   items,
        Dictionary: dictionary,
    }
}

type Expression struct {
    Internal *regexp.Regexp
}

func Compile(pattern string) Expression {
    return Expression{regexp.MustCompile(pattern)}
}

func (self *Expression) Search(content string) map[string]string {
    matches := self.Internal.FindStringSubmatch(content)
    result  := make(map[string]string)

    for i, name := range self.Internal.SubexpNames() {
        result[name] = matches[i]
    }

    return result
}
