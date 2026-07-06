package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL string = "https://httpbin.org/robots.txt"

func main() {
	fmt.Println("Web Request Demo")
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Response:", response.Status)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		panic(err)
	}
	fmt.Println("Response Body:", string(body))
}