package main

import (
    "net/http"
    "encoding/json"
    "fmt"
)

type Page struct {
    PageNumber int
}

func (p *Page) Next() {
    p.PageNumber += 1
}

func (p *Page) Prev() {
    if p.PageNumber > 1 {
        p.PageNumber -= 1
    }
}

func GeneratePage() Page {
    return Page{PageNumber: 1}
}

func main() {
    p := GeneratePage()

    http.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(p)
	fmt.Println(p.PageNumber)
    })

    http.HandleFunc("/next", func(w http.ResponseWriter, req *http.Request) {
        p.Next()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(p)
	fmt.Println(p.PageNumber)
    })

    http.HandleFunc("/prev", func(w http.ResponseWriter, req *http.Request) {
        p.Prev()
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(p)
	fmt.Println(p.PageNumber)
    })

    http.ListenAndServe(":8080", nil)
}
