package main

import "fmt"

type Course struct{
	Name string
	Price int
	Platform string
	Password string
	Tags []string
}

func main() {
	fmt.Println("Json creation and consumption in go")
	courses := []Course{
		{"ReactJS Bootcamp", 299, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 199, "Udemy", "xyz456", []string{"web-dev", "js"}},
		{"Python Bootcamp", 399, "Coursera", "pqr789", []string{"web-dev", "python"}},
		{"JavaScript Bootcamp", 299, "Udemy", "jkl012", nil},

	}
	jsonString := EncodeJson(courses)
	fmt.Println(jsonString)
}

func EncodeJson(courses []Course) string {
	// code to encode the courses slice into JSON string
	return ""
}