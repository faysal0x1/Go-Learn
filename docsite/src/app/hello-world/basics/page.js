import DocLayout from '@/components/DocLayout'
import CodeBlock from '@/components/CodeBlock'

export const metadata = {
    title: 'Hello World - Go Programming Documentation',
    description: 'Learn how to write your first Go program and understand the basics of Go programming language.',
}

export default function HelloWorldBasics() {
    return (
        <DocLayout
            title="Hello World"
            description="Learn how to write your first Go program and understand the basics of Go programming language."
        >
            <h2>Your First Go Program</h2>
            <p>
                Let's start with a simple "Hello, World!" program. This is a traditional first program
                that demonstrates the basic syntax and structure of Go.
            </p>

            <CodeBlock
                code={`package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`}
                showOutput={true}
                output="Hello, World!"
            />

            <h2>Understanding the Code</h2>
            <p>Let's break down each part of the program:</p>

            <h3>1. Package Declaration</h3>
            <p>
                Every Go program starts with a package declaration. The <code>main</code> package is
                special - it's the entry point for an executable program.
            </p>

            <h3>2. Import Statement</h3>
            <p>
                The <code>import</code> statement brings in the <code>fmt</code> package, which
                contains functions for formatted I/O operations.
            </p>

            <h3>3. Main Function</h3>
            <p>
                The <code>main</code> function is where program execution begins. It's a special
                function that must be present in the main package.
            </p>

            <h2>Running the Program</h2>
            <p>To run this program:</p>
            <ol>
                <li>Save the code in a file with a <code>.go</code> extension (e.g., <code>hello.go</code>)</li>
                <li>Open a terminal in the directory containing the file</li>
                <li>Run the command: <code>go run hello.go</code></li>
            </ol>

            <h2>Key Concepts</h2>
            <ul>
                <li>Go programs are organized into packages</li>
                <li>The <code>main</code> package is required for executable programs</li>
                <li>Functions are declared using the <code>func</code> keyword</li>
                <li>Go uses curly braces <code>{ }</code> to define code blocks</li>
                <li>Statements end with a semicolon (though it's usually omitted)</li>
            </ul>

            <div className="mt-8 p-4 bg-blue-50 dark:bg-blue-900 rounded-lg">
                <h3 className="text-lg font-semibold mb-2 text-blue-800 dark:text-blue-200">
                    Next Steps
                </h3>
                <p className="text-blue-700 dark:text-blue-300">
                    Now that you've written your first Go program, try modifying it to print different
                    messages or explore more examples in the Examples section.
                </p>
            </div>
        </DocLayout>
    )
} 