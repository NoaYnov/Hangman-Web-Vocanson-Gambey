package main


import (
	"html/template"
	// "log"
	"net/http"
	"src/hangman"
	"fmt"
)

type User struct {
    Hangman string
    Success bool
	choix1 string
	choix2 string
	choix3 string

}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", difficulté)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)

}

 

func difficulté(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	details := User{
		choix1: r.FormValue("facile"),
		choix2: r.FormValue("moyen"),
		choix3: r.FormValue("difficile"),
		Success: true,
	}
	tmpl1.Execute(w, details)
	Test(r, w)
	
	
	
}


func register(w http.ResponseWriter, r *http.Request) {
	tmp2 := template.Must(template.ParseFiles("register.html"))
	if r.Method != http.MethodPost {
		tmp2.Execute(w, nil)
		return
	}
	details := User{
		Hangman: r.FormValue("letter"),
		Success: true,
		
	}
	tmp2.Execute(w, details)
	Testform(r, w)

	
}



func HangDif(r *http.Request, w http.ResponseWriter) {
	var hang hangman.Hang
	hang.Start(r.FormValue("choice"))
	return
	
}


func Test(r *http.Request, w http.ResponseWriter) {
	fmt.Println(r.FormValue("choice"))
}
  
func Testform(r *http.Request, w http.ResponseWriter) {
	fmt.Println(r.FormValue("letter"))
}