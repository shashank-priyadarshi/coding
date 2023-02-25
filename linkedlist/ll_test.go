package main

import "testing"

func TestInsertAtHead(t *testing.T) {
	// Create a new list with one node
	list := &List{
		data: 1,
		node: nil,
	}

	// Insert a new node at the head
	list.insertAtHead(2)

	// Verify that the new head node has been created correctly
	if list.data != 2 {
		t.Errorf("Expected head node data to be 2, but got %d", list.data)
	}

	if list.node == nil {
		t.Errorf("Expected head node to point to original list, but was nil")
	} else if list.node.data != 1 {
		t.Errorf("Expected head node to point to original list with data 1, but got %d", list.node.data)
	}

	if list.node != list {
		t.Errorf("Expected original list node to point to new head node, but was %v", list.node)
	}
}
