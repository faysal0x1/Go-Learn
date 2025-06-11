import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Slices() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Slices in Go</h2>
            <p className="mb-4">
                Slices in Go are dynamic, flexible views into sequences of data. Unlike arrays, slices are
                reference types that can grow and shrink as needed. They are one of the most commonly used
                data structures in Go.
            </p>
            <p className="mb-4">
                Key features of slices in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Dynamic length</li>
                <li>Reference type (points to underlying array)</li>
                <li>Can grow and shrink</li>
                <li>Zero-based indexing</li>
                <li>Can be multi-dimensional</li>
                <li>Built-in append and copy operations</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Slice Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Creating Slices</h3>
                    <p className="mb-4">
                        Different ways to create and initialize slices:
                    </p>
                    <CodeBlock code={`// Using slice literal
s1 := []int{1, 2, 3, 4, 5}

// Using make() function
s2 := make([]int, 5)      // length=5, capacity=5
s3 := make([]int, 3, 10)  // length=3, capacity=10

// Empty slices
var s4 []int              // nil slice
s5 := []int{}             // empty slice, not nil

// From array
arr := [5]int{10, 20, 30, 40, 50}
s6 := arr[:]              // slice of the entire array`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Basic Slice Operations</h3>
                    <p className="mb-4">
                        Working with slice elements and properties:
                    </p>
                    <CodeBlock code={`nums := []int{10, 20, 30, 40, 50}

// Accessing elements
fmt.Println("First element:", nums[0])
fmt.Println("Last element:", nums[len(nums)-1])

// Modifying elements
nums[1] = 25
fmt.Println("After modification:", nums)

// Length and capacity
fmt.Printf("Length: %d, Capacity: %d\\n", len(nums), cap(nums))`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Slicing Operations</h3>
                    <p className="mb-4">
                        Different ways to create subslices:
                    </p>
                    <CodeBlock code={`original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// Slice[start:end] - elements from index start to end-1
fmt.Println("original[1:4]:", original[1:4])  // [1, 2, 3]

// Slice[:end] - elements from start to end-1
fmt.Println("original[:3]:", original[:3])    // [0, 1, 2]

// Slice[start:] - elements from start to end
fmt.Println("original[6:]:", original[6:])    // [6, 7, 8, 9]

// Slice[:] - all elements
fmt.Println("original[:]:", original[:])      // all elements`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Growing Slices</h3>
                    <p className="mb-4">
                        Using append and copy operations:
                    </p>
                    <CodeBlock code={`// Append to slice
s := []int{1, 2, 3}
s = append(s, 4)                    // [1, 2, 3, 4]
s = append(s, 5, 6, 7)             // [1, 2, 3, 4, 5, 6, 7]

// Append one slice to another
s2 := []int{8, 9, 10}
s = append(s, s2...)               // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

// Copy slices
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)
copied := copy(dst, src)           // copies min(len(dst), len(src)) elements`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Common Slice Patterns</h3>
                    <p className="mb-4">
                        Useful patterns for working with slices:
                    </p>
                    <CodeBlock code={`// Stack operations
stack := []int{}
// Push
stack = append(stack, 1)
stack = append(stack, 2)
// Pop
x, stack := stack[len(stack)-1], stack[:len(stack)-1]

// Removing elements
elements := []string{"a", "b", "c", "d", "e"}
i := 2 // remove index 2 ("c")
elements = append(elements[:i], elements[i+1:]...)

// Filtering elements
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var evens []int
for _, n := range numbers {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Slice Operations</h3>
                    <p className="mb-4">
                        Create a program that performs various operations on a slice: finding the sum, average,
                        and implementing a function to remove duplicates.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Original: [1, 2, 2, 3, 4, 4, 5]
                            After removing duplicates: [1, 2, 3, 4, 5]
                            Sum: 15
                            Average: 3</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Slice Manipulation</h3>
                    <p className="mb-4">
                        Create a program that implements a sliding window operation on a slice, finding the
                        maximum sum of k consecutive elements.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Array: [1, 3, 2, 4, 5, 1, 2]
                            Window size: 3
                            Maximum sum: 11 (from [4, 5, 1])</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Matrix Operations</h3>
                    <p className="mb-4">
                        Create a program that performs operations on a 2D slice (matrix): transpose the matrix
                        and find the sum of each row and column.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Original Matrix:
                            1 2 3
                            4 5 6
                            7 8 9

                            Transposed Matrix:
                            1 4 7
                            2 5 8
                            3 6 9

                            Row sums: 6 15 24
                            Column sums: 12 15 18</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different slice operations in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Create and initialize a slice
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Original slice:", numbers)
    
    // Append elements
    numbers = append(numbers, 6, 7, 8)
    fmt.Println("After append:", numbers)
    
    // Create a subslice
    subslice := numbers[2:5]
    fmt.Println("Subslice:", subslice)
    
    // Modify the subslice (affects original)
    subslice[0] = 99
    fmt.Println("After modification:")
    fmt.Println("Original:", numbers)
    fmt.Println("Subslice:", subslice)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Original slice: [1 2 3 4 5]
                        After append: [1 2 3 4 5 6 7 8]
                        Subslice: [3 4 5]
                        After modification:
                        Original: [1 2 99 4 5 6 7 8]
                        Subslice: [99 4 5]</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Slices"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 