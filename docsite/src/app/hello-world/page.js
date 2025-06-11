import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function HelloWorld() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Introduction to Go</h2>
            <p className="mb-4">
                Go (or Golang) is a statically typed, compiled programming language designed at Google.
                It's known for its simplicity, efficiency, and strong support for concurrent programming.
            </p>
            <p className="mb-4">
                In this section, we'll create our first Go program - the traditional "Hello, World!"
                program. This will help you understand the basic structure of a Go program and how to run it.
            </p>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Hello World Example</h2>
            <p className="mb-4">
                Here's a simple Hello World program in Go:
            </p>
            <CodeBlock code={`package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`} />
            <p className="mt-4">
                Let's break down this program:
            </p>
            <ul className="list-disc pl-6 mt-2 space-y-2">
                <li><code>package main</code> - Every Go program starts with a package declaration</li>
                <li><code>import "fmt"</code> - Imports the format package for printing</li>
                <li><code>func main()</code> - The entry point of the program</li>
                <li><code>fmt.Println()</code> - Prints text to the console</li>
            </ul>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Custom Greeting</h3>
                    <p className="mb-4">
                        Modify the Hello World program to print a custom greeting message.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Welcome to Go Programming!</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Multiple Lines</h3>
                    <p className="mb-4">
                        Create a program that prints multiple lines of text.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Line 1
                            Line 2
                            Line 3</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out the Hello World program in the interactive playground below.
                Modify the code and see the results instantly!
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Try modifying this line!
    fmt.Println("Hello, World!")
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Hello, World!</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Hello World"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 