package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, p, r float64
	fmt.Scan(&a, &b, &c)
	p = (a + b + c) / 2
	r = p * (p - a) * (p - b) * (p - c)
	r = math.Sqrt(r)
	fmt.Printf("%.2f\n", r)

}
