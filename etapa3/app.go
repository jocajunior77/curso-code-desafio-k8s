package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func showGreeting(w http.ResponseWriter, r *http.Request) {
    msg := greeting("Code.education Rocks!")
    w.Write([]byte(msg))
}


func greeting(msg string) (string) {
    return "<b>" + msg + "</b>"
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", showGreeting)
    http.ListenAndServe(":8000", r)
}