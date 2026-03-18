package main

import "fmt"

func main() {
	var qtd0, qtd1 int
	fmt.Scan(&qtd0, &qtd1)
	vetor0 := make([]int, qtd1)
	unicos := make(map[int]bool)
	repetidos := make([]int, 0, qtd1)

	for i := range vetor0 {
		fmt.Scan(&vetor0[i])
	}
	for _, fig := range vetor0 {
		if unicos[fig] {
			repetidos = append(repetidos, fig)
		} else {
			unicos[fig] = true
		}
	}
	if len(repetidos) == 0 {
		fmt.Println("N")
	} else {
		for i, valor := range repetidos {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%v", valor)
		}
		fmt.Println("")
	}
	

	saida := ""
	for i := 1; i <= qtd0; i++ {

		if !unicos[i] {
			saida += fmt.Sprintf("%v ", i)
		}
	}

	if saida == "" {
		fmt.Println("N")
	} else {
		fmt.Println(saida[:len(saida)-1])
	}

	/*vetorRep := make([]int, qtd1)
	aux := -999999999999999999
	j := 0
	for i := 0; i < qtd1; i++ {
		fmt.Scan(&vetor0[i])
		if vetor0[i] == aux {
			vetorRep[j] = aux
			j++
		}
		aux = vetor0[i]

	}
	for i := 0; i < j; i++ {
		fmt.Print(vetorRep[i])
		if i < j-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")

	*/

}
