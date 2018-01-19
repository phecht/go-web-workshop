package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, s)

}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rURI := r.RequestURI
	u, err := url.Parse(rURI)
	if err != nil {
		fmt.Println("Problem:", err)
	}
	fmt.Fprint(w, u.Path)
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye, web")
}

func main() {
	http.Handle("/string", String("I'm a frayed knot"))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bye", byeHandler)
	http.Handle("/bye2", String("Bye, web!"))
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
