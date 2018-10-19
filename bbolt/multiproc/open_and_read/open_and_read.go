package main

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open("../data", 0644, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}
	defer db.Close()
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("bucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s= %s\n", k, v)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to read database: %s", err)
	}
}
