package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func inside(grid [][]int, p Pos) bool {
	return p.l >= 0 && p.l < len(grid) && p.c >= 0 && p.c < len(grid[0])
}

func lerEntrada(scanner *bufio.Scanner) ([][]int, []Pos, int, bool) {
	if !scanner.Scan() {
		return nil, nil, 0, false
	}

	var nl, nc int
	fmt.Sscanf(scanner.Text(), "%d", &nl)

	if !scanner.Scan() {
		return nil, nil, 0, false
	}
	fmt.Sscanf(scanner.Text(), "%d", &nc)

	grid := make([][]int, nl)
	var queue []Pos
	laranjasFrescas := 0

	for l := 0; l < nl; l++ {
		grid[l] = make([]int, nc)
		for c := 0; c < nc; c++ {
			scanner.Scan()
			var val int
			fmt.Sscanf(scanner.Text(), "%d", &val)
			grid[l][c] = val

			if val == 2 {
				queue = append(queue, Pos{l, c})
			} else if val == 1 {
				laranjasFrescas++
			}
		}
	}
	return grid, queue, laranjasFrescas, true

}

func simularAprodrecimento(grid [][]int, queue []Pos, frecas int) int {
	if frecas == 0 {
		return 0
	}

	minutos := 0
	direcoes := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 && frecas > 0 {
		minutos++
		tamanhoNivel := len(queue)
		for i := 0; i < tamanhoNivel; i++ {
			atual := queue[0]
			queue = queue[1:]

			for _, d := range direcoes {
				vizinho := Pos{l: atual.l + d.l, c: atual.c + d.c}
				if inside(grid, vizinho) && grid[vizinho.l][vizinho.c] == 1 {
					grid[vizinho.l][vizinho.c] = 2
					frecas--
					queue = append(queue, vizinho)

				}

			}
		}
	}
	if frecas > 0 {
		return -1
	}
	return minutos

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for {
		grid, queue, laranFresca, ok := lerEntrada(scanner)
		if !ok {
			break
		}
		resultado := simularAprodrecimento(grid, queue, laranFresca)
		fmt.Println(resultado)
	}
}
