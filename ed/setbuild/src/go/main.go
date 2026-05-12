package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Vector {
	return &Vector{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Insert(value int) error {
	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
			return nil
		}
	}

	if v.capacity <= v.size {
		v.capacity *= 2
	}
	pos := 0
	low, high := 0, v.size
	for low < high {
		mid := (low + high) / 2
		if v.data[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	pos = low

	v.data = append(v.data, 0)

	for i := v.size; i > pos; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[pos] = value
	v.size++

	return nil

}

func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
			return true
		}
	}
	return false
}

func (v *Vector) Erase(value int) error {
	stop := true
	for i := 0; i < v.size-1; i++ {
		if value == v.data[i] {
			for j := i; j < v.size-1; j++ {
				v.data[j] = v.data[j+1]
			}
			stop = false
		}
	}
	if stop {
		return fmt.Errorf("value not found")
	}
	v.size--
	return nil
}
func (v *Vector) Show() string {

	txt := "["
	for i := 0; i < v.size; i++ {
		txt += strconv.Itoa(v.data[i])
		if i < v.size-1 {
			txt += ", "
		}
	}
	txt += "]"
	return txt

}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":

			fmt.Println(v.Show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}

		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
