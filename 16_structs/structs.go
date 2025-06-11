package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("=== Go Structs Examples ===\n")

	// 1. Basic struct declaration and usage
	fmt.Println("1. Basic struct:")
	var person Person
	fmt.Printf("Zero value struct: %+v\n", person)

	// Initialize struct
	person = Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Printf("Initialized struct: %+v\n", person)

	// 2. Struct literal initialization
	fmt.Println("\n2. Struct literal initialization:")
	// Named fields
	person1 := Person{
		Name: "Bob",
		Age:  25,
		City: "San Francisco",
	}

	// Positional fields (order matters)
	person2 := Person{"Charlie", 35, "Chicago"}

	fmt.Printf("Named fields: %+v\n", person1)
	fmt.Printf("Positional fields: %+v\n", person2)

	// 3. Anonymous structs
	fmt.Println("\n3. Anonymous structs:")
	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  true,
	}
	fmt.Printf("Anonymous struct: %+v\n", config)

	// 4. Accessing struct fields
	fmt.Println("\n4. Accessing struct fields:")
	fmt.Printf("Name: %s\n", person1.Name)
	fmt.Printf("Age: %d\n", person1.Age)

	// Modifying fields
	person1.Age = 26
	fmt.Printf("Updated age: %d\n", person1.Age)

	// 5. Struct pointers
	fmt.Println("\n5. Struct pointers:")
	personPtr := &person1
	fmt.Printf("Via pointer: %s\n", personPtr.Name)

	// Automatic dereferencing
	personPtr.Age = 27
	fmt.Printf("Updated via pointer: %d\n", person1.Age)

	// 6. Nested structs
	fmt.Println("\n6. Nested structs:")
	emp := Employee{
		Person: Person{
			Name: "David",
			Age:  40,
			City: "Boston",
		},
		ID:       1001,
		Salary:   75000,
		Position: "Engineer",
	}
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Employee name: %s\n", emp.Person.Name)

	// 7. Embedded structs (composition)
	fmt.Println("\n7. Embedded structs:")
	manager := Manager{
		Person: Person{
			Name: "Eve",
			Age:  45,
			City: "Seattle",
		},
		Department: "Engineering",
		TeamSize:   10,
	}

	// Can access embedded fields directly
	fmt.Printf("Manager name: %s\n", manager.Name) // Same as manager.Person.Name
	fmt.Printf("Manager department: %s\n", manager.Department)

	// 8. Methods on structs
	fmt.Println("\n8. Methods on structs:")
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// 9. Pointer receivers vs value receivers
	fmt.Println("\n9. Pointer vs value receivers:")
	counter := Counter{Value: 0}
	fmt.Printf("Initial counter: %d\n", counter.Value)

	counter.IncrementValue() // Value receiver - doesn't modify original
	fmt.Printf("After IncrementValue: %d\n", counter.Value)

	counter.IncrementPointer() // Pointer receiver - modifies original
	fmt.Printf("After IncrementPointer: %d\n", counter.Value)

	// 10. Struct comparison
	fmt.Println("\n10. Struct comparison:")
	p1 := Person{"Alice", 30, "NYC"}
	p2 := Person{"Alice", 30, "NYC"}
	p3 := Person{"Bob", 25, "LA"}

	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n", p1 == p3)

	// 11. Struct copying
	fmt.Println("\n11. Struct copying:")
	original := Person{"Original", 20, "City"}
	copy := original // Creates a copy
	copy.Name = "Copy"

	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copy: %+v\n", copy)

	// 12. Struct tags
	fmt.Println("\n12. Struct tags:")
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	// JSON serialization using tags
	jsonData, _ := json.Marshal(user)
	fmt.Printf("JSON: %s\n", jsonData)

	// Reflection to read tags
	userType := reflect.TypeOf(user)
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			fmt.Printf("Field: %s, JSON tag: %s\n", field.Name, jsonTag)
		}
	}

	// 13. Constructor functions
	fmt.Println("\n13. Constructor functions:")
	newPerson := NewPerson("Alice", 25)
	fmt.Printf("New person: %+v\n", newPerson)

	validatedPerson, err := NewValidatedPerson("", -5)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	validatedPerson, err = NewValidatedPerson("Bob", 30)
	if err == nil {
		fmt.Printf("Valid person: %+v\n", validatedPerson)
	}

	// 14. Struct interfaces
	fmt.Println("\n14. Struct interfaces:")
	var shapes []Shape
	shapes = append(shapes, Rectangle{Width: 4, Height: 5})
	shapes = append(shapes, Circle{Radius: 3})

	for i, shape := range shapes {
		fmt.Printf("Shape %d area: %.2f\n", i+1, shape.Area())
	}

	// 15. Empty struct
	fmt.Println("\n15. Empty struct:")
	// var empty struct{}
	// fmt.Printf("Empty struct size: %d bytes\n", unsafe.Sizeof(empty))

	// Used for signaling
	done := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Signal received via empty struct")

	// 16. Struct field promotion
	fmt.Println("\n16. Field promotion:")
	student := Student{
		Person: Person{Name: "Alice", Age: 20, City: "Boston"},
		School: "MIT",
		GPA:    3.8,
	}

	// Can call promoted methods directly
	student.Introduce() // Method from embedded Person
	fmt.Printf("Student at: %s\n", student.School)

	// 17. Struct visibility and packages
	fmt.Println("\n17. Visibility:")
	// Exported (public) fields start with capital letter
	// unexported (private) fields start with lowercase letter
	visibility := VisibilityExample{
		PublicField:  "Everyone can see this",
		privateField: "Only same package can see this",
	}
	fmt.Printf("Public: %s\n", visibility.PublicField)
	// fmt.Printf("Private: %s\n", visibility.privateField) // Would work in same package

	// 18. Struct with function fields
	fmt.Println("\n18. Struct with function fields:")
	calculator := Calculator{
		Add: func(a, b float64) float64 { return a + b },
		Sub: func(a, b float64) float64 { return a - b },
		Mul: func(a, b float64) float64 { return a * b },
		Div: func(a, b float64) float64 {
			if b != 0 {
				return a / b
			}
			return 0
		},
	}

	fmt.Printf("5 + 3 = %.2f\n", calculator.Add(5, 3))
	fmt.Printf("5 * 3 = %.2f\n", calculator.Mul(5, 3))

	// 19. Struct slices and maps
	fmt.Println("\n19. Struct collections:")
	people := []Person{
		{"Alice", 25, "NYC"},
		{"Bob", 30, "LA"},
		{"Charlie", 35, "Chicago"},
	}

	fmt.Println("People slice:")
	for i, p := range people {
		fmt.Printf("%d: %s (%d) from %s\n", i+1, p.Name, p.Age, p.City)
	}

	// Map with struct values
	cityPopulation := map[string]CityInfo{
		"NYC": {Population: 8000000, Country: "USA"},
		"LA":  {Population: 4000000, Country: "USA"},
	}

	for city, info := range cityPopulation {
		fmt.Printf("%s: %d people in %s\n", city, info.Population, info.Country)
	}

	// 20. Advanced struct patterns
	fmt.Println("\n20. Advanced patterns:")

	// Builder pattern
	builder := NewPersonBuilder()
	builtPerson := builder.
		WithName("Builder").
		WithAge(40).
		WithCity("Portland").
		Build()
	fmt.Printf("Built person: %+v\n", builtPerson)

	// Options pattern
	server := NewServer(
		WithPort(8080),
		WithHost("localhost"),
		WithTimeout(30),
	)
	fmt.Printf("Server config: %+v\n", server)

	// State machine using structs
	fsm := &StateMachine{State: "idle"}
	fsm.Transition("start")
	fsm.Transition("process")
	fsm.Transition("complete")
}

// Type definitions

// 1. Basic struct
type Person struct {
	Name string
	Age  int
	City string
}

// 6. Nested struct
type Employee struct {
	Person   Person // Nested struct
	ID       int
	Salary   float64
	Position string
}

// 7. Embedded struct
type Manager struct {
	Person     // Embedded (anonymous field)
	Department string
	TeamSize   int
}

// 8. Methods on structs
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 9. Pointer vs value receivers
type Counter struct {
	Value int
}

func (c Counter) IncrementValue() {
	c.Value++ // Doesn't modify original
}

func (c *Counter) IncrementPointer() {
	c.Value++ // Modifies original
}

// 12. Struct with tags
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age,omitempty"`
}

// 13. Constructor functions
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
		City: "Unknown",
	}
}

func NewValidatedPerson(name string, age int) (Person, error) {
	if name == "" {
		return Person{}, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return Person{}, fmt.Errorf("age cannot be negative")
	}
	return Person{Name: name, Age: age}, nil
}

// 14. Interfaces
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// 16. Field promotion
type Student struct {
	Person // Embedded
	School string
	GPA    float64
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s, %d years old from %s\n", p.Name, p.Age, p.City)
}

// 17. Visibility
type VisibilityExample struct {
	PublicField  string // Exported
	privateField string // Unexported
}

// 18. Function fields
type Calculator struct {
	Add func(float64, float64) float64
	Sub func(float64, float64) float64
	Mul func(float64, float64) float64
	Div func(float64, float64) float64
}

// 19. Collections
type CityInfo struct {
	Population int
	Country    string
}

// 20. Advanced patterns

// Builder pattern
type PersonBuilder struct {
	person Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func (pb *PersonBuilder) WithName(name string) *PersonBuilder {
	pb.person.Name = name
	return pb
}

func (pb *PersonBuilder) WithAge(age int) *PersonBuilder {
	pb.person.Age = age
	return pb
}

func (pb *PersonBuilder) WithCity(city string) *PersonBuilder {
	pb.person.City = city
	return pb
}

func (pb *PersonBuilder) Build() Person {
	return pb.person
}

// Options pattern
type ServerConfig struct {
	Host    string
	Port    int
	Timeout int
}

type ServerOption func(*ServerConfig)

func WithPort(port int) ServerOption {
	return func(config *ServerConfig) {
		config.Port = port
	}
}

func WithHost(host string) ServerOption {
	return func(config *ServerConfig) {
		config.Host = host
	}
}

func WithTimeout(timeout int) ServerOption {
	return func(config *ServerConfig) {
		config.Timeout = timeout
	}
}

func NewServer(options ...ServerOption) ServerConfig {
	config := ServerConfig{
		Host:    "127.0.0.1",
		Port:    8000,
		Timeout: 10,
	}

	for _, option := range options {
		option(&config)
	}

	return config
}

// State machine
type StateMachine struct {
	State string
}

func (sm *StateMachine) Transition(event string) {
	fmt.Printf("State: %s, Event: %s -> ", sm.State, event)

	switch sm.State {
	case "idle":
		if event == "start" {
			sm.State = "running"
		}
	case "running":
		if event == "process" {
			sm.State = "processing"
		} else if event == "complete" {
			sm.State = "idle"
		}
	case "processing":
		if event == "complete" {
			sm.State = "idle"
		}
	}

	fmt.Printf("New state: %s\n", sm.State)
}
