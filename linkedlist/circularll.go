package linkedlist

type circularLinkedList struct {
	data int
	next *circularLinkedList
	prev *circularLinkedList
}

func CircularLinkedList()                     {}
func (ll *circularLinkedList) insertByIndex() {}
func (ll *circularLinkedList) deleteByIndex() {}
func (ll *circularLinkedList) print()         {}
