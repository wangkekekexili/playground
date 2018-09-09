package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "rsc.io/sqlite"
)

type Page struct {
	Name     string
	DBStatus bool
}

type SearchResult struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   string `json:"year"`
	ID     string `json:"id"`
}

func main() {
	db, err := sql.Open("sqlite3", "dev.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tmpl := template.Must(template.ParseFiles("template/index.gohtml"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := &Page{
			Name:     "gopher",
			DBStatus: false,
		}

		name := r.FormValue("name")
		if name != "" {
			p.Name = name
		}
		p.DBStatus = db.Ping() == nil

		tmpl.ExecuteTemplate(w, "index.gohtml", p)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		results := []SearchResult{
			{"Unlimited Memory", "Kevin", "2014", "B00I3QS1XQ"},
			{"The Extraordinary Life of Sam Hell: A Novel", "Robert", "2018", "B07BNVZDM7"},
		}
		json.NewEncoder(w).Encode(results)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
