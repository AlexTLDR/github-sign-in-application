package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello from \\!\n")
}
