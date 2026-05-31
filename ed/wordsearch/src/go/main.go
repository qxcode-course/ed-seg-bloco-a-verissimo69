package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	rows := len(grid)
	if rows == 0 || len(word) == 0 {
		return false
	}

	cols := len(grid[0])
	gridCounts := make(map[byte]int)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			gridCounts[grid[r][c]]++
		}
	}

	wordCounts := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		wordCounts[word[i]]++
	}

	for char, count := range wordCounts {
		if gridCounts[char] < count {
			return false
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == word[0] {
				if buscarPalavra(grid, word, r, c, 0) {
					return true
				}
			}
		}
	}

	return false
}

func buscarPalavra(grid [][]byte, word string, r, c, index int) bool {
	if index == len(word) {
		return true
	}
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != word[index] {
		return false
	}

	temp := grid[r][c]
	grid[r][c] = '#'
	encontrou := buscarPalavra(grid, word, r+1, c, index+1) ||
		buscarPalavra(grid, word, r-1, c, index+1) ||
		buscarPalavra(grid, word, r, c+1, index+1) ||
		buscarPalavra(grid, word, r, c-1, index+1)

	grid[r][c] = temp
	return encontrou

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 1. Lê a palavra alvo
	//fmt.Println("Digite a palavra secreta:")
	var word string
	if scanner.Scan() {
		word = scanner.Text()
	}

	// 2. Lê as linhas do tabuleiro
	//fmt.Println("Digite as linhas do tabuleiro (pressione ENTER em uma linha vazia para finalizar):")
	grid := make([][]byte, 0)
	for scanner.Scan() {
		linha := scanner.Text()
		// Se o usuário apenas apertar ENTER (linha vazia), encerra a leitura
		if len(linha) == 0 {
			break
		}
		grid = append(grid, []byte(linha))
	}

	// 3. Executa o algoritmo e exibe o resultado
	//fmt.Print("Resultado: ")
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
