package hangman

import (
	"strings"
)

func (w *Hang) CheckLetter(letter string) {
	var tmp bool
	if strings.Contains(w.Letters, letter) {
		return
	}
	for i := 0; i < len(w.Word); i++ {
		if string(w.Word[i]) == letter {
			w.Guess[i] = letter
			tmp = true
		}
	}
	if !tmp {
		w.I++
		w.NbTry--
	}
	w.Letters += letter
}

func (w *Hang) CheckEnd() {
	if w.NbTry <= 0 {
		w.Loop = false
	}
	for i := 0; i < len(w.Word); i++ {
		if w.Guess[i] == "_" {
			return
		}
	}
	w.Loop = false
	w.Win = true
}
