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
	root *Node
	next *Node
	prev *Node
}

type LList struct {
	root *Node
	size int
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}
func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}
func (l *LList) Size() int {
	return l.size
}

func NewLList() *LList {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	sentinel.root = sentinel

	return &LList{root: sentinel, size: 0}
}
func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (l *LList) Insert(node *Node, value int) {
	if node == nil {
		return
	}
	NewNode := &Node{data: value, root: l.root}

	NewNode.next = node
	NewNode.prev = node.prev
	node.prev.next = NewNode
	node.prev = NewNode
	l.size++
}

func (l *LList) Remove(node *Node) *Node {

	if node == nil || node == l.root {
		return nil
	}
	nextNode := node.next
	node.prev.next = node.next
	node.next.prev = node.prev

	l.size--
	if nextNode == l.root {
		return nil
	}
	return nextNode
}
func (l *LList) PushBack(num int) {
	l.Insert(l.root, num)
}
func (l *LList) PushFront(num int) {
	l.Insert(l.root.next, num)
}
func (ll *LList) Front() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.prev
}

func (ll *LList) PopFront() {
	if ll.size > 0 {
		ll.Remove(ll.root.next)
	}
}

// PopBack remove o último elemento real da lista
func (ll *LList) PopBack() {
	if ll.size > 0 {
		ll.Remove(ll.root.prev)
	}
}

func (l *LList) Search(value int) *Node {
	for ele := l.root.next; ele != l.root; ele = ele.next {
		if ele.data == value {
			return ele
		}
	}
	return nil
}

func (l *LList) String() string {
	if l.size == 0 {
		return "[]"
	}
	var elementos []string

	for ele := l.root.next; ele != l.root; ele = ele.next {
		elementos = append(elementos, strconv.Itoa(ele.data))
	}
	return "[" + strings.Join(elementos, ", ") + "]"
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.data)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.data)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.data = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
