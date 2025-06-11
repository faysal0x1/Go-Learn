package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== Go Pointers Examples ===\n")

	// 1. Basic pointer declaration and usage
	fmt.Println("1. Basic pointer operations:")
	var x int = 42
	var p *int = &x // p is a pointer to int, stores address of x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of p (address): %p\n", p)
	fmt.Printf("Value pointed to by p: %d\n", *p)

	// Modifying value through pointer
	*p = 100
	fmt.Printf("After *p = 100, x = %d\n", x)

	// 2. Zero value of pointers
	fmt.Println("\n2. Zero value of pointers:")
	var ptr *int
	fmt.Printf("Zero value pointer: %v\n", ptr)
	fmt.Printf("Is nil: %t\n", ptr == nil)

	// Avoid dereferencing nil pointer (would cause panic)
	// fmt.Println(*ptr) // This would panic!

	// 3. Creating pointers with new()
	fmt.Println("\n3. Creating pointers with new():")
	numPtr := new(int)
	*numPtr = 25
	fmt.Printf("Value: %d, Address: %p\n", *numPtr, numPtr)

	// 4. Pointer to pointer
	fmt.Println("\n4. Pointer to pointer:")
	a := 10
	pa := &a
	ppa := &pa

	fmt.Printf("a = %d\n", a)
	fmt.Printf("*pa = %d\n", *pa)
	fmt.Printf("**ppa = %d\n", **ppa)

	**ppa = 20
	fmt.Printf("After **ppa = 20, a = %d\n", a)

	// 5. Pointers with functions - pass by reference
	fmt.Println("\n5. Pointers with functions:")
	num := 5
	fmt.Printf("Before increment: %d\n", num)
	increment(&num)
	fmt.Printf("After increment: %d\n", num)

	// Swap example
	x1, y1 := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x1, y1)
	swap(&x1, &y1)
	fmt.Printf("After swap: x=%d, y=%d\n", x1, y1)

	// 6. Pointers with structs
	fmt.Println("\n6. Pointers with structs:")
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Original: %+v\n", person)

	updatePerson(&person, "Bob", 35)
	fmt.Printf("Updated: %+v\n", person)

	// 7. Struct pointer shorthand
	fmt.Println("\n7. Struct pointer shorthand:")
	personPtr := &Person{Name: "Charlie", Age: 25}
	// Go allows (*personPtr).Name or personPtr.Name
	fmt.Printf("Name: %s, Age: %d\n", personPtr.Name, personPtr.Age)

	personPtr.Age = 26 // Equivalent to (*personPtr).Age = 26
	fmt.Printf("Updated age: %d\n", personPtr.Age)

	// 8. Pointers with arrays and slices
	fmt.Println("\n8. Pointers with arrays and slices:")
	arr := [3]int{1, 2, 3}
	fmt.Printf("Original array: %v\n", arr)

	modifyArray(&arr)
	fmt.Printf("Modified array: %v\n", arr)

	// Slices are reference types (contain pointer internally)
	slice := []int{4, 5, 6}
	fmt.Printf("Original slice: %v\n", slice)
	modifySlice(slice) // No need for & with slices
	fmt.Printf("Modified slice: %v\n", slice)

	// 9. Returning pointers from functions
	fmt.Println("\n9. Returning pointers from functions:")
	newPerson := createPerson("David", 40)
	fmt.Printf("New person: %+v\n", *newPerson)

	// Local variable can be returned as pointer (Go handles memory)
	localPtr := createLocal()
	fmt.Printf("Local value: %d\n", *localPtr)

	// 10. Pointer arithmetic (limited in Go)
	fmt.Println("\n10. Pointer arithmetic with unsafe:")
	nums := [5]int{10, 20, 30, 40, 50}
	ptr1 := &nums[0]

	fmt.Printf("First element: %d\n", *ptr1)

	// Use unsafe to do pointer arithmetic
	ptr2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr1)) + unsafe.Sizeof(int(0))))
	fmt.Printf("Second element via pointer arithmetic: %d\n", *ptr2)

	// 11. Comparing pointers
	fmt.Println("\n11. Comparing pointers:")
	val1, val2 := 100, 100
	ptr_a := &val1
	ptr_b := &val2
	ptr_c := &val1

	fmt.Printf("ptr_a == ptr_b: %t\n", ptr_a == ptr_b)
	fmt.Printf("ptr_a == ptr_c: %t\n", ptr_a == ptr_c)
	fmt.Printf("*ptr_a == *ptr_b: %t\n", *ptr_a == *ptr_b)

	// 12. Pointers with interfaces
	fmt.Println("\n12. Pointers with interfaces:")
	var writer Writer

	// Pointer receiver
	fw := &FileWriter{filename: "test.txt"}
	writer = fw
	writer.Write("Hello from pointer receiver")

	// Value receiver
	cw := ConsoleWriter{}
	writer = cw
	writer.Write("Hello from value receiver")

	// 13. Method receivers - pointer vs value
	fmt.Println("\n13. Method receivers:")
	rect1 := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Original rectangle: %+v\n", rect1)

	// Value receiver method
	area := rect1.Area()
	fmt.Printf("Area: %.2f\n", area)

	// Pointer receiver method
	rect1.Scale(2)
	fmt.Printf("After scaling: %+v\n", rect1)

	// 14. Pointers in data structures
	fmt.Println("\n14. Pointers in data structures:")
	// Linked list example
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}

	printList(head)

	// 15. Memory allocation patterns
	fmt.Println("\n15. Memory allocation:")
	// Stack allocation (usually)
	stackVar := 42
	fmt.Printf("Stack variable: %d at %p\n", stackVar, &stackVar)

	// Heap allocation (when address escapes)
	heapVar := createHeapVariable()
	fmt.Printf("Heap variable: %d at %p\n", *heapVar, heapVar)

	// 16. Common pointer patterns
	fmt.Println("\n16. Common patterns:")

	// Optional values using pointers
	user1 := User{Name: "Alice", Email: stringPtr("alice@example.com")}
	user2 := User{Name: "Bob", Email: nil}

	printUser(user1)
	printUser(user2)

	// 17. Pointer slices
	fmt.Println("\n17. Pointer slices:")
	people := []*Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	// Modify through pointers
	for _, p := range people {
		p.Age++
	}

	fmt.Println("After aging everyone by 1 year:")
	for _, p := range people {
		fmt.Printf("- %s: %d\n", p.Name, p.Age)
	}

	// 18. Function pointers
	fmt.Println("\n18. Function pointers:")
	var operation func(int, int) int
	operation = add

	fmt.Printf("Function pointer result: %d\n", operation(5, 3))

	// Array of function pointers
	operations := []func(int, int) int{add, subtract, multiply}
	for i, op := range operations {
		fmt.Printf("Operation %d: %d\n", i, op(10, 3))
	}

	// 19. Pointer best practices
	fmt.Println("\n19. Best practices:")

	// Check for nil before dereferencing
	var maybeNil *int
	if maybeNil != nil {
		fmt.Printf("Value: %d\n", *maybeNil)
	} else {
		fmt.Println("Pointer is nil, safe check!")
	}

	// Use pointers for large structs to avoid copying
	largeStruct := LargeStruct{Data: [1000]int{}}
	processLargeStruct(&largeStruct) // Efficient - no copy

	// 20. Common pitfalls and gotchas
	fmt.Println("\n20. Common pitfalls:")

	// Pitfall 1: Loop variable addresses
	var pointers []*int
	values := []int{1, 2, 3}

	// Wrong way
	for _, v := range values {
		pointers = append(pointers, &v) // All point to same variable!
	}
	fmt.Print("Wrong way (all same): ")
	for _, p := range pointers {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()

	// Correct way
	pointers = nil
	for _, v := range values {
		v := v // Create new variable
		pointers = append(pointers, &v)
	}
	fmt.Print("Correct way: ")
	for _, p := range pointers {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
}

// Helper functions and types

// 5. Functions using pointers
func increment(n *int) {
	*n++
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

// 6. Struct for pointer examples
type Person struct {
	Name string
	Age  int
}

func updatePerson(p *Person, name string, age int) {
	p.Name = name
	p.Age = age
}

// 8. Array and slice functions
func modifyArray(arr *[3]int) {
	arr[0] = 100
}

func modifySlice(slice []int) {
	slice[0] = 400
}

// 9. Returning pointers
func createPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func createLocal() *int {
	local := 42 // This will be allocated on heap
	return &local
}

// 12. Interface examples
type Writer interface {
	Write(data string)
}

type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(data string) {
	fmt.Printf("Writing to %s: %s\n", fw.filename, data)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) {
	fmt.Printf("Console: %s\n", data)
}

// 13. Method receiver examples
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 14. Linked list node
type Node struct {
	Value int
	Next  *Node
}

func printList(head *Node) {
	fmt.Print("Linked list: ")
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

// 15. Heap allocation
func createHeapVariable() *int {
	x := 100
	return &x // x escapes to heap
}

// 16. Optional values
type User struct {
	Name  string
	Email *string
}

func stringPtr(s string) *string {
	return &s
}

func printUser(u User) {
	fmt.Printf("User: %s", u.Name)
	if u.Email != nil {
		fmt.Printf(", Email: %s", *u.Email)
	} else {
		fmt.Printf(", Email: not provided")
	}
	fmt.Println()
}

// 18. Function operations
func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }

// 19. Large struct example
type LargeStruct struct {
	Data [1000]int
}

func processLargeStruct(ls *LargeStruct) {
	// Process without copying the large struct
	fmt.Printf("Processing large struct with %d elements\n", len(ls.Data))
}
