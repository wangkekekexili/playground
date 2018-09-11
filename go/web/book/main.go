package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	_ "rsc.io/sqlite"
)

type Page struct {
	Name     string
	DBStatus bool
}

type SearchResult struct {
	Title  string `json:"name" xml:"title,attr"`
	Author string `json:"author" xml:"author,attr"`
	Year   string `json:"year" xml:"hyr,attr"`
	ID     string `json:"id" xml:"owi,attr"`
}

type ClassifySearchResult struct {
	Results []SearchResult `xml:"works>work"`
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
		results, _ := search(r.FormValue("search"))
		json.NewEncoder(w).Encode(results)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func search(query string) ([]SearchResult, error) {
	u := "http://classify.oclc.org/classify2/Classify?summary=true&title=" + url.QueryEscape(query)
	resp, err := http.Get(u)
	if err != nil {
		return nil, errors.Wrap(err, "get from url")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}

	r := &ClassifySearchResult{}
	err = xml.Unmarshal(b, r)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal result")
	}
	return r.Results, nil
}
