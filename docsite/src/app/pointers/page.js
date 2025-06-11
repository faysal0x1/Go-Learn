import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Pointers() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Pointers in Go</h2>
            <p className="mb-4">
                Pointers in Go allow you to store and manipulate memory addresses directly. They are essential for reference semantics, efficient data structures, and certain programming patterns.
            </p>
            <p className="mb-4">Key features of pointers in Go:</p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Store memory addresses of variables</li>
                <li>Enable pass-by-reference to functions</li>
                <li>Support pointer to pointer, pointer to struct, and more</li>
                <li>Allow efficient manipulation of large data structures</li>
                <li>Nil pointers and zero values</li>
                <li>Safe, with limited pointer arithmetic (via unsafe)</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Pointer Examples</h2>
            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Pointer Operations</h3>
                    <CodeBlock code={`var x int = 42
var p *int = &x
fmt.Printf("Value of x: %d\n", x)
fmt.Printf("Address of x: %p\n", &x)
fmt.Printf("Value of p (address): %p\n", p)
fmt.Printf("Value pointed to by p: %d\n", *p)
*p = 100
fmt.Printf("After *p = 100, x = %d\n", x)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Zero Value and new()</h3>
                    <CodeBlock code={`var ptr *int
fmt.Printf("Zero value pointer: %v\n", ptr)
fmt.Printf("Is nil: %t\n", ptr == nil)
numPtr := new(int)
*numPtr = 25
fmt.Printf("Value: %d, Address: %p\n", *numPtr, numPtr)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Pointer to Pointer</h3>
                    <CodeBlock code={`a := 10
pa := &a
ppa := &pa
fmt.Printf("a = %d\n", a)
fmt.Printf("*pa = %d\n", *pa)
fmt.Printf("**ppa = %d\n", **ppa)
**ppa = 20
fmt.Printf("After **ppa = 20, a = %d\n", a)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Pointers with Functions</h3>
                    <CodeBlock code={`func increment(n *int) { *n++ }
num := 5
increment(&num)
fmt.Printf("After increment: %d\n", num)
func swap(a, b *int) { *a, *b = *b, *a }
x, y := 10, 20
swap(&x, &y)
fmt.Printf("After swap: x=%d, y=%d\n", x, y)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Pointers with Structs</h3>
                    <CodeBlock code={`type Person struct { Name string; Age int }
p := Person{"Alice", 30}
updatePerson(&p, "Bob", 35)
func updatePerson(p *Person, name string, age int) {
  p.Name = name; p.Age = age
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">6. Pointers with Arrays and Slices</h3>
                    <CodeBlock code={`arr := [3]int{1,2,3}
modifyArray(&arr)
func modifyArray(arr *[3]int) { arr[0] = 99 }
// Slices are reference types
slice := []int{4,5,6}
modifySlice(slice)
func modifySlice(slice []int) { slice[0] = 100 }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">7. Returning Pointers from Functions</h3>
                    <CodeBlock code={`func createPerson(name string, age int) *Person {
  return &Person{Name: name, Age: age}
}
newPerson := createPerson("David", 40)
fmt.Printf("New person: %+v\n", *newPerson)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">8. Pointer Arithmetic (unsafe)</h3>
                    <CodeBlock code={`import "unsafe"
nums := [5]int{10,20,30,40,50}
ptr1 := &nums[0]
ptr2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr1)) + unsafe.Sizeof(int(0))))
fmt.Printf("Second element: %d\n", *ptr2)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">9. Comparing Pointers</h3>
                    <CodeBlock code={`val1, val2 := 100, 100
ptr_a := &val1
ptr_b := &val2
ptr_c := &val1
fmt.Printf("ptr_a == ptr_b: %t\n", ptr_a == ptr_b)
fmt.Printf("ptr_a == ptr_c: %t\n", ptr_a == ptr_c)
fmt.Printf("*ptr_a == *ptr_b: %t\n", *ptr_a == *ptr_b)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">10. Pointers with Interfaces and Methods</h3>
                    <CodeBlock code={`type Writer interface { Write(data string) }
type FileWriter struct{ filename string }
func (fw *FileWriter) Write(data string) { fmt.Println("FileWriter:", data) }
type ConsoleWriter struct{}
func (cw ConsoleWriter) Write(data string) { fmt.Println("ConsoleWriter:", data) }
var writer Writer
fw := &FileWriter{"test.txt"}
writer = fw
writer.Write("Hello from pointer receiver")
cw := ConsoleWriter{}
writer = cw
writer.Write("Hello from value receiver")`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">11. Method Receivers: Pointer vs Value</h3>
                    <CodeBlock code={`type Rectangle struct { Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (r *Rectangle) Scale(factor float64) { r.Width *= factor; r.Height *= factor }
rect := Rectangle{5,3}
area := rect.Area()
rect.Scale(2)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">12. Linked List with Pointers</h3>
                    <CodeBlock code={`type Node struct { Value int; Next *Node }
head := &Node{Value: 1}
head.Next = &Node{Value: 2}
head.Next.Next = &Node{Value: 3}
printList(head)
func printList(head *Node) {
  for head != nil { fmt.Printf("%d ", head.Value); head = head.Next }
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">13. Optional Values with Pointers</h3>
                    <CodeBlock code={`type User struct { Name string; Email *string }
user1 := User{"Alice", stringPtr("alice@example.com")}
user2 := User{"Bob", nil}
func stringPtr(s string) *string { return &s }
func printUser(u User) {
  if u.Email != nil { fmt.Printf("%s <%s>\n", u.Name, *u.Email) }
  else { fmt.Printf("%s <no email>\n", u.Name) }
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">14. Pointer Slices</h3>
                    <CodeBlock code={`people := []*Person{
  {"Alice", 25}, {"Bob", 30}, {"Charlie", 35},
}
for _, p := range people { p.Age++ }`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">15. Function Pointers</h3>
                    <CodeBlock code={`var operation func(int, int) int
operation = add
fmt.Printf("Result: %d\n", operation(5,3))
operations := []func(int, int) int{add, subtract, multiply}
for i, op := range operations {
  fmt.Printf("Operation %d: %d\n", i, op(10,3))
}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">16. Best Practices & Pitfalls</h3>
                    <CodeBlock code={`// Check for nil before dereferencing
var maybeNil *int
if maybeNil != nil { fmt.Printf("Value: %d\n", *maybeNil) } else { fmt.Println("Pointer is nil!") }
// Use pointers for large structs to avoid copying
largeStruct := LargeStruct{Data: [1000]int{}}
processLargeStruct(&largeStruct)
// Pitfall: loop variable addresses
var pointers []*int
values := []int{1,2,3}
for _, v := range values { pointers = append(pointers, &v) } // Wrong: all point to same variable
// Correct way
pointers = nil
for _, v := range values { v := v; pointers = append(pointers, &v) }`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Swap Function</h3>
                    <p className="mb-4">Write a function that swaps two integers using pointers.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Linked List</h3>
                    <p className="mb-4">Implement a singly linked list using pointers. Add functions to insert, delete, and print nodes.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Optional Email</h3>
                    <p className="mb-4">Create a User struct with an optional email field (pointer to string). Write a function to print user info, handling nil emails gracefully.</p>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">Try out pointer operations in the interactive playground below.</p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main
import "fmt"
func main() {
  x := 42
  p := &x
  fmt.Printf("x: %d, *p: %d\n", x, *p)
  *p = 99
  fmt.Printf("x: %d, *p: %d\n", x, *p)
  arr := [3]int{1,2,3}
  modifyArray(&arr)
  fmt.Printf("Modified array: %v\n", arr)
}
func modifyArray(arr *[3]int) { arr[0] = 100 }`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>x: 42, *p: 42
                        x: 99, *p: 99
                        Modified array: [100 2 3]</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Pointers"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 