package main

type DoublyLinkedList struct {
	data int
	prev *DoublyLinkedList
	next *DoublyLinkedList
}

func doublyLinkedList() {}

func (ll *DoublyLinkedList) insertAtHead(data int)         {}
func (ll *DoublyLinkedList) insertAtTail(data int)         {}
func (ll *DoublyLinkedList) insertByIndex(data, index int) {}
func (ll *DoublyLinkedList) deleteByKey(data int)          {}
func (ll *DoublyLinkedList) deleteByIndex(index int)       {}
func (ll *DoublyLinkedList) print()                        {}
