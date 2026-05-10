package main

import "fmt"

func escada(a int) int {
	if a == 1 || a == 2 {
		return 1
	}
	if a == 3 {
		return 2
	}
	c := escada(a-1) + escada(a-3)
	return c
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(escada(a))
}
