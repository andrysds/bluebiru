package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func healthz(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK!")
}

func main() {
	http.HandleFunc("/healthz", healthz)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
