package main

import (
	"encoding/json"
	"fmt"
)

type Course struct{
	Name string `json:"coursename"`
	Price int `json:"price"`
	Platform string `json:"platform"`
	Password string `json:"-"` //means dont send in json
	Tags []string `json:"tags,omitempty"` //if tags is empty then dont send in json
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
	fmt.Println("==================================")
	isJsonValid := json.Valid([]byte(jsonString))
	if isJsonValid {
		fmt.Println("JSON is valid")
	} else {
		fmt.Println("JSON is not valid")
		panic("invalidJson")
	}
	originalCourses, originalJsonMap := DecodeJson(jsonString)
	fmt.Println(originalCourses)
	fmt.Println(originalJsonMap)
}

func EncodeJson(courses []Course) string {
	// code to encode the courses slice into JSON string
	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "  ") //for indentation
	if err != nil {
		panic(err)
	}
	
	return string(finalJson)
}


func DecodeJson(jsonData string) ([]Course, []map[string]interface{}) {
	// code to decode the JSON string into courses slice
	var courses []Course
	var err error
	// err = json.Unmarshal([]byte(jsonData), &courses)
	// if err != nil {
	// 	panic(err)
	// }

	var myJsonMap []map[string]any
	err = json.Unmarshal([]byte(jsonData), &myJsonMap)
	if err != nil {
		panic(err)
	}
	
	return courses, myJsonMap
}