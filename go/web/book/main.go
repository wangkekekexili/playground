package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "rsc.io/sqlite"
)

type Page struct {
	Name     string
	DBStatus bool
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
