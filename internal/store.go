package store

import (
	"encoding/gob"
	"os"
	"sync"
)

type Database struct {
	data map[string]any
	lock sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{
		data: make(map[string]any),
	}
}

func (db *Database) Set(key string, value any) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.data[key] = value
}

func (db *Database) Get(key string) (any, bool) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	value, exists := db.data[key]
	return value, exists
}

func (db *Database) Persist(fileName string) error {
	db.lock.RLock()
	defer db.lock.RUnlock()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(db.data); err != nil {
		return err
	}
	return nil
}

func (db *Database) Load(fileName string) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&db.data); err != nil {
		return err
	}
	return nil	
}