package datastructures

import (
	"fmt"
)

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
//
// Parâmetros:
//   - x: O valor inteiro a ser armazenado no novo nó.
//
// Exemplo de uso:
//
//	list := LinkedList{}
//	list.Add(10)
//	list.Add(20) // Lista: 10 -> 20 -> nil
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
//
// Exemplo de saída:
//
//	list.Print() // Saída: 10 -> 20 -> nil
func (T *LinkedList) Print() {
	current := T.Initial
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

// RemoveLast remove o último elemento da lista encadeada.
//
// Caso a lista esteja vazia, a função não faz nada.
//
// Exemplo de uso:
//
//	list := LinkedList{}
//	list.Add(10)
//	list.Add(20)
//	list.RemoveLast() // Remove 20, lista fica: 10 -> nil
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

// Get retorna o valor do nó na posição especificada da lista encadeada.
//
// Parâmetros:
//   - index: A posição do nó na lista encadeada. Deve estar dentro dos limites válidos.
//
// Retorna:
//   - O valor do nó na posição especificada.
//
// Pânico:
//   - Se a lista estiver vazia (T.Initial == nil).
//   - Se o índice for negativo ou maior ou igual ao tamanho da lista (fora dos limites).
//
// Exemplo de uso:
//
//	list := LinkedList{}
//	list.Add(10)
//	list.Add(20)
//	fmt.Println(list.Get(1)) // Saída: 20
func (T *LinkedList) Get(index int) int {
	if T.Initial == nil || index < 0 || index >= T.Size {
		panic("out of bounds")
	}

	return T.GetNode(index).Value
}

// GetNode retorna o nó no índice especificado.
//
// Parâmetros:
//   - index: A posição do nó na lista encadeada. Deve estar dentro dos limites válidos.
//
// Retorna:
//   - O ponteiro para o nó no índice especificado.
//   - Retorna nil se o índice for inválido.
//
// Exemplo de uso:
//
//	node := list.GetNode(1)
//	if node != nil {
//	    fmt.Println(node.Value) // Exibe o valor do nó
//	}
func (T *LinkedList) GetNode(index int) *Node {
	if T.Initial == nil || index < 0 || index >= T.Size {
		return nil
	}

	current := T.Initial
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current
}

// toArray converte a lista ligada em um slice de inteiros.
// Se a lista não estiver inicializada (T.Initial == nil), imprime um aviso e retorna nil.
//
// Retorno:
//   - []int: um slice contendo os valores da lista ligada na mesma ordem.
//
// Complexidade: O(n), onde n é o tamanho da lista.
func (T *LinkedList) toArray() []int {
	if T.Initial == nil {
		fmt.Print("linkedlist not initialized")
		return nil
	}
	arr := make([]int, T.Size)
	current := T.Initial
	for i := 0; current != nil; i++ {
		arr[i] = current.Value
		current = current.Next
	}
	return arr
}

