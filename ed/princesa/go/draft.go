package main

import (
	"fmt"
)

// imprimirFila formata a saída exatamente como os casos de teste exigem.
func imprimirFila(vivos []int, posEspada int) {
	fmt.Print("[ ")
	for i, val := range vivos {
		if i == posEspada {
			fmt.Printf("%d> ", val)
		} else {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println("]")
}
func mostrarJogadore(jagadores []bool, espada int){
	fmt.Printf("[")
	for i, valor := true
}

func main() {
	var n, e int
	
	fmt.Scan(&n, &e)
	
	vivos := make([]bool, n)
	for i := 0; i < n; i++ {
		vivos[i] = true
	}
	espada := e
	mostrarJogadore(vivos, espada)


	posEspada := -1
	for i, val := range vivos {
		if val == e {
			posEspada = i
			break
		}
	}

	for len(vivos) > 0 {
		imprimirFila(vivos, posEspada)

		if len(vivos) == 1 {
			break
		}

		// Calcula a posição de quem vai morrer (o próximo da fila circular)
		posMorte := (posEspada + 1) % len(vivos)

		// Remove a pessoa do slice
		// Pega tudo do início até a posMorte e junta com tudo depois da posMorte
		vivos = append(vivos[:posMorte], vivos[posMorte+1:]...)
		posEspada = posMorte % len(vivos)
	}
}
