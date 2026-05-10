package main

import "fmt"

func pedaco(a string) {
	if a == "" {
		return
	}
	pedaco(a[1:])
	fmt.Println(a)
}

func main() {
	var nome string
	fmt.Scan(&nome)
	pedaco(nome)
}
