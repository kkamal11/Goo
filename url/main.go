package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.google.com:8080/search?q=example&oq=example&aqs=chrome..69i57j69i60l3j69i65.1234j0j7&sourceid=chrome&ie=UTF-8"

func main() {
	fmt.Println("Handling Url demo")
	fmt.Println("Url:", myUrl)
	result, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Parsed URL:", result)
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Hostname: ", result.Hostname())
	fmt.Println("Port: ", result.Port())
	fmt.Println("Path: ", result.Path)
	fmt.Println("RawQuery: ", result.RawQuery)
	fmt.Println("Fragment: ", result.Fragment)

	qParams := result.Query()
	fmt.Println("Query Parameters: ", qParams)
	fmt.Printf("Type of qParams: %T\n", qParams)
	fmt.Println("Value of 'q' parameter: ", qParams["q"])
	fmt.Println("Value of 'sourceid' parameter: ", qParams["sourceid"])
	fmt.Println("Value of 'ie' parameter: ", qParams["ie"])
	fmt.Println("Value of 'nonexistent' parameter: ", qParams["nonexistent"])

	for key, value := range qParams {
		fmt.Println("Key:", key, "Value:", value)
	}
}