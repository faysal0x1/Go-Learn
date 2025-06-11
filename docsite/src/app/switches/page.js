import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Switches() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Switch Statements in Go</h2>
            <p className="mb-4">
                Switch statements in Go provide a cleaner way to write multiple if-else conditions. Go's switch
                statement is more flexible than in many other languages, with features like expression switches,
                type switches, and the ability to use multiple values in a case.
            </p>
            <p className="mb-4">
                Key features of switch statements in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Basic value matching</li>
                <li>Multiple values in cases</li>
                <li>Expression switches</li>
                <li>Type switches</li>
                <li>Fallthrough behavior</li>
                <li>Initialization statements</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Switch Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Switch</h3>
                    <p className="mb-4">
                        Simple switch statement matching a value:
                    </p>
                    <CodeBlock code={`day := 3
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
case 4:
    fmt.Println("Thursday")
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
    fmt.Println("Sunday")
default:
    fmt.Println("Invalid day")
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Multiple Values in Case</h3>
                    <p className="mb-4">
                        Switch statement with multiple values in a case:
                    </p>
                    <CodeBlock code={`month := "Feb"
switch month {
case "Jan", "Mar", "May", "Jul", "Aug", "Oct", "Dec":
    fmt.Println("This month has 31 days")
case "Apr", "Jun", "Sep", "Nov":
    fmt.Println("This month has 30 days")
case "Feb":
    fmt.Println("This month has 28 or 29 days")
default:
    fmt.Println("Invalid month")
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Expression Switch</h3>
                    <p className="mb-4">
                        Switch statement without an expression (like if-else chains):
                    </p>
                    <CodeBlock code={`score := 85
switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80:
    fmt.Println("Grade: B")
case score >= 70:
    fmt.Println("Grade: C")
case score >= 60:
    fmt.Println("Grade: D")
default:
    fmt.Println("Grade: F")
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Type Switch</h3>
                    <p className="mb-4">
                        Switch statement for type checking:
                    </p>
                    <CodeBlock code={`var i interface{} = "Hello"
switch v := i.(type) {
case int:
    fmt.Printf("Twice %v is %v\\n", v, v*2)
case string:
    fmt.Printf("String length of %v is %v\\n", v, len(v))
case bool:
    fmt.Printf("Boolean complement of %v is %v\\n", v, !v)
default:
    fmt.Printf("I don't know about type %T\\n", v)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Switch with Fallthrough</h3>
                    <p className="mb-4">
                        Using fallthrough to continue execution to the next case:
                    </p>
                    <CodeBlock code={`num := 5
switch num {
case 5:
    fmt.Println("Five")
    fallthrough
case 6:
    fmt.Println("Six")
    fallthrough
case 7:
    fmt.Println("Seven")
default:
    fmt.Println("Other number")
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
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Day of Week</h3>
                    <p className="mb-4">
                        Create a program that uses a switch statement to determine if a given day is a weekday or weekend.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Day: Monday
                            Type: Weekday
                            Status: Working day</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Grade Calculator</h3>
                    <p className="mb-4">
                        Create a program that uses a switch statement to calculate letter grades based on numerical scores.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Score: 85
                            Grade: B
                            Status: Good</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Type Handler</h3>
                    <p className="mb-4">
                        Create a program that uses a type switch to handle different types of data and perform appropriate operations.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Value: 42
                            Type: int
                            Operation: Doubled = 84</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different switch patterns in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Try different switch patterns here
    day := 3
    switch day {
    case 1, 2, 3, 4, 5:
        fmt.Println("Weekday")
    case 6, 7:
        fmt.Println("Weekend")
    default:
        fmt.Println("Invalid day")
    }
    
    // Try expression switch
    score := 85
    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 80:
        fmt.Println("Grade: B")
    default:
        fmt.Println("Grade: C or below")
    }
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Weekday
                        Grade: B</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Switch Statements"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 