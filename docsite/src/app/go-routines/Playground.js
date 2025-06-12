import Playgrounds from '../../components/Playground';

export default function Playground() {
    return (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">Try out goroutine patterns in the interactive playground below.</p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <Playgrounds initialCode={`package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  // Basic goroutine example
  go func() {
    fmt.Println("Hello from goroutine!")
  }()
  
  // WaitGroup example
  var wg sync.WaitGroup
  wg.Add(1)
  
  go func() {
    defer wg.Done()
    fmt.Println("Worker goroutine running")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Worker goroutine done")
  }()
  
  // Channel example
  ch := make(chan string)
  go func() {
    time.Sleep(50 * time.Millisecond)
    ch <- "Message from goroutine"
  }()
  
  // Main thread waits for goroutines
  wg.Wait()
  msg := <-ch
  fmt.Println("Received:", msg)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Expected Output:</h3>
                    <pre className="text-sm">
                        {`Hello from goroutine!
Worker goroutine running
Worker goroutine done
Received: Message from goroutine`}
                    </pre>
                </div>
            </div>

            <div className="mt-8 bg-gray-50 p-6 rounded-lg">
                <h3 className="text-xl font-semibold mb-4">Try These Modifications:</h3>
                <ul className="list-disc pl-6 space-y-2">
                    <li>Add more goroutines and see how they execute</li>
                    <li>Experiment with buffered channels</li>
                    <li>Add a select statement to handle multiple channels</li>
                    <li>Try implementing a simple worker pool pattern</li>
                    <li>Add error handling with an error channel</li>
                </ul>
            </div>
        </>
    );
}