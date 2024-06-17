package doubly

import "fmt"

// define the Node for doubly linkedlist
type Node struct {
	prev  *Node
	value int
	next  *Node
}

// now define the linked list
type LinkedList struct {
	head *Node
}

// insert the value at the start of the list
func (list *LinkedList) InsertAtStart(value int) {

	// create a new node with value
	newNode := &Node{value: value}

	// check if the head is none
	if list.head == nil {
		list.head = newNode
		return
	}

	// store new node to the temp
	tmp := newNode
	// save head to the temp next
	tmp.next = list.head
	// save tmp prev in head
	list.head.prev = tmp
	// now change the head with the start node
	list.head = tmp
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
		fmt.Printf("%d ⇆ ", current.value)
		// save the next node to the current for keep forwading the loop
		current = current.next
	}

	fmt.Println("nil")

}

// a function check list is empty of not
func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

// count the list
func (list *LinkedList) Count() int {

	// check there is any data in linked list
	if list.IsEmpty() {
		return 0
	}

	// initialize the count
	count := 0

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


// find the node in list
func (list *LinkedList) Find(index int) {

	// check there is any data in linked list
	if list.IsEmpty() {
		fmt.Println("Ooops! There is no data to display. please insert the data.")
		return
	}

	// get the count of the linked list
	count := list.Count()

	// check the count
	if index > count {
		fmt.Println("Ooops! index is not found.")
		return
	}

	// current node to store head
	current := list.head

	// loop through the count
	for range index {
		current = current.next
	}

	// print the current value
	fmt.Printf("So %d was found at index %d\n", current.value, index)

}