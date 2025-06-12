import CodeBlock from '../../components/CodeBlock';

export default function Examples() {
    return (
        <>
            <h2 className="text-2xl font-bold mb-4">Goroutine Examples</h2>
            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Goroutine</h3>
                    <CodeBlock code={`go sayHello("World")
func sayHello(name string) { fmt.Printf("Hello, %s!\\n", name) }`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Multiple Goroutines</h3>
                    <CodeBlock code={`for i := 1; i <= 3; i++ {
  go func(id int) {
    fmt.Printf("Goroutine %d is running\\n", id)
  }(i)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Channels for Communication</h3>
                    <CodeBlock code={`ch := make(chan string)
go func() {
  time.Sleep(50 * time.Millisecond)
  ch <- "Hello from goroutine!"
}()
message := <-ch`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. WaitGroup Synchronization</h3>
                    <CodeBlock code={`var wg sync.WaitGroup
for i := 1; i <= 3; i++ {
  wg.Add(1)
  go worker(i, &wg)
}
wg.Wait()`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Producer-Consumer Pattern</h3>
                    <CodeBlock code={`jobs := make(chan int, 5)
results := make(chan int, 5)

// Start workers
for w := 1; w <= 3; w++ {
  go consumerWorker(w, jobs, results)
}

// Send jobs
for j := 1; j <= 5; j++ {
  jobs <- j
}
close(jobs)

// Collect results
for r := 1; r <= 5; r++ {
  result := <-results
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">6. Fan-out, Fan-in Pattern</h3>
                    <CodeBlock code={`input := generateNumbers(1, 5)
worker1 := squareNumbers(input)
worker2 := squareNumbers(input)
output := fanIn(worker1, worker2)

for result := range output {
  fmt.Printf("Squared: %d\\n", result)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">7. Context Cancellation</h3>
                    <CodeBlock code={`ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
defer cancel()

go longRunningTask(ctx)`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">8. Worker Pool</h3>
                    <CodeBlock code={`jobQueue := make(chan Job, 10)
resultQueue := make(chan Result, 10)

pool := NewWorkerPool(3, jobQueue, resultQueue)
pool.Start()

// Send jobs
for i := 1; i <= 5; i++ {
  jobQueue <- Job{ID: i, Data: i * 2}
}
close(jobQueue)

// Collect results
for i := 1; i <= 5; i++ {
  result := <-resultQueue
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">9. Graceful Shutdown</h3>
                    <CodeBlock code={`server := NewServer()
go server.Start()
time.Sleep(100 * time.Millisecond)
server.Shutdown()`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">10. Rate Limiting</h3>
                    <CodeBlock code={`rateLimiter := time.Tick(100 * time.Millisecond)
for i := 1; i <= 5; i++ {
  <-rateLimiter
  fmt.Printf("Request %d processed\\n", i)
}`} />
                </div>
            </div>
        </>
    );
}