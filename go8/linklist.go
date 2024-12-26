package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// InsertAtEnd adds a new node with the given value at the end of the list
func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Delete removes the first occurrence of the node with the given value
func (list *LinkedList) Delete(data int) {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	// Search for the node to delete
	prev, found := list.SearchForDelete(data)

	if !found {
		fmt.Println("Value not found in the list")
		return
	}

	// If the target is the head node
	if prev == nil {
		list.head = list.head.next
		return
	}

	// Remove the node by adjusting the previous node's next pointer
	prev.next = prev.next.next
}

// SearchElement searches for the first occurrence of the given value in the linked list
// Returns a pointer to the node if found, otherwise returns nil
func (list *LinkedList) SearchElement(data int) *Node {
	current := list.head
	position := 0

	for current != nil {
		if current.data == data {
			fmt.Printf("Element %d found at position %d\n", data, position)
			return current
		}
		current = current.next
		position++
	}

	fmt.Printf("Element %d not found in the list\n", data)
	return nil
}

// SearchForDelete finds the node before the target value and returns it
// Returns the previous node and a boolean indicating if the target node exists
func (list *LinkedList) SearchForDelete(data int) (*Node, bool) {
	if list.head == nil {
		return nil, false
	}

	// If the node to be deleted is the head node
	if list.head.data == data {
		return nil, true
	}

	current := list.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	// If target node is found, return the previous node and true
	if current.next != nil {
		return current, true
	}

	// Target node not found
	return nil, false
}

// UpdateElement updates the first occurrence of the node with the given old value to the new value
func (list *LinkedList) UpdateElement(oldValue, newValue int) {
	node := list.SearchElement(oldValue)
	if node != nil {
		node.data = newValue
		fmt.Printf("Updated element %d to %d\n", oldValue, newValue)
	} else {
		fmt.Printf("Element %d not found, no update performed\n", oldValue)
	}
}

// Display prints all elements of the linked list
func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}

	// Adding elements to the linked list
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)

	fmt.Println("Linked List:")
	list.Display()

	// Deleting an element
	list.Delete(20)
	fmt.Println("After deleting 20:")
	list.Display()

	// Searching for elements
	list.SearchElement(10)
	list.SearchElement(50)

	// Updating an element
	list.UpdateElement(30, 35)
	fmt.Println("After updating 30 to 35:")
	list.Display()

	// Trying to delete an element not in the list
	list.Delete(50)
	fmt.Println("After trying to delete 50:")
	list.Display()
}






