package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var vetor []int
	for _, alem := range vet {
		if alem > 0 {
			vetor = append(vetor, alem)
		}
	}
	return vetor
}

func getCalmWomen(vet []int) []int {
	var vetor []int
	for _, elem := range vet {
		if elem > -10 && elem < 0 {
			vetor = append(vetor, elem)
		}
	}
	return vetor
}

func sortVet(vet []int) []int {
	slices.Sort(vet)
	return vet
}

func sortStress(vet []int) []int {
	var apenasPositivos []int
	var apenasNegativos []int
	for _, elem := range vet {
		if elem >= 0 {
			apenasPositivos = append(apenasPositivos, elem)
		} else {
			apenasNegativos = append(apenasNegativos, elem)
		}
	}
	slices.Sort(apenasPositivos)
	slices.SortFunc(apenasNegativos, func(a, b int) int {
		return b - a
	})
	aux := apenasNegativos[0]
	cont := 0
	for _, elem := range apenasNegativos {
		if elem == aux {
			cont += 1
		}
	} // --- A CORREÇÃO ESTÁ AQUI ---
	// Inicializa o slice com o tamanho total necessário
	total := len(apenasNegativos) + len(apenasPositivos)
	resultado := make([]int, total)

	// Agora o copy vai funcionar porque o 'resultado' tem espaço
	copy(resultado[:cont], apenasNegativos[:cont])

	copy(resultado[cont:], apenasPositivos)

	posicaoFinal := cont + len(apenasPositivos)
	copy(resultado[posicaoFinal:], apenasNegativos[cont:])

	return resultado
}

func reverse(vet []int) []int {

	clone := slices.Clone(vet)

	j := len(clone) - 1

	for i := 0; i < len(clone)/2; i++ {

		clone[i], clone[j] = clone[j], clone[i]

		j--
	}
	return clone

}

func unique(vet []int) []int {
	// Mapa para registrar o que já vimos
	visualizados := make(map[int]bool)
	var resultado []int

	for _, valor := range vet {
		// Se o valor NÃO está no mapa, adicionamos ao resultado
		if !visualizados[valor] {
			visualizados[valor] = true
			resultado = append(resultado, valor)
		}
	}
	return resultado
}
func repeated(vet []int) []int {
    // Mapa para saber se já vimos o número antes
    vistos := make(map[int]bool)
    // Mapa para garantir que não adicionamos o mesmo número duas vezes no resultado
    jaNoResultado := make(map[int]bool)
    var resultado []int

    for _, valor := range vet {
        if vistos[valor] {
            // Se já vimos o número, ele é um repetido.
            // Mas só adicionamos se ele ainda não estiver na nossa lista de repetidos.
            if !jaNoResultado[valor] {
                resultado = append(resultado, valor)
                jaNoResultado[valor] = true
            }
        } else {
            // Primeira vez que o número aparece, apenas marcamos como visto.
            vistos[valor] = true
        }
    }
    
    // Se o vetor for [5,4,3,2,1], o 'vistos' será preenchido, 
    // mas o 'if vistos[valor]' nunca será verdadeiro, retornando [] corretamente.
    return resultado
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
