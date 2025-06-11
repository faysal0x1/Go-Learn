import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function ForLoop() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">For Loops in Go</h2>
            <p className="mb-4">
                In Go, the <code>for</code> loop is the only looping construct available. However, it's versatile enough
                to handle all types of iteration needs. Go's for loop can be used in several different ways to achieve
                different looping patterns.
            </p>
            <p className="mb-4">
                Key features of for loops in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Traditional C-style for loop</li>
                <li>While-style loop</li>
                <li>Infinite loop with break</li>
                <li>Range-based iteration</li>
                <li>Nested loops with labels</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">For Loop Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic For Loop</h3>
                    <p className="mb-4">
                        The traditional C-style for loop with initialization, condition, and increment:
                    </p>
                    <CodeBlock code={`for i := 0; i < 5; i++ {
    fmt.Printf("i = %d\\n", i)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. While-Style Loop</h3>
                    <p className="mb-4">
                        Using for loop as a while loop with only a condition:
                    </p>
                    <CodeBlock code={`j := 0
for j < 5 {
    fmt.Printf("j = %d\\n", j)
    j++
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Range-Based Loop</h3>
                    <p className="mb-4">
                        Iterating over different data structures using range:
                    </p>
                    <CodeBlock code={`// Array iteration
numbers := [5]int{10, 20, 30, 40, 50}
for index, value := range numbers {
    fmt.Printf("index = %d, value = %d\\n", index, value)
}

// Slice iteration
fruits := []string{"apple", "banana", "cherry"}
for index, fruit := range fruits {
    fmt.Printf("index = %d, fruit = %s\\n", index, fruit)
}

// Map iteration
capitals := map[string]string{
    "USA":    "Washington DC",
    "France": "Paris",
    "Japan":  "Tokyo",
}
for country, capital := range capitals {
    fmt.Printf("Capital of %s is %s\\n", country, capital)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Break and Continue</h3>
                    <p className="mb-4">
                        Using break and continue statements in loops:
                    </p>
                    <CodeBlock code={`// Using break
for i := 0; i < 10; i++ {
    if i == 5 {
        fmt.Println("Breaking the loop at i =", i)
        break
    }
    fmt.Printf("i = %d\\n", i)
}

// Using continue
for i := 0; i < 5; i++ {
    if i == 2 {
        fmt.Println("Skipping i =", i)
        continue
    }
    fmt.Printf("i = %d\\n", i)
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
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Number Patterns</h3>
                    <p className="mb-4">
                        Create a program that prints the following number pattern using nested for loops:
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">1
                            1 2
                            1 2 3
                            1 2 3 4
                            1 2 3 4 5</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Array Processing</h3>
                    <p className="mb-4">
                        Create a program that finds the sum and average of numbers in an array using a for loop.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Numbers: [10, 20, 30, 40, 50]
                            Sum: 150
                            Average: 30</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: String Processing</h3>
                    <p className="mb-4">
                        Create a program that counts the number of vowels in a string using a for loop.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">String: "Hello, World!"
                            Number of vowels: 3</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different for loop patterns in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Try different for loop patterns here
    // 1. Basic for loop
    for i := 0; i < 5; i++ {
        fmt.Printf("i = %d\\n", i)
    }
    
    // 2. Range-based loop
    fruits := []string{"apple", "banana", "cherry"}
    for index, fruit := range fruits {
        fmt.Printf("index = %d, fruit = %s\\n", index, fruit)
    }
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>i = 0
                        i = 1
                        i = 2
                        i = 3
                        i = 4
                        index = 0, fruit = apple
                        index = 1, fruit = banana
                        index = 2, fruit = cherry</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="For Loops"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 