package cache

import tori "../"

// In-memory Cache Driver
type InMemoryCacheDriver struct {
    Driver // implements tori.cache.Driver

    Enigma      *tori.Enigma
    Compressed  bool
    MemoryTable map[string][]byte
}

func NewInMemoryCacheDriver(enigma *tori.Enigma, compressed bool) *InMemoryCacheDriver {
    imcd := InMemoryCacheDriver{
        Enigma:      enigma,
        Compressed:  compressed,
        MemoryTable: make(map[string][]byte),
    }

    return &imcd
}

func (self *InMemoryCacheDriver) Load(key string) []byte {
    content, existed := self.MemoryTable[key]

    if !existed {
        return nil
    }

    if !self.Compressed {
        return content
    }

    return (*self.Enigma).Decompress(content)
}

func (self *InMemoryCacheDriver) Save(key string, content []byte) {
    if self.MemoryTable == nil {
        self.MemoryTable = make(map[string][]byte)
    }

    if !self.Compressed {
        self.MemoryTable[key] = content

        return
    }

    compressed           := (*self.Enigma).Compress(content)
    self.MemoryTable[key] = compressed
}
