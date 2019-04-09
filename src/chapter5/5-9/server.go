package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var filePrefix string
var t *template.Template

func init() {
	filePrefix, _ := filepath.Abs("src/chapter5/5-9/templates")
	t1,err:= template.ParseFiles(filePrefix+"/t1.html")
	if err == nil {
		t = t1
	} else {
	   fmt.Println(err)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t.Execute(w,"www")
}
