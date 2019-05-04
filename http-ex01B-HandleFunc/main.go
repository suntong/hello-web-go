package main

import (
	"bufio"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func test0(w http.ResponseWriter, r *http.Request) {
	// Handles top-level page.
	fmt.Fprintf(w, `You are on the home page

Here are the sample urls that you can visit:

- /about
- /images
- and /image


`)
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

func Image(w http.ResponseWriter, r *http.Request) {
	// Open a JPG file.
	f, _ := os.Open("vlc-16x16.png")

	// Read the entire JPG file into memory.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Set the Content Type header.
	w.Header().Set("Content-Type", "image/png")

	// Write the image to the response.
	w.Write(content)
}

func main() {
	// Add handle funcs for 3 pages.
	http.HandleFunc("/", test0)
	http.HandleFunc("/about", test1)
	http.HandleFunc("/images", test2)
	http.HandleFunc("/image", Image)

	// Run the web server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
