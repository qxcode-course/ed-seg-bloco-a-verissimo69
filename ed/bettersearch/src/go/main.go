package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	baixo := 0
	alto := len(slice)

	for baixo < alto {

		meio := baixo + (alto-baixo)/2
		chute := slice[meio]
		if chute == value {
			for i := 0; i < len(slice); i++ {
				if value == slice[i] {
					return true, i
				}
			}
		}
		if chute < value {
			baixo = meio + 1
		} else {
			alto = meio
		}
	}
	return false, baixo
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
