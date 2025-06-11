import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function VariadicFunctions() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Variadic Functions in Go</h2>
            <p className="mb-4">
                Variadic functions in Go can accept a variable number of arguments. They are useful when
                you don't know in advance how many arguments a function will receive.
            </p>
            <p className="mb-4">
                Key features of variadic functions in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Accept variable number of arguments</li>
                <li>Must be the last parameter</li>
                <li>Can accept zero arguments</li>
                <li>Can be passed a slice using ...</li>
                <li>Can work with any type</li>
                <li>Can be combined with other parameters</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Variadic Function Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Variadic Functions</h3>
                    <p className="mb-4">
                        Simple variadic functions for basic operations:
                    </p>
                    <CodeBlock code={`// Basic variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Using the variadic function
fmt.Printf("Sum of no numbers: %d\\n", sum())
fmt.Printf("Sum of 1: %d\\n", sum(1))
fmt.Printf("Sum of 1,2,3: %d\\n", sum(1, 2, 3))

// Passing a slice to a variadic function
numbers := []int{10, 20, 30, 40}
fmt.Printf("Sum of slice: %d\\n", sum(numbers...))`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Variadic with Other Parameters</h3>
                    <p className="mb-4">
                        Combining variadic parameters with regular parameters:
                    </p>
                    <CodeBlock code={`// Variadic parameter must be last
func multiply(multiplier int, numbers ...int) int {
    result := 1
    for _, num := range numbers {
        result *= num
    }
    return result * multiplier
}

// Using the function
result := multiply(2, 1, 2, 3, 4)
fmt.Printf("2 * (1,2,3,4) = %d\\n", result)

// String joining with separator
func joinStrings(separator string, strings ...string) string {
    return strings.Join(strings, separator)
}

message := joinStrings("-", "Hello", "World", "Go")
fmt.Printf("Joined: %s\\n", message)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Variadic with Different Types</h3>
                    <p className="mb-4">
                        Using interface{ } to accept any type of arguments:
                    </p>
                    <CodeBlock code={`// Variadic function accepting any type
func printValues(label string, values ...interface{}) {
    fmt.Printf("%s ", label)
    for i, value := range values {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Printf("%v", value)
    }
    fmt.Println()
}

// Using the function with different types
printValues("Numbers:", 1, 2, 3, 4, 5)
printValues("Mixed:", "hello", 42, 3.14, true)

// Processing different types of numbers
func processNumbers(operation string, numbers ...int) {
    fmt.Printf("%s: ", operation)
    for i, num := range numbers {
        if i > 0 {
            fmt.Print(", ")
        }
        switch operation {
        case "Squares":
            fmt.Printf("%d", num*num)
        case "Cubes":
            fmt.Printf("%d", num*num*num)
        }
    }
    fmt.Println()
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Error Handling in Variadic Functions</h3>
                    <p className="mb-4">
                        Combining variadic functions with error handling:
                    </p>
                    <CodeBlock code={`// Variadic function with error handling
func average(numbers ...float64) (float64, error) {
    if len(numbers) == 0 {
        return 0, errors.New("no numbers provided")
    }
    
    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    return sum / float64(len(numbers)), nil
}

// Using the function
avg, err := average(10, 20, 30, 40)
if err != nil {
    fmt.Printf("Error: %v\\n", err)
} else {
    fmt.Printf("Average: %.2f\\n", avg)
}

// Handling empty arguments
_, err = average()
if err != nil {
    fmt.Printf("Error: %v\\n", err)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Advanced Variadic Patterns</h3>
                    <p className="mb-4">
                        More complex uses of variadic functions:
                    </p>
                    <CodeBlock code={`// Variadic function returning slice
func doubleAll(numbers ...int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = num * 2
    }
    return result
}

// Variadic with structs
type Person struct {
    Name string
    Age  int
}

func createPeople(people ...Person) []Person {
    return people
}

// Variadic with maps
func mergeMaps(maps ...map[string]int) map[string]int {
    result := make(map[string]int)
    for _, m := range maps {
        for k, v := range m {
            result[k] = v
        }
    }
    return result
}

// Using these functions
people := createPeople(
    Person{"Alice", 25},
    Person{"Bob", 30},
)

merged := mergeMaps(
    map[string]int{"a": 1, "b": 2},
    map[string]int{"c": 3, "d": 4},
)`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Statistics Functions</h3>
                    <p className="mb-4">
                        Create a set of variadic functions for statistical calculations: mean, median, mode,
                        and standard deviation.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Data: [2.5, 3.7, 1.2, 4.8, 2.1, 3.3, 2.9]
                            Mean: 2.93
                            Median: 2.9
                            Mode: 2.5
                            Standard Deviation: 1.12</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: String Processing</h3>
                    <p className="mb-4">
                        Create variadic functions for string processing: concatenation with different
                        separators, finding common prefixes/suffixes, and filtering strings based on
                        conditions.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Strings: ["hello", "world", "go", "programming"]
                            Joined with space: "hello world go programming"
                            Joined with comma: "hello,world,go,programming"
                            Common prefix: ""
                            Strings longer than 3: ["hello", "world", "programming"]</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Data Validation</h3>
                    <p className="mb-4">
                        Implement variadic functions for data validation: email validation, number range
                        checking, and type checking.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Emails: ["test@example.com", "invalid-email", "user@domain.org"]
                            Valid emails: ["test@example.com", "user@domain.org"]

                            Numbers: [1, 5, 10, 15, 20]
                            Numbers in range 1-10: [1, 5, 10]

                            Types: [42, "hello", 3.14, true]
                            Integers: [42]</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different variadic function patterns in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Basic variadic function
    fmt.Println("Sum of numbers:")
    fmt.Printf("sum(1, 2, 3) = %d\\n", sum(1, 2, 3))
    fmt.Printf("sum(10, 20, 30, 40) = %d\\n", sum(10, 20, 30, 40))
    
    // Variadic with other parameters
    fmt.Println("\\nString joining:")
    fmt.Printf("joinStrings(\"-\", \"Hello\", \"World\") = %s\\n",
        joinStrings("-", "Hello", "World"))
    
    // Variadic with different types
    fmt.Println("\\nPrinting different types:")
    printValues("Mixed values:", 42, "hello", 3.14, true)
    
    // Processing numbers
    fmt.Println("\\nProcessing numbers:")
    processNumbers("Squares", 1, 2, 3, 4, 5)
    processNumbers("Cubes", 2, 3, 4)
}

func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func joinStrings(separator string, strings ...string) string {
    return strings.Join(strings, separator)
}

func printValues(label string, values ...interface{}) {
    fmt.Printf("%s ", label)
    for i, value := range values {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Printf("%v", value)
    }
    fmt.Println()
}

func processNumbers(operation string, numbers ...int) {
    fmt.Printf("%s: ", operation)
    for i, num := range numbers {
        if i > 0 {
            fmt.Print(", ")
        }
        switch operation {
        case "Squares":
            fmt.Printf("%d", num*num)
        case "Cubes":
            fmt.Printf("%d", num*num*num)
        }
    }
    fmt.Println()
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Sum of numbers:
                        sum(1, 2, 3) = 6
                        sum(10, 20, 30, 40) = 100

                        String joining:
                        joinStrings("-", "Hello", "World") = Hello-World

                        Printing different types:
                        Mixed values: 42, hello, 3.14, true

                        Processing numbers:
                        Squares: 1, 4, 9, 16, 25
                        Cubes: 8, 27, 64</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Variadic Functions"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 