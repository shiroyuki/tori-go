package cache

type Driver interface {
    Load(key string) []byte
    Save(key string, content []byte)
}
