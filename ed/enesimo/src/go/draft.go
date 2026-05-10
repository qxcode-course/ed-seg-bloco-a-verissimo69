package main

import "fmt"

// 1. Sua função de apoio para verificar se é primo
func eh_primo(n int, divisor int) bool {
	if n < 2 {
		return false
	}
	if divisor == 1 {
		return true
	}
	if n%divisor == 0 {
		return false
	}
	return eh_primo(n, divisor-1)
}

// 2. Função auxiliar para chamar o eh_primo de forma fácil
func verificar(n int) bool {
	return eh_primo(n, n/2)
}

// 3. A função principal RECURSIVA para achar o n-ésimo primo
// n: quantos primos faltam encontrar
// atual: o número que estamos testando agora (começa em 2)
func enesimo_primo_rec(n int, atual int) int {
	if verificar(atual) {
		// Se achamos um primo e n era 1, encontramos o que queríamos!
		if n == 1 {
			return atual
		}
		// Se achamos um primo mas n > 1, continuamos procurando o próximo
		// diminuindo o n (pois falta 1 a menos)
		return enesimo_primo_rec(n-1, atual+1)
	}

	// Se o número atual NÃO é primo, apenas passamos para o próximo número
	// mantendo o mesmo n (pois ainda não encontramos um primo novo)
	return enesimo_primo_rec(n, atual+1)
}

// 4. Função que o usuário chama
func enesimo_primo(n int) int {
	return enesimo_primo_rec(n, 2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d\n", enesimo_primo(n))
}
