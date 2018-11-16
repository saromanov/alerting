package bolt

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/saromanov/alerting/structs"
)

// Bolt provides definition for the Bolt instance
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

// Get retruns message by id
func (b *Bolt) Get() (*structs.Message, error) {
	return nil, nil
}

// View shows key-value pairs
func (b *Bolt) View() error {
	return b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("alerts"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
}
