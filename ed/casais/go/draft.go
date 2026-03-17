package main

import "fmt"

func main() {
    var count, valor, pares int
    fmt.Scan(&count)

    estoque := make(map[int]int)

    for i := 0; i < count; i++ {
        fmt.Scan(&valor)
        if estoque[-valor] > 0 {
            pares++           
            estoque[-valor]--  
        } else {
            estoque[valor]++
        }
    }

    fmt.Println(pares)
}