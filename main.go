package main

import (
    "html/template"
    "net/http"
    "os"
)

var index = template.Must(template.ParseFiles("index.html"))
var form = template.Must(template.ParseFiles("form.html"))

func indexHandler(writer http.ResponseWriter, request *http.Request) {
    if request.Method == http.MethodGet {
        index.Execute(writer, nil)
    } else if request.Method == http.MethodPost {
        request.ParseForm()
        data1 := request.Form["router"][0]
        data2 := request.Form["wan"][0]
        data3 := data1 + " " + data2
        form.Execute(writer, data3)
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