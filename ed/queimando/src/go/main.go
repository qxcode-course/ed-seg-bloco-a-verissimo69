package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, startL, startc int) {
	stack := []Pos{}
	stack = append(stack, Pos{l: startL, c: startc})

	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	nl := len(grid)
	if nl == 0 {
		return
	}

	nc := len(grid[0])
	for len(stack) > 0 {

		topIndex := len(stack) - 1
		curr := stack[topIndex]
		stack = stack[:topIndex]

		if curr.l < 0 || curr.l >= nl || curr.c < 0 || curr.c >= nc {
			continue
		}

		if grid[curr.l][curr.c] == '#' {
			grid[curr.l][curr.c] = 'o'
			for i := 0; i < 4; i++ {
				nextL := curr.l + dr[i]
				nextC := curr.c + dc[i]
				if nextL >= 0 && nextL < nl && nextC >= 0 && nextC < nc && grid[nextL][nextC] == '#' {
					stack = append(stack, Pos{l: nextL, c: nextC})
				}
			}
		}
	}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
