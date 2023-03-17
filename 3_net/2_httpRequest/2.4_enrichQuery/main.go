package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	req, err := http.NewRequest(http.MethodGet,
		"https://reqres.in/api/users", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Custom", "Custom")

	q := url.Values{}
	q.Add("page", "2")
	q.Add("pAge", "2")
	req.URL.RawQuery = q.Encode()

	fmt.Println()
	fmt.Println("Headers: ", req.Header)
	fmt.Println()
	fmt.Println("URL: ", req.URL)
}
