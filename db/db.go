package db

type db struct {
	cntMap map[int]int
	keyMap map[string]int
}

// Create new db
func New() *db {
	var newdb db
	newdb.cntMap = make(map[int]int)
	newdb.keyMap = make(map[string]int)
	return &newdb
}

// Set a key value pair in the database, returns error
func (d *db) Set(key string, value int) error {

}

// Get a key value pair in the database, returns value and error
func (d *db) Get(key string) (int, error) {

}

// Unset a key value pair in the database, returns error
func (d *db) Unset(key string) error {

}

// Find number of keys equal to value in the database, returns number and error
func (d *db) NumEqualTo(value int) (int, error) {

}
