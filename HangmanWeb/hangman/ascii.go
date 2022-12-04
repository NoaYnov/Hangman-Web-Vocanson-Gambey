package hangman

import "os"

func SaveJoséToTab() []string {
	var list []string
	var data []byte
	var tmp string
	list = append(list, "\n")
	data, _ = os.ReadFile("doc/josé.txt")
	for i := 0; i < len(data); i++ {
		if data[i] == '=' {
			i += 8
			tmp += "========"
			list = append(list, tmp)
			tmp = ""
		} else if data[i] != '\r' {
			tmp += string(data[i])
		}
	}
	list = append(list, tmp)
	return list
}
