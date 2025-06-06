package main

import "fmt"

func main() {
	const name = "John Doe"
	const age = 30

	const (
		port = 202
		host = "localhost"
	)

	const isEmployed = true

	fmt.Println("Name:", name)

	fmt.Println("Age:", age)
	fmt.Println("Port:", port)
	fmt.Println("Host:", host)
	fmt.Println("Is Employed:", isEmployed)
}
