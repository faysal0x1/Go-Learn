package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== Go Mutex Examples ===\n")

	// 1. Basic Mutex usage
	fmt.Println("1. Basic Mutex:")
	counter := &Counter{}
	var wg sync.WaitGroup

	// Start multiple goroutines that increment counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())

	// 2. RWMutex - Read/Write Mutex
	fmt.Println("\n2. RWMutex (Read/Write Mutex):")
	cache := &Cache{data: make(map[string]string)}

	// Writers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			cache.Set(key, value)
			fmt.Printf("Writer %d: Set %s = %s\n", id, key, value)
		}(i)
	}

	// Readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond) // Let writers finish first
			key := fmt.Sprintf("key%d", id%3)
			value := cache.Get(key)
			fmt.Printf("Reader %d: Get %s = %s\n", id, key, value)
		}(i)
	}

	wg.Wait()

	// 3. Mutex vs Atomic operations
	fmt.Println("\n3. Mutex vs Atomic comparison:")

	// Mutex version
	start := time.Now()
	mutexCounter := &MutexCounter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutexCounter.Increment()
		}()
	}
	wg.Wait()
	mutexDuration := time.Since(start)
	fmt.Printf("Mutex counter: %d, Time: %v\n", mutexCounter.Value(), mutexDuration)

	// Atomic version
	start = time.Now()
	var atomicCounter int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	wg.Wait()
	atomicDuration := time.Since(start)
	fmt.Printf("Atomic counter: %d, Time: %v\n", atomic.LoadInt64(&atomicCounter), atomicDuration)

	// 4. Deadlock prevention
	fmt.Println("\n4. Deadlock prevention:")
	account1 := &Account{ID: 1, Balance: 1000}
	account2 := &Account{ID: 2, Balance: 500}

	// Transfer money safely
	go transfer(account1, account2, 200)
	go transfer(account2, account1, 100)

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Account 1 balance: %.2f\n", account1.GetBalance())
	fmt.Printf("Account 2 balance: %.2f\n", account2.GetBalance())

	// 5. Condition Variables with Mutex
	fmt.Println("\n5. Condition Variables:")
	buffer := &Buffer{maxSize: 5}

	// Producer
	go func() {
		for i := 1; i <= 10; i++ {
			buffer.Put(i)
			fmt.Printf("Produced: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Consumer
	go func() {
		for i := 1; i <= 10; i++ {
			value := buffer.Get()
			fmt.Printf("Consumed: %d\n", value)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)

	// 6. Once - Execute only once
	fmt.Println("\n6. sync.Once:")
	service := &Service{}

	// Multiple goroutines trying to initialize
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			service.Initialize()
			fmt.Printf("Goroutine %d called Initialize\n", id)
		}(i)
	}

	wg.Wait()

	// 7. Pool - Object pooling
	fmt.Println("\n7. sync.Pool:")
	objectPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new object")
			return &PooledObject{ID: rand.Intn(1000)}
		},
	}

	// Use pool objects
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Get object from pool
			obj := objectPool.Get().(*PooledObject)
			fmt.Printf("Worker %d got object %d\n", id, obj.ID)

			// Use object
			time.Sleep(50 * time.Millisecond)

			// Put back to pool
			objectPool.Put(obj)
		}(i)
	}

	wg.Wait()

	// 8. Map with Mutex
	fmt.Println("\n8. Thread-safe Map:")
	safeMap := &SafeMap{data: make(map[string]int)}

	// Writers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			safeMap.Set(key, id*10)
		}(i)
	}

	// Readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			key := fmt.Sprintf("key%d", id)
			if value, exists := safeMap.Get(key); exists {
				fmt.Printf("Map[%s] = %d\n", key, value)
			}
		}(i)
	}

	wg.Wait()

	// 9. Semaphore using channels and mutex
	fmt.Println("\n9. Semaphore pattern:")
	semaphore := NewSemaphore(3) // Allow 3 concurrent operations

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			semaphore.Acquire()
			fmt.Printf("Operation %d started\n", id)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Operation %d finished\n", id)
			semaphore.Release()
		}(i)
	}

	wg.Wait()

	// 10. Reader-Writer problem
	fmt.Println("\n10. Reader-Writer problem:")
	library := &Library{books: make(map[string]string)}

	// Add some books
	library.AddBook("Go Programming", "John Doe")
	library.AddBook("Concurrency", "Jane Smith")

	// Multiple readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			books := []string{"Go Programming", "Concurrency", "Non-existent"}
			book := books[id%len(books)]
			author := library.GetAuthor(book)
			if author != "" {
				fmt.Printf("Reader %d: %s by %s\n", id, book, author)
			} else {
				fmt.Printf("Reader %d: Book '%s' not found\n", id, book)
			}
		}(i)
	}

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		library.AddBook("Patterns", "Gang of Four")
		fmt.Println("Writer: Added new book")
	}()

	wg.Wait()

	// 11. Metrics collection with mutex
	fmt.Println("\n11. Metrics collection:")
	metrics := &Metrics{}

	// Simulate operations
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			start := time.Now()
			// Simulate work
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			duration := time.Since(start)

			metrics.RecordOperation(duration)
		}()
	}

	wg.Wait()
	stats := metrics.GetStats()
	fmt.Printf("Operations: %d, Total time: %v, Average: %v\n",
		stats.Count, stats.Total, stats.Average)

	// 12. Connection pool
	fmt.Println("\n12. Connection pool:")
	pool := NewConnectionPool(3)

	// Simulate multiple clients
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn := pool.Get()
			fmt.Printf("Client %d got connection %d\n", id, conn.ID)

			// Use connection
			time.Sleep(100 * time.Millisecond)

			pool.Put(conn)
		}(i)
	}

	wg.Wait()

	// 13. Event system with mutex
	fmt.Println("\n13. Event system:")
	eventBus := &EventBus{handlers: make(map[string][]EventHandler)}

	// Register handlers
	eventBus.Subscribe("user.created", func(data interface{}) {
		fmt.Printf("Email service: User created - %v\n", data)
	})

	eventBus.Subscribe("user.created", func(data interface{}) {
		fmt.Printf("Analytics: User created - %v\n", data)
	})

	// Publish events
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			userData := fmt.Sprintf("user_%d", id)
			eventBus.Publish("user.created", userData)
		}(i)
	}

	wg.Wait()

	// 14. State machine with mutex
	fmt.Println("\n14. State machine:")
	fsm := &StateMachine{state: "idle"}

	// Multiple state transitions
	transitions := []string{"start", "process", "complete", "reset"}
	for i, transition := range transitions {
		wg.Add(1)
		go func(id int, event string) {
			defer wg.Done()
			time.Sleep(time.Duration(id*50) * time.Millisecond)
			fsm.Transition(event)
		}(i, transition)
	}

	wg.Wait()

	// 15. Rate limiter with mutex
	fmt.Println("\n15. Rate limiter:")
	limiter := NewRateLimiter(3, time.Second) // 3 requests per second

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if limiter.Allow() {
				fmt.Printf("Request %d: Allowed\n", id)
			} else {
				fmt.Printf("Request %d: Rate limited\n", id)
			}
		}(i)
	}

	wg.Wait()

	// 16. Work queue with mutex
	fmt.Println("\n16. Work queue:")
	queue := &WorkQueue{}

	// Producer
	go func() {
		for i := 1; i <= 10; i++ {
			task := Task{ID: i, Data: fmt.Sprintf("task_%d", i)}
			queue.Enqueue(task)
			fmt.Printf("Enqueued task %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Consumer
	go func() {
		for i := 1; i <= 10; i++ {
			if task, ok := queue.Dequeue(); ok {
				fmt.Printf("Processing task %d: %s\n", task.ID, task.Data)
			}
			time.Sleep(80 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)

	// 17. Distributed counter
	fmt.Println("\n17. Distributed counter:")
	distCounter := &DistributedCounter{counters: make(map[string]int)}

	nodes := []string{"node1", "node2", "node3"}
	for _, node := range nodes {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(nodeID string) {
				defer wg.Done()
				distCounter.Increment(nodeID)
			}(node)
		}
	}

	wg.Wait()
	total := distCounter.Total()
	fmt.Printf("Total distributed count: %d\n", total)

	// 18. Circuit breaker with mutex
	fmt.Println("\n18. Circuit breaker:")
	breaker := NewCircuitBreaker(3, 500*time.Millisecond)

	// Simulate operations with failures
	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			err := breaker.Call(func() error {
				// Simulate failure for some operations
				if rand.Float32() < 0.6 {
					return fmt.Errorf("operation failed")
				}
				return nil
			})

			if err != nil {
				fmt.Printf("Operation %d: %v\n", id, err)
			} else {
				fmt.Printf("Operation %d: Success\n", id)
			}
		}(i)
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()

	// 19. Lazy initialization
	fmt.Println("\n19. Lazy initialization:")
	lazy := &LazyResource{}

	// Multiple goroutines trying to access resource
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			resource := lazy.GetResource()
			fmt.Printf("Goroutine %d got resource: %s\n", id, resource.Name)
		}(i)
	}

	wg.Wait()

	// 20. Performance monitoring
	fmt.Println("\n20. Performance monitoring:")
	monitor := &PerformanceMonitor{
		operations: make(map[string]*OperationStats),
	}

	// Simulate different operations
	operations := []string{"database", "api", "cache"}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			op := operations[id%len(operations)]
			start := time.Now()

			// Simulate work
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

			duration := time.Since(start)
			monitor.RecordOperation(op, duration)
		}(i)
	}

	wg.Wait()

	fmt.Println("Performance Statistics:")
	stats := monitor.GetAllStats()
	for op, stat := range stats {
		fmt.Printf("%s: Count=%d, Total=%v, Avg=%v\n",
			op, stat.Count, stat.Total, stat.Average)
	}

	fmt.Printf("\nFinal number of goroutines: %d\n", runtime.NumGoroutine())
}

// 1. Basic Counter with Mutex
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

// 2. Cache with RWMutex
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// 3. Mutex Counter for comparison
type MutexCounter struct {
	mu    sync.Mutex
	value int
}

func (mc *MutexCounter) Increment() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.value++
}

func (mc *MutexCounter) Value() int {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	return mc.value
}

// 4. Account for deadlock prevention
type Account struct {
	mu      sync.Mutex
	ID      int
	Balance float64
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}

func transfer(from, to *Account, amount float64) {
	// Always lock accounts in the same order to prevent deadlock
	if from.ID < to.ID {
		from.mu.Lock()
		defer from.mu.Unlock()
		to.mu.Lock()
		defer to.mu.Unlock()
	} else {
		to.mu.Lock()
		defer to.mu.Unlock()
		from.mu.Lock()
		defer from.mu.Unlock()
	}

	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
		fmt.Printf("Transferred %.2f from account %d to account %d\n",
			amount, from.ID, to.ID)
	}
}

// 5. Buffer with Condition Variables
type Buffer struct {
	mu      sync.Mutex
	cond    *sync.Cond
	items   []int
	maxSize int
}

func (b *Buffer) Put(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.cond == nil {
		b.cond = sync.NewCond(&b.mu)
	}

	// Wait while buffer is full
	for len(b.items) >= b.maxSize {
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	b.cond.Signal() // Notify consumers
}

func (b *Buffer) Get() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.cond == nil {
		b.cond = sync.NewCond(&b.mu)
	}

	// Wait while buffer is empty
	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]
	b.cond.Signal() // Notify producers
	return item
}

// 6. Service with Once
type Service struct {
	once        sync.Once
	initialized bool
}

func (s *Service) Initialize() {
	s.once.Do(func() {
		fmt.Println("Service initialization - this runs only once")
		time.Sleep(100 * time.Millisecond) // Simulate initialization work
		s.initialized = true
	})
}

// 7. Pooled Object
type PooledObject struct {
	ID int
}

// 8. Thread-safe Map
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

// 9. Semaphore
type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, size)}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}

// 10. Library (Reader-Writer)
type Library struct {
	mu    sync.RWMutex
	books map[string]string
}

func (l *Library) AddBook(title, author string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.books[title] = author
}

func (l *Library) GetAuthor(title string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.books[title]
}

// 11. Metrics
type Metrics struct {
	mu        sync.Mutex
	count     int
	totalTime time.Duration
}

type Stats struct {
	Count   int
	Total   time.Duration
	Average time.Duration
}

func (m *Metrics) RecordOperation(duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.count++
	m.totalTime += duration
}

func (m *Metrics) GetStats() Stats {
	m.mu.Lock()
	defer m.mu.Unlock()

	avg := time.Duration(0)
	if m.count > 0 {
		avg = m.totalTime / time.Duration(m.count)
	}

	return Stats{
		Count:   m.count,
		Total:   m.totalTime,
		Average: avg,
	}
}

// 12. Connection Pool
type Connection struct {
	ID int
}

type ConnectionPool struct {
	mu    sync.Mutex
	conns []*Connection
	size  int
}

func NewConnectionPool(size int) *ConnectionPool {
	pool := &ConnectionPool{size: size}
	for i := 0; i < size; i++ {
		pool.conns = append(pool.conns, &Connection{ID: i + 1})
	}
	return pool
}

func (cp *ConnectionPool) Get() *Connection {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	for len(cp.conns) == 0 {
		// In real implementation, you'd use condition variables
		cp.mu.Unlock()
		time.Sleep(10 * time.Millisecond)
		cp.mu.Lock()
	}

	conn := cp.conns[0]
	cp.conns = cp.conns[1:]
	return conn
}

func (cp *ConnectionPool) Put(conn *Connection) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.conns = append(cp.conns, conn)
}

// 13. Event System
type EventHandler func(data interface{})

type EventBus struct {
	mu       sync.RWMutex
	handlers map[string][]EventHandler
}

func (eb *EventBus) Subscribe(event string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mu.RLock()
	handlers := make([]EventHandler, len(eb.handlers[event]))
	copy(handlers, eb.handlers[event])
	eb.mu.RUnlock()

	for _, handler := range handlers {
		go handler(data)
	}
}

// 14. State Machine
type StateMachine struct {
	mu    sync.Mutex
	state string
}

func (sm *StateMachine) Transition(event string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	oldState := sm.state
	switch sm.state {
	case "idle":
		if event == "start" {
			sm.state = "running"
		}
	case "running":
		if event == "process" {
			sm.state = "processing"
		}
	case "processing":
		if event == "complete" {
			sm.state = "completed"
		}
	case "completed":
		if event == "reset" {
			sm.state = "idle"
		}
	}

	fmt.Printf("State transition: %s --[%s]--> %s\n", oldState, event, sm.state)
}

// 15. Rate Limiter
type RateLimiter struct {
	mu        sync.Mutex
	requests  int
	limit     int
	window    time.Duration
	lastReset time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:     limit,
		window:    window,
		lastReset: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.Sub(rl.lastReset) >= rl.window {
		rl.requests = 0
		rl.lastReset = now
	}

	if rl.requests < rl.limit {
		rl.requests++
		return true
	}

	return false
}

// 16. Work Queue
type Task struct {
	ID   int
	Data string
}

type WorkQueue struct {
	mu    sync.Mutex
	tasks []Task
}

func (wq *WorkQueue) Enqueue(task Task) {
	wq.mu.Lock()
	defer wq.mu.Unlock()
	wq.tasks = append(wq.tasks, task)
}

func (wq *WorkQueue) Dequeue() (Task, bool) {
	wq.mu.Lock()
	defer wq.mu.Unlock()

	if len(wq.tasks) == 0 {
		return Task{}, false
	}

	task := wq.tasks[0]
	wq.tasks = wq.tasks[1:]
	return task, true
}

// 17. Distributed Counter
type DistributedCounter struct {
	mu       sync.Mutex
	counters map[string]int
}

func (dc *DistributedCounter) Increment(nodeID string) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.counters[nodeID]++
}

func (dc *DistributedCounter) Total() int {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	total := 0
	for _, count := range dc.counters {
		total += count
	}
	return total
}

// 18. Circuit Breaker
type CircuitBreaker struct {
	mu          sync.Mutex
	failures    int
	maxFailures int
	timeout     time.Duration
	lastFailure time.Time
	state       string // "closed", "open", "half-open"
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       "closed",
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == "open" {
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = "half-open"
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}

	err := fn()

	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()

		if cb.failures >= cb.maxFailures {
			cb.state = "open"
		}
		return err
	}

	cb.failures = 0
	cb.state = "closed"
	return nil
}

// 19. Lazy Resource
type Resource struct {
	Name string
}

type LazyResource struct {
	once     sync.Once
	resource *Resource
}

func (lr *LazyResource) GetResource() *Resource {
	lr.once.Do(func() {
		fmt.Println("Initializing expensive resource...")
		time.Sleep(100 * time.Millisecond)
		lr.resource = &Resource{Name: "Expensive Resource"}
	})
	return lr.resource
}

// 20. Performance Monitor
type OperationStats struct {
	Count   int
	Total   time.Duration
	Average time.Duration
}

type PerformanceMonitor struct {
	mu         sync.Mutex
	operations map[string]*OperationStats
}

func (pm *PerformanceMonitor) RecordOperation(name string, duration time.Duration) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if pm.operations[name] == nil {
		pm.operations[name] = &OperationStats{}
	}

	stats := pm.operations[name]
	stats.Count++
	stats.Total += duration
	stats.Average = stats.Total / time.Duration(stats.Count)
}

func (pm *PerformanceMonitor) GetAllStats() map[string]*OperationStats {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	result := make(map[string]*OperationStats)
	for name, stats := range pm.operations {
		result[name] = &OperationStats{
			Count:   stats.Count,
			Total:   stats.Total,
			Average: stats.Average,
		}
	}
	return result
}
