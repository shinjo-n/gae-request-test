package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("error")
	}

	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}

	w.Write([]byte("success"))
}
