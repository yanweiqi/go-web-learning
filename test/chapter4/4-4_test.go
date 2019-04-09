package chapter4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type ServerSlice struct {
	Server   []Server
	ServerId string
}

var url = "http://127.0.0.1:8080/process"
var encType = "application/x-www-form-urlencoded"

func TestPost(t *testing.T) {

	s1 := ServerSlice{ServerId: "4-4", Server: []Server{{"bj", "127.0.0.1:8080"}}}
	s2 := ServerSlice{ServerId: "4-4", Server: []Server{{"sh", "127.0.0.1:8080"}}}

	b1, _ := json.Marshal(s1)
	b2, _ := json.Marshal(s2)

	var f []string

	f = append(append(f, "p1="), string(b1))
	f = append(append(f, "&p2="), string(b2))
	resp, _ := http.Post(url, encType, strings.NewReader(strings.Join(f, "")))

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	/**
	var m map[string]string
	json.Unmarshal(body, &m)
    for k,v :=range m{
    	fmt.Println(k,v)
	}**/
}
