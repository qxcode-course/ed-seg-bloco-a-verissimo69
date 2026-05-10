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

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, 0, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}
func (v *Vector) PushBack(value int) {
	v.data = append(v.data, value)
	v.size = len(v.data)
	v.capacity = cap(v.data)
}
func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return -1, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}
func (v *Vector) Capacity() int {
	return v.capacity
}
func (v *Vector) Set(index int, value int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}
func (v *Vector) Clear() {
	v.data = v.data[:0]
	v.size = 0
}
func (v *Vector) Reserve(newCapacity int) {
	v.capacity = newCapacity
}

func (v *Vector) PopBack() error {
	if v.size == 0 {
		return fmt.Errorf("vector is empty")
	}
	v.data = v.data[:v.size-1]
	v.size--
	return nil
}

func (v *Vector) Insert(index, value int) error {
	if index > v.capacity || index < 0 {
		return fmt.Errorf("index out of range")

	}
	v.data = append(v.data, 0)
	v.size++
	for i := v.size - 1; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.capacity = cap(v.data)
	return nil
}

func (v *Vector) Erase(index int) error {
	if index > v.capacity || index < 0 {
		return fmt.Errorf("index out of range")

	}
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	v.capacity = cap(v.data)
	return nil
}
func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return i
		}
	}
	return -1
}
func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return true
		}
	}
	return false
}

func (v *Vector) Slice(a, b int) string {
	var newVetor []int
	if b < 0 {
		b = v.size
	}
	for i := b; i < v.size - 1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size --
	for i:= a ; i<b; i++{
		newVetor = append(newVetor, v.data[i])
	}
	resul := "["
	resul += Join(newVetor, ", ") + "]"
	return resul
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}
func (v *Vector) Status() string {
	var txt string
	txt = fmt.Sprintf("size:%v capacity:%v", v.size, v.capacity)
	return txt

}
func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"

}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.String())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
