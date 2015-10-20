package util

// Since we can't set an int to nil, need to make a type which allows
// us to say, before x happened "name" was set to nil (meaning it didn't exist)
type value struct {
	val   int
	isNil bool
}

// Struct to store all destructive modifications of a transaction
type transItem struct {
	trans map[string]value
	next  *transItem
}

// Linked list stack implementation to store transItems
type TransStack struct {
	head *transItem
	tail *transItem
	cnt  int
}

// Makes new Transaction Stack and return pointer
func New() *TransStack {
	var newTS TransStack
	newTS.head = nil
	newTS.tail = nil
	newTS.cnt = 0
	return &newTS
}

// When doing a destructive modification, use StoreVal to save value of
func (t *TransStack) StoreVal() {

}
