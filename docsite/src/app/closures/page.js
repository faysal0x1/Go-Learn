import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Closures() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Closures in Go</h2>
            <p className="mb-4">
                Closures in Go are functions that capture and maintain references to variables from their
                surrounding scope. They are powerful tools for creating functions with state and
                implementing various programming patterns.
            </p>
            <p className="mb-4">
                Key features of closures in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Capture variables from outer scope</li>
                <li>Maintain state between function calls</li>
                <li>Can return functions</li>
                <li>Support for multiple independent instances</li>
                <li>Can modify captured variables</li>
                <li>Useful for various design patterns</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Closure Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Closures</h3>
                    <p className="mb-4">
                        Simple closures capturing variables and returning functions:
                    </p>
                    <CodeBlock code={`// Basic closure capturing a variable
x := 10
increment := func() {
    x++
    fmt.Printf("x is now: %d\\n", x)
}

increment() // x is now: 11
increment() // x is now: 12

// Closure returning a function
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

counter := makeCounter()
fmt.Printf("Counter: %d\\n", counter()) // 1
fmt.Printf("Counter: %d\\n", counter()) // 2
fmt.Printf("Counter: %d\\n", counter()) // 3`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Closures with Parameters</h3>
                    <p className="mb-4">
                        Closures that take parameters and maintain state:
                    </p>
                    <CodeBlock code={`// Closure with parameters
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

func makeMultiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}

// Using the closures
adder := makeAdder(5)
fmt.Printf("5 + 3 = %d\\n", adder(3))    // 8
fmt.Printf("5 + 7 = %d\\n", adder(7))    // 12

multiplier := makeMultiplier(4)
fmt.Printf("4 * 6 = %d\\n", multiplier(6)) // 24`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. State Management</h3>
                    <p className="mb-4">
                        Using closures to manage state and create encapsulated objects:
                    </p>
                    <CodeBlock code={`// Calculator with state
func makeCalculator() struct {
    add       func(float64)
    subtract  func(float64)
    multiply  func(float64)
    divide    func(float64)
    getResult func() float64
} {
    result := 0.0
    return struct {
        add       func(float64)
        subtract  func(float64)
        multiply  func(float64)
        divide    func(float64)
        getResult func() float64
    }{
        add: func(x float64) {
            result += x
        },
        subtract: func(x float64) {
            result -= x
        },
        multiply: func(x float64) {
            result *= x
        },
        divide: func(x float64) {
            if x != 0 {
                result /= x
            }
        },
        getResult: func() float64 {
            return result
        },
    }
}

// Using the calculator
calc := makeCalculator()
calc.add(10)
calc.multiply(3)
calc.subtract(5)
fmt.Printf("Result: %.2f\\n", calc.getResult())`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Common Patterns</h3>
                    <p className="mb-4">
                        Practical patterns using closures:
                    </p>
                    <CodeBlock code={`// Logger with closure
func makeLogger(level string) func(string) {
    return func(message string) {
        fmt.Printf("[%s] %s\\n", level, message)
    }
}

// Filter with closure
func filter(slice []int, predicate func(int) bool) []int {
    result := []int{}
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// Using the patterns
logger := makeLogger("INFO")
logger("Application started")

numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evens := filter(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Printf("Even numbers: %v\\n", evens)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Advanced Patterns</h3>
                    <p className="mb-4">
                        More complex closure patterns:
                    </p>
                    <CodeBlock code={`// Memoization with closure
func makeFibonacci() func(int) int {
    cache := make(map[int]int)
    return func(n int) int {
        if val, ok := cache[n]; ok {
            return val
        }
        if n <= 1 {
            return n
        }
        cache[n] = makeFibonacci()(n-1) + makeFibonacci()(n-2)
        return cache[n]
    }
}

// Middleware pattern
func withLogging(handler func(string)) func(string) {
    return func(msg string) {
        fmt.Printf("Logging: %s\\n", msg)
        handler(msg)
    }
}

// Using the patterns
fib := makeFibonacci()
fmt.Printf("Fibonacci(10): %d\\n", fib(10))

loggedHandler := withLogging(func(msg string) {
    fmt.Printf("Processing: %s\\n", msg)
})
loggedHandler("Hello World")`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Bank Account</h3>
                    <p className="mb-4">
                        Create a closure-based bank account that supports deposits, withdrawals, and balance
                        checking. Include overdraft protection and transaction history.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Initial balance: $100.00
                            Deposit: $50.00
                            Withdraw: $30.00
                            Withdraw: $150.00 (failed - insufficient funds)
                            Current balance: $120.00
                            Transaction history:
                            - Deposit: +$50.00
                            - Withdraw: -$30.00
                            - Withdraw: -$150.00 (failed)</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Event System</h3>
                    <p className="mb-4">
                        Implement an event system using closures that supports subscribing to events,
                        emitting events, and handling multiple subscribers.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Event: user_login
                            - Handler 1: User logged in: alice@example.com
                            - Handler 2: Welcome back, alice@example.com

                            Event: user_logout
                            - Handler 1: User logged out: alice@example.com
                            - Handler 2: Goodbye, alice@example.com</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Validation System</h3>
                    <p className="mb-4">
                        Create a validation system using closures that can combine multiple validators and
                        apply them to different types of data.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Validating email: test@example.com
                            - Length check: passed
                            - Format check: passed
                            - Domain check: passed
                            Result: Valid

                            Validating email: invalid
                            - Length check: passed
                            - Format check: failed
                            Result: Invalid</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different closure patterns in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Basic counter
    counter := makeCounter()
    fmt.Println("Counter:")
    fmt.Printf("Count: %d\\n", counter())
    fmt.Printf("Count: %d\\n", counter())
    fmt.Printf("Count: %d\\n", counter())
    
    // Calculator
    calc := makeCalculator()
    fmt.Println("\\nCalculator:")
    calc.add(10)
    calc.multiply(3)
    calc.subtract(5)
    fmt.Printf("Result: %.2f\\n", calc.getResult())
    
    // Logger
    fmt.Println("\\nLogger:")
    logger := makeLogger("INFO")
    logger("Application started")
    logger("Processing data")
    
    // Filter
    fmt.Println("\\nFilter:")
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evens := filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Printf("Even numbers: %v\\n", evens)
}

func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func makeCalculator() struct {
    add       func(float64)
    multiply  func(float64)
    subtract  func(float64)
    getResult func() float64
} {
    result := 0.0
    return struct {
        add       func(float64)
        multiply  func(float64)
        subtract  func(float64)
        getResult func() float64
    }{
        add: func(x float64) {
            result += x
        },
        multiply: func(x float64) {
            result *= x
        },
        subtract: func(x float64) {
            result -= x
        },
        getResult: func() float64 {
            return result
        },
    }
}

func makeLogger(level string) func(string) {
    return func(message string) {
        fmt.Printf("[%s] %s\\n", level, message)
    }
}

func filter(slice []int, predicate func(int) bool) []int {
    result := []int{}
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Counter:
                        Count: 1
                        Count: 2
                        Count: 3

                        Calculator:
                        Result: 25.00

                        Logger:
                        [INFO] Application started
                        [INFO] Processing data

                        Filter:
                        Even numbers: [2 4 6 8 10]</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Closures"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 