package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Router struct {
	Model string
	Wan   string
}

var index = template.Must(template.ParseFiles("index.html"))
var form = template.Must(template.ParseFiles("form.html"))

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		index.Execute(writer, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		router := request.Form["router"][0]
		wan := request.Form["wan"][0]
		data := router + " " + wan

		r1 := Router{
			Model: router,
			Wan:   wan,
		}
		json_data, err := json.Marshal(r1)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(json_data))
		form.Execute(writer, data)
	}

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
