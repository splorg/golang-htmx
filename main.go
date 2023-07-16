package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello world!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
    films := map[string][]Film{
      "Films": {
        {Title: "Tropa de Elite", Director: "José Padilha"},
        {Title: "O Auto da Compadecida", Director: "Guel Arraes"},
      },
    }
		tmpl.Execute(w, films)
	}

  h2 := func (w http.ResponseWriter, r *http.Request) {
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
  }

	http.HandleFunc("/", h1)
  http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
