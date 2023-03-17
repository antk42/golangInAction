package main

import (
	"fmt"
	"os/exec"
)

func main() {
	RunMultipleArgs()
}

func RunMultipleArgs() {
	prg := "echo"

	arg1 := "arg1"
	arg2 := "arg2"
	arg3 := "arg3"
	arg4 := "arg4"
	arg5 := "arg5"

	cmd := exec.Command(prg, arg1, prg, arg2, arg3, arg4, arg5)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))
}
