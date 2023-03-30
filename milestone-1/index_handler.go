package main

import (
	"errors"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
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

// func setCookieHandler(w http.ResponseWriter, r *http.Request) {

// }

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}
	w.Write([]byte(cookie.Value))
}
