package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", IndexHandler)

	http.HandleFunc("/json", jsonHandlerFunc)

	//	ポート番号を環境変数から取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//	サービスを開始
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	if err := r.ParseForm(); err != nil {
		fmt.Println("error")
	}

	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}

	w.Write([]byte("success"))
}

func jsonHandlerFunc(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v\n", jsonBody)
	for k, v := range jsonBody {
		fmt.Printf("%v : %v\n", k, v)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
