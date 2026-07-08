package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dentroImagem(image [][]int, r, c int) bool {
	return r >= 0 && r < len(image) && c >= 0 && c < len(image[0])
}
func podePintar(image [][]int, r, c, corOriginal int) bool {
	return dentroImagem(image, r, c) && image[r][c] == corOriginal
}
func pintar(image [][]int, r, c, color int) {
	image[r][c] = color
}

func flood(image [][]int, r, c, corOriginal, color int) {
	if !podePintar(image, r, c, corOriginal) {
		return
	}

	pintar(image, r, c, color)

	flood(image, r+1, c, corOriginal, color)
	flood(image, r-1, c, corOriginal, color)
	flood(image, r, c+1, corOriginal, color)
	flood(image, r, c-1, corOriginal, color)
}

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	corOriginal := image[sr][sc]
	if corOriginal == color {
		return image
	}
	flood(image, sr, sc, corOriginal, color)
	return image
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
