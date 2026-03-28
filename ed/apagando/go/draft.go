package main

import (
	"fmt"
)

func main() {
	var n, m, id int

	// 1. Lê a quantidade de pessoas na fila inicial
	fmt.Scan(&n)

	// Cria um slice para guardar a fila mantendo a ordem original
	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	// 2. Lê a quantidade de pessoas que saíram
	fmt.Scan(&m)

	// Cria um map de booleanos para registrar quem saiu.
	// Usamos map[int]bool porque a busca nele é instantânea O(1).
	sairam := make(map[int]bool)
	for i := 0; i < m; i++ {
		fmt.Scan(&id)
		sairam[id] = true // Marca que este ID saiu da fila
	}

	// 3. Percorre a fila original na ordem correta
	for _, pessoa := range fila {
		// Se a pessoa NÃO está no map de quem saiu, nós a imprimimos
		if !sairam[pessoa] {
			fmt.Printf("%v ", pessoa)
		}
	}
	fmt.Println() // Quebra de linha no final para manter a saída limpa
}