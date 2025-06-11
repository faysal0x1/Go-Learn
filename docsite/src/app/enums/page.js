import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Enums() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Enums in Go</h2>
            <p className="mb-4">
                Go does not have a built-in enum type, but you can create enumerated constants using <code>const</code> and <code>iota</code>. Enums are useful for representing a set of related named values, such as states, options, or categories.
            </p>
            <p className="mb-4">Key features of enums in Go:</p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Use <code>const</code> and <code>iota</code> for integer enums</li>
                <li>Support for custom values and string enums</li>
                <li>Attach methods to enum types for validation and string conversion</li>
                <li>Bit flag enums for permissions and options</li>
                <li>Enum iteration, parsing, and JSON serialization</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Enum Examples</h2>
            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic iota Enum</h3>
                    <CodeBlock code={`type Color int
const (
  Red Color = iota
  Green
  Blue
)
fmt.Printf("Red: %d, Green: %d, Blue: %d\n", Red, Green, Blue)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Enum with Custom Values</h3>
                    <CodeBlock code={`type Size int
const (
  Small Size = 10
  Medium Size = 20
  Large Size = 30
)
fmt.Printf("Small: %d, Medium: %d, Large: %d\n", Small, Medium, Large)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">3. String Enums</h3>
                    <CodeBlock code={`type Day string
const (
  Monday Day = "monday"
  Tuesday Day = "tuesday"
  Wednesday Day = "wednesday"
)
fmt.Printf("Monday: %s\n", Monday)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Enum with Methods</h3>
                    <CodeBlock code={`type Status int
const (
  StatusActive Status = iota + 1
  StatusInactive
  StatusPending
)
func (s Status) String() string {
  switch s {
  case StatusActive: return "active"
  case StatusInactive: return "inactive"
  case StatusPending: return "pending"
  default: return "unknown"
  }
}
status := StatusActive
fmt.Printf("Status: %s\n", status)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Bit Flag Enums</h3>
                    <CodeBlock code={`type Permission int
const (
  Read Permission = 1 << iota
  Write
  Execute
)
permissions := Read | Write
fmt.Printf("Has read: %t\n", permissions&Read != 0)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">6. Enum Iteration</h3>
                    <CodeBlock code={`func AllStatuses() []Status {
  return []Status{StatusActive, StatusInactive, StatusPending}
}
for _, status := range AllStatuses() {
  fmt.Printf("- %s\n", status)
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">7. Enum Parsing and Validation</h3>
                    <CodeBlock code={`func ParseStatus(s string) (Status, error) {
  switch strings.ToLower(s) {
  case "active": return StatusActive, nil
  case "inactive": return StatusInactive, nil
  case "pending": return StatusPending, nil
  default: return 0, fmt.Errorf("invalid status")
  }
}
status, err := ParseStatus("active")`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">8. Enum in Switch Statements</h3>
                    <CodeBlock code={`func handleStatus(status Status) {
  switch status {
  case StatusActive:
    fmt.Println("Active!")
  case StatusInactive:
    fmt.Println("Inactive!")
  case StatusPending:
    fmt.Println("Pending!")
  }
}
handleStatus(StatusActive)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">9. JSON Serialization</h3>
                    <CodeBlock code={`type User struct {
  Name   string ` + '`json:"name"`' + `
  Status Status ` + '`json:"status"`' + `
}
user := User{"Alice", StatusActive}
data, _ := json.Marshal(user)
fmt.Printf("JSON: %s\n", data)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">10. Enum Maps and Slices</h3>
                    <CodeBlock code={`statusDescriptions := map[Status]string{
  StatusActive: "User is active",
  StatusInactive: "User is inactive",
}
for status, desc := range statusDescriptions {
  fmt.Printf("%s: %s\n", status, desc)
}`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: HTTP Status Enum</h3>
                    <p className="mb-4">Create an enum for HTTP status codes with methods for code, message, and success check. Use it in a function that prints status info.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Enum Parsing</h3>
                    <p className="mb-4">Write a function that parses a string into a Day enum and returns an error for invalid input.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Bit Flags</h3>
                    <p className="mb-4">Implement a permissions system using bit flag enums. Write functions to add, remove, and check permissions.</p>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">Try out enum patterns in the interactive playground below.</p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main
import (
  "fmt"
  "strings"
)
type Status int
const (
  StatusActive Status = iota + 1
  StatusInactive
  StatusPending
)
func (s Status) String() string {
  switch s {
  case StatusActive: return "active"
  case StatusInactive: return "inactive"
  case StatusPending: return "pending"
  default: return "unknown"
  }
}
func ParseStatus(s string) (Status, error) {
  switch strings.ToLower(s) {
  case "active": return StatusActive, nil
  case "inactive": return StatusInactive, nil
  case "pending": return StatusPending, nil
  default: return 0, fmt.Errorf("invalid status")
  }
}
func main() {
  status, err := ParseStatus("active")
  if err == nil {
    fmt.Printf("Parsed: %s\n", status)
  }
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Parsed: active</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Enums"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 