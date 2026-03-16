package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	vencedorIndice := -1
	melhorPontuacao := math.MaxInt32

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		if a >= 10 && b >= 10 {
			pontuacao := int(math.Abs(float64(a - b)))
			if pontuacao < melhorPontuacao {
				melhorPontuacao = pontuacao
				vencedorIndice = i
			}
		}
	}

	if vencedorIndice == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(vencedorIndice)
	}
}