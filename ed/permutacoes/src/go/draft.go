package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func dfsPermutar(chars []rune, usado []bool, atual []rune) {
	if len(atual) == len(chars) {
		fmt.Println(string(atual))
		return
	}

	for i := 0; i < len(chars); i++ {
		if usado[i] {
			continue
		}

		// Decisão
		usado[i] = true
		atual = append(atual, chars[i])

		// Recursão (vai mais fundo na árvore de decisões)
		dfsPermutar(chars, usado, atual)

		// Backtracking (desfaz a decisão para tentar o próximo caractere do loop)
		usado[i] = false
		atual = atual[:len(atual)-1]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	s := scanner.Text()

	// 1. Converter a string para um slice de caracteres (runes ou bytes)
	chars := []rune(s)

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	usado := make([]bool, len(chars))
	var atual []rune
	dfsPermutar(chars, usado, atual)

}
