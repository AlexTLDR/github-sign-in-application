package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello from \\!\n")
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("cookie set!"))

}
