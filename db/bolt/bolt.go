package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/saromanov/alerting/structs"
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

// Set stores new alert message
func (b *Bolt) Set(m *structs.Message) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("alerts"))
		r, err := m.Marshal()
		if err != nil {
			return err
		}
		return b.Put([]byte(m.Namespace), r)
	})
}
