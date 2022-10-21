package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{\"state\": \"UP\"}`)
}

func main() {
	db := database{"shoes": 104.40, "socks": 9.8}
	mux := http.NewServeMux()
	mux.Handle("/health", http.HandlerFunc(health))
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/", db.list)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
