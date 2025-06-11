import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Arrays() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Arrays in Go</h2>
            <p className="mb-4">
                Arrays in Go are fixed-length sequences of elements of the same type. The size of an array
                is part of its type, making them different from slices which are dynamic in size.
            </p>
            <p className="mb-4">
                Key features of arrays in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Fixed length</li>
                <li>Zero-based indexing</li>
                <li>Type includes size</li>
                <li>Value type (copied when passed)</li>
                <li>Can be multi-dimensional</li>
                <li>Can be initialized with values</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Array Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Array Declaration and Initialization</h3>
                    <p className="mb-4">
                        Different ways to declare and initialize arrays:
                    </p>
                    <CodeBlock code={`// Declare an array of 5 integers
var numbers [5]int

// Initialize array with values
scores := [5]int{95, 87, 98, 93, 85}

// Let compiler count the elements
days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

fmt.Println("Numbers:", numbers)  // [0 0 0 0 0]
fmt.Println("Scores:", scores)    // [95 87 98 93 85]
fmt.Println("Days:", days)        // [Monday Tuesday Wednesday Thursday Friday]`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Accessing and Modifying Elements</h3>
                    <p className="mb-4">
                        Working with array elements:
                    </p>
                    <CodeBlock code={`fruits := [3]string{"apple", "banana", "cherry"}

// Access elements
fmt.Println("First fruit:", fruits[0])  // apple
fmt.Println("Second fruit:", fruits[1]) // banana

// Modify elements
fruits[2] = "orange"
fmt.Println("Modified array:", fruits)  // [apple banana orange]

// Get array length
fmt.Println("Array length:", len(fruits))  // 3`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Iterating Over Arrays</h3>
                    <p className="mb-4">
                        Different ways to iterate over arrays:
                    </p>
                    <CodeBlock code={`numbers := [4]int{10, 20, 30, 40}

// Using for loop with index
for i := 0; i < len(numbers); i++ {
    fmt.Printf("numbers[%d] = %d\\n", i, numbers[i])
}

// Using range
for index, value := range numbers {
    fmt.Printf("numbers[%d] = %d\\n", index, value)
}

// Using range with blank identifier
for _, value := range numbers {
    fmt.Printf("Value: %d\\n", value)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Multi-dimensional Arrays</h3>
                    <p className="mb-4">
                        Working with 2D arrays:
                    </p>
                    <CodeBlock code={`// 2D array (3x3)
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Access elements
fmt.Println("Center element:", matrix[1][1])  // 5

// Iterate over 2D array
for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[i]); j++ {
        fmt.Printf("%d ", matrix[i][j])
    }
    fmt.Println()
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Array Comparison</h3>
                    <p className="mb-4">
                        Comparing arrays in Go:
                    </p>
                    <CodeBlock code={`a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{3, 2, 1}

fmt.Println("a == b:", a == b)  // true
fmt.Println("a == c:", a == c)  // false`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Array Operations</h3>
                    <p className="mb-4">
                        Create a program that performs basic operations on an array: finding the sum, average, and maximum value.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Array: [10, 20, 30, 40, 50]
                            Sum: 150
                            Average: 30
                            Maximum: 50</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Matrix Operations</h3>
                    <p className="mb-4">
                        Create a program that performs operations on a 2D array (matrix): finding the sum of each row and column.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Matrix:
                            1 2 3
                            4 5 6
                            7 8 9

                            Row sums: 6 15 24
                            Column sums: 12 15 18</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Array Rotation</h3>
                    <p className="mb-4">
                        Create a program that rotates an array by a given number of positions.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Original: [1, 2, 3, 4, 5]
                            Rotated by 2: [4, 5, 1, 2, 3]</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different array operations in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Try different array operations here
    numbers := [5]int{10, 20, 30, 40, 50}
    
    // Print array
    fmt.Println("Array:", numbers)
    
    // Calculate sum
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Println("Sum:", sum)
    
    // Find maximum
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    fmt.Println("Maximum:", max)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Array: [10 20 30 40 50]
                        Sum: 150
                        Maximum: 50</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Arrays"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 