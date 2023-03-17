package main

import (
	"fmt"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Println(w, "Hello there!")
		_, _ = fmt.Fprintln(w, "Hello there!")

	}

	http.HandleFunc("/", helloHandler)

	_ = http.ListenAndServe(":8080", nil)

}
