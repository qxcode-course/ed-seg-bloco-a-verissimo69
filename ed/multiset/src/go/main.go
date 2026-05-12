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

func procuararIndece(vector []int, value int) int {
	low, high := 0, len(vector)
	for low < high {
		medio := (low + high) / 2
		if vector[medio] < value {
			low = medio + 1
		} else {
			high = medio
		}
	}
	return low
}

func NewMultiSet(value int) *Vector {
	return &Vector{
		data:     make([]int, 0, value),
		size:     0,
		capacity: value,
	}
}

func (v *Vector) Insert(value int) {
	count := procuararIndece(v.data, value)
	if v.capacity == v.size {
		v.capacity *= 2
	}

	v.data = append(v.data, 0)

	for i := v.size; i > count; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[count] = value
	v.size++
}

func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size-1; i++ {
		if value == v.data[i] {
			return true
		}
	}
	return false
}

func (v *Vector) Erase(value int) error {
	ind := -1
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			ind = i
			break
		}
	}
	if ind == -1 {
		return fmt.Errorf("value not found")
	}
	for i := ind; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	v.data = v.data[:v.size]
	return nil
}

func Count(vetor []int, value int) (int, error) {
	ind := 0
	for i := 0; i < len(vetor); i++ {
		if vetor[i] == value {
			ind++
		}
	}
	return ind, nil
}
func Unique(vetor []int) (int, error) {
	ind := 1
	for i := 0; i < len(vetor)-1; i++ {
		if vetor[i] != vetor[i+1] {
			ind++
		}
	}
	if len(vetor) == 0 {
		ind = 0
	}
	return ind, nil
}

func (v *Vector) Clear() {
	v.data = v.data[:0]
	v.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	v := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			v = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			txt := "["
			txt += Join(v.data, ", ")
			txt += "]"
			fmt.Println(txt)

		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(v.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			ind, err := Count(v.data, value)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(ind)
		case "unique":
			ind, err := Unique(v.data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(ind)
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
