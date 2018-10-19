package main

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func main() {
	for {
		db, err := bbolt.Open("../data", 0644, nil)
		if err != nil {
			log.Fatalf("failed to open db: %s", err)
		}
		err = db.Update(func(tx *bbolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("bucket"))
			if err != nil {
				return fmt.Errorf("failed to get bucket: %s", err)
			}
			if err := b.Put([]byte(fmt.Sprint(time.Now())), []byte("data")); err != nil {
				return fmt.Errorf("failed to put data: %s", err)
			}
			return nil
		})
		if err != nil {
			log.Fatalf("failed to update data: %s", err)
		}
		fmt.Println("wrote")
		db.Close()
		time.Sleep(100 * time.Millisecond)
	}
}
