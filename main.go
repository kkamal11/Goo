package main

import (
	"fmt"
	"time"
)


func add(a int, b int) int{
	return a + b
}

func demo(){
	fmt.Println("Hello, World!")
	var age int
	var name string
	var salary float64
	var isAdmin bool

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(salary)
	fmt.Println(isAdmin)



	    var x int
    y := 10
    const z = 20

    x = y

	fmt.Println("---------")
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Println(x,y,z)

	var str string = "Hello"
	fmt.Println((str[0]))

	fmt.Println(add(2,3))

	fmt.Println("-----------")
	var ptr *string = &str
	ptr2 := &str
	fmt.Println(ptr, ptr2)
}

func run(){
	time.Sleep(2 * time.Second)
	fmt.Println("Run")
}

func run2(){
	time.Sleep(5 * time.Second)
	fmt.Println("Run2")
}

func goRoutine(){
	start := time.Now()
	go run()
	go run2()
	end := time.Now()
	fmt.Println("Time Taken ", (end.Second() - start.Second()), "seconds")
}

func main() {
	// demo()
	goRoutine()
}