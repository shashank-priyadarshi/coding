package main

import "fmt"

type SinglyLinkedList struct {
	data int
	node *SinglyLinkedList
}

func singlyLinkedList() {
	ll := &SinglyLinkedList{
		data: 1,
		node: nil,
	}
	ll = ll.insertAtHead(0)
	fmt.Println("printing after inserting element at head")
	ll.print()
	ll.insertAtTail(3)
	fmt.Println("printing after inserting element at tail")
	ll.print()
	ll.insertByIndex(2, 2)
	ll.insertByIndex(2, 5)
	fmt.Println("printing after inserting element at index 2")
	ll.print()
	ll.deleteByKey(3)
	ll.deleteByKey(7)
	fmt.Println("printing after deleting element")
	ll.print()
	ll.deleteByIndex(4)
	ll.deleteByIndex(2)
	fmt.Println("printing after deleting by index")
	ll.print()
}

func (ll *SinglyLinkedList) insertAtHead(data int) *SinglyLinkedList {
	// create a head node
	// add addr of linked list as head.node
	// edit linked list head to addr of head
	return &SinglyLinkedList{
		data: data,
		node: ll,
	}
}
func (ll *SinglyLinkedList) insertAtTail(data int) {
	// iterate over linked list
	// check if head.node is nil, meaning end of list
	// add address of new node to head.node
	head := ll
	for {
		if head.node == nil {
			head.node = &SinglyLinkedList{
				data: data,
			}
			break
		}
		head = head.node
	}
}
func (ll *SinglyLinkedList) insertByIndex(data, index int) {
	// iterate over list
	// check if iteration equals index
	// prev node will refer to addr of new node, new node will refer to addr of curr node
	iter, head, prev := 0, ll, &SinglyLinkedList{}
	for {
		if iter == index {
			new_node := SinglyLinkedList{
				data: data,
				node: head,
			}
			if prev.node == nil {
				ll = &new_node
			} else {
				prev.node = &new_node
			}
			break
		}
		if head.node == nil {
			fmt.Println("index out of range")
			break
		}
		prev = head
		iter++
		head = head.node
	}
}
func (ll *SinglyLinkedList) deleteByKey(key int) {
	// iterate over list
	// if curr.data == key
	// prev.node will point to curr.node
	head, prev := ll, &SinglyLinkedList{}
	for {
		if head.data == key {
			prev.node = head.node
			break
		}
		if head.node == nil {
			fmt.Println("key not found in list")
			break
		}
		prev = head
		head = head.node
	}
}
func (ll *SinglyLinkedList) deleteByIndex(index int) {
	// iterate over list
	// if iter == index
	// prev.node will point to curr.node
	iter, head, prev := 0, ll, &SinglyLinkedList{}
	for {
		if iter == index {
			if prev.node == nil {
				ll = head.node
			} else {
				prev.node = head.node
			}
			break
		}
		if head.node == nil {
			fmt.Println("index out of range")
			break
		}
		iter++
		prev = head
		head = head.node
	}
}
func (ll *SinglyLinkedList) print() {
	head := ll
	for {
		fmt.Print(head.data)
		if head.node == nil {
			fmt.Println()
			break
		}
		head = head.node
	}
}
