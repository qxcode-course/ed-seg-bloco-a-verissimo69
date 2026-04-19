package main

import (
	"fmt"
)

func embuahhh(pen *Pen, dista float64) {
	pen.SetRGB(100+dista, 0, 0)
	if dista < 10 {
		return
	}
	//pen.SetLineWidth(dista)
	pen.Walk(dista)
	pen.Right(90)
	embuahhh(pen, dista-5)

}

func main() {
	pen := NewPen(500, 500)
	pen.SetPosition(100, 400)
	pen.SetHeading(90)
	side := 300.0
	embuahhh(pen, side)
	pen.SavePNG("Quadrado.png")
	fmt.Println("PNG file created successfully.")
}
