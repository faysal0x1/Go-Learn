package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go Channels Examples ===\n")

	// 1. Basic channel operations
	fmt.Println("1. Basic channel operations:")
	ch := make(chan string)

	go func() {
		ch <- "Hello, Channel!"
	}()

	message := <-ch
	fmt.Printf("Received: %s\n", message)

	// 2. Buffered channels
	fmt.Println("\n2. Buffered channels:")
	bufferedCh := make(chan int, 3)

	// Can send up to buffer size without blocking
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	fmt.Printf("Buffer length: %d, capacity: %d\n", len(bufferedCh), cap(bufferedCh))

	// Receive values
	for i := 0; i < 3; i++ {
		value := <-bufferedCh
		fmt.Printf("Received: %d\n", value)
	}

	// 3. Channel closing and range
	fmt.Println("\n3. Channel closing and range:")
	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // Important: close channel when done
	}()

	// Range over channel until closed
	for num := range numbers {
		fmt.Printf("Number: %d\n", num)
	}

	// 4. Select statement
	fmt.Println("\n4. Select statement:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Printf("Received %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received %s\n", msg2)
	}

	// 5. Non-blocking channel operations
	fmt.Println("\n5. Non-blocking operations:")
	nonBlockingCh := make(chan int, 1)

	// Non-blocking send
	select {
	case nonBlockingCh <- 42:
		fmt.Println("Sent 42")
	default:
		fmt.Println("Could not send")
	}

	// Non-blocking receive
	select {
	case value := <-nonBlockingCh:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("Could not receive")
	}

	// 6. Channel directions (send-only, receive-only)
	fmt.Println("\n6. Channel directions:")
	dataCh := make(chan string, 2)

	go sendOnly(dataCh)
	go receiveOnly(dataCh)

	time.Sleep(100 * time.Millisecond)

	// 7. Worker pool with channels
	fmt.Println("\n7. Worker pool:")
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- Job{ID: j, Data: j * 2}
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Job %d result: %d\n", result.JobID, result.Value)
	}

	// 8. Fan-out pattern
	fmt.Println("\n8. Fan-out pattern:")
	input := make(chan int)
	output1 := make(chan int)
	output2 := make(chan int)
	output3 := make(chan int)

	// Fan-out: one input to multiple outputs
	go fanOut(input, output1, output2, output3)

	// Send data
	go func() {
		for i := 1; i <= 6; i++ {
			input <- i
		}
		close(input)
	}()

	// Collect from all outputs
	var wg sync.WaitGroup
	wg.Add(3)

	go collectResults("Output1", output1, &wg)
	go collectResults("Output2", output2, &wg)
	go collectResults("Output3", output3, &wg)

	wg.Wait()

	// 9. Fan-in pattern
	fmt.Println("\n9. Fan-in pattern:")
	in1 := make(chan string)
	in2 := make(chan string)
	in3 := make(chan string)

	merged := fanIn(in1, in2, in3)

	// Send data to different inputs
	go func() {
		in1 <- "from input 1"
		in2 <- "from input 2"
		in3 <- "from input 3"
		close(in1)
		close(in2)
		close(in3)
	}()

	// Receive merged data
	for msg := range merged {
		fmt.Printf("Merged: %s\n", msg)
	}

	// 10. Pipeline pattern
	fmt.Println("\n10. Pipeline pattern:")
	// Stage 1: Generate numbers
	generator := generateNumbers(1, 5)

	// Stage 2: Square numbers
	squarer := squareNumbers(generator)

	// Stage 3: Add prefix
	formatter := addPrefix(squarer, "Square:")

	// Consume final output
	for result := range formatter {
		fmt.Println(result)
	}

	// 11. Rate limiting with channels
	fmt.Println("\n11. Rate limiting:")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Rate limiter: allow one request per 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter // Wait for rate limiter
		fmt.Printf("Request %d processed at %s\n", req, time.Now().Format("15:04:05"))
	}

	// 12. Timeout with channels
	fmt.Println("\n12. Timeout operations:")
	slowCh := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		slowCh <- "slow response"
	}()

	select {
	case response := <-slowCh:
		fmt.Printf("Received: %s\n", response)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation timed out")
	}

	// 13. Channel multiplexing
	fmt.Println("\n13. Channel multiplexing:")
	mux := NewChannelMux()

	// Add channels to multiplexer
	ch_a := make(chan string, 1)
	ch_b := make(chan string, 1)

	mux.Add("channel_a", ch_a)
	mux.Add("channel_b", ch_b)

	// Send data
	ch_a <- "message from A"
	ch_b <- "message from B"

	// Receive multiplexed data
	for i := 0; i < 2; i++ {
		msg := mux.Receive()
		fmt.Printf("Mux received: %s from %s\n", msg.Data, msg.Channel)
	}

	// 14. Graceful shutdown with channels
	fmt.Println("\n14. Graceful shutdown:")
	server := NewServer()

	go server.Start()
	time.Sleep(200 * time.Millisecond)

	server.Stop()
	fmt.Println("Server stopped gracefully")

	// 15. Channel as semaphore
	fmt.Println("\n15. Channel semaphore:")
	semaphore := make(chan struct{}, 2) // Allow 2 concurrent operations

	var wg2 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()

			semaphore <- struct{}{} // Acquire
			fmt.Printf("Operation %d started\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Operation %d finished\n", id)
			<-semaphore // Release
		}(i)
	}

	wg2.Wait()

	// 16. Error propagation through channels
	fmt.Println("\n16. Error propagation:")
	errorCh := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		go riskyOperation(i, errorCh)
	}

	for i := 1; i <= 3; i++ {
		if err := <-errorCh; err != nil {
			fmt.Printf("Error from operation: %v\n", err)
		} else {
			fmt.Printf("Operation %d succeeded\n", i)
		}
	}

	// 17. Channel orchestration
	fmt.Println("\n17. Channel orchestration:")
	orchestrator := NewOrchestrator()

	// Start multiple workers
	for i := 1; i <= 3; i++ {
		orchestrator.AddWorker(fmt.Sprintf("worker-%d", i))
	}

	// Send tasks
	for i := 1; i <= 5; i++ {
		orchestrator.SendTask(Task{ID: i, Data: fmt.Sprintf("task-%d", i)})
	}

	orchestrator.Stop()

	// 18. Broadcast pattern
	fmt.Println("\n18. Broadcast pattern:")
	broadcaster := NewBroadcaster()

	// Subscribe multiple listeners
	listener1 := broadcaster.Subscribe("listener1")
	listener2 := broadcaster.Subscribe("listener2")
	listener3 := broadcaster.Subscribe("listener3")

	// Start listeners
	go listenForBroadcasts("Listener1", listener1)
	go listenForBroadcasts("Listener2", listener2)
	go listenForBroadcasts("Listener3", listener3)

	// Broadcast messages
	broadcaster.Broadcast("Hello everyone!")
	broadcaster.Broadcast("Important update!")

	time.Sleep(100 * time.Millisecond)
	broadcaster.Close()

	// 19. Channel with context
	fmt.Println("\n19. Channel with context:")
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	resultCh := make(chan string)
	go longRunningTask(ctx, resultCh)

	select {
	case result := <-resultCh:
		fmt.Printf("Task result: %s\n", result)
	case <-ctx.Done():
		fmt.Printf("Task cancelled: %v\n", ctx.Err())
	}

	// 20. Advanced channel patterns
	fmt.Println("\n20. Advanced patterns:")

	// Circuit breaker pattern
	breaker := NewCircuitBreaker(3, 100*time.Millisecond)

	for i := 1; i <= 10; i++ {
		result, err := breaker.Call(func() (interface{}, error) {
			if rand.Float32() < 0.6 { // 60% failure rate
				return nil, fmt.Errorf("operation failed")
			}
			return fmt.Sprintf("success-%d", i), nil
		})

		if err != nil {
			fmt.Printf("Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("Call %d succeeded: %v\n", i, result)
		}
	}
}

// Helper types and functions

// Worker pool types
type Job struct {
	ID   int
	Data int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(50 * time.Millisecond)
		results <- Result{JobID: job.ID, Value: job.Data * job.Data}
	}
}

// Channel directions
func sendOnly(ch chan<- string) {
	ch <- "data1"
	ch <- "data2"
	close(ch)
}

func receiveOnly(ch <-chan string) {
	for data := range ch {
		fmt.Printf("Received: %s\n", data)
	}
}

// Fan-out pattern
func fanOut(input <-chan int, outputs ...chan<- int) {
	for value := range input {
		for i, output := range outputs {
			if i == value%len(outputs) {
				output <- value
				break
			}
		}
	}

	for _, output := range outputs {
		close(output)
	}
}

func collectResults(name string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		fmt.Printf("%s received: %d\n", name, value)
	}
}

// Fan-in pattern
func fanIn(inputs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for value := range ch {
				out <- value
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Pipeline pattern
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

func addPrefix(in <-chan int, prefix string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for num := range in {
			out <- fmt.Sprintf("%s %d", prefix, num)
		}
	}()
	return out
}

// Channel multiplexer
type ChannelMessage struct {
	Channel string
	Data    string
}

type ChannelMux struct {
	channels map[string]<-chan string
	output   chan ChannelMessage
}

func NewChannelMux() *ChannelMux {
	return &ChannelMux{
		channels: make(map[string]<-chan string),
		output:   make(chan ChannelMessage),
	}
}

func (mux *ChannelMux) Add(name string, ch <-chan string) {
	mux.channels[name] = ch
	go func(name string, ch <-chan string) {
		for data := range ch {
			mux.output <- ChannelMessage{Channel: name, Data: data}
		}
	}(name, ch)
}

func (mux *ChannelMux) Receive() ChannelMessage {
	return <-mux.output
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
			fmt.Printf("Server heartbeat at %s\n", time.Now().Format("15:04:05"))
		case <-s.quit:
			fmt.Println("Server shutting down...")
			close(s.done)
			return
		}
	}
}

func (s *Server) Stop() {
	close(s.quit)
	<-s.done
}

// Error operations
func riskyOperation(id int, errCh chan<- error) {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	if rand.Float32() < 0.3 { // 30% chance of error
		errCh <- fmt.Errorf("operation %d failed", id)
	} else {
		errCh <- nil
	}
}

// Task orchestration
type Task struct {
	ID   int
	Data string
}

type Orchestrator struct {
	tasks   chan Task
	workers map[string]chan Task
	quit    chan struct{}
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		tasks:   make(chan Task, 10),
		workers: make(map[string]chan Task),
		quit:    make(chan struct{}),
	}
}

func (o *Orchestrator) AddWorker(name string) {
	workerCh := make(chan Task, 5)
	o.workers[name] = workerCh

	go func(name string, tasks <-chan Task) {
		for task := range tasks {
			fmt.Printf("Worker %s processing task %d: %s\n", name, task.ID, task.Data)
			time.Sleep(50 * time.Millisecond)
		}
	}(name, workerCh)

	go func() {
		for {
			select {
			case task := <-o.tasks:
				workerCh <- task
			case <-o.quit:
				close(workerCh)
				return
			}
		}
	}()
}

func (o *Orchestrator) SendTask(task Task) {
	o.tasks <- task
}

func (o *Orchestrator) Stop() {
	close(o.quit)
	close(o.tasks)
	time.Sleep(100 * time.Millisecond) // Allow workers to finish
}

// Broadcast pattern
type Broadcaster struct {
	listeners map[string]chan string
	input     chan string
	quit      chan struct{}
}

func NewBroadcaster() *Broadcaster {
	b := &Broadcaster{
		listeners: make(map[string]chan string),
		input:     make(chan string),
		quit:      make(chan struct{}),
	}

	go b.run()
	return b
}

func (b *Broadcaster) Subscribe(name string) <-chan string {
	ch := make(chan string, 10)
	b.listeners[name] = ch
	return ch
}

func (b *Broadcaster) Broadcast(message string) {
	select {
	case b.input <- message:
	case <-b.quit:
	}
}

func (b *Broadcaster) run() {
	for {
		select {
		case message := <-b.input:
			for _, listener := range b.listeners {
				select {
				case listener <- message:
				default: // Skip if listener is full
				}
			}
		case <-b.quit:
			for _, listener := range b.listeners {
				close(listener)
			}
			return
		}
	}
}

func (b *Broadcaster) Close() {
	close(b.quit)
}

func listenForBroadcasts(name string, ch <-chan string) {
	for message := range ch {
		fmt.Printf("%s received broadcast: %s\n", name, message)
	}
}

// Context operations
func longRunningTask(ctx context.Context, result chan<- string) {
	select {
	case <-time.After(200 * time.Millisecond):
		result <- "task completed"
	case <-ctx.Done():
		return
	}
}

// Circuit breaker pattern
type CircuitBreaker struct {
	failures    int
	maxFailures int
	timeout     time.Duration
	lastFailure time.Time
	state       string // "closed", "open", "half-open"
	mutex       sync.Mutex
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       "closed",
	}
}

func (cb *CircuitBreaker) Call(fn func() (interface{}, error)) (interface{}, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state == "open" {
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = "half-open"
		} else {
			return nil, fmt.Errorf("circuit breaker is open")
		}
	}

	result, err := fn()

	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()

		if cb.failures >= cb.maxFailures {
			cb.state = "open"
		}
		return nil, err
	}

	cb.failures = 0
	cb.state = "closed"
	return result, nil
}
