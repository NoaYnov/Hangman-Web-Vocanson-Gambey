


package main


import (
	"html/template"
	"log"
	"net/http"
	"src/hangman"
	// "fmt"
)



//link css



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
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl1.Execute(w, nil)
            return
        }
        details := User{
            // Hangman:	r.FormValue("hangman"),
            // Success:	true,
			choix1: r.FormValue("facile"),
			choix2: r.FormValue("moyen"),
			choix3: r.FormValue("difficile"),
        }
		// register(r)
        tmpl1.Execute(w, details)
		
	})
	tmpl2 := template.Must(template.ParseFiles("register.html"))
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl2.Execute(w, nil)
            return
        }
		details := User{
			Hangman:	r.FormValue("hangman"),
			Success:	true,
		}
        
        tmpl2.Execute(w, details)

	})






	log.Fatal(http.ListenAndServe(":8080", nil))
}





func HangDif(r *http.Request) {
	var hang hangman.Hang
	hang.InitHang(r.FormValue("choice"))
	// fmt.Println(choix)


}