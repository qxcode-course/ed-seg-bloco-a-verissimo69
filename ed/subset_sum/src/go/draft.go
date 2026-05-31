package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func encontrarSoma(vetor []int, alvo, indice, somaAtual int) bool{
    if somaAtual == alvo {
        return true
    }

    if somaAtual> alvo || indice >= len(vetor){
        return false
    }
    if encontrarSoma(vetor,alvo, indice+1, somaAtual+vetor[indice]){
        return true
    }
    if encontrarSoma(vetor, alvo, indice +1, somaAtual) {
        return true
    }
    return false

}


func main() {
   reader := bufio.NewReader(os.Stdin)

	var n, k int
	// Lê n (quantidade de elementos) e k (soma alvo)
	fmt.Fscan(reader, &n, &k)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	// Otimização: Ordenar os números permite fazer podas eficientes
	sort.Ints(nums)

	if encontrarSoma(nums, k, 0, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}