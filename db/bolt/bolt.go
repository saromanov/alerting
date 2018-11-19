package bolt

import (
	"encoding/json"
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
func (b *Bolt) Get(id string) (*structs.Message, error) {
	var msg *structs.Message
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("alerts"))
		v := bucket.Get([]byte(id))
		err := json.Unmarshal(v, &msg)
		if err != nil {
			return fmt.Errorf("unable to unmarshal message: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get message by id: %v", err)
	}
	return msg, nil
}

// View shows key-value pairs
func (b *Bolt) View() ([]*structs.Message, error) {
	result := []*structs.Message
	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("alerts"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			err := json.Unmarshal(v, &v)
			if err != nil {
				continue
			}
			result = append(result, v)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to search alerts: %v", err)
	}
	return result, nil
}
