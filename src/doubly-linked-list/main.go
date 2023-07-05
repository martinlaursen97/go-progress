package main

//////////////////////
// LinkedList in Go //
//////////////////////

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

type LinkedList[T any] interface {
	Add(value T)
	Size() int
	Print()
	Pop()
}

type DoublyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func main() {
	ll := DoublyLinkedList[int]{}

	ll.Add(1)
	ll.Add(3)
	ll.Add(-5)

	fmt.Printf("Size: %d\n", ll.Size())

	ll.Add(5)
	ll.Add(1)
	ll.Add(-9)

	fmt.Printf("Size: %d\n", ll.Size())
	ll.Print()

	ll.Pop()
	ll.Pop()
	ll.Pop()

	fmt.Printf("Size: %d\n", ll.Size())
	ll.Print()
}

func (ll *DoublyLinkedList[T]) Add(value T) {
	nN := Node[T]{
		Data: value,
	}

	if ll.Head == nil {
		ll.Head = &nN
		ll.Tail = &nN
		return
	}

	ll.Tail.Next = &nN
	nN.Prev = ll.Tail
	ll.Tail = &nN
}

func (ll *DoublyLinkedList[T]) Size() int {
	if ll.Head == nil {
		return 0
	}

	sum := 0
	current := ll.Head
	for current != nil {
		current = current.Next
		sum += 1
	}
	return sum
}

func (ll *DoublyLinkedList[T]) Print() {
	if ll.Head == nil {
		return
	}

	current := ll.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (ll *DoublyLinkedList[T]) Pop() {
	if ll.Head == nil {
		return
	}
	if ll.Tail == nil {
		ll.Head = nil
	}

	ll.Tail = ll.Tail.Prev
	ll.Tail.Next = nil
}
