package main

import (
	"fmt"
	"strings"
)

type DNode[T comparable] struct {
	Value            T
	next, prev, root *DNode[T]
}

func (n *DNode[T]) Next() *DNode[T] {
	if n == n.root {
		return n
	}
	return n.next
}

func (n *DNode[T]) Prev() *DNode[T] {
	if n == n.root {
		return n
	}
	return n.prev
}

type DList[T comparable] struct {
	root *DNode[T]
	size int
}

func NewDList[T comparable]() *DList[T] {
	root := &DNode[T]{}
	root.next = root
	root.prev = root
	root.root = root
	return &DList[T]{root: root, size: 0}
}

func (l *DList[T]) PushBack(value T) {
	l.Insert(l.root, value)
}

func (l *DList[T]) PopBack() {
	if l.size == 0 {
		return
	}
	l.Erase(l.root.prev)
}

func (l *DList[T]) PopFront() {
	if l.size == 0 {
		return
	}
	l.Erase(l.root.next)
}

func (l *DList[T]) PushFront(value T) {
	l.Insert(l.root.next, value)
}

func (l *DList[T]) Insert(it *DNode[T], value T) *DNode[T] {
	n := &DNode[T]{Value: value, root: l.root}
	n.prev = it.prev
	n.next = it
	it.prev.next = n
	it.prev = n
	l.size++
	return n
}

// Erase agora limpa as referências do nó removido
func (l *DList[T]) Erase(it *DNode[T]) {
	if it == l.root || it == nil {
		return
	}
	// Religa os vizinhos
	it.prev.next = it.next
	it.next.prev = it.prev

	// Limpa os ponteiros do nó isolado
	it.next = nil
	it.prev = nil
	it.root = nil

	l.size--
}

func (l *DList[T]) String() string {
	values := []string{}
	for n := l.root.next; n != l.root; n = n.next {
		values = append(values, fmt.Sprint(n.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func (l *DList[T]) Size() int {
	return l.size
}

func (l *DList[T]) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *DList[T]) Front() *DNode[T] {
	return l.root.next
}

func (l *DList[T]) Back() *DNode[T] {
	return l.root.prev
}

func (l *DList[T]) End() *DNode[T] {
	return l.root
}

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var txt []string
	for i := l.Front(); i != l.End(); i = i.next {
		if i == sword {
			txt = append(txt, fmt.Sprintf("%v>", i.Value))
		} else {
			txt = append(txt, fmt.Sprintf("%v", i.Value))
		}
	}
	return "[ " + strings.Join(txt, " ") + " ]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil || l.size == 0 {
		return nil
	}

	proximo := it.next
	if proximo == l.End() {
		proximo = proximo.next
	}

	return proximo
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)

	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}

	sword := l.Front()

	// CORREÇÃO AQUI: Trocado for range por laço clássico para compatibilidade
	for i := 0; i < chosen-1; i++ {
		sword = Next(l, sword)
	}

	// CORREÇÃO AQUI: Trocado for range por laço clássico para compatibilidade
	for i := 0; i < qtd-1; i++ {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}

	// Imprime o último sobrevivente
	fmt.Println(ToStr(l, sword))
}
