import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Constants() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Constants in Go</h2>
            <p className="mb-4">
                Constants in Go are fixed values that cannot be changed after they are declared. They are useful for defining
                values that should remain constant throughout the program's execution.
            </p>
            <p className="mb-4">
                Key features of constants in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Must be initialized at declaration time</li>
                <li>Can be of any basic type</li>
                <li>Can be declared individually or in groups</li>
                <li>Cannot be modified after declaration</li>
                <li>Can be used in expressions and calculations</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Constant Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Single Constant Declaration</h3>
                    <p className="mb-4">
                        Constants can be declared individually using the <code>const</code> keyword:
                    </p>
                    <CodeBlock code={`const name = "John Doe"
const age = 30

fmt.Println("Name:", name)
fmt.Println("Age:", age)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Multiple Constants Declaration</h3>
                    <p className="mb-4">
                        Constants can be declared in a group using parentheses:
                    </p>
                    <CodeBlock code={`const (
    port = 202
    host = "localhost"
)

fmt.Println("Port:", port)
fmt.Println("Host:", host)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Boolean Constants</h3>
                    <p className="mb-4">
                        Constants can be of boolean type:
                    </p>
                    <CodeBlock code={`const isEmployed = true
fmt.Println("Is Employed:", isEmployed)`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Basic Constants</h3>
                    <p className="mb-4">
                        Create a program that declares and uses constants for a simple configuration.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Database Configuration:
                            Host: localhost
                            Port: 5432
                            User: admin
                            SSL Enabled: true</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Mathematical Constants</h3>
                    <p className="mb-4">
                        Create a program that uses mathematical constants to calculate the area and circumference of a circle.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Circle with radius 5:
                            Area: 78.54
                            Circumference: 31.42</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Status Codes</h3>
                    <p className="mb-4">
                        Create a program that defines HTTP status codes as constants and uses them in a simple response function.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Response Status: 200 OK
                            Error Status: 404 Not Found
                            Server Error: 500 Internal Server Error</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different constant declarations and their usage in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Try different constant declarations here
    const (
        name = "John"
        age = 30
        isActive = true
    )
    
    fmt.Printf("Name: %s\\n", name)
    fmt.Printf("Age: %d\\n", age)
    fmt.Printf("Is Active: %v\\n", isActive)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Name: John
                        Age: 30
                        Is Active: true</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Constants"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 