package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Int   int
	Float float64
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:  "foo bar",
		Int:   25,
		Float: 3.1416,
		Slice: []string{"a", "b", "c"},
		Map: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
