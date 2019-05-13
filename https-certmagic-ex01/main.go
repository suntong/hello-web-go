package main

import (
	"net/http"

	"github.com/mholt/certmagic"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, encrypted world!"))
	})
	certmagic.HTTPS([]string{"example.com", "www.example.com"}, handler)
}
