import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Structs() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Structs in Go</h2>
            <p className="mb-4">
                Structs are composite types in Go that group together zero or more fields with varying types. They are the foundation for building complex data models, supporting composition, methods, and more.
            </p>
            <p className="mb-4">Key features of structs in Go:</p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Group related data fields</li>
                <li>Support for methods and interfaces</li>
                <li>Composition via embedding (anonymous fields)</li>
                <li>Struct tags for metadata (e.g., JSON)</li>
                <li>Zero values, copying, and comparison</li>
                <li>Can be used with pointers, slices, and maps</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Struct Examples</h2>
            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Struct Declaration and Usage</h3>
                    <CodeBlock code={`type Person struct { Name string; Age int; City string }
var person Person
person = Person{Name: "Alice", Age: 30, City: "New York"}
fmt.Printf("Person: %+v\n", person)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Struct Literal Initialization</h3>
                    <CodeBlock code={`person1 := Person{Name: "Bob", Age: 25, City: "San Francisco"}
person2 := Person{"Charlie", 35, "Chicago"}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Anonymous Structs</h3>
                    <CodeBlock code={`config := struct { Host string; Port int; SSL bool }{
  Host: "localhost", Port: 8080, SSL: true,
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Accessing and Modifying Fields</h3>
                    <CodeBlock code={`fmt.Printf("Name: %s\n", person1.Name)
person1.Age = 26`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Struct Pointers</h3>
                    <CodeBlock code={`personPtr := &person1
personPtr.Age = 27 // Automatic dereferencing`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">6. Nested and Embedded Structs</h3>
                    <CodeBlock code={`type Employee struct {
  Person   // Embedded
  ID       int
  Salary   float64
  Position string
}
emp := Employee{Person: Person{"David", 40, "Boston"}, ID: 1001, Salary: 75000, Position: "Engineer"}
fmt.Printf("Employee name: %s\n", emp.Name)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">7. Methods on Structs</h3>
                    <CodeBlock code={`type Rectangle struct { Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
rect := Rectangle{5, 3}
fmt.Printf("Area: %.2f\n", rect.Area())`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">8. Pointer vs Value Receivers</h3>
                    <CodeBlock code={`type Counter struct { Value int }
func (c Counter) IncrementValue() { c.Value++ }
func (c *Counter) IncrementPointer() { c.Value++ }
c := Counter{0}
c.IncrementValue() // No effect
c.IncrementPointer() // Modifies original`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">9. Struct Comparison and Copying</h3>
                    <CodeBlock code={`p1 := Person{"Alice", 30, "NYC"}
p2 := Person{"Alice", 30, "NYC"}
p3 := Person{"Bob", 25, "LA"}
fmt.Printf("p1 == p2: %t\n", p1 == p2)
copy := p1
copy.Name = "Copy"`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">10. Struct Tags and JSON</h3>
                    <CodeBlock code={`type User struct {
  ID    int    ` + '`json:"id"`' + `
  Name  string ` + '`json:"name"`' + `
  Email string ` + '`json:"email"`' + `
  Age   int    ` + '`json:"age,omitempty"`' + `
}
user := User{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30}
jsonData, _ := json.Marshal(user)
fmt.Printf("JSON: %s\n", jsonData)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">11. Constructor Functions</h3>
                    <CodeBlock code={`func NewPerson(name string, age int) Person { return Person{Name: name, Age: age} }
newPerson := NewPerson("Alice", 25)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">12. Structs and Interfaces</h3>
                    <CodeBlock code={`type Shape interface { Area() float64 }
type Circle struct { Radius float64 }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
shapes := []Shape{Rectangle{4,5}, Circle{3}}
for _, shape := range shapes { fmt.Printf("Area: %.2f\n", shape.Area()) }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">13. Empty Structs</h3>
                    <CodeBlock code={`done := make(chan struct{})
go func() { done <- struct{}{} }()
<-done`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">14. Field Promotion</h3>
                    <CodeBlock code={`type Student struct {
  Person // Embedded
  School string
  GPA    float64
}
student := Student{Person: Person{"Alice", 20, "Boston"}, School: "MIT", GPA: 3.8}
student.Introduce() // Method from embedded Person`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">15. Struct Slices and Maps</h3>
                    <CodeBlock code={`people := []Person{
  {"Alice", 25, "NYC"},
  {"Bob", 30, "LA"},
  {"Charlie", 35, "Chicago"},
}
for _, p := range people { fmt.Println(p.Name) }`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Employee Directory</h3>
                    <p className="mb-4">Create a slice of Employee structs. Write functions to add, remove, and search employees by name or ID.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: JSON Serialization</h3>
                    <p className="mb-4">Define a struct with JSON tags. Write code to serialize and deserialize it using the encoding/json package.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Shape Interface</h3>
                    <p className="mb-4">Implement a Shape interface with Area() and Perimeter() methods. Create Rectangle and Circle types that satisfy the interface.</p>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">Try out struct operations in the interactive playground below.</p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main
import (
  "fmt"
  "encoding/json"
)
type Person struct { Name string; Age int; City string }
func main() {
  p := Person{"Alice", 30, "NYC"}
  fmt.Printf("Person: %+v\n", p)
  data, _ := json.Marshal(p)
  fmt.Printf("JSON: %s\n", data)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Person: {Name:Alice Age:30 City:NYC}
                        JSON: {"Name":"Alice","Age":30,"City":"NYC"}</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Structs"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 