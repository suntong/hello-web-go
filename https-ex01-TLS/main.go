package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var tls = false

	flag.BoolVar(&tls, "tls", false, "Set to true to enable HTTPS rather than HTTP")
	flag.Parse()

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/ping", ping)

	var err error
	if tls {
		log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
		err = http.ListenAndServeTLS(":10443", "server.crt", "server.key", nil)
	} else {
		log.Printf("About to listen on 8080. Go to http://127.0.0.1:8080/")
		err = http.ListenAndServe(":8080", nil)
	}
	log.Fatal(err)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is an example https server.\n"))
}
