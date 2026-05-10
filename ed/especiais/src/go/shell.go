package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func verifcar(num int, vet []int) bool {
	for i := 0; i < len(vet); i++ {
		if num == vet[i] {
			return true
		}
	}
	return false
}

func occurr(vet []int) []Pair {
	var newPair Pair
	var vetorUax, vetorNew []int
	var count int
	var resul []Pair
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 {
			vetorNew = append(vetorNew, vet[i]*-1)
		} else if vet[i] > 0 {
			vetorNew = append(vetorNew, vet[i])
		}
	}
	for i := 0; i < len(vet); i++ {

		if verifcar(vetorNew[i], vetorUax) {
			continue
		}
		vetorUax = append(vetorUax, vetorNew[i])
		count = 0
		for j := 0; j < len(vetorNew); j++ {
			if vetorNew[i] == vetorNew[j] {
				count++
			}
		}
		newPair.One = vetorNew[i]
		newPair.Two = count
		resul = append(resul, newPair)

	}
	sort.Slice(resul, func(i, j int) bool {
		return resul[i].One < resul[j].One
	})
	return resul
}

func teams(vet []int) []Pair {

	var newPair Pair
	var vetorNew []int
	var resul []Pair
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 {
			vetorNew = append(vetorNew, vet[i]*-1)
		} else if vet[i] > 0 {
			vetorNew = append(vetorNew, vet[i])
		}
	}
	for i := 0; i < len(vet); {
		estresse := vetorNew[i]
		count := 0
		j := i

		for j < len(vetorNew) && vetorNew[j] == estresse {
			count++
			j++
		}
		newPair.One = estresse
		newPair.Two = count
		resul = append(resul, newPair)
		i = j

	}

	return resul
}

func mnext(vet []int) []int {
	newVetor := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {

			temMulherDoLado := false

			if i > 0 {
				if vet[i-1] < 0 {
					temMulherDoLado = true
				}
			}
			if i < len(vet)-1 {
				if vet[i+1] < 0 {
					temMulherDoLado = true
				}
			}
			if temMulherDoLado == true {
				newVetor[i] = 1
			}
		}
		// Se for mulher ou homem sozinho, o valor continua 0 (já padrão do make)
	}

	return newVetor

}

func alone(vet []int) []int {
	newVetor := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {

		if vet[i] > 0 {
			sohomem := true

			if i > 0 {
				if vet[i-1] < 0 {
					sohomem = false
				}
			}
			if i < len(vet)-1 {
				if vet[i+1] < 0 {
					sohomem = false
				}
			}

			if sohomem {
				newVetor[i] = 1
			}
		}
	}
	return newVetor
}

func couple(vet []int) int {
	copia := make([]int, len(vet))
	copy(copia, vet)
	cont := 0
	for i := 0; i < len(vet); i++ {
		if copia[i] == 0 {
			continue
		}
		aux := vet[i]
		for j := i; j < len(vet); j++ {
			if (aux*-1) == vet[j] && i != j {
				cont++
				vet[i] = 0
				vet[j] = 0
				break
			}
		}
	}
	return cont
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	// 1. Verificação de segurança: se a sequência for maior, não cabe na fila
	if len(seq) > len(vet) {
		return -1
	}

	// 2. Percorrer a fila original
	// Usamos o limite (tamanho da fila - tamanho da sequência) para não "cair" do vetor
	for i := 0; i <= len(vet)-len(seq); i++ {

		// Esta variável vai nos dizer se a tentativa atual deu certo
		bateuTudo := true

		// 3. Comparar cada elemento da sequência a partir da posição 'i'
		for j := 0; j < len(seq); j++ {
			// Comparamos o elemento da fila (i + deslocamento j) com a sequência
			if vet[i+j] != seq[j] {
				bateuTudo = false
				break // Se um elemento for diferente, essa posição 'i' não serve
			}
		}

		// 4. Se chegamos aqui e 'bateuTudo' ainda é true, encontramos a primeira ocorrência!
		if bateuTudo {
			return i
		}
	}

	// 5. Se os laços acabarem e nada for retornado, a sequência não existe na fila
	return -1
}

func erase(vet []int, posicoes []int) []int {
	// 1. Criamos um mapa de quem deve ser removido
	// Usamos um slice de booleanos do mesmo tamanho do vetor original
	deveRemover := make([]bool, len(vet))

	// 2. Marcamos as posições que recebemos como 'true'
	for i := 0; i < len(posicoes); i++ {
		indiceParaApagar := posicoes[i]

		// Verificação de segurança: só marca se o índice existir no vetor
		if indiceParaApagar >= 0 && indiceParaApagar < len(vet) {
			deveRemover[indiceParaApagar] = true
		}
	}

	// 3. Criamos o vetor de resultado
	resul := make([]int, 0)

	// 4. Percorremos o vetor original e só guardamos quem NÃO foi marcado
	for i := 0; i < len(vet); i++ {
		if deveRemover[i] == false {
			resul = append(resul, vet[i])
		}
	}

	return resul
}

func clear(vet []int, valor int) []int {
	// Criamos o vetor de resultado vazio
	// Usamos o append, então começamos com tamanho 0
	resul := make([]int, 0)

	for i := 0; i < len(vet); i++ {
		// Pegamos o nível de stress atual (valor absoluto)
		stressAtual := vet[i]
		if stressAtual < 0 {
			stressAtual = stressAtual * -1
		}

		// Se o stress da pessoa for diferente do valor que queremos apagar,
		// ela pode ficar na fila.
		if stressAtual != valor {
			resul = append(resul, vet[i]) // Guardamos o valor original (com sinal)
		}
	}

	return resul
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
