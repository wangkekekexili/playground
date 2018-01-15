package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("test", 0666, &bolt.Options{Timeout: time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("first"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("hello"), []byte("world"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("first"))
		bytes := b.Get([]byte("test"))
		fmt.Println(bytes)
		bytes = b.Get([]byte("hello"))
		fmt.Println(string(bytes))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
