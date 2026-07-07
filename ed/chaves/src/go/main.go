package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ==========================================
// 1. DEFINIÇÃO DA FILA GENÉRICA
// ==========================================
type Queue[T any] struct {
	items *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: list.New()}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items.PushBack(item)
}

func (q *Queue[T]) Dequeue() T {
	if q.items.Len() == 0 {
		var zero T
		return zero
	}
	front := q.items.Front().Value.(T)
	q.items.Remove(q.items.Front())
	return front
}

func (q *Queue[T]) IsEmpty() bool {
	return q.items.Len() == 0
}

func (q *Queue[T]) Size() int {
	return q.items.Len()
}

func inicializarTimes() *Queue[string] {
	fila := NewQueue[string]()
	for c := 'A'; c <= 'P'; c++ {
		fila.Enqueue(string(c))
	}
	return fila
}

func processarPartida(linha string, fila *Queue[string]) {
	gols := strings.Fields(linha)
	if len(gols) < 2 {
		return
	}

	golsEsquerda, _ := strconv.Atoi(gols[0])
	golsDireita, _ := strconv.Atoi(gols[1])

	// Remove os dois times que estão jogando nesta chave
	timeEsquerda := fila.Dequeue()
	timeDireita := fila.Dequeue()

	// O vencedor volta para o fim da fila para a próxima fase
	if golsEsquerda > golsDireita {
		fila.Enqueue(timeEsquerda)
	} else {
		fila.Enqueue(timeDireita)
	}
}

// executarTorneio gerencia o fluxo de leitura da entrada e execução dos jogos
func executarTorneio(fila *Queue[string]) {
	scanner := bufio.NewScanner(os.Stdin)

	// Enquanto houver mais de um time, o torneio continua
	for fila.Size() > 1 {
		if !scanner.Scan() {
			break
		}
		linha := scanner.Text()
		if strings.TrimSpace(linha) == "" {
			continue
		}

		processarPartida(linha, fila)
	}
}

func main() {
	fila := inicializarTimes()
	executarTorneio(fila)

	fmt.Println(fila.Dequeue())
}
