package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Basic variable declaration with var keyword
	var age int
	age = 30
	fmt.Println("Age:", age)

	// 2. Variable declaration with initialization
	var name string = "John Doe"
	fmt.Println("Name:", name)

	// 3. Type inference - compiler determines type
	var salary = 50000.50
	fmt.Printf("Salary: %.2f (Type: %T)\n", salary, salary)

	// 4. Short variable declaration (only inside functions)
	count := 10
	fmt.Println("Count:", count)

	// 5. Multiple variable declaration
	var a, b, c int = 1, 2, 3
	fmt.Println("Multiple values:", a, b, c)

	// 6. Multiple types in one declaration
	var (
		active bool   = true
		email  string = "john@example.com"
		age2   int    = 25
	)
	fmt.Println("Active:", active)
	fmt.Println("Email:", email)
	fmt.Println("Age2:", age2)

	// 7. Short declaration for multiple variables
	x, y := 5, "hello"
	fmt.Printf("x: %v (Type: %T), y: %v (Type: %T)\n", x, x, y, y)

	// 8. Constants
	const Pi = 3.14159
	const (
		StatusOK       = 200
		StatusNotFound = 404
	)
	fmt.Println("Pi:", Pi)
	fmt.Println("Status codes:", StatusOK, StatusNotFound)

	// 9. Zero values (default values)
	var i int
	var f float64
	var b_ool bool
	var s string
	fmt.Printf("Zero values: %v, %v, %v, %q\n", i, f, b_ool, s)

	// 10. Type conversion
	var i1 int = 42
	var f1 float64 = float64(i1)
	var u uint = uint(f1)
	fmt.Println("Type conversion:", i1, f1, u)

	// 11. Numeric constants
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)
	fmt.Println("Small constant:", Small)

	// 12. Shadowing variables
	x = 10
	{
		x := 20 // This shadows the outer x
		fmt.Println("Inner x:", x)
	}
	fmt.Println("Outer x:", x)

	// 13. Unused variables - Go will not compile with unused variables
	// Uncomment to see error:
	// var unused int

	// 14. Using math package with variables
	var radius = 5.0
	area := math.Pi * radius * radius
	fmt.Printf("Circle area with radius %.1f: %.2f\n", radius, area)
}
