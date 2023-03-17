package main

import (
	"fmt"
	"os"
)

func main() {
	read("test.txt")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(filepath string) {
	dat, err := os.ReadFile(filepath)
	check(err)
	fmt.Print(string(dat))
}
