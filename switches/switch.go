package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go Switch Statement Examples")
	fmt.Println("===========================")

	// 1. Basic switch statement
	fmt.Println("\n1. Basic switch statement:")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// 2. Switch with multiple values in a case
	fmt.Println("\n2. Switch with multiple values in a case:")
	month := "Feb"
	switch month {
	case "Jan", "Mar", "May", "Jul", "Aug", "Oct", "Dec":
		fmt.Println("This month has 31 days")
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println("This month has 30 days")
	case "Feb":
		fmt.Println("This month has 28 or 29 days")
	default:
		fmt.Println("Invalid month")
	}

	// 3. Switch with no expression (like if-else chains)
	fmt.Println("\n3. Switch with no expression:")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	// 4. Switch with initialization statement
	fmt.Println("\n4. Switch with initialization statement:")
	switch n := 2; n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}
	// Note: n is not accessible here

	// 5. Switch with fallthrough
	fmt.Println("\n5. Switch with fallthrough:")
	num := 5
	switch num {
	case 5:
		fmt.Println("Five")
		fallthrough
	case 6:
		fmt.Println("Six")
		fallthrough
	case 7:
		fmt.Println("Seven")
	default:
		fmt.Println("Other number")
	}

	// 6. Switch with expressions in case
	fmt.Println("\n6. Switch with expressions in case:")
	x := 42
	switch {
	case x < 0:
		fmt.Println("Negative")
	case x >= 0 && x < 10:
		fmt.Println("Single digit")
	case x >= 10 && x < 100:
		fmt.Println("Double digit")
	default:
		fmt.Println("Large number")
	}

	// 7. Type switch
	fmt.Println("\n7. Type switch:")
	var i interface{} = "Hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("String length of %v is %v\n", v, len(v))
	case bool:
		fmt.Printf("Boolean complement of %v is %v\n", v, !v)
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}

	// 8. Switch with time
	fmt.Println("\n8. Switch with time:")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// 9. Switch on OS
	fmt.Println("\n9. Switch on operating system:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	case "windows":
		fmt.Println("Running on Windows")
	default:
		fmt.Printf("Running on %s\n", os)
	}

	// 10. Switch with break
	fmt.Println("\n10. Switch with break:")
	for i := 0; i < 5; i++ {
		switch i {
		case 2:
			fmt.Println("Breaking at 2")
			break // Breaks out of the switch, not the for loop
		case 3:
			fmt.Println("This won't print for i=3")
		}
		fmt.Printf("After switch: i = %d\n", i)
	}

	// 11. Breaking out of outer loop from switch
	fmt.Println("\n11. Breaking out of outer loop from switch:")
OuterLoop:
	for i := 0; i < 5; i++ {
		switch i {
		case 3:
			fmt.Println("Breaking outer loop at i=3")
			break OuterLoop
		}
		fmt.Printf("In outer loop: i = %d\n", i)
	}

	// 12. Switch with complex initialization
	fmt.Println("\n12. Switch with complex initialization:")
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// 13. Empty switch (useful for scoping variables)
	fmt.Println("\n13. Empty switch with variables:")
	switch user, err := getUser(1); {
	case err != nil:
		fmt.Printf("Error: %v\n", err)
	case user.Admin:
		fmt.Printf("User %s is an admin\n", user.Name)
	default:
		fmt.Printf("User %s is not an admin\n", user.Name)
	}
}

// Helper struct and function for example 13
type User struct {
	ID    int
	Name  string
	Admin bool
}

func getUser(id int) (User, error) {
	// In a real app, this might fetch from a database
	return User{ID: id, Name: "John", Admin: true}, nil
}
