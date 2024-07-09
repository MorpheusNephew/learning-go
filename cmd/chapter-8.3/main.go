package main

import (
	"fmt"
	"strings"
)

/*
Write a generic singly linked list data type.
Each element can hold a comparable value and has a pointer to the next element in the list.
The methods to implement are as follows:

// adds a new element to the end of the linked list
Add(T)
// adds an element at the specified position in the linked list
Insert(T, int)
// returns the position of the supplied value, -1 if it's not present
Index (T) int
*/
func main() {
	linkedList := MakeLinkedList(4)

	linkedList.Add(8)
	linkedList.Add(2)
	linkedList.Add(11)
	linkedList.Add(19)
	fmt.Println(linkedList)
	linkedList.Insert(5, 2)
	fmt.Println(linkedList)
	fmt.Println(linkedList.Index(11))
	fmt.Println(linkedList.Index(5))
}

func MakeLinkedList[T comparable](val T) *Node[T] {
	return &Node[T]{Val: val}
}

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) Add(val T) {
	if n.Next != nil {
		n.Next.Add(val)
	} else {
		n.Next = &Node[T]{
			Val: val,
		}
	}
}

func (n *Node[T]) Insert(val T, index int) {
	if index == 0 {
		next := &Node[T]{Val: n.Val, Next: n.Next}
		n.Next = next
		n.Val = val
	} else if n.Next != nil {
		n.Next.Insert(val, index-1)
	} else {
		n.Next = &Node[T]{Val: val}
	}
}

func (n *Node[T]) Index(val T) int {
	index := 0

	current := n

	for current != nil {
		if current.Val == val {
			return index
		}

		current = current.Next
		index++
	}

	return -1
}

func (n *Node[T]) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprint(n.Val))

	next := n.Next

	for next != nil {
		builder.WriteString(fmt.Sprint("->", next.Val))

		next = next.Next
	}

	return builder.String()
}
