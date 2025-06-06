package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go For Loop Examples")
	fmt.Println("===================")

	// 1. Basic for loop with initialization, condition, and increment
	fmt.Println("\n1. Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// 2. For loop as a while loop (only condition)
	fmt.Println("\n2. For loop as a while loop:")
	j := 0
	for j < 5 {
		fmt.Printf("j = %d\n", j)
		j++
	}

	// 3. Infinite loop with break
	fmt.Println("\n3. Infinite loop with break:")
	counter := 0
	for {
		fmt.Printf("counter = %d\n", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	// 4. For loop with range
	fmt.Println("\n4. For loop with range:")

	// a. Iterating over an array
	fmt.Println("a. Iterating over an array:")
	numbers := [5]int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}

	// b. Iterating over a slice
	fmt.Println("\nb. Iterating over a slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("index = %d, fruit = %s\n", index, fruit)
	}

	// c. Iterating over a map
	fmt.Println("\nc. Iterating over a map:")
	capitals := map[string]string{
		"USA":    "Washington DC",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	for country, capital := range capitals {
		fmt.Printf("Capital of %s is %s\n", country, capital)
	}

	// d. Iterating over a string (by rune)
	fmt.Println("\nd. Iterating over a string (by rune):")
	message := "Hello, 世界"
	for index, char := range message {
		fmt.Printf("index = %d, char = %c, unicode = %U\n", index, char, char)
	}

	// 5. For loop with break and continue
	fmt.Println("\n5. For loop with break and continue:")

	// a. Using break
	fmt.Println("a. Using break:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at i =", i)
			break
		}
		fmt.Printf("i = %d\n", i)
	}

	// b. Using continue
	fmt.Println("\nb. Using continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Skipping i =", i)
			continue
		}
		fmt.Printf("i = %d\n", i)
	}

	// 6. Nested for loops
	fmt.Println("\n6. Nested for loops:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	// 7. For loop with labels for breaking out of nested loops
	fmt.Println("\n7. For loop with labels:")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("Breaking out of both loops when i =", i, "and j =", j)
				break OuterLoop
			}
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	// 8. For loop with multiple variables
	fmt.Println("\n8. For loop with multiple variables:")
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}

	// 9. For loop with blank identifier (_)
	fmt.Println("\n9. For loop with blank identifier:")

	// a. Ignoring the index
	fmt.Println("a. Ignoring the index:")
	colors := []string{"red", "green", "blue"}
	for _, color := range colors {
		fmt.Println(color)
	}

	// b. Ignoring the value
	fmt.Println("\nb. Ignoring the value:")
	for i, _ := range colors {
		fmt.Printf("Index %d\n", i)
	}

	// 10. For-each style loop
	fmt.Println("\n10. For-each style loop:")
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for _, day := range weekdays {
		fmt.Println(day)
	}

	// 11. Using range with channel
	fmt.Println("\n11. Using range with channel (simulated):")
	// Usually with channels, we'd do something like:
	// for item := range channel {
	//     fmt.Println(item)
	// }

	// Simulating it with a slice:
	ch := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // Important to close the channel for range to work
	}()

	// Range over channel
	for num := range ch {
		fmt.Println("Received:", num)
	}

	// 12. For loop to implement custom repetition
	fmt.Println("\n12. Custom repetition:")
	repeatString("Go", 5)
}

// Helper function to demonstrate custom repetition
func repeatString(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(s, " ")
	}
	fmt.Println()
}
