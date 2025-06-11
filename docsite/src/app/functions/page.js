import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Functions() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Functions in Go</h2>
            <p className="mb-4">
                Functions in Go are first-class citizens that can be assigned to variables, passed as
                arguments, and returned from other functions. They are a fundamental building block of
                Go programs.
            </p>
            <p className="mb-4">
                Key features of functions in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Multiple return values</li>
                <li>Named return values</li>
                <li>Variadic functions</li>
                <li>First-class functions</li>
                <li>Closures</li>
                <li>Methods (functions with receivers)</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Function Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Functions</h3>
                    <p className="mb-4">
                        Different ways to define and use functions:
                    </p>
                    <CodeBlock code={`// Function with no parameters or return value
func greet() {
    fmt.Println("Hello, World!")
}

// Function with parameter
func greetPerson(name string) {
    fmt.Printf("Hello, %s!\\n", name)
}

// Function with return value
func add(a, b int) int {
    return a + b
}

// Function with multiple return values
func calculate(a, b int) (int, int) {
    return a + b, a * b
}

// Named return values
func rectangleStats(length, width int) (area, perimeter int) {
    area = length * width
    perimeter = 2 * (length + width)
    return // naked return
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Variadic Functions</h3>
                    <p className="mb-4">
                        Functions that can accept a variable number of arguments:
                    </p>
                    <CodeBlock code={`// Variadic function
func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Using variadic functions
fmt.Printf("Sum of 1,2,3: %d\\n", sumAll(1, 2, 3))
fmt.Printf("Sum of 1,2,3,4,5: %d\\n", sumAll(1, 2, 3, 4, 5))

// Passing a slice to a variadic function
numbers := []int{10, 20, 30}
fmt.Printf("Sum of slice: %d\\n", sumAll(numbers...))`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Functions as Values</h3>
                    <p className="mb-4">
                        Using functions as first-class citizens:
                    </p>
                    <CodeBlock code={`// Function as variable
var operation func(int, int) int
operation = add
fmt.Printf("Using function variable: %d\\n", operation(7, 8))

// Anonymous function
multiply := func(a, b int) int {
    return a * b
}
fmt.Printf("Anonymous function result: %d\\n", multiply(4, 5))

// Higher-order function
func mapInts(slice []int, fn func(int) int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Using higher-order function
nums := []int{1, 2, 3, 4, 5}
doubled := mapInts(nums, func(x int) int { return x * 2 })`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Closures</h3>
                    <p className="mb-4">
                        Functions that capture their surrounding state:
                    </p>
                    <CodeBlock code={`// Function returning function (closure)
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

// Using closure
adder := makeAdder(10)
fmt.Printf("Adder(5): %d\\n", adder(5))
fmt.Printf("Adder(3): %d\\n", adder(3))

// Counter closure
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Using counter
counter := makeCounter()
fmt.Printf("Counter: %d\\n", counter())
fmt.Printf("Counter: %d\\n", counter())`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Error Handling</h3>
                    <p className="mb-4">
                        Working with errors in functions:
                    </p>
                    <CodeBlock code={`// Function with error handling
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Using error handling
quotient, err := divide(10, 2)
if err != nil {
    fmt.Printf("Error: %v\\n", err)
} else {
    fmt.Printf("10 / 2 = %.2f\\n", quotient)
}

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateAge(age int) (int, error) {
    if age < 0 {
        return 0, ValidationError{
            Field:   "age",
            Message: "age cannot be negative",
        }
    }
    return age, nil
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
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Calculator Functions</h3>
                    <p className="mb-4">
                        Create a set of calculator functions that perform basic arithmetic operations. Include
                        functions for addition, subtraction, multiplication, and division with proper error
                        handling.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">5 + 3 = 8
                            10 - 4 = 6
                            6 * 7 = 42
                            20 / 5 = 4
                            Error: division by zero</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Function Composition</h3>
                    <p className="mb-4">
                        Create a program that demonstrates function composition. Implement functions that can
                        be combined to create more complex operations.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Original number: 5
                            After addOne: 6
                            After multiplyByTwo: 12
                            After square: 144</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Higher-Order Functions</h3>
                    <p className="mb-4">
                        Implement a set of higher-order functions for working with slices: map, filter, and
                        reduce. Use these functions to process a list of numbers.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Original: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
                            Evens: [2, 4, 6, 8, 10]
                            Squares: [4, 16, 36, 64, 100]
                            Sum: 220</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different function patterns in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Basic function
    fmt.Println("Basic function:")
    greet("Alice")
    
    // Function with multiple return values
    fmt.Println("\\nMultiple return values:")
    sum, product := calculate(4, 6)
    fmt.Printf("Sum: %d, Product: %d\\n", sum, product)
    
    // Variadic function
    fmt.Println("\\nVariadic function:")
    total := sumAll(1, 2, 3, 4, 5)
    fmt.Printf("Sum of numbers: %d\\n", total)
    
    // Closure
    fmt.Println("\\nClosure:")
    counter := makeCounter()
    fmt.Printf("Counter: %d\\n", counter())
    fmt.Printf("Counter: %d\\n", counter())
}

func greet(name string) {
    fmt.Printf("Hello, %s!\\n", name)
}

func calculate(a, b int) (int, int) {
    return a + b, a * b
}

func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Basic function:
                        Hello, Alice!

                        Multiple return values:
                        Sum: 10, Product: 24

                        Variadic function:
                        Sum of numbers: 15

                        Closure:
                        Counter: 1
                        Counter: 2</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Functions"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
}