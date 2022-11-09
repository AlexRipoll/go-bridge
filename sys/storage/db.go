package storage

type Db interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	GetAllKeys() [][]byte
}
