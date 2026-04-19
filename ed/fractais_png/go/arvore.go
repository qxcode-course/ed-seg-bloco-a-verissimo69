package main

import (
	"fmt"
)

func arvore(pen *Pen, dista float64) {

	ang := 25.0
	fator := 0.82
	if dista < 10 {
		return
	}
	//pen.SetLineWidth(dista)
	pen.Walk(dista)
	pen.Right(ang)
	arvore(pen, dista*fator)
	pen.Left(2 * ang)
	arvore(pen, dista*fator)
	pen.Right(ang)
	pen.Walk(-dista)

}

func main() {
	pen := NewPen(1000, 1000)
	pen.SetPosition(500, 1000)
	pen.SetHeading(90)
	side := 100.0
	arvore(pen, side)
	pen.SavePNG("arvore.png")
	fmt.Println("PNG file created successfully.")
}
