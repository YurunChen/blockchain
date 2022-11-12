package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("test.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败！")
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建一个bucket!")
			}
		}
		bucket.Put([]byte("1111"), []byte("hello"))
		bucket.Put([]byte("2222"), []byte("world"))
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket不存在，请检查！！")
		}
		v1 := bucket.Get([]byte("1111"))
		v2 := bucket.Get([]byte("2222"))
		fmt.Printf("数据1：%s\n", v1)
		fmt.Printf("数据2：%s\n", v2)
		return nil
	})

}
