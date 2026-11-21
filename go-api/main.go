package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            fmt.Fprint(w, "OK")
        } else {
            http.NotFound(w, r)
        }
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "OK")
    })

    fmt.Println("Go server running on http://0.0.0.0:8080")
    http.ListenAndServe("0.0.0.0:8080", nil)
}