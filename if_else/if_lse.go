package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// SHow All Examples of if else statement

	// 1. Basic if-else
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// 2. if-else if-else chains
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// 3. if with initialization (short statement)
	if num := 9; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}
	// Note: num is not accessible here

	// 4. if with logical operators
	a, b := 5, 10
	if a > 0 && b > 0 {
		fmt.Println("Both a and b are positive")
	}

	if a > 10 || b > 5 {
		fmt.Println("Either a is greater than 10 or b is greater than 5")
	}

	// 5. Nested if statements
	num := 15
	if num > 10 {
		fmt.Println("num is greater than 10")
		if num%2 == 0 {
			fmt.Println("And it's even")
		} else {
			fmt.Println("And it's odd")
		}
	}

	// 6. If with error handling (common Go pattern)
	str := "42"
	if val, err := strconv.Atoi(str); err == nil {
		fmt.Printf("String converted to integer: %d\n", val)
	} else {
		fmt.Printf("Conversion error: %s\n", err)
	}

	// 7. Using if for conditional assignment
	hour := time.Now().Hour()
	greeting := ""
	if hour < 12 {
		greeting = "Good morning"
	} else if hour < 18 {
		greeting = "Good afternoon"
	} else {
		greeting = "Good evening"
	}
	fmt.Println(greeting)
}
