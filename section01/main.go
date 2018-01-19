package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func doGet() {
	req, err := http.NewRequest("GET", "https://golang.org", nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	fmt.Println(res.Status)

}

func main() {
	// try changing the value of this url
	res, err := http.Get("https://golang.org/")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode == 404 {
		fmt.Println(res.Status)
	} else {
		defer res.Body.Close()
		_, err := io.Copy(os.Stdout, res.Body)
		if err != nil {
			log.Fatal(err)
		}

	}
	doGet()
}

// curl https://http-methods.appspot.com/pdevhecht56/MessageIsTheMedium
