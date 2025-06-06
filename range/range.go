package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Go Range Examples ===")

	// 1. Range over slices - getting index and value
	fmt.Println("1. Range over slice (index and value):")
	fruits := []string{"apple", "banana", "cherry", "date"}
	for i, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", i, fruit)
	}

	// 2. Range over slices - getting only values
	fmt.Println("\n2. Range over slice (values only):")
	for _, fruit := range fruits {
		fmt.Printf("Fruit: %s\n", fruit)
	}

	// 3. Range over slices - getting only indices
	fmt.Println("\n3. Range over slice (indices only):")
	for i := range fruits {
		fmt.Printf("Index: %d\n", i)
	}

	// 4. Range over arrays
	fmt.Println("\n4. Range over array:")
	numbers := [5]int{10, 20, 30, 40, 50}
	for i, num := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, num)
	}

	// 5. Range over maps
	fmt.Println("\n5. Range over map:")
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// 6. Range over map keys only
	fmt.Println("\n6. Range over map (keys only):")
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}

	// 7. Range over strings (runes)
	fmt.Println("\n7. Range over string (runes):")
	text := "Hello, 世界"
	for i, r := range text {
		fmt.Printf("Index %d: %c (Unicode: %U)\n", i, r, r)
	}

	// 8. Range over string (bytes vs runes)
	fmt.Println("\n8. String bytes vs runes:")
	str := "café"
	fmt.Printf("String: %s (length: %d bytes)\n", str, len(str))
	fmt.Println("Ranging over runes:")
	for i, r := range str {
		fmt.Printf("  Byte index %d: %c\n", i, r)
	}

	// 9. Range over channels
	fmt.Println("\n9. Range over channel:")
	ch := make(chan int, 5)
	// Send some values
	for i := 1; i <= 5; i++ {
		ch <- i * 10
	}
	close(ch) // Must close channel to end range loop

	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}

	// 10. Range with break and continue
	fmt.Println("\n10. Range with break and continue:")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Even numbers up to first number > 6:")
	for _, num := range nums {
		if num > 6 {
			break // Exit the loop
		}
		if num%2 != 0 {
			continue // Skip odd numbers
		}
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 11. Nested range loops
	fmt.Println("\n11. Nested range loops:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, val)
		}
	}

	// 12. Range over slice of structs
	fmt.Println("\n12. Range over slice of structs:")
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	for i, person := range people {
		fmt.Printf("Person %d: %s (%d years old)\n", i+1, person.Name, person.Age)
	}

	// 13. Modifying slice elements during range
	fmt.Println("\n13. Modifying slice elements:")
	scores := []int{85, 90, 78, 92, 88}
	fmt.Println("Original scores:", scores)

	// This doesn't modify the original slice
	fmt.Println("Trying to modify via range variable (doesn't work):")
	for _, score := range scores {
		score = score + 5 // This doesn't modify the original
	}
	fmt.Println("Scores after range modification:", scores)

	// This modifies the original slice
	fmt.Println("Modifying via index:")
	for i := range scores {
		scores[i] = scores[i] + 5
	}
	fmt.Println("Scores after index modification:", scores)

	// 14. Range over map with complex values
	fmt.Println("\n14. Range over map with slice values:")
	hobbies := map[string][]string{
		"Alice": {"reading", "swimming", "coding"},
		"Bob":   {"gaming", "football", "movies"},
	}
	for name, hobbyList := range hobbies {
		fmt.Printf("%s's hobbies: %s\n", name, strings.Join(hobbyList, ", "))
	}

	// 15. Range with function calls
	fmt.Println("\n15. Range with function results:")
	for i, word := range getWords() {
		fmt.Printf("Word %d: %s\n", i+1, word)
	}

	// 16. Range over empty collections
	fmt.Println("\n16. Range over empty collections:")
	var emptySlice []int
	var emptyMap map[string]int

	fmt.Println("Ranging over empty slice:")
	for i, v := range emptySlice {
		fmt.Printf("This won't print: %d, %d\n", i, v)
	}
	fmt.Println("Empty slice range completed")

	fmt.Println("Ranging over empty map:")
	for k, v := range emptyMap {
		fmt.Printf("This won't print: %s, %d\n", k, v)
	}
	fmt.Println("Empty map range completed")

	// 17. Range with labeled loops
	fmt.Println("\n17. Range with labeled loops (break outer):")
	matrix2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
outer:
	for i, row := range matrix2 {
		for j, val := range row {
			fmt.Printf("Checking matrix2[%d][%d] = %d\n", i, j, val)
			if val == 5 {
				fmt.Println("Found 5, breaking out of both loops")
				break outer
			}
		}
	}

	// 18. Range with pointers
	fmt.Println("\n18. Range with pointers:")
	ptrs := []*int{}
	vals := []int{10, 20, 30}
	for i := range vals {
		ptrs = append(ptrs, &vals[i])
	}
	for i, ptr := range ptrs {
		fmt.Printf("ptrs[%d] points to value: %d\n", i, *ptr)
	}

	// 19. Performance consideration - range vs traditional loop
	fmt.Println("\n19. Range vs traditional loop:")
	largeSlice := make([]int, 5)
	for i := range largeSlice {
		largeSlice[i] = i * i
	}

	fmt.Println("Using range:")
	for i, v := range largeSlice {
		fmt.Printf("largeSlice[%d] = %d\n", i, v)
	}

	fmt.Println("Using traditional loop:")
	for i := 0; i < len(largeSlice); i++ {
		fmt.Printf("largeSlice[%d] = %d\n", i, largeSlice[i])
	}

	// 20. Range with interface slices
	fmt.Println("\n20. Range with interface slice:")
	mixed := []interface{}{42, "hello", 3.14, true}
	for i, item := range mixed {
		fmt.Printf("Item %d: %v (Type: %T)\n", i, item, item)
	}
}

// Helper function for example 15
func getWords() []string {
	return []string{"Go", "is", "awesome", "for", "programming"}
}
