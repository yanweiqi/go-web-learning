package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	//s, _ := json.Marshal(r.Form)
	//fmt.Println(string(s))
	fmt.Fprintln(w,"(0)",r.Form)
	fmt.Fprintln(w,"(1)",r.FormValue("p1"))
	fmt.Fprintln(w,"(2)",r.PostFormValue("p2"))
	fmt.Fprintln(w,"(3)",r.PostForm)
	fmt.Fprintln(w,"(4)",r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/process",process)
	server.ListenAndServe()
}
