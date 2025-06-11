import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Generics() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Generics in Go</h2>
            <p className="mb-4">
                Generics in Go allow you to write flexible and reusable code by parameterizing types. Introduced in Go 1.18, generics enable you to define functions, types, and data structures that work with any type, while maintaining type safety.
            </p>
            <p className="mb-4">Key features of generics in Go:</p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Type parameters for functions, types, and methods</li>
                <li>Type constraints for restricting allowed types</li>
                <li>Support for generic data structures and algorithms</li>
                <li>Type inference in many cases</li>
                <li>Improved code reuse and safety</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Generic Examples</h2>
            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Basic Generic Function</h3>
                    <CodeBlock code={`func Max[T comparable](a, b T) T {
  if a > b { return a } else { return b }
}
fmt.Println(Max(5, 3))
fmt.Println(Max("hello", "world"))`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Multiple Type Parameters</h3>
                    <CodeBlock code={`type Pair[T, U any] struct { First T; Second U }
func MakePair[T, U any](first T, second U) Pair[T, U] {
  return Pair[T, U]{First: first, Second: second}
}
pair := MakePair("Alice", 30)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Generic Slice Operations</h3>
                    <CodeBlock code={`func First[T any](slice []T) T { return slice[0] }
func Contains[T comparable](slice []T, item T) bool {
  for _, v := range slice { if v == item { return true } }
  return false
}
ints := []int{1,2,3}
fmt.Println(First(ints))
fmt.Println(Contains(ints, 2))`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Generic Data Structures</h3>
                    <CodeBlock code={`type Stack[T any] struct { items []T }
func (s *Stack[T]) Push(item T) { s.items = append(s.items, item) }
func (s *Stack[T]) Pop() T {
  n := len(s.items)
  item := s.items[n-1]
  s.items = s.items[:n-1]
  return item
}
stack := Stack[int]{}
stack.Push(10); stack.Push(20)
fmt.Println(stack.Pop())`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Type Constraints</h3>
                    <CodeBlock code={`type Numeric interface {
  int | int8 | int16 | int32 | int64 |
  uint | uint8 | uint16 | uint32 | uint64 |
  float32 | float64
}
func Sum[T Numeric](values []T) T {
  var total T
  for _, v := range values { total += v }
  return total
}
fmt.Println(Sum([]int{1,2,3}))`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">6. Generic Interfaces</h3>
                    <CodeBlock code={`type Container[T any] interface {
  Add(item T)
  Get(index int) T
  Size() int
}
type SliceContainer[T any] struct { items []T }
func (sc *SliceContainer[T]) Add(item T) { sc.items = append(sc.items, item) }
func (sc *SliceContainer[T]) Get(i int) T { return sc.items[i] }
func (sc *SliceContainer[T]) Size() int { return len(sc.items) }
var c Container[string] = &SliceContainer[string]{}
c.Add("hello")`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">7. Functional Programming with Generics</h3>
                    <CodeBlock code={`func Map[T, U any](slice []T, fn func(T) U) []U {
  result := make([]U, len(slice))
  for i, v := range slice { result[i] = fn(v) }
  return result
}
fmt.Println(Map([]int{1,2,3}, func(x int) int { return x*2 }))`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">8. Generic Optional Type</h3>
                    <CodeBlock code={`type Optional[T any] struct { value T; hasValue bool }
func Some[T any](value T) Optional[T] { return Optional[T]{value, true} }
func None[T any]() Optional[T] { return Optional[T]{hasValue: false} }
opt := Some(42)
fmt.Println(opt.value, opt.hasValue)`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">9. Generic Tree Structure</h3>
                    <CodeBlock code={`type TreeNode[T any] struct {
  Value T
  Left, Right *TreeNode[T]
}
tree := &TreeNode[int]{Value: 10}
tree.Left = &TreeNode[int]{Value: 5}`} />
                </div>
                <div>
                    <h3 className="text-xl font-semibold mb-2">10. Generic Result Type</h3>
                    <CodeBlock code={`type Result[T any] struct { value T; err error }
func Ok[T any](value T) Result[T] { return Result[T]{value: value} }
func Err[T any](err error) Result[T] { return Result[T]{err: err} }
res := Ok(10)
if res.err == nil { fmt.Println(res.value) }`} />
                </div>
            </div>
        </>
    );

    const exercises = (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Generic Stack</h3>
                    <p className="mb-4">Implement a generic stack with Push, Pop, and Peek methods. Test it with int and string types.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Generic Map Function</h3>
                    <p className="mb-4">Write a generic Map function that applies a transformation to each element of a slice and returns a new slice.</p>
                </div>
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Type Constraints</h3>
                    <p className="mb-4">Create a generic function that finds the minimum value in a slice, using a type constraint for ordered types.</p>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">Try out generics in the interactive playground below.</p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main
import "fmt"
func Max[T comparable](a, b T) T {
  if a > b { return a } else { return b }
}
func main() {
  fmt.Println(Max(5, 3))
  fmt.Println(Max("hello", "world"))
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>5
                        world</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Generics"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 