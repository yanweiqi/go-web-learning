package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/write",writeExample)
	http.HandleFunc("/header",headExample)
	http.HandleFunc("/redirect",writeHeaderExample)
	http.HandleFunc("/json",jsonExample)
	server.ListenAndServe()
}

func writeHeaderExample(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Location","http://google.com")
	w.WriteHeader(302)
}

func writeExample(w http.ResponseWriter,r *http.Request){
	html :=`<html> 
           <head>
           <title>Go Web Programming</title>
           </head>
           <body>
           <h1>jkjkjkj</H1>
           </html>`
	w.Write([]byte(html))
}

func headExample(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(501)
	fmt.Fprintln(w,"No such server,try next door")
}

func jsonExample(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	post := &Post{
		User:"Sau Sheong",
		Threads :[]string{"first","second","third"},
	}
	json,_:= json.Marshal(post)
	w.Write(json)
}

type Post struct {
	User    string
	Threads []string
}