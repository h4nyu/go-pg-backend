package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page. Welcome!")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homePage)
}

