package main

import (
	"html/template"
	// "log"
	"fmt"
	"net/http"
	"src/hangman"
)

type User struct {
	Hang    *hangman.Hang
	Hangman string
	Success bool
	choix1  string
	choix2  string
	choix3  string
}

type Input struct {
	Diff string
	Name string
}

func main() {
	var i Input
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", i.difficulté)
	http.HandleFunc("/register", i.register)
	http.ListenAndServe(":8080", nil)

}

func (i *Input) difficulté(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	details := User{
		choix1:  r.FormValue("facile"),
		choix2:  r.FormValue("moyen"),
		choix3:  r.FormValue("difficile"),
		Success: true,
	}
	tmpl1.Execute(w, details)
	i.Diff = r.FormValue("choice")
	Test(r, w)
}

func (i *Input) register(w http.ResponseWriter, r *http.Request) {
	var hang hangman.Hang
	tmp2 := template.Must(template.ParseFiles("register.html"))
	if r.Method != http.MethodPost {
		tmp2.Execute(w, nil)
		return
	}
	details := User{
		Hangman: r.FormValue("letter"),
		Hang:    hang.Start(r, w, i.Diff),
	}
	tmp2.Execute(w, details)
	Testform(r, w)
}

func Test(r *http.Request, w http.ResponseWriter) {
	fmt.Println(r.FormValue("choice"))
}

func Testform(r *http.Request, w http.ResponseWriter) {
	fmt.Println(r.FormValue("letter"))
}
