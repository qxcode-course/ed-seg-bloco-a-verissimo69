package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. DEFINIÇÃO DA FILA GENÉRICA (O código precisa disso para compilar)
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

// 2. FUNÇÃO PRINCIPAL
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fila := NewQueue[string]()

	// Coloca as 16 letras (de 'A' até 'P') na fila
	for c := 'A'; c <= 'P'; c++ {
		fila.Enqueue(string(c))
	}

	for fila.Size() > 1 {
		if !scanner.Scan() {
			break
		}
		linha := scanner.Text()
		if strings.TrimSpace(linha) == "" {
			continue
		}

		// Lê os gols do jogo atual
		gols := strings.Fields(linha)
		if len(gols) < 2 {
			continue
		}
		golsEsquerda, _ := strconv.Atoi(gols[0])
		golsDireita, _ := strconv.Atoi(gols[1])

		// Retira os dois times que se enfrentam nesta chave
		timeEsquerda := fila.Dequeue()
		timeDireita := fila.Dequeue()

		// O vencedor volta para o fim da fila para a próxima fase
		if golsEsquerda > golsDireita {
			fila.Enqueue(timeEsquerda)
		} else {
			fila.Enqueue(timeDireita)
		}
	}

	// O último que restar na fila é o grande campeão
	fmt.Println(fila.Dequeue())
}
