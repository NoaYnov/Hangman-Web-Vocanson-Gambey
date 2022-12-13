package hangman

import (
	"fmt"
	"net/http"
	"os"
)

func (hang *Hang) Init(r *http.Request, w http.ResponseWriter, diff string) *Hang {
	hang.InitHang(diff)
	hang.GetRandomLetter()
	fmt.Println(hang.Word)
	hang.SaveFile()
	return hang
}

func (hang *Hang) Start(r *http.Request, w http.ResponseWriter) *Hang {
	var letter string
	var tmp rune
	hang.GetData()
	fmt.Println("il vous reste", hang.NbTry, "essais")
	fmt.Println(hang.Guess)
	fmt.Print("Choose a letter: ")
	letter = r.FormValue("letter")
	fmt.Println(letter)
	if len(letter) == 1 {
		tmp = rune(letter[0])
		if tmp >= 97 && tmp <= 122 {
			if tmp >= 65 && tmp <= 90 {
				tmp += 32
				letter = ""
				letter += string(tmp)
			}
			hang.CheckLetter(letter)
			hang.CheckEnd()
		} else {
			fmt.Println("Not a letter")
		}
	} else if len(letter) == len(hang.Word) {
		if letter == hang.Word {
			for i := 0; i < len(hang.Word); i++ {
				hang.Guess[i] = string(hang.Word[i])
			}
			fmt.Println("you win")
			hang.Loop = false
			hang.Win = true
		} else {
			if hang.NbTry <= 2 {
				fmt.Println("You lose, the word was :", hang.Word)
				hang.Loop = false
			} else {
				hang.I += 2
				hang.NbTry -= 2
				fmt.Print("wrong guess, ", hang.NbTry, "/10 tries left\n")
			}
		}
	} else {
		fmt.Println("You must enter a letter or a guess")
	}
	fmt.Println(hang.Guess)
	fmt.Println(hang.José[hang.I])
	hang.SaveFile()
	if !hang.Loop {
		os.Remove("save.txt")
	}
	return hang
}
