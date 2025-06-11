import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Range() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Range in Go</h2>
            <p className="mb-4">
                The `range` keyword in Go is used to iterate over elements in various data structures.
                It provides a clean and efficient way to traverse arrays, slices, maps, strings, and channels.
            </p>
            <p className="mb-4">
                Key features of range in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Iterates over arrays and slices</li>
                <li>Iterates over maps (keys and values)</li>
                <li>Iterates over strings (runes)</li>
                <li>Iterates over channels</li>
                <li>Supports index/value pairs</li>
                <li>Can be used with break and continue</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Range Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Range over Slices and Arrays</h3>
                    <p className="mb-4">
                        Different ways to iterate over slices and arrays:
                    </p>
                    <CodeBlock code={`// Range over slice with index and value
fruits := []string{"apple", "banana", "cherry"}
for i, fruit := range fruits {
    fmt.Printf("Index %d: %s\\n", i, fruit)
}

// Range over slice with value only
for _, fruit := range fruits {
    fmt.Printf("Fruit: %s\\n", fruit)
}

// Range over slice with index only
for i := range fruits {
    fmt.Printf("Index: %d\\n", i)
}

// Range over array
numbers := [5]int{10, 20, 30, 40, 50}
for i, num := range numbers {
    fmt.Printf("numbers[%d] = %d\\n", i, num)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Range over Maps</h3>
                    <p className="mb-4">
                        Iterating over maps in different ways:
                    </p>
                    <CodeBlock code={`ages := map[string]int{
    "Alice":   25,
    "Bob":     30,
    "Charlie": 35,
}

// Range over map with key and value
for name, age := range ages {
    fmt.Printf("%s is %d years old\\n", name, age)
}

// Range over map keys only
for name := range ages {
    fmt.Printf("Name: %s\\n", name)
}

// Range over map with complex values
hobbies := map[string][]string{
    "Alice": {"reading", "swimming", "coding"},
    "Bob":   {"gaming", "football", "movies"},
}
for name, hobbyList := range hobbies {
    fmt.Printf("%s's hobbies: %s\\n", name, strings.Join(hobbyList, ", "))
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Range over Strings</h3>
                    <p className="mb-4">
                        Working with strings using range:
                    </p>
                    <CodeBlock code={`// Range over string (runes)
text := "Hello, 世界"
for i, r := range text {
    fmt.Printf("Index %d: %c (Unicode: %U)\\n", i, r, r)
}

// String bytes vs runes
str := "café"
fmt.Printf("String: %s (length: %d bytes)\\n", str, len(str))
fmt.Println("Ranging over runes:")
for i, r := range str {
    fmt.Printf("  Byte index %d: %c\\n", i, r)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Range over Channels</h3>
                    <p className="mb-4">
                        Using range with channels:
                    </p>
                    <CodeBlock code={`// Create and fill a channel
ch := make(chan int, 5)
for i := 1; i <= 5; i++ {
    ch <- i * 10
}
close(ch) // Must close channel to end range loop

// Range over channel
for value := range ch {
    fmt.Printf("Received: %d\\n", value)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Advanced Range Patterns</h3>
                    <p className="mb-4">
                        More complex range patterns:
                    </p>
                    <CodeBlock code={`// Nested range loops
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
for i, row := range matrix {
    for j, val := range row {
        fmt.Printf("matrix[%d][%d] = %d\\n", i, j, val)
    }
}

// Range with break and continue
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
for _, num := range nums {
    if num > 6 {
        break // Exit the loop
    }
    if num%2 != 0 {
        continue // Skip odd numbers
    }
    fmt.Printf("%d ", num)
}

// Range with labeled loops
outer:
for i, row := range matrix {
    for j, val := range row {
        if val == 5 {
            fmt.Println("Found 5, breaking out of both loops")
            break outer
        }
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
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Matrix Operations</h3>
                    <p className="mb-4">
                        Create a program that performs operations on a 2D matrix using range loops. Implement
                        functions to find the sum of each row and column, and to find the maximum value in
                        each row.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Matrix:
                            1 2 3
                            4 5 6
                            7 8 9

                            Row sums: 6 15 24
                            Column sums: 12 15 18
                            Row maximums: 3 6 9</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Text Analysis</h3>
                    <p className="mb-4">
                        Create a program that analyzes text using range loops. Count the frequency of each
                        character, find the most common character, and identify all unique characters.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Text: "Hello, 世界!"
                            Character frequencies:
                            H: 1
                            e: 1
                            l: 2
                            o: 1
                            ,: 1
                            世: 1
                            界: 1
                            !: 1

                            Most common character: 'l' (2 occurrences)
                            Unique characters: 8</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Student Records</h3>
                    <p className="mb-4">
                        Create a program that manages student records using maps and slices. Use range loops
                        to calculate averages, find the highest score, and identify students who need
                        improvement.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Student Records:
                            Alice: [95, 88, 92]
                            Bob: [78, 85, 80]
                            Charlie: [90, 95, 88]

                            Averages:
                            Alice: 91.67
                            Bob: 81.00
                            Charlie: 91.00

                            Highest score: 95 (Charlie in Science)
                            Need improvement: Bob</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different range operations in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Create a slice of numbers
    numbers := []int{10, 20, 30, 40, 50}
    fmt.Println("Original numbers:", numbers)
    
    // Range over slice with index and value
    fmt.Println("\\nIterating with index and value:")
    for i, num := range numbers {
        fmt.Printf("numbers[%d] = %d\\n", i, num)
    }
    
    // Range over slice with value only
    fmt.Println("\\nIterating with value only:")
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Printf("Sum of numbers: %d\\n", sum)
    
    // Range over slice with index only
    fmt.Println("\\nIterating with index only:")
    for i := range numbers {
        numbers[i] *= 2
    }
    fmt.Println("Numbers after doubling:", numbers)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Original numbers: [10 20 30 40 50]

                        Iterating with index and value:
                        numbers[0] = 10
                        numbers[1] = 20
                        numbers[2] = 30
                        numbers[3] = 40
                        numbers[4] = 50

                        Iterating with value only:
                        Sum of numbers: 150

                        Iterating with index only:
                        Numbers after doubling: [20 40 60 80 100]</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Range"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 