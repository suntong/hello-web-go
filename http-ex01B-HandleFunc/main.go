package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func test0(w http.ResponseWriter, r *http.Request) {
	// Handles top-level page.
	fmt.Fprintf(w, "You are on the home page")
}

func test1(w http.ResponseWriter, r *http.Request) {
	// Handles about page.
	// ... Get the path from the URL of the request.
	path := html.EscapeString(r.URL.Path)
	fmt.Fprintf(w, "Now you are on: %q", path)
}

func test2(w http.ResponseWriter, r *http.Request) {
	// Handles "images" page.
	fmt.Fprintf(w, "Image page")
}

func main() {
	// Add handle funcs for 3 pages.
	http.HandleFunc("/", test0)
	http.HandleFunc("/about", test1)
	http.HandleFunc("/images", test2)

	// Run the web server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
