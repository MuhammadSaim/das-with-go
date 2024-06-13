package main

import (
	"fmt"

	"github.com/MuhammadSaim/dsa-with-go/linkedlist/doubly"
	"github.com/MuhammadSaim/dsa-with-go/linkedlist/singly"
)

func main() {

	// initialize the linkedlist
	linkedList := &singly.LinkedList{}

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

	// find in list
	linkedList.Find(3)

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

	// delete from the end
	linkedList.DeleteFromEnd()
	linkedList.DeleteFromEnd()
	linkedList.DeleteFromEnd()
	linkedList.DeleteFromEnd()

	// display the list
	linkedList.Display()

	// insert at end
	linkedList.InsertAtEnd(10)
	linkedList.InsertAtEnd(12)
	linkedList.InsertAtEnd(13)

	// check linked list is cycle
	linkedList.DetectCycle()

	print("\n\nDoubly Linked List\n\n")

	// doubly linked list
	doublyLinkedList := &doubly.LinkedList{}

	// display the list
	doublyLinkedList.Display()

	// insret at start
	doublyLinkedList.InsertAtStart(4)
	doublyLinkedList.InsertAtStart(5)

	// display the list
	doublyLinkedList.Display()

}
