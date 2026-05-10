package main

import "fmt"

func verificar(a, b int) bool {
	if a < 2 {
		return false
	}
	if b == 1 {
		return true
	}
	if a%b == 0 {
		return false
	}
	return verificar(a, b-1)
}

func vetor(n, atual int, lista []int) []int {
	if len(lista) == n {
		return lista
	}
	if verificar(atual, atual/2) {
		lista = append(lista, atual)
	}
	return vetor(n, atual+1, lista)
}

func main() {
	var a int
	fmt.Scan(&a)
	vetorI := vetor(a, 2, []int{})
	aux := "["
	for i := 0; i < len(vetorI); i++ {
		aux += fmt.Sprint(vetorI[i])
		if i+1 < len(vetorI) {
			aux += ", "
		}
	}
	aux += "]"
	fmt.Println(aux)
}
