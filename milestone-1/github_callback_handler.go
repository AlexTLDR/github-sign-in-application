package main

import (
	"fmt"
	"net/http"
)

func githubCallbackHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello from /github/callback!\n")
}
