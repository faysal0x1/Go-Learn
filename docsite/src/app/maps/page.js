import SectionLayout from '../../components/SectionLayout';
import CodeBlock from '../../components/CodeBlock';

export default function Maps() {
    const introduction = (
        <>
            <h2 className="text-2xl font-bold mb-4">Maps in Go</h2>
            <p className="mb-4">
                Maps in Go are hash tables that store key-value pairs. They are one of the most versatile
                data structures in Go, allowing you to store and retrieve values by their associated keys.
            </p>
            <p className="mb-4">
                Key features of maps in Go:
            </p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Key-value storage</li>
                <li>Dynamic size</li>
                <li>Reference type</li>
                <li>Fast lookups</li>
                <li>Unordered iteration</li>
                <li>Can store complex types</li>
            </ul>
        </>
    );

    const examples = (
        <>
            <h2 className="text-2xl font-bold mb-4">Map Examples</h2>

            <div className="space-y-8">
                <div>
                    <h3 className="text-xl font-semibold mb-2">1. Creating and Initializing Maps</h3>
                    <p className="mb-4">
                        Different ways to create and initialize maps:
                    </p>
                    <CodeBlock code={`// Declare an empty map
var emptyMap map[string]int

// Initialize using make
scores := make(map[string]int)
scores["Alice"] = 95
scores["Bob"] = 80

// Initialize using map literal
people := map[string]int{
    "John":  30,
    "Maria": 25,
    "Peter": 40,
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">2. Basic Map Operations</h3>
                    <p className="mb-4">
                        Working with map elements:
                    </p>
                    <CodeBlock code={`// Accessing elements
aliceScore := scores["Alice"]
fmt.Println("Alice's score:", aliceScore)

// Checking if key exists
bobScore, exists := scores["Bob"]
if exists {
    fmt.Println("Bob's score:", bobScore)
}

// Adding/updating elements
scores["Charlie"] = 85  // Add new
scores["Alice"] = 97    // Update existing

// Deleting elements
delete(scores, "Bob")

// Getting map length
fmt.Println("Number of entries:", len(scores))`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">3. Iterating Over Maps</h3>
                    <p className="mb-4">
                        Different ways to iterate over maps:
                    </p>
                    <CodeBlock code={`// Iterate over key-value pairs
for name, age := range people {
    fmt.Printf("%s is %d years old\\n", name, age)
}

// Iterate over keys only
for name := range people {
    fmt.Println(name)
}

// Iterate over values only
for _, age := range people {
    fmt.Println(age)
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">4. Complex Map Types</h3>
                    <p className="mb-4">
                        Maps can store complex types as values:
                    </p>
                    <CodeBlock code={`// Map with map values
studentScores := map[string]map[string]int{
    "Alice": {
        "Math":    95,
        "Science": 92,
        "English": 88,
    },
    "Bob": {
        "Math":    80,
        "Science": 85,
        "English": 90,
    },
}

// Map with slice values
interests := map[string][]string{
    "Alice": {"Reading", "Swimming", "Coding"},
    "Bob":   {"Gaming", "Football", "Movies"},
}`} />
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-2">5. Common Map Operations</h3>
                    <p className="mb-4">
                        Useful operations for working with maps:
                    </p>
                    <CodeBlock code={`// Copying a map
scoresCopy := make(map[string]int)
for k, v := range scores {
    scoresCopy[k] = v
}

// Merging maps
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"b": 3, "c": 4}
for k, v := range map2 {
    map1[k] = v
}

// Finding keys by value
ages := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 30}
var thirtyYearOlds []string
for name, age := range ages {
    if age == 30 {
        thirtyYearOlds = append(thirtyYearOlds, name)
    }
}

// Word frequency counter
text := "the quick brown fox jumps over the lazy dog"
wordFreq := make(map[string]int)
for _, word := range strings.Fields(text) {
    wordFreq[word]++
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
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Student Records</h3>
                    <p className="mb-4">
                        Create a program that manages student records using maps. Each student should have a
                        name, age, and a map of subject scores. Implement functions to add students, update
                        scores, and calculate averages.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Student: Alice
                            Age: 20
                            Scores: Math=95, Science=92, English=88
                            Average: 91.67</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Word Analysis</h3>
                    <p className="mb-4">
                        Create a program that analyzes text using maps. Count word frequencies, find the most
                        common words, and identify words that appear only once.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Text: "the quick brown fox jumps over the lazy dog"
                            Word frequencies:
                            the: 2
                            quick: 1
                            brown: 1
                            fox: 1
                            jumps: 1
                            over: 1
                            lazy: 1
                            dog: 1

                            Most common word: "the" (2 occurrences)
                            Unique words: 7</pre>
                    </div>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Cache Implementation</h3>
                    <p className="mb-4">
                        Implement a simple cache using maps. The cache should store key-value pairs with an
                        expiration time. Include functions to add, get, and remove items from the cache.
                    </p>
                    <div className="bg-gray-100 p-4 rounded">
                        <p className="font-medium">Expected Output:</p>
                        <pre className="mt-2">Adding items to cache...
                            Getting item: value1
                            Item expired: value2
                            Cache size: 1</pre>
                    </div>
                </div>
            </div>
        </>
    );

    const playground = (
        <>
            <h2 className="text-2xl font-bold mb-4">Interactive Playground</h2>
            <p className="mb-4">
                Try out different map operations in the interactive playground below.
            </p>
            <div className="bg-gray-50 p-6 rounded-lg">
                <CodeBlock code={`package main

import "fmt"

func main() {
    // Create and initialize a map
    scores := map[string]int{
        "Alice": 95,
        "Bob":   80,
        "Charlie": 85,
    }
    fmt.Println("Original scores:", scores)
    
    // Add a new entry
    scores["David"] = 90
    fmt.Println("After adding David:", scores)
    
    // Update an existing entry
    scores["Alice"] = 97
    fmt.Println("After updating Alice:", scores)
    
    // Check if a key exists
    if score, exists := scores["Bob"]; exists {
        fmt.Println("Bob's score:", score)
    }
    
    // Delete an entry
    delete(scores, "Charlie")
    fmt.Println("After deleting Charlie:", scores)
}`} />
                <div className="mt-4 p-4 bg-gray-100 rounded">
                    <h3 className="font-medium mb-2">Output:</h3>
                    <pre>Original scores: map[Alice:95 Bob:80 Charlie:85]
                        After adding David: map[Alice:95 Bob:80 Charlie:85 David:90]
                        After updating Alice: map[Alice:97 Bob:80 Charlie:85 David:90]
                        Bob's score: 80
                        After deleting Charlie: map[Alice:97 Bob:80 David:90]</pre>
                </div>
            </div>
        </>
    );

    return (
        <SectionLayout
            title="Maps"
            introduction={introduction}
            examples={examples}
            exercises={exercises}
            playground={playground}
        />
    );
} 