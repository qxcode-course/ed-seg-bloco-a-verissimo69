package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(board [][]byte, visited [][]bool, r, c int) {
	rows := len(board)
	cols := len(board[0])

	if r < 0 || r >= rows || c < 0 || c >= cols {
		return
	}
	if board[r][c] != 'O' || visited[r][c] {
		return
	}

	visited[r][c] = true

	dfs(board, visited, r+1, c)
	dfs(board, visited, r-1, c)
	dfs(board, visited, r, c+1)
	dfs(board, visited, r, c-1)

}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {

	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for c := 0; c < cols; c++ {
		if board[0][c] == 'O' {
			dfs(board, visited, 0, c)
		}
		if board[rows-1][c] == 'O' {
			dfs(board, visited, rows-1, c)
		}
	}

	for r := 0; r < rows; r++ {
		if board[r][0] == 'O' {
			dfs(board, visited, r, 0)
		}
		if board[r][cols-1] == 'O' {
			dfs(board, visited, r, cols-1)
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' && !visited[r][c] {
				board[r][c] = 'X'
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
