package util

// Struct store all destructive modifications of a transaction
type transItem struct {
	trans map[string]int
	next  *transItem
}

// Linked list stack implementation to store transItems
type TransStack struct {
	head  *TransactionItem
	tail  *TransactionItem
	count int
}
