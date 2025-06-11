import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function GoRoutines() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Goroutines in Go</h2>
            <p className="mb-4">
                Goroutines are lightweight threads managed by the Go runtime. They enable concurrent programming, allowing you to run multiple functions independently and efficiently.
            </p>
            <p className="mb-4">Key features of goroutines in Go:</p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Start with the <code>go</code> keyword</li>
                <li>Efficient and lightweight (thousands per program)</li>
                <li>Communicate via channels</li>
                <li>Synchronization with WaitGroup, Mutex, and atomic operations</li>
                <li>Support for context cancellation and timeouts</li>
                <li>Patterns: worker pool, pipeline, fan-out/fan-in, producer-consumer</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Goroutine Examples</h2>
            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Goroutine</h3>
                    <CodeBlock code={`go sayHello("World")
func sayHello(name string) { fmt.Printf("Hello, %s!\n", name) }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Goroutines with Channels</h3>
                    <CodeBlock code={`ch := make(chan string)
go func() { ch <- "Hello from goroutine!" }()
msg := <-ch
fmt.Println(msg)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Buffered Channels</h3>
                    <CodeBlock code={`bufferedCh := make(chan int, 3)
go func() { bufferedCh <- 1; bufferedCh <- 2; close(bufferedCh) }()
for v := range bufferedCh { fmt.Println(v) }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">4. WaitGroup Synchronization</h3>
                    <CodeBlock code={`var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); fmt.Println("Done!") }()
wg.Wait()`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Select Statement</h3>
                    <CodeBlock code={`ch1 := make(chan string)
ch2 := make(chan string)
go func() { ch1 <- "one" }()
go func() { ch2 <- "two" }()
select {
case msg1 := <-ch1: fmt.Println(msg1)
case msg2 := <-ch2: fmt.Println(msg2)
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">6. Producer-Consumer Pattern</h3>
                    <CodeBlock code={`jobs := make(chan int, 5)
results := make(chan int, 5)
go func() { for j := range jobs { results <- j * 2 } }()
for j := 1; j <= 5; j++ { jobs <- j }
close(jobs)
for r := 1; r <= 5; r++ { fmt.Println(<-results) }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">7. Context Cancellation</h3>
                    <CodeBlock code={`ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
go func() {
  select {
  case <-ctx.Done(): fmt.Println("Cancelled!")
  }
}()
time.Sleep(2 * time.Second)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">8. Mutex for Shared State</h3>
                    <CodeBlock code={`var mu sync.Mutex
counter := 0
go func() { mu.Lock(); counter++; mu.Unlock() }()`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">9. Worker Pool</h3>
                    <CodeBlock code={`type Job struct { ID int }
type Result struct { JobID, Value int }
jobQueue := make(chan Job, 10)
resultQueue := make(chan Result, 10)
for w := 1; w <= 3; w++ {
  go func() {
    for job := range jobQueue {
      resultQueue <- Result{JobID: job.ID, Value: job.ID * 2}
    }
  }()
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">10. Graceful Shutdown</h3>
                    <CodeBlock code={`type Server struct { quit chan struct{} }
func (s *Server) Start() { <-s.quit }
func (s *Server) Shutdown() { close(s.quit) }
server := &Server{quit: make(chan struct{})}
go server.Start()
time.Sleep(100 * time.Millisecond)
server.Shutdown()`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Parallel Sum</h3>
                    <p className="mb-4">Write a function that splits a slice into N parts and sums each part in a separate goroutine, then combines the results.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Pipeline Pattern</h3>
                    <p className="mb-4">Implement a pipeline with three stages: generate numbers, square them, and double them, using goroutines and channels.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Graceful Shutdown</h3>
                    <p className="mb-4">Create a server that can be started and gracefully shut down using channels and goroutines.</p>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">Try out goroutine patterns in the interactive playground below.</p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main
import (
  "fmt"
  "sync"
)
func main() {
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    fmt.Println("Hello from goroutine!")
  }()
  wg.Wait()
  fmt.Println("Main done.")
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Hello from goroutine!
                        Main done.</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Go Routines"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 