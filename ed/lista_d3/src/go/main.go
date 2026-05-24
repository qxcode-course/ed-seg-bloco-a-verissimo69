package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}
func equals(list1, list2 *LList) bool {
	if list1.size != list1.size {
		return false
	}
	for i1, i2 := list1.root.next, list2.root.next; i1 != list1.root || i2 != list2.root; i1, i2 = i1.next, i2.next {
		if i1.Value != i2.Value {
			return false
		}
	}
	return true
}

func addsorted(list *LList, value int) {
	curr := list.root.next
	for curr != list.root && curr.Value < value {
		curr = curr.next
	}

	newNode := &Node{
		Value: value,
		next:  curr,
		prev:  curr.prev,
		root:  curr.root,
	}

	curr.prev.next = newNode
	curr.prev = newNode
}
func merge(list1, list2 *LList) *LList {
	result := NewLList()
	// Iniciamos a recursão apenas com o primeiro nó real de cada lista
	mergeRec(list1.root.next, list2.root.next, result)
	return result
}

func mergeRec(n1, n2 *Node, result *LList) {
	// 1. Condição de parada: ambas as listas chegaram no sentinela
	if n1 == n1.root && n2 == n2.root {
		return
	}

	// 2. Escolhemos pegar o valor do n2 se:
	// - A lista 1 já acabou (n1 == n1.root) OU
	// - A lista 2 não acabou E o valor de n2 é menor que n1
	if n1 == n1.root || (n2 != n2.root && n2.Value < n1.Value) {
		result.PushBack(n2.Value)
		mergeRec(n1, n2.next, result) // Avança apenas n2
	} else {
		// 3. Caso contrário, com certeza devemos pegar o n1
		result.PushBack(n1.Value)
		mergeRec(n1.next, n2, result) // Avança apenas n1
	}
}

func (l *LList) reverse() []int {
	var ele []int
	for i := l.root.prev; i != l.root; i = i.prev {
		ele = append(ele, i.Value)
	}
	return ele

}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func Stringg(list []int) string {
	var ele string
	ele += "["
	for i := 0; i < len(list); i++ {
		ele += strconv.Itoa(list[i])
		if i < len(list)-1 {
			ele += ", "
		}
	}

	ele += "]"
	return ele
}

func String(l *LList) string {
	var elements []string

	// Percorre a lista ignorando o nó sentinela (root)
	for curr := l.root.next; curr != l.root; curr = curr.next {
		elements = append(elements, strconv.Itoa(curr.Value))
	}

	// Formata a saída unindo os elementos com vírgula e espaço, cercados por colchetes
	return "[" + strings.Join(elements, ", ") + "]"
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(String(lla))
		case "reverse":
			lla := str2list(args[1])
			list := lla.reverse()
			fmt.Println(Stringg(list))
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(String(merged))
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
