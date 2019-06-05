package main

import (
	"fmt"
	"net/http"
	"strings"
)

type specificHandler struct {
	Thing string
}

func (h *specificHandler) handle(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	h := &specificHandler{Thing: "Hello world!"}
	http.HandleFunc("/", h.handle)
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
