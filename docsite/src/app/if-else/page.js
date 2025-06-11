import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function IfElse() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">If-Else Statements in Go</h2>
            <p className="mb-4">
                If-else statements in Go are used for conditional execution of code blocks. Go's if-else syntax
                is similar to other C-like languages but has some unique features, such as the ability to include
                initialization statements within the if condition.
            </p>
            <p className="mb-4">
                Key features of if-else in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Basic if-else conditions</li>
                <li>If-else if-else chains</li>
                <li>Initialization statements in if conditions</li>
                <li>Logical operators (&&, ||)</li>
                <li>Nested if statements</li>
                <li>Error handling patterns</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">If-Else Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic If-Else</h3>
                    <p className="mb-4">
                        Simple if-else statement with a condition:
                    </p>
                    <CodeBlock code={`x := 10
if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is not greater than 5")
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. If-Else If-Else Chain</h3>
                    <p className="mb-4">
                        Multiple conditions using if-else if-else:
                    </p>
                    <CodeBlock code={`score := 85
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else if score >= 60 {
    fmt.Println("Grade: D")
} else {
    fmt.Println("Grade: F")
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. If with Initialization</h3>
                    <p className="mb-4">
                        Go allows initialization statements within if conditions:
                    </p>
                    <CodeBlock code={`if num := 9; num < 10 {
    fmt.Println("num is less than 10")
} else {
    fmt.Println("num is not less than 10")
}
// Note: num is not accessible here`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Logical Operators</h3>
                    <p className="mb-4">
                        Using logical operators in conditions:
                    </p>
                    <CodeBlock code={`a, b := 5, 10
if a > 0 && b > 0 {
    fmt.Println("Both a and b are positive")
}

if a > 10 || b > 5 {
    fmt.Println("Either a is greater than 10 or b is greater than 5")
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Error Handling Pattern</h3>
                    <p className="mb-4">
                        Common Go pattern for error handling:
                    </p>
                    <CodeBlock code={`str := "42"
if val, err := strconv.Atoi(str); err == nil {
    fmt.Printf("String converted to integer: %d\\n", val)
} else {
    fmt.Printf("Conversion error: %s\\n", err)
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
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Number Classification</h3>
                    <p className="mb-4">
                        Create a program that classifies a number as positive, negative, or zero, and also determines if it's even or odd.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Number: 7
                            Classification: Positive
                            Type: Odd</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Age Verification</h3>
                    <p className="mb-4">
                        Create a program that verifies if a person is eligible to vote based on their age and citizenship status.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Age: 20
                            Citizen: true
                            Status: Eligible to vote</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Password Strength Checker</h3>
                    <p className="mb-4">
                        Create a program that checks the strength of a password based on its length and character types.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Password: "GoLang123!"
                            Strength: Strong
                            Reason: Contains uppercase, lowercase, numbers, and special characters</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different if-else patterns in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Try different if-else patterns here
    age := 20
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
    
    // Try if-else if chain
    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C or below")
    }
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Adult
                        Grade: B</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="If-Else Statements"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 