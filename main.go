package datastructures

import "fmt"

// Node representa um elemento da lista encadeada, armazenando um valor inteiro e um ponteiro para o próximo nó.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList representa uma lista encadeada com um ponteiro para o primeiro nó e um contador de tamanho.
type LinkedList struct {
	Initial *Node
	Size    int
}

// Add adiciona um novo nó ao final da lista encadeada.
func (T *LinkedList) Add(x int) {
	newNode := &Node{Value: x}

	if T.Initial == nil {
		T.Initial = newNode
	} else {
		current := T.Initial
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	T.Size++
}

// Print exibe os elementos da lista encadeada no formato "valor -> valor -> ... -> nil".
func (T *LinkedList) Print() {
	current := T.Initial
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

// RemoveLast remove o último elemento da lista encadeada.
func (T *LinkedList) RemoveLast() {
	if T.Initial == nil {
		return
	}

	// Caso a lista tenha apenas um elemento
	if T.Initial.Next == nil {
		T.Initial = nil
		T.Size--
		return
	}

	current := T.Initial
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
	T.Size--
}

// Get retorna o nó no índice especificado.
// Se o índice for inválido, retorna nil.
func (T *LinkedList) Get(index int) *Node {
	if T.Initial == nil || index < 0 || index >= T.Size {
		return nil
	}

	current := T.Initial
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current
}
