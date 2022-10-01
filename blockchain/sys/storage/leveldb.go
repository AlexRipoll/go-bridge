package storage

import "github.com/syndtr/goleveldb/leveldb"

type levelDB struct {
	db *leveldb.DB
}

func NewLevelDB(path string) (*levelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return 	&levelDB{db: db}, nil
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
