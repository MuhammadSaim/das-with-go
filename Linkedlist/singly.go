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

// delete the node from the end of the list
func (list *LinkedList) DeleteFromEnd() {
	// check there is any data in linked list
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// check if there is only head
	if list.head.next == nil {
		list.head = nil
		return
	}

	// temp node to store head
	temp := list.head

	// loop through the list to reaach the end of the list
	for temp.next.next != nil {
		temp = temp.next
	}

	// delete the node from the end
	temp.next = nil
}

// count the list
func (list *LinkedList) Count() int {

	// check there is any data in linked list
	if list.IsEmpty() {
		return 0
	}

	// initialize the count
	count := 1

	// temp node to store head
	temp := list.head

	// loop through the list
	for temp.next != nil {
		count++
		temp = temp.next
	}

	// return the counter
	return count
}

// reverse the linkedlist
func (list *LinkedList) Reverse() {

	// check there is any data in linked list
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// current node to store head
	current := list.head

	// check there is only a head then return the head no need to reverse the list
	if list.head.next == nil {
		return
	}

	// initialize the prev
	var prev *Node

	// loop through the list
	for current != nil {

		// store the next node
		next := current.next

		// save the prev node to the current node
		current.next = prev

		// move prev to the current node
		prev = current

		// move temp to next node
		// for further travercels
		current = next

	}

	list.head = prev

}

// linearly search value into a list
func (list *LinkedList) LinearSearch(value int) {

	// check there is any data in linked list
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// current node to store head
	current := list.head

	// found to set initialy false
	found := false

	// count the index
	index := 0

	// loop through the list
	for current != nil {

		// check the value in the list
		if current.value == value {
			found = true
			break // as soon we found break the loop
		}

		// update the index to 1 at a time
		index++

		// update the current with the next ndoe for
		// keep iterating
		current = current.next
	}

	// check found is true if it is means we found over value
	if found {
		fmt.Printf("%d is found at %d index\n", value, index)
		return
	}

	fmt.Printf("%d is not found\n", value)
}

// get the middle element of the list
func (list *LinkedList) Middle() {

	// check there is any data in linked list
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// create a two varibale which are holding
	// the head of the list we will talk about laters when we use those
	slow := list.head
	fast := list.head

	// loop through the list and find the middle
	// until fast is null and fast's next node is null
	for fast != nil && fast.next != nil {

		// now increment the slow 1 time
		slow = slow.next

		// and for the fast we will increment at 2x
		// so when fast get the nil slow will be at middle of the list
		fast = fast.next.next
	}

	// Now slow is middle just display the value
	fmt.Printf("Middle of the list is %d\n", slow.value)
}

// remove the duplicate from a list
func (list *LinkedList) RemoveDuplicate() {

	// check there is any data in linked list
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// current node to store head
	current := list.head

	// initialize map for memorize the element
	seen := make(map[int]bool)

	// set value of current node to hash map
	seen[current.value] = true

	// loop through the list
	for current.next != nil {

		// if value found in hashmap means we already visited it
		if seen[current.next.value] {

			// we have to think one step ahead
			// and if we found the value then next node will
			// be replace with the next node
			current.next = current.next.next

		} else {
			// register the next node value if not found and register to the map
			seen[current.next.value] = true
			// at the end rotate the current with the next for loop through
			current = current.next
		}

	}

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
	linkedList.InsertAtEnd(15)
	linkedList.InsertAtEnd(11)
	linkedList.InsertAtEnd(10)
	linkedList.InsertAtEnd(23)
	linkedList.InsertAtEnd(14)
	linkedList.InsertAtEnd(78)
	linkedList.InsertAtEnd(11)
	linkedList.InsertAtEnd(90)
	linkedList.InsertAtEnd(14)

	// display the list
	linkedList.Display()

	// remove duplicates
	linkedList.RemoveDuplicate()

	// display the list
	linkedList.Display()

	// middle of the list
	linkedList.Middle()

	// linear search
	linkedList.LinearSearch(12)
	linkedList.LinearSearch(190)

	// display the list
	linkedList.Display()

	// reverse the linkedlist
	linkedList.Reverse()

	// display the list
	linkedList.Display()

	// count the list
	fmt.Printf("Total nodes in the list are: %d\n", linkedList.Count())

	//delete by value
	linkedList.DeleteByValue(12)
	linkedList.DeleteByValue(4)

	// display the list
	linkedList.Display()

	// delete from the start
	linkedList.DeleteFromStart()
	linkedList.DeleteFromStart()

	// delete from the end
	linkedList.DeleteFromEnd()
	linkedList.DeleteFromEnd()
	linkedList.DeleteFromEnd()

	// display the list
	linkedList.Display()

	// count the list
	fmt.Printf("Total nodes in the list are: %d\n", linkedList.Count())

}
