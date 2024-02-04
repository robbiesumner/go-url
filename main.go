package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robbiesumner/go-url/db_store"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	dbClient := db_store.NewDBClient()
	err := dbClient.Set("foo1", "bar1")
	fmt.Println(err)
	err = dbClient.Set("foo1", "bar2")
	fmt.Println(err)
	val, err := dbClient.Get("foo2")
	fmt.Println(err, val)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
