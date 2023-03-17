package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	RunAppWithSimpleArgs()
}

func RunAppWithSimpleArgs() {
	cmd := exec.Command("tr", "a-z", "A-Z") // команда может вызываться со списком аргументов

	cmd.Stdin = strings.NewReader("hello there!")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("translated phrase: %q\n", out.String()) //выводим результат
}
