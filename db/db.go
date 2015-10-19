package db

import "errors"

type Database struct {
	cntMap map[int]int
	keyMap map[string]int
}

// Create new db
func New() *Database {
	var newdb Database
	newdb.cntMap = make(map[int]int)
	newdb.keyMap = make(map[string]int)
	return &newdb
}

// Set a key value pair in the database, returns error
func (d *Database) Set(key string, value int) error {
	d.keyMap[key] = value
	d.cntMap[value] += 1
	return nil
}

// Get a key value pair in the database, returns value and error
func (d *Database) Get(key string) (int, error) {
	v, ok := d.keyMap[key]
	if !ok {
		return 0, errors.New("ERROR: Key does not exist")
	} else {
		return v, nil
	}
}

// Unset a key value pair in the database, returns error
func (d *Database) Unset(key string) error {
	if v, ok := d.keyMap[key]; ok {
		delete(d.keyMap, key)
		d.cntMap[v] -= 1
		return nil
	} else {
		return errors.New("ERROR: Key does not exist")
	}
}

// Find number of keys equal to value in the database, returns number and error
func (d *Database) NumEqualTo(value int) (int, error) {
	if v, ok := d.cntMap[value]; ok {
		return v, nil
	} else {
		return 0, errors.New("ERROR: Key does not exist")
	}
}
