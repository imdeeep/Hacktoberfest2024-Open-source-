package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Delete(value int) {
	if ll.head == nil {
		return
	}
	if ll.head.value == value {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}

	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Display()

	ll.Delete(2)
	ll.Display()

	ll.Delete(1)
	ll.Display()

	ll.Delete(3)
	ll.Display()
}
