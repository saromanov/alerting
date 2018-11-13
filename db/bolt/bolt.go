package bolt

import (
	"github.com/boltdb/bolt"
)

type Bolt struct {
	db *bolt.DB
}

// New initialize new Bolt
func New() *Bolt {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	return &Bolt{
		db: db,
	}
}
