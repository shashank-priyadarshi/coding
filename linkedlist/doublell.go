package linkedlist

import "fmt"

type doublyLinkedList struct {
	data int
	prev *doublyLinkedList
	next *doublyLinkedList
}

func DoublyLinkedList() {
	ll := &doublyLinkedList{data: 1, prev: nil, next: nil}
	ll = ll.insertAtHead(0)
	ll.insertAtTail(2)
	ll = ll.insertByIndex(12345, 0)
	ll = ll.insertByIndex(11, 1)
	ll = ll.insertByIndex(123, 11)
	ll.deleteByKey(2)
	ll.deleteByIndex(2)
	ll.print()
}

func (ll *doublyLinkedList) insertAtHead(data int) *doublyLinkedList {
	// create new node with prev empty, next as current
	// point ll to new node
	ll.prev = &doublyLinkedList{
		data: data,
		prev: nil,
		next: ll,
	}
	return ll.prev
}

func (ll *doublyLinkedList) insertAtTail(data int) {
	curr := ll
	for {
		if curr.next == nil {
			curr.next = &doublyLinkedList{
				data: data,
				prev: curr,
				next: nil,
			}
			break
		}
		curr = curr.next
	}
}

func (ll *doublyLinkedList) insertByIndex(data, index int) *doublyLinkedList {
	// create new node
	// iterate over list, once index reached, insert
	if index <= 0 {
		ll.prev = &doublyLinkedList{
			data: data,
			prev: nil,
			next: ll,
		}
		return ll.prev
	}
	iter := 1
	ll = ll.next
	for {
		if iter == index {
			// node in current iteration will become next node
			// ancestor of current node will be the ancestor to new node
			// new node will become the ancestor to current node
			newNode := &doublyLinkedList{
				data: data,
				prev: ll.prev,
				next: ll,
			}
			ll = newNode
			break
		}
		if ll.next == nil {
			fmt.Println("index out of bounds")
			break
		}
		iter++
		ll = ll.next
	}
	return ll
}

func (ll *doublyLinkedList) deleteByKey(data int) {}

func (ll *doublyLinkedList) deleteByIndex(index int) {}

func (ll *doublyLinkedList) print() {
	curr := ll
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}
