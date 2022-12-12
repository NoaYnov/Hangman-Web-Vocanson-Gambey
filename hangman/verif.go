package hangman

import (
	"fmt"
	"strings"
)

func (w *Hang) CheckLetter(letter string) {
	var tmp bool
	if strings.Contains(w.Letters, letter) {
		fmt.Println("Already used")
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
		fmt.Print("Wrong letter, ", w.NbTry, "/10 tries left\n")
	}
	w.Letters += letter
}

func (w *Hang) CheckEnd() {
	if w.NbTry <= 0 {
		w.Loop = false
		fmt.Println("You lose, the word was", w.Word)
		w.Loop = false
	}
	for i := 0; i < len(w.Word); i++ {
		if w.Guess[i] == "_" {
			return
		}
	}
	fmt.Println("you win")
	w.Loop = false
}
