package main

type CircularLinkedList struct {
	data int
	next *CircularLinkedList
	prev *CircularLinkedList
}

func circularLinkedList()                     {}
func (ll *CircularLinkedList) insertByIndex() {}
func (ll *CircularLinkedList) deleteByIndex() {}
func (ll *CircularLinkedList) print()         {}
