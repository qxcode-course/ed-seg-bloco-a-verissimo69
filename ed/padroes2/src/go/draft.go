package main

import "fmt"

func padroes(n int) int {
	if n == 1 {
		return 3
	}
	a := n*2 + 1
	return padroes(n-1) + a
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(padroes(n))
}
