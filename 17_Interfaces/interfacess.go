package main

import (
    "fmt"
    "io"
    "math"
    "sort"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("=== Go Interfaces Examples ===")

    // 1. Basic interface definition and implementation
    fmt.Println("1. Basic interface:")
    var shape Shape
    
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}
    
    shape = circle
    fmt.Printf("Circle area: %.2f", shape.Area())
    
    shape = rectangle
    fmt.Printf("Rectangle area: %.2f", shape.Area())

    // 2. Interface with multiple methods
    fmt.Println("2. Interface with multiple methods:")
    var animal Animal
    
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    
    animal = dog
    animal.Speak()
    animal.Move()
    
    animal = cat
    animal.Speak()
    animal.Move()

    // 3. Empty interface (interface{})
    fmt.Println("3. Empty interface:")
    var anything interface{}
    
    anything = 42
    fmt.Printf("Int: %v", anything)
    
    anything = "Hello"
    fmt.Printf("String: %v", anything)
    
    anything = []int{1, 2, 3}
    fmt.Printf("Slice: %v", anything)

    // 4. Type assertion
    fmt.Println("4. Type assertion:")
    var i interface{} = "Hello, World!"
    
    // Safe type assertion
    str, ok := i.(string)
    if ok {
        fmt.Printf("String value: %s", str)
    } else {
        fmt.Println("Not a string")
    }
    
    // Direct type assertion (can panic if wrong)
    str2 := i.(string)
    fmt.Printf("Direct assertion: %s", str2)

    // 5. Type switch
    fmt.Println("5. Type switch:")
    values := []interface{}{42, "hello", 3.14, true, []int{1, 2, 3}}
    
    for _, v := range values {
        switch val := v.(type) {
        case int:
            fmt.Printf("Integer: %d", val)
        case string:
            fmt.Printf("String: %s", val)
        case float64:
            fmt.Printf("Float: %.2f", val)
        case bool:
            fmt.Printf("Boolean: %t", val)
        default:
            fmt.Printf("Unknown type: %T = %v", val, val)
        }
    }

    // 6. Interface composition
    fmt.Println("6. Interface composition:")
    var rw ReadWriter = &Buffer{data: []byte("initial data")}
    
    // Read operation
    data := make([]byte, 5)
    n, _ := rw.Read(data)
    fmt.Printf("Read %d bytes: %s", n, string(data))
    
    // Write operation
    n, _ = rw.Write([]byte(" more"))
    fmt.Printf("Wrote %d bytes", n)

    // 7. Interface values and nil
    fmt.Println("7. Interface values and nil:")
    var w Writer
    fmt.Printf("Nil interface: %v", w == nil)
    
    var f *File
    w = f // Interface holds nil pointer
    fmt.Printf("Interface with nil pointer: %v", w == nil)
    fmt.Printf("Underlying value is nil: %v", f == nil)

    // 8. Polymorphism with interfaces
    fmt.Println("8. Polymorphism:")
    shapes := []Shape{
        Circle{Radius: 3},
        Rectangle{Width: 4, Height: 5},
        Triangle{Base: 6, Height: 8},
    }
    
    totalArea := 0.0
    for i, shape := range shapes {
        area := shape.Area()
        fmt.Printf("Shape %d area: %.2f", i+1, area)
        totalArea += area
    }
    fmt.Printf("Total area: %.2f", totalArea)

    // 9. Interface satisfaction
    fmt.Println("9. Interface satisfaction:")
    var printer Printer
    
    // Both types implement Printer interface
    printer = Document{Content: "Hello World"}
    printer.Print()
    
    printer = Photo{Filename: "vacation.jpg"}
    printer.Print()

    // 10. Standard library interfaces
    fmt.Println("10. Standard library interfaces:")
    
    // Stringer interface
    person := Person{Name: "Alice", Age: 30}
    fmt.Printf("Person: %s", person) // Uses String() method
    
    // Sort interface
    numbers := Numbers{64, 17, 238, 42, 89}
    fmt.Printf("Before sorting: %v", numbers)
    sort.Sort(numbers)
    fmt.Printf("After sorting: %v", numbers)

    // 11. Interface embedding
    fmt.Println("11. Interface embedding:")
    var db Database = &SQLDatabase{connected: true}
    
    // Can call methods from embedded interfaces
    db.Connect()
    db.Query("SELECT * FROM users")
    db.Transaction(func() {
        fmt.Println("Executing transaction")
    })

    // 12. Method sets and pointer receivers
    fmt.Println("12. Method sets:")
    var counter Counter
    
    // Value can call both pointer and value receiver methods
    counter = ValueCounter{count: 0}
    fmt.Printf("Value counter: %d", counter.Count())
    counter.Increment()
    fmt.Printf("After increment: %d", counter.Count())
    
    // Pointer can also call both
    pc := &PointerCounter{count: 10}
    counter = pc
    fmt.Printf("Pointer counter: %d", counter.Count())
    counter.Increment()
    fmt.Printf("After increment: %d", counter.Count())

    // 13. Interface type assertion patterns
    fmt.Println("13. Type assertion patterns:")
    var obj interface{} = &Document{Content: "Test document"}
    
    // Check if implements specific interface
    if printer, ok := obj.(Printer); ok {
        fmt.Println("Object implements Printer:")
        printer.Print()
    }
    
    // Check for specific type
    if doc, ok := obj.(*Document); ok {
        fmt.Printf("It's a document: %s", doc.Content)
    }

    // 14. Interface with generic behavior
    fmt.Println("14. Generic behavior with interfaces:")
    var container Container
    
    container = &SliceContainer{items: []interface{}{1, "hello", 3.14}}
    fmt.Printf("Container size: %d", container.Size())
    container.Add("new item")
    fmt.Printf("After adding: size = %d", container.Size())

    // 15. Error interface
    fmt.Println("15. Error interface:")
    err := divide(10, 0)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
        
        // Check for custom error type
        if mathErr, ok := err.(*MathError); ok {
            fmt.Printf("Math error code: %d", mathErr.Code)
        }
    }
    
    result, err := divide(10, 2)
    if err == nil {
        fmt.Printf("Result: %.2f", result)
    }

    // 16. Interface with context
    fmt.Println("16. Interface with context:")
    services := []Service{
        &EmailService{},
        &SMSService{},
        &PushService{},
    }
    
    for _, service := range services {
        service.Start()
        service.Process("Hello")
        service.Stop()
        fmt.Println("---")
    }

    // 17. Interface for testing (dependency injection)
    fmt.Println("17. Dependency injection:")
    // Production logger
    prodLogger := &FileLogger{filename: "app.log"}
    app1 := NewApplication(prodLogger)
    app1.Run()
    
    // Test logger
    testLogger := &MockLogger{}
    app2 := NewApplication(testLogger)
    app2.Run()

    // 18. Interface duck typing
    fmt.Println("18. Duck typing:")
    // If it walks like a duck and quacks like a duck, it's a duck
    things := []interface{}{
        Duck{},
        Robot{},
        Person{Name: "Bob", Age: 25},
    }
    
    for _, thing := range things {
        if walker, ok := thing.(Walker); ok {
            walker.Walk()
        }
        if talker, ok := thing.(Talker); ok {
            talker.Talk()
        }
    }

    // 19. Interface with reflection
    fmt.Println("19. Interface with reflection:")
    var processor Processor = &DataProcessor{}
    processInterface(processor)

    // 20. Advanced interface patterns
    fmt.Println("20. Advanced patterns:")
    
    // Functional interface pattern
    var handler Handler = HandlerFunc(func(data string) {
        fmt.Printf("Handling: %s", data)
    })
    handler.Handle("test data")
    
    // Adapter pattern
    var payment PaymentProcessor = &PayPalAdapter{
        paypal: &PayPalAPI{},
    }
    payment.ProcessPayment(100.0)
}

// Type definitions and implementations

// 1. Basic interface
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Triangle struct {
    Base, Height float64
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

// 2. Interface with multiple methods
type Animal interface {
    Speak()
    Move()
}

type Dog struct {
    Name string
}

func (d Dog) Speak() {
    fmt.Printf("%s says: Woof!", d.Name)
}

func (d Dog) Move() {
    fmt.Printf("%s runs", d.Name)
}

type Cat struct {
    Name string
}

func (c Cat) Speak() {
    fmt.Printf("%s says: Meow!", c.Name)
}

func (c Cat) Move() {
    fmt.Printf("%s prowls", c.Name)
}

// 6. Interface composition
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}

type Buffer struct {
    data []byte
}

func (b *Buffer) Read(p []byte) (int, error) {
    if len(b.data) == 0 {
        return 0, io.EOF
    }
    n := copy(p, b.data)
    b.data = b.data[n:]
    return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
    b.data = append(b.data, p...)
    return len(p), nil
}

// 7. Interface values and nil
type File struct {
    name string
}

func (f *File) Write(data []byte) (int, error) {
    if f == nil {
        return 0, fmt.Errorf("file is nil")
    }
    fmt.Printf("Writing to %s: %s", f.name, string(data))
    return len(data), nil
}

// 9. Interface satisfaction
type Printer interface {
    Print()
}

type Document struct {
    Content string
}

func (d Document) Print() {
    fmt.Printf("Printing document: %s", d.Content)
}

type Photo struct {
    Filename string
}

func (p Photo) Print() {
    fmt.Printf("Printing photo: %s", p.Filename)
}

// 10. Standard library interfaces
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

type Numbers []int

func (n Numbers) Len() int           { return len(n) }
func (n Numbers) Less(i, j int) bool { return n[i] < n[j] }
func (n Numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// 11. Interface embedding
type Connector interface {
    Connect()
    Disconnect()
}

type Querier interface {
    Query(sql string) interface{}
}

type Transactor interface {
    Transaction(func())
}

type Database interface {
    Connector
    Querier
    Transactor
}

type SQLDatabase struct {
    connected bool
}

func (db *SQLDatabase) Connect() {
    db.connected = true
    fmt.Println("Connected to SQL database")
}

func (db *SQLDatabase) Disconnect() {
    db.connected = false
    fmt.Println("Disconnected from SQL database")
}

func (db *SQLDatabase) Query(sql string) interface{} {
    fmt.Printf("Executing query: %s", sql)
    return nil
}

func (db *SQLDatabase) Transaction(fn func()) {
    fmt.Println("Starting transaction")
    fn()
    fmt.Println("Committing transaction")
}

// 12. Method sets
type Counter interface {
    Count() int
    Increment()
}

type ValueCounter struct {
    count int
}

func (vc ValueCounter) Count() int {
    return vc.count
}

func (vc ValueCounter) Increment() {
    vc.count++ // This won't modify the original!
}

type PointerCounter struct {
    count int
}

func (pc *PointerCounter) Count() int {
    return pc.count
}

func (pc *PointerCounter) Increment() {
    pc.count++
}

// 14. Generic behavior
type Container interface {
    Add(item interface{})
    Size() int
}

type SliceContainer struct {
    items []interface{}
}

func (sc *SliceContainer) Add(item interface{}) {
    sc.items = append(sc.items, item)
}

func (sc *SliceContainer) Size() int {
    return len(sc.items)
}

// 15. Error interface
type MathError struct {
    Op  string
    Err string
    Code int
}

func (e *MathError) Error() string {
    return fmt.Sprintf("math error in %s: %s", e.Op, e.Err)
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, &MathError{
            Op:   "divide",
            Err:  "division by zero",
            Code: 1001,
        }
    }
    return a / b, nil
}

// 16. Service interface
type Service interface {
    Start()
    Stop()
    Process(data string)
}

type EmailService struct{}

func (es *EmailService) Start() {
    fmt.Println("Email service started")
}

func (es *EmailService) Stop() {
    fmt.Println("Email service stopped")
}

func (es *EmailService) Process(data string) {
    fmt.Printf("Sending email: %s", data)
}

type SMSService struct{}

func (ss *SMSService) Start() {
    fmt.Println("SMS service started")
}

func (ss *SMSService) Stop() {
    fmt.Println("SMS service stopped")
}

func (ss *SMSService) Process(data string) {
    fmt.Printf("Sending SMS: %s", data)
}

type PushService struct{}

func (ps *PushService) Start() {
    fmt.Println("Push service started")
}

func (ps *PushService) Stop() {
    fmt.Println("Push service stopped")
}

func (ps *PushService) Process(data string) {
    fmt.Printf("Sending push notification: %s", data)
}

// 17. Dependency injection
type Logger interface {
    Log(message string)
}

type FileLogger struct {
    filename string
}

func (fl *FileLogger) Log(message string) {
    fmt.Printf("[FILE] %s: %s", fl.filename, message)
}

type MockLogger struct{}

func (ml *MockLogger) Log(message string) {
    fmt.Printf("[MOCK] %s", message)
}

type Application struct {
    logger Logger
}

func NewApplication(logger Logger) *Application {
    return &Application{logger: logger}
}

func (app *Application) Run() {
    app.logger.Log("Application started")
}

// 18. Duck typing
type Walker interface {
    Walk()
}

type Talker interface {
    Talk()
}

type Duck struct{}

func (d Duck) Walk() {
    fmt.Println("Duck waddles")
}

func (d Duck) Talk() {
    fmt.Println("Duck quacks")
}

type Robot struct{}

func (r Robot) Walk() {
    fmt.Println("Robot moves mechanically")
}

func (r Robot) Talk() {
    fmt.Println("Robot speaks digitally")
}

func (p Person) Walk() {
    fmt.Printf("%s walks", p.Name)
}

func (p Person) Talk() {
    fmt.Printf("%s talks", p.Name)
}

// 19. Reflection with interfaces
type Processor interface {
    Process()
}

type DataProcessor struct{}

func (dp *DataProcessor) Process() {
    fmt.Println("Processing data")
}

func processInterface(p Processor) {
    fmt.Printf("Processing with type: %T", p)
    p.Process()
}

// 20. Advanced patterns
type Handler interface {
    Handle(data string)
}

type HandlerFunc func(string)

func (hf HandlerFunc) Handle(data string) {
    hf(data)
}

// Adapter pattern
type PaymentProcessor interface {
    ProcessPayment(amount float64)
}

type PayPalAPI struct{}

func (pp *PayPalAPI) MakePayment(amount float64) {
    fmt.Printf("PayPal payment: $%.2f", amount)
}

type PayPalAdapter struct {
    paypal *PayPalAPI
}

func (ppa *PayPalAdapter) ProcessPayment(amount float64) {
    ppa.paypal.MakePayment(amount)
}