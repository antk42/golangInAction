package main

import (
	"log"
	"os/exec"
)

func main() {
	RunSimpleApp()
}

func RunSimpleApp() {
	cmd := exec.Command("firefox")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
