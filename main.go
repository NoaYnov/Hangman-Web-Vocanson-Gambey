package main

import (
	"html/template"
	// "log"
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
	Name    string
}

func main() {
	var u User
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", u.Difficulte)
	http.HandleFunc("/register", u.register)
	http.ListenAndServe(":8080", nil)

}

func (u *User) Difficulte(w http.ResponseWriter, r *http.Request) {
	var hang hangman.Hang
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	details := User{
		choix1:  r.FormValue("facile"),
		choix2:  r.FormValue("moyen"),
		choix3:  r.FormValue("difficile"),
		Hang:    hang.Init(r, w, r.FormValue("choice")),
		Success: true,
	}
	tmpl1.Execute(w, details)

}

func (u *User) register(w http.ResponseWriter, r *http.Request) {
	var hang hangman.Hang
	tmp2 := template.Must(template.ParseFiles("register.html"))
	details := User{
		Hangman: r.FormValue("letter"),
		Hang:    hang.Start(r, w),
	}
	tmp2.Execute(w, details)

}


