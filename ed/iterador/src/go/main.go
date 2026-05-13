package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data     []int
	index    int
	step     int
	isCyclic bool
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1, step: 1}
}

func (i *Iterator) HasNext() bool {
	if i.isCyclic {
		return len(i.data) > 0
	}
	if i.step == 1 {
		return i.index < len(i.data)-1
	}
	return i.index > 0
}

func (i *Iterator) Next() int {
	if i.isCyclic {
		i.index = (i.index + 1) % len(i.data)
		return i.data[i.index]
	}

	i.index += i.step
	return i.data[i.index]
}

func (i *MyList) ReverseIterator() *Iterator {
	return &Iterator{
		data:  i.data,
		index: len(i.data),
		step:  -1,
	}
}

func (i *MyList) CyclicIterator() *Iterator {
	return &Iterator{
		data:     i.data,
		index:    -1,
		isCyclic: true,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterator()
			for range qtd {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		}
	}

}
