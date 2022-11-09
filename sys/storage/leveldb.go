package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type levelDB struct {
	db *leveldb.DB
}

func NewLevelDB(path string) (Db, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &levelDB{db: db}, nil
}

func (l levelDB) Close() {
	l.db.Close()
}

func (l levelDB) Get(key []byte) ([]byte, error) {
	return l.db.Get(key, nil)
}

func (l levelDB) Put(key, value []byte) error {
	return l.db.Put(key, value, nil)
}

func (l levelDB) Delete(key []byte) error {
	return l.db.Delete(key, nil)
}

func (l levelDB) GetAllKeys() [][]byte {
	var keys [][]byte
	iterator := l.db.NewIterator(&util.Range{
		Start: nil,
		Limit: nil,
	}, nil)
	defer iterator.Release()

	for iterator.Next() {
		keys = append(keys, iterator.Key())
	}

	return keys
}