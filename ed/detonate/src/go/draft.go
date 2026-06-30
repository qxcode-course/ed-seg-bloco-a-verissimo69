package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bomb struct {
	x, y, r int
}

func canDetonate(b1, b2 Bomb) bool {
	dx := b1.x - b2.x
	dy := b1.y - b2.y
	return (dx*dx)+(dy*dy) <= b1.r*b1.r
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)

	// Função inline para calcular a distância e verificar o alcance
	canDetonate := func(i, j int) bool {
		x1, y1, r1 := int64(bombs[i][0]), int64(bombs[i][1]), int64(bombs[i][2])
		x2, y2 := int64(bombs[j][0]), int64(bombs[j][1])

		dx := x1 - x2
		dy := y1 - y2
		return (dx*dx)+(dy*dy) <= r1*r1
	}

	// Construção do Grafo Direcionado
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && canDetonate(i, j) {
				adj[i] = append(adj[i], j)
			}
		}
	}

	maxBombs := 0

	// BFS para cada bomba
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		queue := []int{i}
		visited[i] = true
		count := 0

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			count++

			for _, next := range adj[curr] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		if count > maxBombs {
			maxBombs = count
		}
	}
	return maxBombs
}
func main() {
	// Scanner configurado para ler palavra por palavra (número por número)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for {
		// 1. Lê a quantidade de bombas (N)
		if !scanner.Scan() {
			break // Fim da entrada (EOF)
		}
		var n int
		fmt.Sscanf(scanner.Text(), "%d", &n)

		// 2. Lê a quantidade de colunas (M) -> Sempre será 3 neste problema
		if !scanner.Scan() {
			break
		}
		var m int
		fmt.Sscanf(scanner.Text(), "%d", &m)

		// 3. Aloca e lê a matriz de bombas baseado em N e M
		bombs := make([][]int, n)
		for i := 0; i < n; i++ {
			bombs[i] = make([]int, m)
			for j := 0; j < m; j++ {
				if scanner.Scan() {
					fmt.Sscanf(scanner.Text(), "%d", &bombs[i][j])
				}
			}
		}

		// 4. Executa a lógica e imprime o resultado do caso atual
		resultado := maximumDetonation(bombs)
		fmt.Println(resultado)
	}
}
