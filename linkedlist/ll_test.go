package main

import "testing"

func TestInsertAtHead(t *testing.T) {
	// Create a new list with one node
	list := &List{
		data: 1,
		node: nil,
	}

	originalHead := list

	// Insert a new node at the head
	list = list.insertAtHead(2)

	// Verify that the new head node has been created correctly
	if list.data != 2 {
		t.Errorf("Expected head node data to be 2, but got %d", list.data)
	}

	if list.node == nil {
		t.Errorf("Expected head node to point to original list, but was nil")
	} else if list.node.data != 1 {
		t.Errorf("Expected head node to point to original list with data 1, but got %d", list.node.data)
	}

	if list.node != originalHead {
		t.Errorf("Expected original list node to point to new head node, but was %v", list.node)
	}
}

func TestInsertAtTail(t *testing.T) {
	// Create a new list with one node
	list := &List{
		data: 1,
		node: nil,
	}

	// Insert a new node at the tail
	list.insertAtTail(2)

	// Verify that the new tail node has been created correctly
	if list.node == nil {
		t.Errorf("Expected tail node to be created, but was nil")
	} else if list.node.data != 2 {
		t.Errorf("Expected tail node to have data 2, but got %d", list.node.data)
	}

	// Verify that current head node is not replaced with new node
	if list.data != 1 {
		t.Errorf("Expected head node to remain unchanged, but got %d", list.data)
	}
}

func TestInsertByIndex(t *testing.T) {
	// Create a new list with one node
	list := &List{
		data: 1,
		node: nil,
	}

	// Insert a new node at the tail
	list.insertAtTail(2)

	// Insert a new node at index 1
	list.insertByIndex(3, 1)

	// Verify that the new node has been inserted correctly
	if list.node == nil {
		t.Errorf("Expected node to be created, but was nil")
	} else if list.node.data != 3 {
		t.Errorf("Expected node to have data 3, but got %d", list.node.data)
	}

	// Verify that the new node points to the original tail node
	if list.node.node == nil {
		t.Errorf("Expected node to point to original tail node, but was nil")
	} else if list.node.node.data != 2 {
		t.Errorf("Expected node to point to original tail node with data 2, but got %d", list.node.node.data)
	}

	// Verify that current head node points to original head node
	if list.data != 1 {
		t.Errorf("Expected head node to remain unchanged, but got %d", list.data)
	}
}

func TestDeleteByKey(t *testing.T) {
	// Create a new list with one node
	list := &List{
		data: 1,
		node: nil,
	}

	// Insert a new node at the tail
	list.insertAtTail(2)

	// Verify that new node has been inserted at tail
	if list.node == nil {
		t.Errorf("Expected node to be created, but was nil")
	} else if list.node.data != 2 {
		t.Errorf("Expected node to have data 2, but got %d", list.node.data)
	}

	// Delete node with key 2
	list.deleteByKey(2)

	// Verify that the node has been deleted correctly
	if list.node != nil {
		t.Errorf("Expected node to be deleted, but was %v", list.node)
	}

	// Verify that current head node points to original head node
	if list.data != 1 {
		t.Errorf("Expected head node to remain unchanged, but got %d", list.data)
	}

	// Delete node with key 1
	list.deleteByKey(1)

	// Verify that the node has been deleted correctly
	if list != nil {
		t.Errorf("Expected node to be deleted, but was %v", list)
	}
}

func TestDeleteByIndex(t *testing.T) {
	// Create a new list with one node
	list := &List{
		data: 1,
		node: nil,
	}

	// Insert a new node at the tail
	list.insertAtTail(2)

	// Verify that new node has been inserted at tail
	if list.node == nil {
		t.Errorf("Expected node to be created, but was nil")
	} else if list.node.data != 2 {
		t.Errorf("Expected node to have data 2, but got %d", list.node.data)
	}

	// Delete node at index 1
	list.deleteByIndex(1)

	// Verify that the node has been deleted correctly
	if list.node != nil {
		t.Errorf("Expected node to be deleted, but was %v", list.node)
	}

	// Verify that current head node points to original head node
	if list.data != 1 {
		t.Errorf("Expected head node to remain unchanged, but got %d", list.data)
	}

	// Delete node at index 0
	list.deleteByIndex(0)

	// Verify that the node has been deleted correctly
	if list != nil {
		t.Errorf("Expected node to be deleted, but was %v", list)
	}
}

func TestPrintList(t *testing.T) {
	// Create a new list with one node
	list := &List{
		data: 1,
		node: nil,
	}

	// Insert a new node at the tail
	list.insertAtTail(2)

	// Verify that new node has been inserted at tail
	if list.node == nil {
		t.Errorf("Expected node to be created, but was nil")
	} else if list.node.data != 2 {
		t.Errorf("Expected node to have data 2, but got %d", list.node.data)
	}

	// Print the list
	list.print()
}
