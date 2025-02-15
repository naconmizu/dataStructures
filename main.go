package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Initial *Node
	Size    int
}

func (T *LinkedList) Add(x int) {
	newNode := &Node{Value: x}

	// Se a lista estiver vazia
	if T.Initial == nil {
		T.Initial = newNode
	} else {
		// Percorre até o último nó
		current := T.Initial
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	T.Size++
}

func (T *LinkedList) Print() {
	current := T.Initial
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func (T *LinkedList) removeLast() {
	if T.Initial == nil {
		return
	}
	current := T.Initial
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
	T.Size--

}
func (T *LinkedList) get(index int) *Node {
	if T.Initial == nil {
		panic("LinkedList not initialized or empty")
	}
	if index > T.Size || index < 0 {
		panic("Index out of bounds")
	}

	current := T.Initial
	for i := 0; i < index; i++ {
    current = current.Next
  }
	return current

}

func main() {
	list := LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()
	list.removeLast()
	list.Print()

}
