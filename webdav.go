package main

import (
	"flag"
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	var address string
	flag.StringVar(&address, "a", "0.0.0.0:7000", "Address to listen to.")
	flag.Parse()

	handler := &webdav.Handler{
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
	}

	http.ListenAndServe(address, handler)
}
