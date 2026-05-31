package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var l int

	// Lê a string e o valor limite L
	if _, err := fmt.Scan(&s, &l); err != nil {
		return
	}

	chars := []rune(s)
	atual := make([]int, len(chars))

	// Converte a string para o nosso vetor de inteiros (-1 para os pontos '.')
	for i := 0; i < len(chars); i++ {
		if chars[i] != '.' {
			atual[i] = int(chars[i] - '0')
		} else {
			atual[i] = -1
		}
	}

	// Dispara o Backtracking iniciando do índice 0
	if resolver(atual, l, 0) {
		// Se encontrou a solução, reconverte para string e imprime
		var resultado strings.Builder
		for _, num := range atual {
			resultado.WriteByte(byte(num) + '0')
		}
		fmt.Println(resultado.String())
	}
}

// Função de Backtracking
func resolver(atual []int, l, indice int) bool {
	// Caso Base: se chegamos ao fim do vetor, encontramos a única solução válida!
	if indice == len(atual) {
		return true
	}

	// Se a posição atual JÁ TIVER um número fixo de fábrica, apenas valida e avança
	if atual[indice] != -1 {
		if podeColocar(atual, l, indice, atual[indice]) {
			return resolver(atual, l, indice+1)
		}
		return false // Se o número fixo violar a regra de distância, o caminho é inválido
	}

	// Se for um espaço vazio (-1), tenta colocar cada dígito de 0 até L
	for num := 0; num <= l; num++ {
		if podeColocar(atual, l, indice, num) {
			// Decisão: coloca o número
			atual[indice] = num

			// Recursão: vai para o próximo índice. Se der certo até o final, retorna true
			if resolver(atual, l, indice+1) {
				return true
			}

			// Backtracking: Se deu errado mais para a frente, limpa a posição e tenta o próximo 'num'
			atual[indice] = -1
		}
	}

	// Se testou todos os números de 0 a L e nenhum serviu, volta na árvore de decisão
	return false
}

// Função que valida a distância mínima de L casas à esquerda
func podeColocar(atual []int, l, indice, num int) bool {
	// Verifica as posições anteriores. Não precisamos olhar além de 'l' casas para trás.
	for i := 1; i <= l; i++ {
		anterior := indice - i
		if anterior >= 0 {
			if atual[anterior] == num {
				return false // Encontrou o mesmo número perto demais!
			}
		}
	}
	return true
}
