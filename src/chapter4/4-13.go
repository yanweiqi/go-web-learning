package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	COOKIE_MAX_MAX_AGE = time.Hour * 24 / time.Second
	maxAge             = int(COOKIE_MAX_MAX_AGE)
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "go web programing",
		HttpOnly: true,
		MaxAge:   maxAge,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
		MaxAge:   maxAge,
	}
	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}
