package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := "./files_io/test.txt"
	readFromFile(path)
	writeToFile(path)
	appendToFile(path)
}

func appendToFile(path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("\nAppending from ./file_io/main.go")
	if err != nil {
		panic(err)
	}
}

func writeToFile(path string) {
	data := []byte("Writing from ./files_io/main.go")
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
}

func readFromFile(path string) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
