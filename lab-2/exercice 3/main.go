package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Bonjour, Docker!")
    })

    fmt.Println("Serveur démarré sur le port 8080")
    http.ListenAndServe(":8080", nil)
}
