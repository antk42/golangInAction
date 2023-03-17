package main

import (
	"os"
)

func main() {
	write("test.txt", "HELLO THERE!")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(filepath string, message string) {
	d1 := []byte(message)
	err := os.WriteFile(filepath, d1, 0644)
	check(err)
}
