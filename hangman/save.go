package hangman

import (
	"encoding/json"
	"fmt"
	"os"
)

func (w *Hang) SaveFile() {
	var err error
	var file *os.File
	var u []byte
	os.Remove("save.txt")
	file, err = os.OpenFile("save.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	u, _ = json.Marshal(w)
	file.Write(u)
}

func (w *Hang) GetData() bool {
	var err error
	var data []byte
	data, err = os.ReadFile("save.txt")
	if err != nil {
		fmt.Println(err)
		return false
	}
	json.Unmarshal(data, w)
	return true
}
