package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	sentinel *Node
}

func NewLList() *List {
	s := &Node{data: 0}
	s.next = s
	s.prev = s
	return &List{sentinel: s}
}

func (l *List) PushBack(num int) {
	n := &Node{
		data: num,
	}
	last := l.sentinel.prev

	l.sentinel.prev = n
	n.next = l.sentinel

	n.prev = last
	last.next = n
}

func (l *List) PushFront(num int) {
	n := &Node{
		data: num,
	}

	first := l.sentinel.next

	l.sentinel.next = n
	n.prev = l.sentinel

	n.next = first
	first.prev = n
}

func (l *List) Clear() {

	l.sentinel.next = l.sentinel
	l.sentinel.prev = l.sentinel

}

func (l *List) Size() int {

	count := 0
	for i := l.sentinel.next; i != l.sentinel; i = i.next {
		count++
	}
	return count
}

func (l *List) PopFront() {
	first := l.sentinel.next

	if first == l.sentinel { // lista vazia
		return
	}

	next := first.next

	// remover 'first'
	l.sentinel.next = next
	next.prev = l.sentinel

	// limpar ponteiros do nó removido (opcional, mas seguro)
	first.next = nil
	first.prev = nil
}

func (l *List) PopBack() {
	last := l.sentinel.prev

	if last == l.sentinel {
		return
	}
	prev := last.prev

	prev.next = l.sentinel
	l.sentinel.prev = prev

	last.next = nil
	last.prev = nil

}

func (l *List) String() string {
	txt := "["
	for i := l.sentinel.next; i != l.sentinel; i = i.next {
		txt += fmt.Sprintf("%d, ", i.data)
	}
	if len(txt) > 1 {
		txt = txt[:len(txt)-2]
	}
	txt += "]"
	return txt

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
