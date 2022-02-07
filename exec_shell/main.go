package main

import (
	"os"
	"os/exec"
)

func main() {
	file, err := os.OpenFile("./exec_shell/output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	c := exec.Command("top")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}
