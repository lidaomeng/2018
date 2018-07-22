package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type dataBase map[string]dollars

func (db dataBase) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := dataBase{"shoes": 50, "socks": 3}

	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
