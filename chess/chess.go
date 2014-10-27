package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"html/template"
	"log"
)
const (

TEMPLATE_DIR = "/Users/dong/go/src/github.com/lishaodong/Maze/chess/"

)



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/gui", func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
		t, err := template.ParseFiles(TEMPLATE_DIR + "index.html")
		if err != nil {
			log.Fatal(err)
		}
		index := index{Title: "首页"}
		t.Execute(w, index)
		}
	})

	mux.HandleFunc("/chess/", func (w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				t, err := template.ParseFiles(TEMPLATE_DIR + "chess.html")
				if err != nil {
					log.Fatal(err)
				}
				index := index{Title: "首页"}
				t.Execute(w, index)
			}
		})


	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

type index struct {

Title string

}

type HomeHandler struct {

}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t, err := template.ParseFiles(TEMPLATE_DIR + "login.html")

		if err != nil {

			log.Fatal(err)

		}

		index := index{Title: "首页"}

		t.Execute(w, index)

	}

}






