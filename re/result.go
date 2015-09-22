package re

// Result interface
type Result interface {
    Index(i int)   *string
    Key(k string)  interface{}
    HasAny()       bool
    Count()        int
    CountIndices() int
    CountKeys()    int
}
