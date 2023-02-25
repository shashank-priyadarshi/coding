package main

import "fmt"

type List struct {
	data int
	node *List
}

func main() {
	ll := List{
		data: 1,
	}
	ll.insertAtHead(0)
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

func (ll *List) insertAtHead(data int) {
	// create a head node
	// add addr of linked list as head.node
	// edit linked list head to addr of head
	ll = &List{
		data: data,
		node: ll,
	}
}
func (ll *List) insertAtTail(data int) {
	// iterate over linked list
	// check if head.node is nil, meaning end of list
	// add address of new node to head.node
	head := ll
	for {
		if head.node == nil {
			head.node = &List{
				data: data,
			}
			break
		}
		head = head.node
	}
}
func (ll *List) insertByIndex(data, index int) {
	// iterate over list
	// check if iteration equals index
	// prev node will refer to addr of new node, new node will refer to addr of curr node
	iter, head, prev := 0, ll, &List{}
	for {
		if iter == index {
			new_node := List{
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
func (ll *List) deleteByKey(key int) {
	// iterate over list
	// if curr.data == key
	// prev.node will point to curr.node
	head, prev := ll, &List{}
	for {
		if head.data == key {
			if prev.node == nil {
				ll = head.node
			} else {
				prev.node = head.node
			}
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
func (ll *List) deleteByIndex(index int) {
	// iterate over list
	// if iter == index
	// prev.node will point to curr.node
	iter, head, prev := 0, ll, &List{}
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
func (ll *List) print() {
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
