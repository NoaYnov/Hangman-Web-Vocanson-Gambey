package hangman

import (
	"net/http"
	"os"
)

func (hang *Hang) Init(r *http.Request, w http.ResponseWriter, diff string) *Hang {
	hang.InitHang(diff)
	hang.GetRandomLetter()
	hang.SaveFile()
	return hang
}

func (hang *Hang) Start(r *http.Request, w http.ResponseWriter) *Hang {
	var letter string
	var tmp rune
	hang.GetData()
	letter = r.FormValue("letter")
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
		}
	} else if len(letter) == len(hang.Word) {
		if letter == hang.Word {
			for i := 0; i < len(hang.Word); i++ {
				hang.Guess[i] = string(hang.Word[i])
			}
			hang.Loop = false
			hang.Win = true
		} else {
			if hang.NbTry <= 2 {
				hang.Loop = false
			} else {
				hang.I += 2
				hang.NbTry -= 2

			}
		}
	} else {
	}

	hang.SaveFile()
	if !hang.Loop {
		os.Remove("save.txt")
	}
	return hang
}
