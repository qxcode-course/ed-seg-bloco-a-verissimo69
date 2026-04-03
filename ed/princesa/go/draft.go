package main

import (
	"fmt"
)

func print_jog(jogadores []int, espada int) {
	fmt.Print("[ ")
	for i, elem := range jogadores {
		if elem == 0 {
			continue
		}
		fmt.Print(elem)
		if i == espada {
			fmt.Print(">")
		}
		fmt.Printf(" ")
	}
	fmt.Print("]\n")
}

func procurarVivo(jogadores []int, espada int) int {
	for {
		espada = (espada + 1) % len(jogadores)

		if jogadores[espada] != 0 {
			return espada
		}
	}
}

// imprimirFila formata a saída exatamente como os casos de teste exigem.
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	jogadores := make([]int, 0, a)
	for i := 1; i <= a; i++ {
		jogadores = append(jogadores, i)
	}
	b -= 1
	for range a - 1 {
		print_jog(jogadores, b)
		vaiMorrer := procurarVivo(jogadores, b)
		jogadores[vaiMorrer] = 0
		b = procurarVivo(jogadores, b)
	}
	print_jog(jogadores, b)

}
