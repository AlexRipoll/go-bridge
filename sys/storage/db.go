package storage

type Storage interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
}
