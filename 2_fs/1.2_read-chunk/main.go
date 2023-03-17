package main

import (
	"bufio"
	"fmt"
	"io"
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
	f, err := os.Open(filepath)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	b2 := make([]byte, 3)
	_, err = f.Read(b2)
	check(err)
	fmt.Printf("%v\n", string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(15)
	check(err)
	fmt.Printf("15 bytes: %s\n", string(b4))

	f.Close()
}
