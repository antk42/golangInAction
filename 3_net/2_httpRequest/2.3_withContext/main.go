package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := &http.Client{Timeout: 5 * time.Second}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second/2)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://www.google.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("StatusCode: ", resp.StatusCode)
}
