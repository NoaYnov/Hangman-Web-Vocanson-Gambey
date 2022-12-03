


package main


import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)


//link css



type User struct {
    // Hangman string
    // Success bool
	choix1 string
	choix2 string
	choix3 string

}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	// tmpl2 := template.Must(template.ParseFiles("register.html"))
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
		register(r)
        tmpl1.Execute(w, details)
		
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func register(r *http.Request) {
	fmt.Println(r.FormValue("choice"))
		//write in info.txt
		file, err := os.OpenFile("info.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if _, err = file.WriteString(r.FormValue("choice")); err != nil {
			panic(err)
		}
		fmt.Println("write in info.txt")




		


}