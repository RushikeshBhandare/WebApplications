package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

type DogData struct {
	Name string
	Age  int
}

type Friends struct {
	Name []string
}

type data struct {
	User
	DogData
	Friends
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	AllData := data{
		User{
			Name: "Rushikesh Bhandare",
			Age:  21,
		},
		DogData{
			Name: "moti",
			Age:  3,
		},
		Friends{
			Name: []string{"Ram", "Sham", "Guru", "Govid"},
		},
	}
	err = t.Execute(os.Stdout, AllData)
	if err != nil {
		panic(err)
	}

}
