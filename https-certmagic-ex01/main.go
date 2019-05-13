package main

import (
	"net/http"
	"os"

	"github.com/mholt/certmagic"
)

var (
	serverDNS = os.Getenv("ACME_SERVER_DNS") // Note: set to your real host
	CAEmail   = os.Getenv("ACME_CA_EMAIL")   // ACME CA account email address
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, encrypted world!"))
	})

	// read and agree to your CA's legal documents
	certmagic.Default.Agreed = true

	// provide an email address
	certmagic.Default.Email = CAEmail

	// uncomment to use the staging endpoint while we're developing
	//certmagic.Default.CA = certmagic.LetsEncryptStagingCA

	// tlsConfig, err := certmagic.TLS([]string{serverDNS})
	// check(err)

	// Serving HTTP handlers with HTTPS
	err := certmagic.HTTPS([]string{serverDNS}, handler)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
