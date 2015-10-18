package db

import "github.com/HuKeping/rbtree"

// Node to build our cntTree (implemented as rbtree)
type cntNode struct {
	val int
	cnt int
}

type db struct {
	cntTree rbtree.Rbtree
	keyMap  map[string]int
}

// Less function required by Item interface in rbtree
func (n cntNode) Less(than rbtree.Item) bool {
	return n.val < than.(cntNode).val
}

// Create new db
func New() *db {
	var newdb db
	newdb.cntTree = *rbtree.New()
	newdb.keyMap = make(map[string]int)
	return &newdb
}

// Set a key value pair in the database, returns error
func (d *db) Set(key string, value int) error {

}

// Get a key value pair in the database, returns value and error
func (d *db) Get(key string) (int, error) {

}

// Unset a key value pair in the database, returns erro
func (d *db) Unset(key string) error {

}

// Find number of keys equal to value in the database, returns number and error
func (d *db) NumEqualTo(value int) (int, error) {

}
