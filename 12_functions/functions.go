package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("=== Go Functions Examples ===\n")

	// 1. Basic function call
	fmt.Println("1. Basic function:")
	greet()

	// 2. Function with parameters
	fmt.Println("\n2. Function with parameters:")
	greetPerson("Alice")

	// 3. Function with return value
	fmt.Println("\n3. Function with return value:")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// 4. Function with multiple parameters and return values
	fmt.Println("\n4. Multiple parameters and return values:")
	sum, product := calculate(4, 6)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// 5. Named return values
	fmt.Println("\n5. Named return values:")
	area, perimeter := rectangleStats(5, 3)
	fmt.Printf("Rectangle - Area: %d, Perimeter: %d\n", area, perimeter)

	// 6. Variadic functions
	fmt.Println("\n6. Variadic functions:")
	fmt.Printf("Sum of 1,2,3: %d\n", sumAll(1, 2, 3))
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", sumAll(1, 2, 3, 4, 5))

	numbers := []int{10, 20, 30}
	fmt.Printf("Sum of slice: %d\n", sumAll(numbers...))

	// 7. Functions as variables
	fmt.Println("\n7. Functions as variables:")
	var operation func(int, int) int
	operation = add
	fmt.Printf("Using function variable: %d\n", operation(7, 8))

	// 8. Anonymous functions
	fmt.Println("\n8. Anonymous functions:")
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("Anonymous function result: %d\n", multiply(4, 5))

	// 9. Higher-order functions
	fmt.Println("\n9. Higher-order functions:")
	nums := []int{1, 2, 3, 4, 5}
	doubled := mapInts(nums, func(x int) int { return x * 2 })
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("Doubled: %v\n", doubled)

	// 10. Function returning function
	fmt.Println("\n10. Function returning function:")
	adder := makeAdder(10)
	fmt.Printf("Adder(5): %d\n", adder(5))
	fmt.Printf("Adder(3): %d\n", adder(3))

	// 11. Closures
	fmt.Println("\n11. Closures:")
	counter := makeCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())

	// 12. Error handling
	fmt.Println("\n12. Error handling:")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 13. Defer statements
	fmt.Println("\n13. Defer statements:")
	deferExample()

	// 14. Panic and recover
	fmt.Println("\n14. Panic and recover:")
	recoverExample()

	// 15. Methods (functions with receivers)
	fmt.Println("\n15. Methods:")
	p := Person{Name: "Bob", Age: 30}
	p.introduce()
	p.haveBirthday()
	p.introduce()

	// 16. Pointer receivers
	fmt.Println("\n16. Pointer receivers:")
	c := Circle{Radius: 5}
	fmt.Printf("Circle area: %.2f\n", c.area())
	c.scale(2)
	fmt.Printf("After scaling, area: %.2f\n", c.area())

	// 17. Interface implementation
	fmt.Println("\n17. Interface implementation:")
	var shape Shape = Circle{Radius: 3}
	fmt.Printf("Shape area: %.2f\n", shape.area())

	rect := Rectangle{Width: 4, Height: 6}
	shape = rect
	fmt.Printf("Shape area: %.2f\n", shape.area())

	// 18. Function composition
	fmt.Println("\n18. Function composition:")
	addOne := func(x int) int { return x + 1 }
	multiplyByTwo := func(x int) int { return x * 2 }

	composed := compose(multiplyByTwo, addOne)
	fmt.Printf("Composed function (2 * (5 + 1)): %d\n", composed(5))

	// 19. Recursive functions
	fmt.Println("\n19. Recursive functions:")
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Fibonacci of 8: %d\n", fibonacci(8))

	// 20. Function with multiple error types
	fmt.Println("\n20. Custom error types:")
	_, err = validateAge(-5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	age, err := validateAge(25)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Valid age: %d\n", age)
	}

	// 21. Function options pattern
	fmt.Println("\n21. Function options pattern:")
	server := NewServer(
		WithPort(8080),
		WithHost("localhost"),
		WithTimeout(30),
	)
	fmt.Printf("Server config: %+v\n", server)

	// 22. Generic functions (Go 1.18+)
	fmt.Println("\n22. Generic functions:")
	fmt.Printf("Max of 5 and 3: %d\n", max(5, 3))
	fmt.Printf("Max of 2.5 and 1.8: %.1f\n", max(2.5, 1.8))
	fmt.Printf("Max of 'hello' and 'world': %s\n", max("hello", "world"))

	// 23. Functional programming patterns
	fmt.Println("\n23. Functional programming:")
	numbers2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(numbers2, func(x int) bool { return x%2 == 0 })
	squares := mapInts(evens, func(x int) int { return x * x })
	sum := reduce(squares, 0, func(acc, x int) int { return acc + x })

	fmt.Printf("Original: %v\n", numbers2)
	fmt.Printf("Evens: %v\n", evens)
	fmt.Printf("Squares: %v\n", squares)
	fmt.Printf("Sum: %d\n", sum)

	// 24. Benchmark-style function
	fmt.Println("\n24. Measuring function performance:")
	result2 := measureTime("Fibonacci calculation", func() int {
		return fibonacci(20)
	})
	fmt.Printf("Result: %d\n", result2)
}

// Basic functions

// 1. Function with no parameters or return value
func greet() {
	fmt.Println("Hello, World!")
}

// 2. Function with parameter
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 3. Function with return value
func add(a, b int) int {
	return a + b
}

// 4. Function with multiple return values
func calculate(a, b int) (int, int) {
	return a + b, a * b
}

// 5. Named return values
func rectangleStats(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

// 6. Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 9. Higher-order function
func mapInts(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 10. Function returning function
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// 11. Closure example
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 12. Error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 13. Defer example
func deferExample() {
	fmt.Println("Start of function")
	defer fmt.Println("This will be printed last")
	defer fmt.Println("This will be printed second to last")
	fmt.Println("End of function")
}

// 14. Panic and recover
func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("Something went wrong!")
	fmt.Println("This will never be printed")
}

// 15. Methods and structs
type Person struct {
	Name string
	Age  int
}

func (p Person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

func (p *Person) haveBirthday() {
	p.Age++
	fmt.Printf("%s just turned %d!\n", p.Name, p.Age)
}

// 16. Methods with pointer receivers
type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) scale(factor float64) {
	c.Radius *= factor
}

// 17. Interfaces
type Shape interface {
	area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

// 18. Function composition
func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// 19. Recursive functions
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 20. Custom error types
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}

func validateAge(age int) (int, error) {
	if age < 0 {
		return 0, ValidationError{
			Field:   "age",
			Message: "age cannot be negative",
		}
	}
	if age > 150 {
		return 0, ValidationError{
			Field:   "age",
			Message: "age cannot be greater than 150",
		}
	}
	return age, nil
}

// 21. Function options pattern
type Server struct {
	Host    string
	Port    int
	Timeout int
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout int) ServerOption {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{
		Host:    "127.0.0.1",
		Port:    8000,
		Timeout: 10,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

// 22. Generic functions
func max[T comparable](a, b T) T {
	if strings.Compare(fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)) > 0 {
		return a
	}
	return b
}

// 23. Functional programming helpers
func filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func reduce(slice []int, initial int, reducer func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// 24. Performance measurement
func measureTime[T any](name string, fn func() T) T {
	fmt.Printf("Starting %s...\n", name)
	start := len(fmt.Sprintf("%d", 1)) // Simple time simulation
	result := fn()
	fmt.Printf("Completed %s\n", name)
	return result
}
