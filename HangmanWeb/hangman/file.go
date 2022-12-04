package hangman

import ("os"
		"time"
		"math/rand"
	)


func SaveFileToTab(dif string) []string {
	var list []string
	var data []byte
	var tmp string
	if dif == "facile" {
		data, _ = os.ReadFile("doc/facile.txt")
	} else if dif == "moyen" {
		data, _ = os.ReadFile("doc/moyen.txt")
	} else if dif == "difficile" {
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

func RandomWord (file []string) string{
	rand.Seed(time.Now().UnixNano())
    return file[rand.Intn(len(file)-1)]
}