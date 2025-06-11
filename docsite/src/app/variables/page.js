import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Variables() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Variables in Go</h2>
            <p className="mb-4">
                Variables in Go are used to store values of different types. Go provides several ways to declare and initialize variables,
                each with its own use case and benefits.
            </p>
            <p className="mb-4">
                Key features of variables in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Static typing with type inference</li>
                <li>Multiple declaration styles</li>
                <li>Zero values for uninitialized variables</li>
                <li>Block scoping and shadowing</li>
                <li>Type conversion capabilities</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Variable Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Variable Declaration</h3>
                    <p className="mb-4">
                        Variables can be declared using the <code>var</code> keyword:
                    </p>
                    <CodeBlock code={`var age int
age = 30
fmt.Println("Age:", age)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Declaration with Initialization</h3>
                    <p className="mb-4">
                        Variables can be declared and initialized in one statement:
                    </p>
                    <CodeBlock code={`var name string = "John Doe"
fmt.Println("Name:", name)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Type Inference</h3>
                    <p className="mb-4">
                        Go can infer the type based on the initial value:
                    </p>
                    <CodeBlock code={`var salary = 50000.50
fmt.Printf("Salary: %.2f (Type: %T)\\n", salary, salary)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Short Variable Declaration</h3>
                    <p className="mb-4">
                        Inside functions, you can use the <code>:=</code> operator for short declaration:
                    </p>
                    <CodeBlock code={`count := 10
fmt.Println("Count:", count)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Multiple Variable Declaration</h3>
                    <p className="mb-4">
                        You can declare multiple variables at once:
                    </p>
                    <CodeBlock code={`var a, b, c int = 1, 2, 3
fmt.Println("Multiple values:", a, b, c)`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Variable Types</h3>
                    <p className="mb-4">
                        Create a program that declares variables of different types (int, float64, string, bool) and prints their values and types.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Number: 42 (Type: int)
                            Price: 19.99 (Type: float64)
                            Name: Alice (Type: string)
                            IsActive: true (Type: bool)</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Type Conversion</h3>
                    <p className="mb-4">
                        Write a program that demonstrates type conversion between different numeric types.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Original: 42 (int)
                            As float64: 42.000000
                            As uint: 42</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Multiple Variables</h3>
                    <p className="mb-4">
                        Create a program that declares and initializes multiple variables of different types in a single declaration block.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Name: John
                            Age: 30
                            Score: 95.5
                            IsStudent: true</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different variable declarations and operations in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Try different variable declarations here
    var name string = "John"
    age := 30
    score := 95.5
    
    fmt.Printf("Name: %s\\n", name)
    fmt.Printf("Age: %d\\n", age)
    fmt.Printf("Score: %.1f\\n", score)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Name: John
                        Age: 30
                        Score: 95.5</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Variables"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 