package main

import (
	"fmt"
)

// defining the Node for the linkedlist
type Node struct {
	value int
	next  *Node
}

// now define the linked list
type LinkedList struct {
	head *Node
}

// insert the value beginning of the list
func (list *LinkedList) InsertAtStart(value int) {
	// create a new node with the given value
	newNode := &Node{value: value}
	// now save the head address to the new next for making new node head
	newNode.next = list.head
	// now assign list head to the newNode
	list.head = newNode
}

// insert the value at the end of the list
func (list *LinkedList) InsertAtEnd(value int) {

	// create a new node
	newNode := &Node{value: value}

	// check if the linked list is empty insert the node at the start
	if list.IsEmpty() {
		list.head = newNode
		return
	}

	// save the value of the head to the temp node
	current := list.head

	// now loop through the list untile we found next fill be null
	for current.next != nil {
		current = current.next
	}

	// now last node is in current so we can save the address of the new node to the current next
	current.next = newNode

}

// display the linked list
func (list *LinkedList) Display() {

	// check there is data or linked list exists
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// save list head in temp var
	current := list.head

	// loop through the list
	for current != nil {
		// display the value with the link arrow
		fmt.Printf("%d -> ", current.value)
		// save the next node to the current for keep forwading the loop
		current = current.next
	}

	fmt.Println("nil")

}

// a function check list is empty of not
func (list *LinkedList) IsEmpty() bool {
	if list.head != nil {
		return false
	}
	return true
}

// delete by value delete the node by their value
func (list *LinkedList) DeleteByValue(value int) {

	// save the head into temp variable
	current := list.head

	// check if head has a value don't need to loop through
	if list.head.value == value {
		list.head = list.head.next
		return
	}

	// loop thorugh the linked list and find for a value
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	// delete the found node
	if current.next != nil {
		current.next = current.next.next
	}

}

// delete the node from the start of the list
func (list *LinkedList) DeleteFromStart() {
	// check there is any data in linked list
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// temp node to store head
	temp := list.head

	// check there is only one node if there is then delete the node
	if temp.next == nil {
		list.head = nil
	}

	// save the next node address to the head
	list.head = temp.next
}

func main() {

	// initialize the linkedlist
	linkedList := &LinkedList{}

	// display the list
	linkedList.Display()

	// insret at start
	linkedList.InsertAtStart(4)
	linkedList.InsertAtStart(5)

	// insert at end
	linkedList.InsertAtEnd(10)
	linkedList.InsertAtEnd(12)
	linkedList.InsertAtEnd(13)

	// display the list
	linkedList.Display()

	//delete by value
	linkedList.DeleteByValue(12)
	linkedList.DeleteByValue(4)

	// display the list
	linkedList.Display()

	// delete from the start
	linkedList.DeleteFromStart()
	linkedList.DeleteFromStart()

	// display the list
	linkedList.Display()

}
