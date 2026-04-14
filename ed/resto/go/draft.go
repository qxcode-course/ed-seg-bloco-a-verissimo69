package main

import "fmt"

func dividirResto(a int) {
	if a == 1 {
		fmt.Println(0, 1)
		return
	}
    dividirResto(a / 2)
	fmt.Println(a/2,a%2)
}

func main() {
	var a int
	fmt.Scan(&a)
	dividirResto(a)
}
