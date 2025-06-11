package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Slices Examples")
	fmt.Println("==================")

	// 1. Creating slices
	fmt.Println("\n1. Creating Slices:")

	// a. Using slice literal
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice literal:", s1)

	// b. Using make() function
	s2 := make([]int, 5) // length=5, capacity=5
	fmt.Println("make([]int, 5):", s2)

	// c. With specific capacity
	s3 := make([]int, 3, 10) // length=3, capacity=10
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	// d. Empty slice
	var s4 []int // nil slice
	fmt.Printf("nil slice: %v, len=%d, cap=%d, isNil=%t\n", s4, len(s4), cap(s4), s4 == nil)

	// e. Empty slice with literal
	s5 := []int{} // empty slice, not nil
	fmt.Printf("empty slice: %v, len=%d, cap=%d, isNil=%t\n", s5, len(s5), cap(s5), s5 == nil)

	// f. From array
	arr := [5]int{10, 20, 30, 40, 50}
	s6 := arr[:] // slice of the entire array
	fmt.Println("Slice from array:", s6)

	// 2. Basic slice operations
	fmt.Println("\n2. Basic Slice Operations:")

	// a. Accessing elements
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("First element:", nums[0])
	fmt.Println("Last element:", nums[len(nums)-1])

	// b. Modifying elements
	nums[1] = 25
	fmt.Println("After modification:", nums)

	// c. Length and capacity
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// 3. Slicing operations
	fmt.Println("\n3. Slicing Operations:")

	// a. Taking subslices
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original:", original)

	// b. Slice[start:end] - elements from index start to end-1
	fmt.Println("original[1:4]:", original[1:4]) // elements 1,2,3

	// c. Slice[:end] - elements from start to end-1
	fmt.Println("original[:3]:", original[:3]) // elements 0,1,2

	// d. Slice[start:] - elements from start to end
	fmt.Println("original[6:]:", original[6:]) // elements 6,7,8,9

	// e. Slice[:] - all elements
	fmt.Println("original[:]:", original[:]) // all elements

	// 4. Growing slices
	fmt.Println("\n4. Growing Slices:")

	// a. Append to slice
	s7 := []int{1, 2, 3}
	fmt.Println("Original:", s7)

	s7 = append(s7, 4)
	fmt.Println("After append(s7, 4):", s7)

	// b. Append multiple elements
	s7 = append(s7, 5, 6, 7)
	fmt.Println("After append(s7, 5, 6, 7):", s7)

	// c. Append one slice to another
	s8 := []int{8, 9, 10}
	s7 = append(s7, s8...) // ... is the spread operator
	fmt.Println("After append(s7, s8...):", s7)

	// d. Copy slices
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)    // destination slice with length 3
	copied := copy(dst, src) // copies min(len(dst), len(src)) elements
	fmt.Printf("copy(dst, src): %v, copied=%d\n", dst, copied)

	// 5. Multi-dimensional slices
	fmt.Println("\n5. Multi-dimensional Slices:")

	// a. 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D slice:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// b. Creating a 2D slice with make
	grid := make([][]int, 3)
	for i := range grid {
		grid[i] = make([]int, 2)
		grid[i][0] = i
		grid[i][1] = i * 2
	}
	fmt.Println("Dynamic 2D slice:")
	for _, row := range grid {
		fmt.Println(row)
	}

	// 6. Common slice patterns
	fmt.Println("\n6. Common Slice Patterns:")

	// a. Slice as a stack
	stack := []int{}
	// Push
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println("Stack after pushes:", stack)

	// Pop
	x, stack := stack[len(stack)-1], stack[:len(stack)-1]
	fmt.Printf("Popped %d, Stack: %v\n", x, stack)

	// b. Removing elements
	// From the middle (preserving order)
	elements := []string{"a", "b", "c", "d", "e"}
	i := 2 // remove index 2 ("c")
	elements = append(elements[:i], elements[i+1:]...)
	fmt.Println("After removing middle element:", elements)

	// c. Filtering elements
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens []int
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Println("Even numbers:", evens)

	// 7. Performance considerations
	fmt.Println("\n7. Performance Considerations:")

	// a. Preallocating capacity for better performance
	fmt.Println("Preallocating capacity:")

	withPrealloc := make([]int, 0, 1000)
	withoutPrealloc := []int{}

	// Example to show how we'd use them:
	for i := 0; i < 10; i++ {
		withPrealloc = append(withPrealloc, i)
		withoutPrealloc = append(withoutPrealloc, i)
	}

	fmt.Printf("withPrealloc: len=%d, cap=%d\n", len(withPrealloc), cap(withPrealloc))
	fmt.Printf("withoutPrealloc: len=%d, cap=%d\n", len(withoutPrealloc), cap(withoutPrealloc))

	// b. Slice memory sharing
	original2 := []int{1, 2, 3, 4, 5}
	slice1 := original2[1:3]
	fmt.Println("original2:", original2)
	fmt.Println("slice1:", slice1)

	// Modifying shared memory
	slice1[0] = 99
	fmt.Println("After slice1[0] = 99:")
	fmt.Println("original2:", original2)
	fmt.Println("slice1:", slice1)

	// c. Breaking the reference with full slice expression
	original3 := []int{1, 2, 3, 4, 5}
	slice2 := original3[1:3:3] // third parameter limits capacity
	fmt.Printf("slice2: %v, cap=%d\n", slice2, cap(slice2))

	// This append will create a new backing array since we're at capacity
	slice2 = append(slice2, 100)
	fmt.Println("After append to slice2:")
	fmt.Println("original3:", original3) // not affected
	fmt.Println("slice2:", slice2)
}
