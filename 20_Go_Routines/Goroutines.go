package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== Go Goroutines Examples ===")

	// 1. Basic goroutine
	fmt.Println("1. Basic goroutine:")
	go sayHello("World")
	time.Sleep(100 * time.Millisecond) // Wait for goroutine to complete
	fmt.Println("Main function continues")

	// 2. Multiple goroutines
	fmt.Println("2. Multiple goroutines:")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d is running\n", id)
		}(i) // Pass i as parameter to avoid closure issue
	}
	time.Sleep(100 * time.Millisecond)

	// 3. Goroutines with channels for communication
	fmt.Println("\n3. Goroutines with channels:")
	ch := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "Hello from goroutine!"
	}()

	message := <-ch
	fmt.Printf("Received: %s\n", message)

	// 4. Buffered channels
	fmt.Println("\n4. Buffered channels:")
	bufferedCh := make(chan int, 3)

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Sending %d\n", i)
			bufferedCh <- i
			time.Sleep(10 * time.Millisecond)
		}
		close(bufferedCh)
	}()

	for value := range bufferedCh {
		fmt.Printf("Received %d\n", value)
	}

	// 5. WaitGroup for synchronization
	fmt.Println("\n5. WaitGroup synchronization:")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed")

	// 6. Select statement with channels
	fmt.Println("\n6. Select statement:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Printf("Received %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received %s\n", msg2)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// 7. Producer-Consumer pattern
	fmt.Println("\n7. Producer-Consumer pattern:")
	jobs := make(chan int, 5)
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
		fmt.Printf("Result: %d\n", result)
	}

	// 8. Fan-out, Fan-in pattern
	fmt.Println("\n8. Fan-out, Fan-in pattern:")
	input := generateNumbers(1, 5)

	// Fan-out to multiple workers
	worker1 := squareNumbers(input)
	worker2 := squareNumbers(input)
	worker3 := squareNumbers(input)

	// Fan-in results
	output := fanIn(worker1, worker2, worker3)

	for result := range output {
		fmt.Printf("Squared: %d\n", result)
	}

	// 9. Context for cancellation
	fmt.Println("\n9. Context cancellation:")
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	go longRunningTask(ctx)
	time.Sleep(300 * time.Millisecond)

	// 10. Mutex for shared state
	fmt.Println("\n10. Mutex for shared state:")
	counter := &Counter{}
	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg2.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())

	// 11. Atomic operations
	fmt.Println("\n11. Atomic operations:")
	var atomicCounter int64
	var wg3 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}

	wg3.Wait()
	fmt.Printf("Atomic counter value: %d\n", atomic.LoadInt64(&atomicCounter))

	// 12. Rate limiting
	fmt.Println("\n12. Rate limiting:")
	rateLimiter := time.Tick(100 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		<-rateLimiter
		fmt.Printf("Request %d processed at %s\n", i, time.Now().Format("15:04:05.000"))
	}

	// 13. Pipeline pattern
	fmt.Println("\n13. Pipeline pattern:")
	numbers := make(chan int)
	squared := make(chan int)
	doubled := make(chan int)

	// Stage 1: Generate numbers
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
	}()

	// Stage 2: Square numbers
	go func() {
		defer close(squared)
		for num := range numbers {
			squared <- num * num
		}
	}()

	// Stage 3: Double numbers
	go func() {
		defer close(doubled)
		for num := range squared {
			doubled <- num * 2
		}
	}()

	// Consume final results
	for result := range doubled {
		fmt.Printf("Pipeline result: %d\n", result)
	}

	// 14. Worker pool
	fmt.Println("\n14. Worker pool:")
	jobQueue := make(chan Job, 10)
	resultQueue := make(chan Result, 10)

	// Start worker pool
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
		fmt.Printf("Job %d result: %d\n", result.JobID, result.Value)
	}

	// 15. Graceful shutdown
	fmt.Println("\n15. Graceful shutdown:")
	server := NewServer()

	go server.Start()
	time.Sleep(100 * time.Millisecond)

	server.Shutdown()
	fmt.Println("Server shutdown complete")

	// 16. Error handling in goroutines
	fmt.Println("\n16. Error handling:")
	errCh := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		go riskyOperation(i, errCh)
	}

	for i := 1; i <= 3; i++ {
		if err := <-errCh; err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Operation %d succeeded\n", i)
		}
	}

	// 17. Channel directions
	fmt.Println("\n17. Channel directions:")
	data := make(chan string, 2)

	go sendData(data)
	go receiveData(data)

	time.Sleep(100 * time.Millisecond)

	// 18. Goroutine leaks prevention
	fmt.Println("\n18. Preventing goroutine leaks:")
	done := make(chan bool)

	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Printf("Tick at %s\n", time.Now().Format("15:04:05.000"))
			case <-done:
				fmt.Println("Ticker stopped")
				return
			}
		}
	}()

	time.Sleep(200 * time.Millisecond)
	done <- true
	time.Sleep(50 * time.Millisecond)

	// 19. Goroutine monitoring
	fmt.Println("\n19. Goroutine monitoring:")
	fmt.Printf("Number of goroutines before: %d\n", runtime.NumGoroutine())

	var wg4 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	fmt.Printf("Number of goroutines during: %d\n", runtime.NumGoroutine())
	wg4.Wait()
	time.Sleep(10 * time.Millisecond) // Allow goroutines to cleanup
	fmt.Printf("Number of goroutines after: %d\n", runtime.NumGoroutine())

	// 20. Advanced patterns
	fmt.Println("\n20. Advanced patterns:")

	// Semaphore pattern for limiting concurrency
	semaphore := make(chan struct{}, 2) // Allow only 2 concurrent operations
	var wg5 sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg5.Add(1)
		go func(id int) {
			defer wg5.Done()

			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			fmt.Printf("Limited operation %d started\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Limited operation %d finished\n", id)
		}(i)
	}

	wg5.Wait()
	fmt.Println("All limited operations completed")
}

// Helper functions and types

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func consumerWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(50 * time.Millisecond)
		results <- job * 2
	}
}

// Fan-out, Fan-in pattern
func generateNumbers(start, end int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i <= end; i++ {
			out <- i
		}
	}()
	return out
}

func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for value := range c {
				out <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Task completed normally")
	case <-ctx.Done():
		fmt.Printf("Task cancelled: %v\n", ctx.Err())
	}
}

// Thread-safe counter
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Worker pool
type Job struct {
	ID   int
	Data int
}

type Result struct {
	JobID int
	Value int
}

type WorkerPool struct {
	workers     int
	jobQueue    <-chan Job
	resultQueue chan<- Result
	wg          sync.WaitGroup
}

func NewWorkerPool(workers int, jobQueue <-chan Job, resultQueue chan<- Result) *WorkerPool {
	return &WorkerPool{
		workers:     workers,
		jobQueue:    jobQueue,
		resultQueue: resultQueue,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}

	go func() {
		wp.wg.Wait()
		close(wp.resultQueue)
	}()
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobQueue {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(50 * time.Millisecond)
		wp.resultQueue <- Result{JobID: job.ID, Value: job.Data * job.Data}
	}
}

// Server with graceful shutdown
type Server struct {
	quit chan struct{}
	done chan struct{}
}

func NewServer() *Server {
	return &Server{
		quit: make(chan struct{}),
		done: make(chan struct{}),
	}
}

func (s *Server) Start() {
	fmt.Println("Server starting...")

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Server tick at %s\n", time.Now().Format("15:04:05.000"))
		case <-s.quit:
			fmt.Println("Server shutting down...")
			close(s.done)
			return
		}
	}
}

func (s *Server) Shutdown() {
	close(s.quit)
	<-s.done
}

func riskyOperation(id int, errCh chan<- error) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	if rand.Float32() < 0.3 { // 30% chance of error
		errCh <- fmt.Errorf("operation %d failed", id)
	} else {
		errCh <- nil
	}
}

// Channel directions
func sendData(ch chan<- string) {
	ch <- "data1"
	ch <- "data2"
	close(ch)
}

func receiveData(ch <-chan string) {
	for data := range ch {
		fmt.Printf("Received: %s\n", data)
	}
}
