package main

import "fmt"

func padrao(n int) int {
	if n == 1 {
		return 20
	}
	return padrao(n-1) + 8
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(padrao(n))
}
