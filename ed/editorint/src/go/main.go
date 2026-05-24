package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	newLine := NewList[rune]()
	for e.cursor != e.line.Value.End() {
		val := e.cursor.Value
		next := e.cursor.Next()

		e.line.Value.Erase(e.cursor)
		newLine.PushBack(val)

		e.cursor = next
	}

	e.lines.Insert(e.line.Next(), newLine)
	e.line = e.line.Next()
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.cursor.Next()
	} else if e.line.Next() != e.lines.End() {
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front() // Começa no início da nova linha
	}
}

func (e *Editor) KeyUp() {
	if e.line != e.lines.Front() {
		// Descobre a posição atual do cursor na linha
		x := e.line.Value.IndexOf(e.cursor)

		// Sobe a linha
		e.line = e.line.Prev()
		e.cursor = e.line.Value.Front()

		// Avança o cursor na nova linha até chegar no 'x' original ou no fim da linha
		for i := 0; i < x && e.cursor != e.line.Value.End(); i++ {
			e.cursor = e.cursor.Next()
		}
	}
}

func (e *Editor) KeyDown() {
	if e.line.Next() != e.lines.End() {
		// Descobre a posição atual do cursor na linha
		x := e.line.Value.IndexOf(e.cursor)

		// Desce a linha
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front()

		// Avança o cursor na nova linha até chegar no 'x' original ou no fim da linha
		for i := 0; i < x && e.cursor != e.line.Value.End(); i++ {
			e.cursor = e.cursor.Next()
		}
	}
}
func (e *Editor) KeyBackspace() {
	// 1. Caso comum: Apagar caratere à esquerda
	if e.cursor != e.line.Value.Front() {
		toDelete := e.cursor.Prev()
		e.line.Value.Erase(toDelete)
		return
	}

	// 2. Caso de início de linha: Unir com a linha de cima
	if e.line != e.lines.Front() {
		prevLineNode := e.line.Prev()
		lineAcima := prevLineNode.Value
		lineAtual := e.line.Value

		// Se a linha atual tiver conteúdo, "costuramos" ela na de cima
		if lineAtual.Size() > 0 {
			ultimoAcima := lineAcima.Back()    // Último nó real de cima
			primeiroAtual := lineAtual.Front() // Primeiro nó real de baixo
			ultimoAtual := lineAtual.Back()    // Último nó real de baixo

			// Fazemos a ponte entre as duas listas
			ultimoAcima.next = primeiroAtual
			primeiroAtual.prev = ultimoAcima

			// Fechamos a nova lista unida com o root da linha de cima
			ultimoAtual.next = lineAcima.End()
			lineAcima.End().prev = ultimoAtual

			// Atualizamos o tamanho da lista de cima
			lineAcima.size += lineAtual.size
		}

		// Guardamos a posição para o cursor (onde era o antigo End da linha de cima)
		e.cursor = lineAcima.Front()
		// (Opcional) Se quiser que o cursor fique exatamente na emenda:
		// novoCursor = primeiroAtual (se a linha atual não estava vazia)

		// Removemos a linha atual da lista de linhas
		linhaParaRemover := e.line
		e.line = prevLineNode
		e.cursor = lineAcima.End() // Coloca o cursor na "emenda" (final da linha anterior)
		e.lines.Erase(linhaParaRemover)
	}
}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.line.Value.Erase(e.cursor)
		return
	}

	if e.line != e.lines.Front() {
		prevLineNode := e.line.Prev()
		lineAcima := prevLineNode.Value
		lineAtual := e.line.Value

		// Se a linha atual tiver conteúdo, "costuramos" ela na de cima
		if lineAtual.Size() > 0 {
			ultimoAcima := lineAcima.Back()    // Último nó real de cima
			primeiroAtual := lineAtual.Front() // Primeiro nó real de baixo
			ultimoAtual := lineAtual.Back()    // Último nó real de baixo

			// Fazemos a ponte entre as duas listas
			ultimoAcima.next = primeiroAtual
			primeiroAtual.prev = ultimoAcima

			// Fechamos a nova lista unida com o root da linha de cima
			ultimoAtual.next = lineAcima.End()
			lineAcima.End().prev = ultimoAtual

			// Atualizamos o tamanho da lista de cima
			lineAcima.size += lineAtual.size
		}

		// Guardamos a posição para o cursor (onde era o antigo End da linha de cima)
		e.cursor = lineAcima.Front()
		// (Opcional) Se quiser que o cursor fique exatamente na emenda:
		// novoCursor = primeiroAtual (se a linha atual não estava vazia)

		// Removemos a linha atual da lista de linhas
		linhaParaRemover := e.line
		e.line = prevLineNode
		e.cursor = lineAcima.End() // Coloca o cursor na "emenda" (final da linha anterior)
		e.lines.Erase(linhaParaRemover)
	}

}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
