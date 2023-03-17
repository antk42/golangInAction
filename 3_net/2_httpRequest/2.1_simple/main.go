package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	fmt.Println("StatusCode: ", res.StatusCode)
}
