package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Go Generics Examples ===\n")

	// 1. Basic generic function
	fmt.Println("1. Basic generic function:")
	fmt.Printf("Max of 5 and 3: %d\n", Max(5, 3))
	fmt.Printf("Max of 2.5 and 1.8: %.1f\n", Max(2.5, 1.8))
	fmt.Printf("Max of 'hello' and 'world': %s\n", Max("hello", "world"))

	// 2. Generic function with multiple type parameters
	fmt.Println("\n2. Multiple type parameters:")
	pair := MakePair("Alice", 30)
	fmt.Printf("Pair: %+v\n", pair)

	swapped := SwapPair(pair)
	fmt.Printf("Swapped: %+v\n", swapped)

	// 3. Generic slice operations
	fmt.Println("\n3. Generic slice operations:")
	ints := []int{1, 2, 3, 4, 5}
	strings := []string{"a", "b", "c"}

	fmt.Printf("First int: %v\n", First(ints))
	fmt.Printf("First string: %v\n", First(strings))

	fmt.Printf("Last int: %v\n", Last(ints))
	fmt.Printf("Contains 3: %t\n", Contains(ints, 3))
	fmt.Printf("Contains 'x': %t\n", Contains(strings, "x"))

	// 4. Generic data structures
	fmt.Println("\n4. Generic data structures:")

	// Generic stack
	stack := NewStack[int]()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Printf("Stack size: %d\n", stack.Size())
	fmt.Printf("Popped: %v\n", stack.Pop())
	fmt.Printf("Peek: %v\n", stack.Peek())

	// Generic map
	cache := NewCache[string, int]()
	cache.Set("one", 1)
	cache.Set("two", 2)

	if value, ok := cache.Get("one"); ok {
		fmt.Printf("Cache value: %d\n", value)
	}

	// 5. Type constraints
	fmt.Println("\n5. Type constraints:")
	fmt.Printf("Sum of ints: %d\n", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("Sum of floats: %.2f\n", Sum([]float64{1.1, 2.2, 3.3}))

	// 6. Custom constraints
	fmt.Println("\n6. Custom constraints:")
	fmt.Printf("Min of ints: %d\n", Min([]int{5, 2, 8, 1, 9}))
	fmt.Printf("Min of strings: %s\n", Min([]string{"zebra", "apple", "banana"}))

	// 7. Struct type parameters
	fmt.Println("\n7. Generic structs:")
	intList := NewLinkedList[int]()
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)

	fmt.Printf("LinkedList: ")
	intList.ForEach(func(value int) {
		fmt.Printf("%d ", value)
	})
	fmt.Println()

	// 8. Generic interfaces
	fmt.Println("\n8. Generic interfaces:")
	var container Container[string] = &SliceContainer[string]{}
	container.Add("hello")
	container.Add("world")

	fmt.Printf("Container size: %d\n", container.Size())
	fmt.Printf("Container items: ")
	for i := 0; i < container.Size(); i++ {
		fmt.Printf("%s ", container.Get(i))
	}
	fmt.Println()

	// 9. Type inference
	fmt.Println("\n9. Type inference:")
	// Go can infer types in many cases
	result1 := Map([]int{1, 2, 3}, func(x int) int { return x * 2 })
	fmt.Printf("Mapped ints: %v\n", result1)

	result2 := Map([]string{"a", "b", "c"}, strings.ToUpper)
	fmt.Printf("Mapped strings: %v\n", result2)

	// 10. Generic methods
	fmt.Println("\n10. Generic methods:")
	processor := &DataProcessor[string]{}
	processor.Process("hello")
	processor.Process("world")

	results := processor.GetResults()
	fmt.Printf("Processed results: %v\n", results)

	// 11. Functional programming with generics
	fmt.Println("\n11. Functional programming:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)

	doubled := Map(evens, func(x int) int { return x * 2 })
	fmt.Printf("Doubled evens: %v\n", doubled)

	sum := Reduce(doubled, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Sum: %d\n", sum)

	// 12. Generic validation
	fmt.Println("\n12. Generic validation:")
	validator := NewValidator[int]()
	validator.AddRule(func(x int) bool { return x > 0 }, "must be positive")
	validator.AddRule(func(x int) bool { return x < 100 }, "must be less than 100")

	if err := validator.Validate(50); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("50 is valid")
	}

	if err := validator.Validate(-5); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	// 13. Generic channels
	fmt.Println("\n13. Generic channels:")
	ch := make(chan int, 3)
	go func() {
		SendToChannel(ch, 1, 2, 3)
	}()

	values := ReceiveFromChannel(ch, 3)
	fmt.Printf("Received values: %v\n", values)

	// 14. Generic comparison
	fmt.Println("\n14. Generic comparison:")
	fmt.Printf("5 == 5: %t\n", Equal(5, 5))
	fmt.Printf("'a' == 'b': %t\n", Equal("a", "b"))

	// 15. Generic optional type
	fmt.Println("\n15. Generic optional:")
	some := Some(42)
	none := None[int]()

	fmt.Printf("Some value: %v\n", some.Get())
	fmt.Printf("Some has value: %t\n", some.HasValue())
	fmt.Printf("None has value: %t\n", none.HasValue())

	// 16. Generic tree structure
	fmt.Println("\n16. Generic tree:")
	tree := &TreeNode[int]{Value: 10}
	tree.Left = &TreeNode[int]{Value: 5}
	tree.Right = &TreeNode[int]{Value: 15}

	fmt.Printf("Tree values (in-order): ")
	tree.InOrder(func(value int) {
		fmt.Printf("%d ", value)
	})
	fmt.Println()

	// 17. Generic sorting
	fmt.Println("\n17. Generic sorting:")
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	SortBy(people, func(p Person) int { return p.Age })
	fmt.Printf("Sorted by age: %v\n", people)

	// 18. Generic result type
	fmt.Println("\n18. Generic result type:")
	divResult := Divide(10, 2)
	if divResult.IsOk() {
		fmt.Printf("Division result: %.2f\n", divResult.Unwrap())
	}

	divByZero := Divide(10, 0)
	if divByZero.IsError() {
		fmt.Printf("Division error: %s\n", divByZero.Error())
	}

	// 19. Generic builder pattern
	fmt.Println("\n19. Generic builder:")
	config := NewBuilder[Config]().
		Set("host", "localhost").
		Set("port", "8080").
		Set("debug", "true").
		Build()

	fmt.Printf("Config: %+v\n", config)

	// 20. Advanced generic patterns
	fmt.Println("\n20. Advanced patterns:")

	// Generic state machine
	fsm := NewStateMachine("idle")
	fsm.AddTransition("idle", "start", "running")
	fsm.AddTransition("running", "stop", "idle")

	fmt.Printf("Initial state: %s\n", fsm.CurrentState())
	fsm.Transition("start")
	fmt.Printf("After start: %s\n", fsm.CurrentState())
	fsm.Transition("stop")
	fmt.Printf("After stop: %s\n", fsm.CurrentState())
}

// 1. Basic generic function with comparable constraint
func Max[T comparable](a, b T) T {
	if fmt.Sprintf("%v", a) > fmt.Sprintf("%v", b) {
		return a
	}
	return b
}

// 2. Multiple type parameters
type Pair[T, U any] struct {
	First  T
	Second U
}

func MakePair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

func SwapPair[T, U any](p Pair[T, U]) Pair[U, T] {
	return Pair[U, T]{First: p.Second, Second: p.First}
}

// 3. Generic slice operations
func First[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[0]
}

func Last[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[len(slice)-1]
}

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// 4. Generic data structures
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	var zero T
	if len(s.items) == 0 {
		return zero
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *Stack[T]) Peek() T {
	var zero T
	if len(s.items) == 0 {
		return zero
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Generic cache
type Cache[K comparable, V any] struct {
	data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{data: make(map[K]V)}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, ok := c.data[key]
	return value, ok
}

// 5. Type constraints for numeric types
type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Sum[T Numeric](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// 6. Custom constraints
type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

func Min[T Ordered](values []T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// 7. Generic linked list
type LinkedList[T any] struct {
	head *ListNode[T]
	size int
}

type ListNode[T any] struct {
	value T
	next  *ListNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Add(value T) {
	newNode := &ListNode[T]{value: value}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}

func (ll *LinkedList[T]) ForEach(fn func(T)) {
	current := ll.head
	for current != nil {
		fn(current.value)
		current = current.next
	}
}

// 8. Generic interfaces
type Container[T any] interface {
	Add(item T)
	Get(index int) T
	Size() int
}

type SliceContainer[T any] struct {
	items []T
}

func (sc *SliceContainer[T]) Add(item T) {
	sc.items = append(sc.items, item)
}

func (sc *SliceContainer[T]) Get(index int) T {
	return sc.items[index]
}

func (sc *SliceContainer[T]) Size() int {
	return len(sc.items)
}

// 9. Generic map function
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 10. Generic methods
type DataProcessor[T any] struct {
	results []T
}

func (dp *DataProcessor[T]) Process(data T) {
	dp.results = append(dp.results, data)
}

func (dp *DataProcessor[T]) GetResults() []T {
	return dp.results
}

// 11. Functional programming
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// 12. Generic validation
type Validator[T any] struct {
	rules []ValidationRule[T]
}

type ValidationRule[T any] struct {
	check   func(T) bool
	message string
}

func NewValidator[T any]() *Validator[T] {
	return &Validator[T]{}
}

func (v *Validator[T]) AddRule(check func(T) bool, message string) {
	v.rules = append(v.rules, ValidationRule[T]{check: check, message: message})
}

func (v *Validator[T]) Validate(value T) error {
	for _, rule := range v.rules {
		if !rule.check(value) {
			return fmt.Errorf("validation failed: %s", rule.message)
		}
	}
	return nil
}

// 13. Generic channels
func SendToChannel[T any](ch chan<- T, values ...T) {
	for _, v := range values {
		ch <- v
	}
	close(ch)
}

func ReceiveFromChannel[T any](ch <-chan T, count int) []T {
	var result []T
	for i := 0; i < count; i++ {
		if value, ok := <-ch; ok {
			result = append(result, value)
		}
	}
	return result
}

// 14. Generic comparison
func Equal[T comparable](a, b T) bool {
	return a == b
}

// 15. Generic optional type
type Optional[T any] struct {
	value    T
	hasValue bool
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{value: value, hasValue: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{hasValue: false}
}

func (o Optional[T]) HasValue() bool {
	return o.hasValue
}

func (o Optional[T]) Get() T {
	if !o.hasValue {
		panic("optional has no value")
	}
	return o.value
}

// 16. Generic tree
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (tn *TreeNode[T]) InOrder(fn func(T)) {
	if tn == nil {
		return
	}
	if tn.Left != nil {
		tn.Left.InOrder(fn)
	}
	fn(tn.Value)
	if tn.Right != nil {
		tn.Right.InOrder(fn)
	}
}

// 17. Generic sorting
type Person struct {
	Name string
	Age  int
}

func SortBy[T any, K Ordered](slice []T, keyFn func(T) K) {
	sort.Slice(slice, func(i, j int) bool {
		return keyFn(slice[i]) < keyFn(slice[j])
	})
}

// 18. Generic result type
type Result[T any] struct {
	value T
	err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsError() bool {
	return r.err != nil
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic("result is error")
	}
	return r.value
}

func (r Result[T]) Error() string {
	if r.err == nil {
		return ""
	}
	return r.err.Error()
}

func Divide(a, b float64) Result[float64] {
	if b == 0 {
		return Err[float64](fmt.Errorf("division by zero"))
	}
	return Ok(a / b)
}

// 19. Generic builder pattern
type Builder[T any] struct {
	data map[string]string
}

func NewBuilder[T any]() *Builder[T] {
	return &Builder[T]{data: make(map[string]string)}
}

func (b *Builder[T]) Set(key, value string) *Builder[T] {
	b.data[key] = value
	return b
}

func (b *Builder[T]) Build() T {
	var result T
	// This is a simplified example - in practice you'd use reflection
	// or specific build logic for each type
	return result
}

type Config struct {
	Host  string
	Port  int
	Debug bool
}

// 20. Generic state machine
type StateMachine[T comparable] struct {
	currentState T
	transitions  map[T]map[string]T
}

func NewStateMachine[T comparable](initialState T) *StateMachine[T] {
	return &StateMachine[T]{
		currentState: initialState,
		transitions:  make(map[T]map[string]T),
	}
}

func (sm *StateMachine[T]) AddTransition(from T, event string, to T) {
	if sm.transitions[from] == nil {
		sm.transitions[from] = make(map[string]T)
	}
	sm.transitions[from][event] = to
}

func (sm *StateMachine[T]) Transition(event string) bool {
	if transitions, ok := sm.transitions[sm.currentState]; ok {
		if nextState, ok := transitions[event]; ok {
			sm.currentState = nextState
			return true
		}
	}
	return false
}

func (sm *StateMachine[T]) CurrentState() T {
	return sm.currentState
}
