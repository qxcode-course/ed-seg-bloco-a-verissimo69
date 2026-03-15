package main

/*
"crianca" se menor que 12 (não use o ç),
- "jovem" se menor que 18,
- "adulto" se menor que 65,
- "idoso" se menor que 1000,
- "mumia" caso contrario (não ponha o acento).
 Uma frase no formato "`<nome>` eh `<classificação>`"
*/
import (
	"fmt"
)

func main() {

	var nome string
	var idade int16
	fmt.Scan(&nome, &idade)
	switch {
	case idade < 12:
		fmt.Printf("%s eh crianca\n", nome)
	case idade < 18:
		fmt.Printf("%s eh jovem\n", nome)
	case idade < 65:
		fmt.Printf("%s eh adulto\n", nome)
	case idade < 1000:
		fmt.Printf("%s eh idoso\n", nome)
	case idade > 1000:
		fmt.Printf("%s eh mumia\n", nome)

	}
}
