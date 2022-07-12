package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Karthikeya",
		Age:  6,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	user = User{
		Name: "Ishaan",
		Age:  3,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	fmt.Println()
}
