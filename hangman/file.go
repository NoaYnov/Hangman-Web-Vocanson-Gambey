package hangman

import (
	"math/rand"
	"os"
	"time"
)

func (w *Hang) SaveFileToTab(dif string) []string {
	var list []string
	var data []byte
	var tmp string
	if dif == "Facile" {
		data, _ = os.ReadFile("doc/facile.txt")
	} else if dif == "Moyen" {
		data, _ = os.ReadFile("doc/moyen.txt")
	} else if dif == "Difficile" {
		data, _ = os.ReadFile("doc/difficile.txt")
	}
	for _, letter := range data {
		if letter == '\n' {
			list = append(list, tmp)
			tmp = ""
		} else if letter != '\r' {
			tmp += string(letter)
		}
	}
	list = append(list, tmp)
	return list
}

func RandomWord(file []string) string {
	rand.Seed(time.Now().UnixNano())
	return file[rand.Intn(len(file)-1)]
}
