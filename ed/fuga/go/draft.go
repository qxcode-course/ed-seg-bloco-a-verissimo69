package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d int16
	fmt.Scan(&a, &b, &c, &d)

	if d == -1 {
		if math.Abs(float64(a-b)) > math.Abs(float64(a-c)) || (a > b && b > c) {
			fmt.Printf("S\n")
		} else {
			fmt.Printf("N\n")
		}
	} else {
		if (math.Abs(float64(a-c)) < math.Abs(float64(a-b))) && (a > b || b > c) && b > c {
			fmt.Println("N")
		} else {
			fmt.Println("S")
		}
	}

}
