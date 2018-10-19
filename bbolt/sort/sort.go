package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"

	"go.etcd.io/bbolt"
)

func int2key(i uint64) []byte {
	b := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func key2int(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func main() {
	db, err := bbolt.Open("data", 0644, nil)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucket([]byte("test"))
		if err != nil {
			return fmt.Errorf("failed to create bucket: %s", err)
		}
		for i := 0; i < 10; i++ {
			idx := rand.Intn(10000)
			println(idx)
			key := int2key(uint64(idx))
			if err := b.Put(key, []byte("value")); err != nil {
				return fmt.Errorf("failed to put: %s", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to update: %s", err)
	}
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("test"))
		c := b.Cursor()
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			fmt.Printf("data= %d: %s\n", key2int(k), v)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to view: %s", err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		if err := tx.DeleteBucket([]byte("test")); err != nil {
			return fmt.Errorf("failed to remove bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to delete: %s", err)
	}
}
