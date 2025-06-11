package main

import (
    "fmt"
    "math"
    "sort"
    "strings"
    "time"
)

func main() {
    fmt.Println("=== Go Closures Examples ===\n")

    // 1. Basic closure - capturing variables
    fmt.Println("1. Basic closure:")
    x := 10
    increment := func() {
        x++
        fmt.Printf("x is now: %d\n", x)
    }
    
    increment()
    increment()
    fmt.Printf("Final x value: %d\n", x)

    // 2. Closure returning a function
    fmt.Println("\n2. Closure returning a function:")
    counter := makeCounter()
    fmt.Printf("Counter: %d\n", counter())
    fmt.Printf("Counter: %d\n", counter())
    fmt.Printf("Counter: %d\n", counter())

    // 3. Multiple independent closures
    fmt.Println("\n3. Multiple independent closures:")
    counter1 := makeCounter()
    counter2 := makeCounter()
    
    fmt.Printf("Counter1: %d\n", counter1())
    fmt.Printf("Counter2: %d\n", counter2())
    fmt.Printf("Counter1: %d\n", counter1())
    fmt.Printf("Counter2: %d\n", counter2())

    // 4. Closure with parameters
    fmt.Println("\n4. Closure with parameters:")
    adder := makeAdder(5)
    fmt.Printf("5 + 3 = %d\n", adder(3))
    fmt.Printf("5 + 7 = %d\n", adder(7))

    multiplier := makeMultiplier(4)
    fmt.Printf("4 * 6 = %d\n", multiplier(6))

    // 5. Closure modifying captured variable
    fmt.Println("\n5. Closure modifying captured variable:")
    balance := 100.0
    deposit := func(amount float64) {
        balance += amount
        fmt.Printf("Deposited %.2f, new balance: %.2f\n", amount, balance)
    }
    withdraw := func(amount float64) bool {
        if balance >= amount {
            balance -= amount
            fmt.Printf("Withdrew %.2f, new balance: %.2f\n", amount, balance)
            return true
        }
        fmt.Printf("Insufficient funds. Current balance: %.2f\n", balance)
        return false
    }
    
    deposit(50.0)
    withdraw(30.0)
    withdraw(150.0)

    // 6. Closure in loops (common gotcha)
    fmt.Println("\n6. Closure in loops - wrong way:")
    var funcs []func()
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Printf("Wrong: %d\n", i) // All will print 3
        })
    }
    for _, f := range funcs {
        f()
    }

    fmt.Println("Closure in loops - correct way:")
    var funcs2 []func()
    for i := 0; i < 3; i++ {
        i := i // Create new variable in each iteration
        funcs2 = append(funcs2, func() {
            fmt.Printf("Correct: %d\n", i)
        })
    }
    for _, f := range funcs2 {
        f()
    }

    // 7. Closure with multiple captured variables
    fmt.Println("\n7. Multiple captured variables:")
    calculator := makeCalculator()
    calculator.add(10)
    calculator.multiply(3)
    calculator.subtract(5)
    fmt.Printf("Final result: %.2f\n", calculator.getResult())

    // 8. Closure for configuration
    fmt.Println("\n8. Closure for configuration:")
    logger := makeLogger("INFO")
    logger("Application started")
    logger("Processing data")

    errorLogger := makeLogger("ERROR")
    errorLogger("Something went wrong")

    // 9. Closure for filtering
    fmt.Println("\n9. Closure for filtering:")
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    isEven := func(n int) bool { return n%2 == 0 }
    evens := filter(numbers, isEven)
    fmt.Printf("Even numbers: %v\n", evens)

    greaterThan5 := func(n int) bool { return n > 5 }
    large := filter(numbers, greaterThan5)
    fmt.Printf("Numbers > 5: %v\n", large)

    // 10. Closure for mapping
    fmt.Println("\n10. Closure for mapping:")
    square := func(n int) int { return n * n }
    squares := mapInts(numbers[:5], square)
    fmt.Printf("Squares: %v\n", squares)

    // 11. Closure with state machine
    fmt.Println("\n11. State machine with closures:")
    fsm := makeStateMachine()
    fsm("start")
    fsm("process")
    fsm("finish")
    fsm("invalid")

    // 12. Closure for memoization
    fmt.Println("\n12. Memoization with closures:")
    fibMemo := makeFibonacci()
    fmt.Printf("Fibonacci(10): %d\n", fibMemo(10))
    fmt.Printf("Fibonacci(15): %d\n", fibMemo(15))
    fmt.Printf("Fibonacci(10): %d (cached)\n", fibMemo(10))

    // 13. Closure for retry logic
    fmt.Println("\n13. Retry logic with closures:")
    unreliableOperation := makeUnreliableOperation()
    retryableOp := makeRetryable(unreliableOperation, 3)
    result, err := retryableOp()
    if err != nil {
        fmt.Printf("Operation failed: %v\n", err)
    } else {
        fmt.Printf("Operation succeeded: %s\n", result)
    }

    // 14. Closure for middleware pattern
    fmt.Println("\n14. Middleware pattern:")
    handler := func(msg string) {
        fmt.Printf("Processing: %s\n", msg)
    }
    
    loggedHandler := withLogging(handler)
    timedHandler := withTiming(loggedHandler)
    
    timedHandler("Hello World")

    // 15. Closure for event handlers
    fmt.Println("\n15. Event handlers:")
    eventSystem := makeEventSystem()
    
    // Register event handlers using closures
    eventSystem.on("user_login", func(data interface{}) {
        fmt.Printf("User logged in: %v\n", data)
    })
    
    eventSystem.on("user_logout", func(data interface{}) {
        fmt.Printf("User logged out: %v\n", data)
    })
    
    eventSystem.emit("user_login", "alice@example.com")
    eventSystem.emit("user_logout", "alice@example.com")

    // 16. Closure for validation
    fmt.Println("\n16. Validation with closures:")
    validators := []func(string) error{
        makeMinLengthValidator(3),
        makeMaxLengthValidator(20),
        makeContainsValidator("@"),
    }
    
    testEmails := []string{"a", "test@example.com", "verylongemailaddressthatexceedslimit@example.com"}
    for _, email := range testEmails {
        fmt.Printf("Validating '%s': ", email)
        valid := true
        for _, validate := range validators {
            if err := validate(email); err != nil {
                fmt.Printf("Invalid - %v", err)
                valid = false
                break
            }
        }
        if valid {
            fmt.Print("Valid")
        }
        fmt.Println()
    }

    // 17. Closure for sorting
    fmt.Println("\n17. Custom sorting with closures:")
    people := []Person{
        {"Alice", 30, 75000},
        {"Bob", 25, 60000},
        {"Charlie", 35, 80000},
    }
    
    // Sort by age
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })
    fmt.Printf("Sorted by age: %v\n", people)
    
    // Sort by salary (descending)
    sort.Slice(people, func(i, j int) bool {
        return people[i].Salary > people[j].Salary
    })
    fmt.Printf("Sorted by salary (desc): %v\n", people)

    // 18. Closure for generators
    fmt.Println("\n18. Generator pattern:")
    primeGen := makePrimeGenerator()
    fmt.Print("First 10 primes: ")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", primeGen())
    }
    fmt.Println()

    // 19. Closure for caching
    fmt.Println("\n19. Caching with closures:")
    expensiveFunc := makeExpensiveFunction()
    cachedFunc := makeCache(expensiveFunc)
    
    fmt.Printf("First call: %d\n", cachedFunc(5))
    fmt.Printf("Second call (cached): %d\n", cachedFunc(5))
    fmt.Printf("Different input: %d\n", cachedFunc(3))

    // 20. Closure for debouncing
    fmt.Println("\n20. Debouncing with closures:")
    action := func() {
        fmt.Println("Action executed!")
    }
    
    debouncedAction := makeDebouncer(action, 100*time.Millisecond)
    
    // Rapid calls - only last one should execute
    debouncedAction()
    debouncedAction()
    debouncedAction()
    
    time.Sleep(200 * time.Millisecond) // Wait for debounce
}

// Helper functions and types

// 2. Basic counter
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 4. Adder and multiplier
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

// 7. Calculator with multiple operations
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
            fmt.Printf("Added %.2f, result: %.2f\n", x, result)
        },
        subtract: func(x float64) {
            result -= x
            fmt.Printf("Subtracted %.2f, result: %.2f\n", x, result)
        },
        multiply: func(x float64) {
            result *= x
            fmt.Printf("Multiplied by %.2f, result: %.2f\n", x, result)
        },
        divide: func(x float64) {
            if x != 0 {
                result /= x
                fmt.Printf("Divided by %.2f, result: %.2f\n", x, result)
            }
        },
        getResult: func() float64 {
            return result
        },
    }
}

// 8. Logger
func makeLogger(level string) func(string) {
    return func(message string) {
        timestamp := time.Now().Format("15:04:05")
        fmt.Printf("[%s] %s: %s\n", timestamp, level, message)
    }
}

// 9. Filter function
func filter(slice []int, predicate func(int) bool) []int {
    var result []int
    for _, item := range slice {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

// 10. Map function
func mapInts(slice []int, transform func(int) int) []int {
    result := make([]int, len(slice))
    for i, item := range slice {
        result[i] = transform(item)
    }
    return result
}

// 11. State machine
func makeStateMachine() func(string) {
    state := "idle"
    
    return func(event string) {
        fmt.Printf("State: %s, Event: %s -> ", state, event)
        
        switch state {
        case "idle":
            if event == "start" {
                state = "running"
            }
        case "running":
            if event == "process" {
                state = "processing"
            } else if event == "finish" {
                state = "idle"
            }
        case "processing":
            if event == "finish" {
                state = "idle"
            }
        }
        
        fmt.Printf("New state: %s\n", state)
    }
}

// 12. Memoized fibonacci
func makeFibonacci() func(int) int {
    cache := make(map[int]int)
    
    var fib func(int) int
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        
        if val, exists := cache[n]; exists {
            return val
        }
        
        result := fib(n-1) + fib(n-2)
        cache[n] = result
        return result
    }
    
    return fib
}

// 13. Unreliable operation and retry
func makeUnreliableOperation() func() (string, error) {
    attempts := 0
    return func() (string, error) {
        attempts++
        if attempts < 3 {
            return "", fmt.Errorf("operation failed (attempt %d)", attempts)
        }
        return "success", nil
    }
}

func makeRetryable(operation func() (string, error), maxRetries int) func() (string, error) {
    return func() (string, error) {
        var lastError error
        for i := 0; i < maxRetries; i++ {
            result, err := operation()
            if err == nil {
                return result, nil
            }
            lastError = err
            fmt.Printf("Retry %d failed: %v\n", i+1, err)
        }
        return "", lastError
    }
}

// 14. Middleware pattern
func withLogging(handler func(string)) func(string) {
    return func(msg string) {
        fmt.Printf("LOG: Starting to process '%s'\n", msg)
        handler(msg)
        fmt.Printf("LOG: Finished processing '%s'\n", msg)
    }
}

func withTiming(handler func(string)) func(string) {
    return func(msg string) {
        start := time.Now()
        handler(msg)
        duration := time.Since(start)
        fmt.Printf("TIMING: Took %v\n", duration)
    }
}

// 15. Event system
type EventSystem struct {
    handlers map[string][]func(interface{})
}

func makeEventSystem() *EventSystem {
    return &EventSystem{
        handlers: make(map[string][]func(interface{})),
    }
}

func (es *EventSystem) on(event string, handler func(interface{})) {
    es.handlers[event] = append(es.handlers[event], handler)
}

func (es *EventSystem) emit(event string, data interface{}) {
    if handlers, exists := es.handlers[event]; exists {
        for _, handler := range handlers {
            handler(data)
        }
    }
}

// 16. Validators
func makeMinLengthValidator(minLength int) func(string) error {
    return func(s string) error {
        if len(s) < minLength {
            return fmt.Errorf("must be at least %d characters", minLength)
        }
        return nil
    }
}

func makeMaxLengthValidator(maxLength int) func(string) error {
    return func(s string) error {
        if len(s) > maxLength {
            return fmt.Errorf("must be at most %d characters", maxLength)
        }
        return nil
    }
}

func makeContainsValidator(substring string) func(string) error {
    return func(s string) error {
        if !strings.Contains(s, substring) {
            return fmt.Errorf("must contain '%s'", substring)
        }
        return nil
    }
}

// 17. Person struct for sorting
type Person struct {
    Name   string
    Age    int
    Salary int
}

// 18. Prime generator
func makePrimeGenerator() func() int {
    primes := []int{2}
    index := 0
    
    return func() int {
        if index < len(primes) {
            result := primes[index]
            index++
            return result
        }
        
        // Generate next prime
        candidate := primes[len(primes)-1] + 1
        for {
            isPrime := true
            for _, p := range primes {
                if p*p > candidate {
                    break
                }
                if candidate%p == 0 {
                    isPrime = false
                    break
                }
            }
            if isPrime {
                primes = append(primes, candidate)
                index++
                return candidate
            }
            candidate++
        }
    }
}

// 19. Expensive function and cache
func makeExpensiveFunction() func(int) int {
    return func(n int) int {
        // Simulate expensive computation
        time.Sleep(10 * time.Millisecond)
        return n * n * n
    }
}

func makeCache(fn func(int) int) func(int) int {
    cache := make(map[int]int)
    
    return func(n int) int {
        if result, exists := cache[n]; exists {
            fmt.Print("(cached) ")
            return result
        }
        
        result := fn(n)
        cache[n] = result
        return result
    }
}

// 20. Debouncer
func makeDebouncer(fn func(), delay time.Duration) func() {
    var timer *time.Timer
    
    return func() {
        if timer != nil {
            timer.Stop()
        }
        
        timer = time.AfterFunc(delay, fn)
    }
}