package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Books   []Book `json:"books"`
}

func main() {
	bytes := []byte(`{"name":"Muslim","surname":"Yarashov","age":19,"books":[{"name":"Atomic Habits","author":"James Clear"},{"name":"Steve Jobs","author":"Walter Isaacson"}]}`)
	var data User

	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}
	fmt.Println(data.Books)
}

func serialize() {
	var books []Book
	books = append(books,
		Book{"Atomic Habits", "James Clear"},
		Book{"Steve Jobs", "Walter Isaacson"})

	user := User{
		Name:    "Muslim",
		Surname: "Yarashov",
		Age:     19,
		Books:   books,
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(userJson))
}
