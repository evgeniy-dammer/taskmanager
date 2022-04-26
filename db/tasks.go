package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func intToByte(val int) []byte {
	b := make([]byte, 8)

	binary.BigEndian.PutUint64(b, uint64(val))

	return b
}

func byteToInt(val []byte) int {
	return int(binary.BigEndian.Uint64(val))
}

func CreateTask(task string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		id64, err := bucket.NextSequence()

		id = int(id64)

		if err != nil {
			fmt.Println(err)
		}

		key := intToByte(int(id64))

		return bucket.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			tasks = append(tasks, Task{
				Key:   byteToInt(key),
				Value: string(value),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		return bucket.Delete(intToByte(key))
	})
}

func InitDB(dbPath string) error {
	var err error

	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
