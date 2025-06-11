package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("=== Go Variadic Functions Examples ===\n")

	// 1. Basic variadic function
	fmt.Println("1. Basic variadic function:")
	fmt.Printf("Sum of no numbers: %d\n", sum())
	fmt.Printf("Sum of 1: %d\n", sum(1))
	fmt.Printf("Sum of 1,2,3: %d\n", sum(1, 2, 3))
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", sum(1, 2, 3, 4, 5))

	// 2. Passing slice to variadic function
	fmt.Println("\n2. Passing slice to variadic function:")
	numbers := []int{10, 20, 30, 40}
	fmt.Printf("Sum of slice %v: %d\n", numbers, sum(numbers...))

	// 3. Variadic with other parameters
	fmt.Println("\n3. Variadic with other parameters:")
	result1 := multiply(2, 1, 2, 3, 4)
	fmt.Printf("2 * (1,2,3,4) = %d\n", result1)

	// 4. Variadic string functions
	fmt.Println("\n4. Variadic string functions:")
	message1 := joinStrings("-", "Hello", "World", "Go")
	fmt.Printf("Joined: %s\n", message1)

	message2 := joinStrings(" | ")
	fmt.Printf("Empty join: '%s'\n", message2)

	// 5. Finding min and max
	fmt.Println("\n5. Finding min and max:")
	fmt.Printf("Min of 5,2,8,1,9: %d\n", min(5, 2, 8, 1, 9))
	fmt.Printf("Max of 5,2,8,1,9: %d\n", max(5, 2, 8, 1, 9))

	// 6. Variadic with different types
	fmt.Println("\n6. Variadic with different types:")
	printValues("Numbers:", 1, 2, 3, 4, 5)
	printValues("Mixed:", "hello", 42, 3.14, true)

	// 7. Collecting arguments into slice
	fmt.Println("\n7. Collecting and processing arguments:")
	processNumbers("Squares", 1, 2, 3, 4, 5)
	processNumbers("Cubes", 2, 3, 4)

	// 8. Variadic function with error handling
	fmt.Println("\n8. Variadic with error handling:")
	avg1, err := average(10, 20, 30, 40)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Average of 10,20,30,40: %.2f\n", avg1)
	}

	_, err = average()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 9. Variadic function returning slice
	fmt.Println("\n9. Variadic returning slice:")
	doubled := doubleAll(1, 2, 3, 4, 5)
	fmt.Printf("Doubled values: %v\n", doubled)

	// 10. Printf-style variadic
	fmt.Println("\n10. Printf-style formatting:")
	formatAndPrint("User %s is %d years old and has %.2f in account", "Alice", 25, 1234.56)
	formatAndPrint("Today is %s", "Monday")

	// 11. Variadic with structs
	fmt.Println("\n11. Variadic with structs:")
	people := createPeople(
		Person{"Alice", 25},
		Person{"Bob", 30},
		Person{"Charlie", 35},
	)
	fmt.Printf("Created %d people\n", len(people))
	for _, p := range people {
		fmt.Printf("- %s (%d)\n", p.Name, p.Age)
	}

	// 12. Nested variadic calls
	fmt.Println("\n12. Nested variadic calls:")
	total := sumOfSums(
		[]int{1, 2, 3},
		[]int{4, 5},
		[]int{6, 7, 8, 9},
	)
	fmt.Printf("Sum of all sums: %d\n", total)

	// 13. Variadic with channels
	fmt.Println("\n13. Variadic with channels:")
	ch := make(chan int, 10)
	sendToChannel(ch, 1, 2, 3, 4, 5)
	close(ch)

	fmt.Print("Received from channel: ")
	for value := range ch {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// 14. Variadic with maps
	fmt.Println("\n14. Variadic with maps:")
	merged := mergeMaps(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"c": 3, "d": 4},
		map[string]int{"e": 5},
	)
	fmt.Printf("Merged map: %v\n", merged)

	// 15. Statistics functions
	fmt.Println("\n15. Statistics functions:")
	data := []float64{2.5, 3.7, 1.2, 4.8, 2.1, 3.3, 2.9}
	stats := calculateStats(data...)
	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Stats: %+v\n", stats)

	// 16. Variadic validation
	fmt.Println("\n16. Variadic validation:")
	validEmails := validateEmails("test@example.com", "invalid-email", "user@domain.org")
	fmt.Printf("Valid emails: %v\n", validEmails)

	// 17. Building SQL-like queries
	fmt.Println("\n17. SQL-like query building:")
	query1 := buildQuery("users", "name", "email", "age")
	fmt.Printf("Query 1: %s\n", query1)

	query2 := buildQuery("products", "id", "name", "price", "category")
	fmt.Printf("Query 2: %s\n", query2)

	// 18. Variadic with function parameters
	fmt.Println("\n18. Variadic with function parameters:")
	operations := []func(int) int{
		func(x int) int { return x * 2 },
		func(x int) int { return x + 10 },
		func(x int) int { return x * x },
	}
	result := applyOperations(5, operations...)
	fmt.Printf("Applied operations to 5: %d\n", result)

	// 19. Collecting unique values
	fmt.Println("\n19. Collecting unique values:")
	unique := getUnique(1, 2, 3, 2, 4, 3, 5, 1, 6)
	fmt.Printf("Unique values: %v\n", unique)

	// 20. Variadic with interfaces
	fmt.Println("\n20. Variadic with interfaces:")
	shapes := createShapes(
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 3},
	)

	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
}

// 1. Basic variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 3. Variadic with other parameters (variadic must be last)
func multiply(multiplier int, numbers ...int) int {
	result := 1
	for _, num := range numbers {
		result *= num
	}
	return result * multiplier
}

// 4. Variadic string function
func joinStrings(separator string, strings ...string) string {
	return strings.Join(strings, separator)
}

// 5. Min and Max functions
func min(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	minimum := numbers[0]
	for _, num := range numbers[1:] {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

func max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	maximum := numbers[0]
	for _, num := range numbers[1:] {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

// 6. Variadic with interface{} (any type)
func printValues(label string, values ...interface{}) {
	fmt.Printf("%s ", label)
	for i, value := range values {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%v", value)
	}
	fmt.Println()
}

// 7. Processing arguments
func processNumbers(operation string, numbers ...int) {
	fmt.Printf("%s: ", operation)
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(", ")
		}
		switch operation {
		case "Squares":
			fmt.Printf("%d", num*num)
		case "Cubes":
			fmt.Printf("%d", num*num*num)
		default:
			fmt.Printf("%d", num)
		}
	}
	fmt.Println()
}

// 8. Variadic with error handling
func average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("cannot calculate average of empty list")
	}

	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers)), nil
}

// 9. Variadic returning slice
func doubleAll(numbers ...int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num * 2
	}
	return result
}

// 10. Printf-style formatting
func formatAndPrint(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	fmt.Printf("Formatted: %s\n", formatted)
}

// 11. Variadic with structs
type Person struct {
	Name string
	Age  int
}

func createPeople(people ...Person) []Person {
	return people
}

// 12. Nested variadic calls
func sumOfSums(slices ...[]int) int {
	total := 0
	for _, slice := range slices {
		total += sum(slice...)
	}
	return total
}

// 13. Variadic with channels
func sendToChannel(ch chan<- int, values ...int) {
	for _, value := range values {
		ch <- value
	}
}

// 14. Variadic with maps
func mergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// 15. Statistics
type Stats struct {
	Count   int
	Sum     float64
	Average float64
	Min     float64
	Max     float64
}

func calculateStats(values ...float64) Stats {
	if len(values) == 0 {
		return Stats{}
	}

	stats := Stats{
		Count: len(values),
		Min:   values[0],
		Max:   values[0],
	}

	for _, value := range values {
		stats.Sum += value
		if value < stats.Min {
			stats.Min = value
		}
		if value > stats.Max {
			stats.Max = value
		}
	}

	stats.Average = stats.Sum / float64(stats.Count)
	return stats
}

// 16. Validation function
func validateEmails(emails ...string) []string {
	var validEmails []string
	for _, email := range emails {
		if strings.Contains(email, "@") && strings.Contains(email, ".") {
			validEmails = append(validEmails, email)
		}
	}
	return validEmails
}

// 17. SQL-like query builder
func buildQuery(table string, columns ...string) string {
	if len(columns) == 0 {
		return fmt.Sprintf("SELECT * FROM %s", table)
	}
	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), table)
}

// 18. Variadic with function parameters
func applyOperations(value int, operations ...func(int) int) int {
	result := value
	for _, operation := range operations {
		result = operation(result)
	}
	return result
}

// 19. Unique values
func getUnique(values ...int) []int {
	seen := make(map[int]bool)
	var unique []int

	for _, value := range values {
		if !seen[value] {
			seen[value] = true
			unique = append(unique, value)
		}
	}
	return unique
}

// 20. Variadic with interfaces
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func createShapes(shapes ...Shape) []Shape {
	return shapes
}
