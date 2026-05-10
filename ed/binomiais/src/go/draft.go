package main

import "fmt"

func binomiais(a, b int) int {
	if b == 0 || b == a {
		return 1
	}
	c := binomiais(a-1, b-1) + binomiais(a-1, b)
	return c

}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(binomiais(a, b))
}
