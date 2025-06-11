import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Interfaces() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interfaces in Go</h2>
            <p className="mb-4">
                Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here. Interfaces enable polymorphism, abstraction, and decoupling in Go programs.
            </p>
            <p className="mb-4">Key features of interfaces in Go:</p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Define method sets for types to implement</li>
                <li>Support for implicit implementation (no explicit 'implements')</li>
                <li>Enable polymorphism and abstraction</li>
                <li>Empty interface <code>interface&#123;&#125;</code> for any type</li>
                <li>Type assertion and type switch for dynamic behavior</li>
                <li>Interface composition and embedding</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interface Examples</h2>
            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Interface Implementation</h3>
                    <CodeBlock code={`type Shape interface { Area() float64 }
type Circle struct { Radius float64 }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
type Rectangle struct { Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
var shape Shape
shape = Circle{5}
fmt.Printf("Circle area: %.2f\n", shape.Area())
shape = Rectangle{4,6}
fmt.Printf("Rectangle area: %.2f\n", shape.Area())`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Interface with Multiple Methods</h3>
                    <CodeBlock code={`type Animal interface { Speak(); Move() }
type Dog struct { Name string }
func (d Dog) Speak() { fmt.Println("Woof!") }
func (d Dog) Move() { fmt.Println("Dog runs") }
dog := Dog{"Buddy"}
var animal Animal = dog
animal.Speak(); animal.Move()`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Empty Interface and Type Assertion</h3>
                    <CodeBlock code={`var anything interface{}
anything = 42
fmt.Printf("Int: %v\n", anything)
anything = "Hello"
fmt.Printf("String: %v\n", anything)
// Type assertion
str, ok := anything.(string)
if ok { fmt.Printf("String value: %s\n", str) }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Type Switch</h3>
                    <CodeBlock code={`values := []interface{}{42, "hello", 3.14, true}
for _, v := range values {
  switch val := v.(type) {
  case int:
    fmt.Printf("Integer: %d\n", val)
  case string:
    fmt.Printf("String: %s\n", val)
  case float64:
    fmt.Printf("Float: %.2f\n", val)
  case bool:
    fmt.Printf("Boolean: %t\n", val)
  default:
    fmt.Printf("Unknown type: %T\n", val)
  }
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Interface Composition</h3>
                    <CodeBlock code={`type Reader interface { Read([]byte) (int, error) }
type Writer interface { Write([]byte) (int, error) }
type ReadWriter interface { Reader; Writer }
// Buffer implements both
var rw ReadWriter = &Buffer{data: []byte("initial data")}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">6. Polymorphism with Interfaces</h3>
                    <CodeBlock code={`shapes := []Shape{Circle{3}, Rectangle{4,5}}
for i, shape := range shapes {
  fmt.Printf("Shape %d area: %.2f\n", i+1, shape.Area())
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">7. Standard Library Interfaces</h3>
                    <CodeBlock code={`type Stringer interface { String() string }
type Person struct { Name string; Age int }
func (p Person) String() string { return fmt.Sprintf("%s (%d)", p.Name, p.Age) }
p := Person{"Alice", 30}
fmt.Printf("Person: %s\n", p)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">8. Interface Embedding</h3>
                    <CodeBlock code={`type Connector interface { Connect() }
type Querier interface { Query(sql string) interface{} }
type Database interface { Connector; Querier }
type SQLDatabase struct { connected bool }
func (db *SQLDatabase) Connect() { db.connected = true }
func (db *SQLDatabase) Query(sql string) interface{} { return nil }
var db Database = &SQLDatabase{}
db.Connect(); db.Query("SELECT * FROM users")`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">9. Error Interface</h3>
                    <CodeBlock code={`type MathError struct { Op, Err string; Code int }
func (e *MathError) Error() string { return fmt.Sprintf("%s: %s (code %d)", e.Op, e.Err, e.Code) }
func divide(a, b float64) (float64, error) {
  if b == 0 { return 0, &MathError{"divide", "division by zero", 1} }
  return a / b, nil
}
result, err := divide(10, 0)
if err != nil { fmt.Println(err) }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">10. Dependency Injection with Interfaces</h3>
                    <CodeBlock code={`type Logger interface { Log(message string) }
type Application struct { logger Logger }
func (app *Application) Run() { app.logger.Log("Running app") }
prodLogger := &FileLogger{"app.log"}
app := Application{logger: prodLogger}
app.Run()`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Shape Interface</h3>
                    <p className="mb-4">Implement a Shape interface with Area() and Perimeter() methods. Create Rectangle and Circle types that satisfy the interface.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Type Switch</h3>
                    <p className="mb-4">Write a function that takes an interface{ } and prints its type and value using a type switch.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Dependency Injection</h3>
                    <p className="mb-4">Create a Logger interface and two implementations: one that logs to the console and one that logs to a slice. Write a function that accepts a Logger and logs messages.</p>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">Try out interface patterns in the interactive playground below.</p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main
import "fmt"
type Shape interface { Area() float64 }
type Rectangle struct { Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
func main() {
  var s Shape = Rectangle{4, 5}
  fmt.Printf("Area: %.2f\n", s.Area())
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Area: 20.00</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Interfaces"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 