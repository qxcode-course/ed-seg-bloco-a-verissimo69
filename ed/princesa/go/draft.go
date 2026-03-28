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

func main() {
	var n, e int

	// Lê a entrada padão
	fmt.Scan(&n, &e)
	

	// Cria o slice com as pessoas de 1 até N
	vivos := make([]int, n)
	for i := 0; i < n; i++ {
		vivos[i] = i + 1
	}

	// Encontra o índice inicial da pessoa com a espada (E)
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

		// A espada passa para a próxima pessoa.
		// Como o elemento foi removido, o próximo assume exatamente o índice 'posMorte'
		posEspada = posMorte % len(vivos)
	}
}