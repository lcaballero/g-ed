package db

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"sync"
)

const dbFileName = "data.db"

type DB struct {
	access   *sync.Mutex
	closed   bool
	filename string
	data     map[string]interface{}
}

func LoadFromFile(filename string) (*DB, error) {
	db := &DB{
		access:   &sync.Mutex{},
		filename: filename,
		data:     make(map[string]interface{}),
	}
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		err = db.Save()
	} else {
		err = db.Load()
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Open() (*DB, error) {
	return LoadFromFile(dbFileName)
}

func (d *DB) Set(key string, val interface{}) error {
	d.access.Lock()
	defer d.access.Unlock()

	if d.closed {
		return errors.New("cannot write to closed db")
	}

	if key == "" {
		return errors.New("cannot save data to '' (empty) key")
	}
	d.data[key] = val
	return nil
}

func (d *DB) Has(key string) bool {
	d.access.Lock()
	defer d.access.Unlock()

	_, ok := d.data[key]
	return ok
}

func (d *DB) Close() error {
	if d.closed {
		return nil
	}

	err := d.Save()
	d.closed = true // probably better as atomic
	return err
}

func (d *DB) remove() error {
	err := os.Remove(d.filename)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Save() error {
	d.access.Lock()
	defer d.access.Unlock()

	if d.closed {
		return nil
	}

	bin, err := json.MarshalIndent(d.data, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(d.filename, bin, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Load() error {
	bin, err := ioutil.ReadFile(d.filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bin, &d.data)
	if err != nil {
		return err
	}
	return nil
}
