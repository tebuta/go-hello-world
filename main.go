package main

import (
    "net/http"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/ping" {
        fmt.Fprintf(w, "OK")
    } else {
        fmt.Fprintf(w, "HelloWorld")
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/ping", handler)
    http.ListenAndServe(":8080", nil)
}
