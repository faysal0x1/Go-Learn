package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 1. Basic Hello World
	fmt.Println("Hello, World!")

	// 2. Hello World with formatting
	fmt.Printf("Hello, %s!\n", "World")

	// 3. Hello World with variables
	greeting := "Hello"
	target := "World"
	fmt.Println(greeting + ", " + target + "!")

	// 4. Hello World with user input
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)

	// 5. Hello World with current time
	currentTime := time.Now()
	fmt.Printf("Hello, World! Current time is: %s\n", currentTime.Format("2006-01-02 15:04:05"))

	// 6. Hello World with command line arguments
	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s!\n", os.Args[1])
	} else {
		fmt.Println("Hello, World! (No name provided)")
	}

	// 7. Multiple Hello World messages
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		fmt.Printf("Hello, %s!\n", name)
	}

	// 8. Hello World in different languages
	greetings := map[string]string{
		"English": "Hello, World!",
		"Spanish": "¡Hola, Mundo!",
		"French":  "Bonjour, le Monde!",
		"German":  "Hallo, Welt!",
		"Italian": "Ciao, Mondo!",
	}

	for language, greeting := range greetings {
		fmt.Printf("%s: %s\n", language, greeting)
	}

	// 9. Hello World with error handling
	err := printHello("Go Developer")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 10. Hello World using a struct
	person := Person{Name: "Gopher", Language: "Go"}
	person.SayHello()

	// 11. Hello World with goroutines
	fmt.Println("\nAsync Hello World:")
	done := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello from goroutine!")
		done <- true
	}()

	fmt.Println("Hello from main!")
	<-done // Wait for goroutine to complete

	// 12. Hello World with function types
	sayHello := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	sayHello("Function Variable")

	// 13. Hello World with interfaces
	var greeter Greeter = EnglishGreeter{}
	greeter.Greet("Interface User")

	// 14. Hello World with constants
	const DefaultGreeting = "Hello"
	const DefaultTarget = "World"
	fmt.Printf("%s, %s!\n", DefaultGreeting, DefaultTarget)

	// 15. Hello World with type conversion
	var count int = 42
	fmt.Printf("Hello, World! This is message number %d\n", count)
	fmt.Printf("Hello, World! Count as string: %s\n", fmt.Sprintf("%d", count))
}

// Function for error handling example
func printHello(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	fmt.Printf("Hello, %s!\n", name)
	return nil
}

// Struct for Hello World
type Person struct {
	Name     string
	Language string
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I love %s!\n", p.Name, p.Language)
}

// Interface example
type Greeter interface {
	Greet(name string)
}

type EnglishGreeter struct{}

func (e EnglishGreeter) Greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

type SpanishGreeter struct{}

func (s SpanishGreeter) Greet(name string) {
	fmt.Printf("¡Hola, %s!\n", name)
}
