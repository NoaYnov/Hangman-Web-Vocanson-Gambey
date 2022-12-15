package hangman

import (
	"math/rand"
	"time"
)

func (w *Hang) GetRandomLetter() {
	for i := 0; i != w.NbLetter/2-1; {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Intn(w.NbLetter - 1)
		if w.Guess[tmp] == "_" {
			i++
			w.Guess[tmp] = string(w.Word[tmp])
		}
	}
}
