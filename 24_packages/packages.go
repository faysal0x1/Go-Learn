package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	// External package (you'd need to install: go mod init && go get github.com/gorilla/mux)
	// "github.com/gorilla/mux"
)

func main() {
	fmt.Println("=== Go Packages Examples ===\n")

	// 1. Standard library packages
	fmt.Println("1. Standard Library Packages:")
	standardLibraryExamples()

	// 2. String manipulation packages
	fmt.Println("\n2. String Packages:")
	stringPackageExamples()

	// 3. Math and numeric packages
	fmt.Println("\n3. Math Packages:")
	mathPackageExamples()

	// 4. Time and date packages
	fmt.Println("\n4. Time Packages:")
	timePackageExamples()

	// 5. I/O packages
	fmt.Println("\n5. I/O Packages:")
	ioPackageExamples()

	// 6. Encoding packages
	fmt.Println("\n6. Encoding Packages:")
	encodingPackageExamples()

	// 7. Crypto packages
	fmt.Println("\n7. Crypto Packages:")
	cryptoPackageExamples()

	// 8. Network packages
	fmt.Println("\n8. Network Packages:")
	networkPackageExamples()

	// 9. File and path packages
	fmt.Println("\n9. File/Path Packages:")
	filePathPackageExamples()

	// 10. Regular expression packages
	fmt.Println("\n10. Regex Packages:")
	regexPackageExamples()

	// 11. Sort package
	fmt.Println("\n11. Sort Package:")
	sortPackageExamples()

	// 12. Sync packages
	fmt.Println("\n12. Sync Packages:")
	syncPackageExamples()

	// 13. Context package
	fmt.Println("\n13. Context Package:")
	contextPackageExamples()

	// 14. Error handling packages
	fmt.Println("\n14. Error Packages:")
	errorPackageExamples()

	// 15. Custom package creation
	fmt.Println("\n15. Custom Package Creation:")
	customPackageExamples()

	// 16. Package initialization
	fmt.Println("\n16. Package Initialization:")
	packageInitExamples()

	// 17. Package visibility
	fmt.Println("\n17. Package Visibility:")
	packageVisibilityExamples()

	// 18. Package documentation
	fmt.Println("\n18. Package Documentation:")
	packageDocumentationExamples()

	// 19. Package testing
	fmt.Println("\n19. Package Testing:")
	packageTestingExamples()

	// 20. Advanced package patterns
	fmt.Println("\n20. Advanced Package Patterns:")
	advancedPackagePatterns()
}

// 1. Standard library packages
func standardLibraryExamples() {
	// fmt package
	fmt.Printf("Printf: %d + %d = %d\n", 5, 3, 5+3)
	fmt.Fprintf(os.Stdout, "Fprintf to stdout\n")

	// log package
	log.Println("This is a log message")

	// os package
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Printf("Hostname: %s\n", hostname)
	}

	// strconv package
	str := strconv.Itoa(42)
	fmt.Printf("Integer to string: %s\n", str)

	num, err := strconv.Atoi("123")
	if err == nil {
		fmt.Printf("String to integer: %d\n", num)
	}
}

// 2. String manipulation packages
func stringPackageExamples() {
	text := "Hello, World! Go is awesome!"

	// strings package
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("Index of 'World': %d\n", strings.Index(text, "World"))
	fmt.Printf("Replace 'Go' with 'Golang': %s\n", strings.Replace(text, "Go", "Golang", 1))
	fmt.Printf("Split by space: %v\n", strings.Split(text, " "))
	fmt.Printf("ToUpper: %s\n", strings.ToUpper(text))
	fmt.Printf("ToLower: %s\n", strings.ToLower(text))
	fmt.Printf("TrimSpace: '%s'\n", strings.TrimSpace("  spaced  "))

	// String builder for efficient concatenation
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString(fmt.Sprintf("Line %d\n", i+1))
	}
	fmt.Printf("StringBuilder result:\n%s", builder.String())
}

// 3. Math and numeric packages
func mathPackageExamples() {
	// math package
	fmt.Printf("Pi: %.6f\n", math.Pi)
	fmt.Printf("E: %.6f\n", math.E)
	fmt.Printf("Sqrt(16): %.2f\n", math.Sqrt(16))
	fmt.Printf("Pow(2, 8): %.0f\n", math.Pow(2, 8))
	fmt.Printf("Sin(Pi/2): %.6f\n", math.Sin(math.Pi/2))
	fmt.Printf("Cos(Pi): %.6f\n", math.Cos(math.Pi))
	fmt.Printf("Max(10, 20): %.0f\n", math.Max(10, 20))
	fmt.Printf("Min(10, 20): %.0f\n", math.Min(10, 20))
	fmt.Printf("Abs(-5): %.0f\n", math.Abs(-5))
	fmt.Printf("Ceil(4.2): %.0f\n", math.Ceil(4.2))
	fmt.Printf("Floor(4.8): %.0f\n", math.Floor(4.8))
	fmt.Printf("Round(4.6): %.0f\n", math.Round(4.6))

	// math/rand package
	fmt.Printf("Random int: %d\n", generateRandomInt(1, 100))
	fmt.Printf("Random float: %.4f\n", generateRandomFloat())
}

func generateRandomInt(min, max int) int {
	return min + int(rand.Int63n(int64(max-min+1)))
}

func generateRandomFloat() float64 {
	return rand.Float64()
}

// 4. Time and date packages
func timePackageExamples() {
	now := time.Now()
	fmt.Printf("Current time: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Unix timestamp: %d\n", now.Unix())

	// Time parsing and formatting
	layout := "2006-01-02 15:04:05"
	parsed, err := time.Parse(layout, "2023-12-25 10:30:00")
	if err == nil {
		fmt.Printf("Parsed time: %s\n", parsed.Format("January 2, 2006"))
	}

	// Duration calculations
	future := now.Add(24 * time.Hour)
	fmt.Printf("24 hours from now: %s\n", future.Format("2006-01-02 15:04:05"))

	duration := future.Sub(now)
	fmt.Printf("Duration: %s\n", duration)

	// Time zones
	utc := now.UTC()
	fmt.Printf("UTC time: %s\n", utc.Format("2006-01-02 15:04:05 MST"))

	// Timer and ticker examples
	fmt.Println("Starting 3-second timer...")
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Timer expired!")
	}()

	// Ticker example (stopped quickly for demo)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		count := 0
		for range ticker.C {
			count++
			fmt.Printf("Tick %d\n", count)
			if count >= 3 {
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(4 * time.Second) // Wait for timer and ticker
}

// 5. I/O packages
func ioPackageExamples() {
	// bytes package
	var buffer bytes.Buffer
	buffer.WriteString("Hello ")
	buffer.WriteString("World!")
	fmt.Printf("Buffer content: %s\n", buffer.String())

	// bufio package
	text := "Line 1\nLine 2\nLine 3"
	reader := bufio.NewReader(strings.NewReader(text))

	fmt.Println("Reading line by line:")
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				fmt.Printf("  %s", line)
			}
			break
		}
		if err != nil {
			break
		}
		fmt.Printf("  %s", line)
	}

	// io.Copy example
	source := strings.NewReader("Data to copy")
	var dest bytes.Buffer
	copied, err := io.Copy(&dest, source)
	if err == nil {
		fmt.Printf("\nCopied %d bytes: %s\n", copied, dest.String())
	}
}

// 6. Encoding packages
func encodingPackageExamples() {
	// JSON encoding
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}

	person := Person{Name: "Alice", Age: 30, City: "New York"}

	// Marshal to JSON
	jsonData, err := json.Marshal(person)
	if err == nil {
		fmt.Printf("JSON: %s\n", jsonData)
	}

	// Unmarshal from JSON
	var parsedPerson Person
	err = json.Unmarshal(jsonData, &parsedPerson)
	if err == nil {
		fmt.Printf("Parsed: %+v\n", parsedPerson)
	}

	// Base64 encoding
	data := "Hello, World!"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("Base64 encoded: %s\n", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		fmt.Printf("Base64 decoded: %s\n", decoded)
	}
}

// 7. Crypto packages
func cryptoPackageExamples() {
	// MD5 hashing
	data := "Hello, World!"
	hash := md5.Sum([]byte(data))
	fmt.Printf("MD5 hash: %x\n", hash)

	// Random bytes generation
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err == nil {
		fmt.Printf("Random bytes: %x\n", randomBytes)
	}

	// Generate random string
	randomString := generateRandomString(10)
	fmt.Printf("Random string: %s\n", randomString)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	return string(bytes)
}

// 8. Network packages
func networkPackageExamples() {
	// URL parsing
	rawURL := "https://example.com:8080/path?param1=value1&param2=value2"
	parsedURL, err := url.Parse(rawURL)
	if err == nil {
		fmt.Printf("Scheme: %s\n", parsedURL.Scheme)
		fmt.Printf("Host: %s\n", parsedURL.Host)
		fmt.Printf("Path: %s\n", parsedURL.Path)
		fmt.Printf("Query: %s\n", parsedURL.RawQuery)
	}

	// HTTP client example (mock)
	fmt.Println("HTTP client example (mock response):")
	mockHTTPResponse()

	// URL encoding
	values := url.Values{}
	values.Add("name", "John Doe")
	values.Add("email", "john@example.com")
	encoded := values.Encode()
	fmt.Printf("URL encoded: %s\n", encoded)
}

func mockHTTPResponse() {
	// In a real scenario, you'd make an actual HTTP request
	fmt.Println("  Status: 200 OK")
	fmt.Println("  Content-Type: application/json")
	fmt.Println("  Body: {\"message\": \"Hello, World!\"}")
}

// 9. File and path packages
func filePathPackageExamples() {
	// filepath package
	path := "/home/user/documents/file.txt"
	fmt.Printf("Base: %s\n", filepath.Base(path))
	fmt.Printf("Dir: %s\n", filepath.Dir(path))
	fmt.Printf("Ext: %s\n", filepath.Ext(path))

	// Join paths
	joined := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Printf("Joined path: %s\n", joined)

	// Clean path
	messyPath := "home//user/../user/./documents/file.txt"
	cleaned := filepath.Clean(messyPath)
	fmt.Printf("Cleaned path: %s\n", cleaned)

	// Check if path is absolute
	fmt.Printf("Is absolute: %t\n", filepath.IsAbs(path))

	// Match patterns
	matched, err := filepath.Match("*.txt", "file.txt")
	if err == nil {
		fmt.Printf("Pattern match: %t\n", matched)
	}
}

// 10. Regular expression packages
func regexPackageExamples() {
	text := "Contact us at support@example.com or sales@company.org"

	// Compile regex pattern
	emailRegex, err := regexp.Compile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return
	}

	// Find all matches
	emails := emailRegex.FindAllString(text, -1)
	fmt.Printf("Found emails: %v\n", emails)

	// Replace matches
	anonymized := emailRegex.ReplaceAllString(text, "***@***.***")
	fmt.Printf("Anonymized: %s\n", anonymized)

	// Match with groups
	phoneRegex := regexp.MustCompile(`(\d{3})-(\d{3})-(\d{4})`)
	phone := "Call me at 555-123-4567"
	matches := phoneRegex.FindStringSubmatch(phone)
	if len(matches) > 0 {
		fmt.Printf("Full match: %s\n", matches[0])
		fmt.Printf("Area code: %s\n", matches[1])
		fmt.Printf("Exchange: %s\n", matches[2])
		fmt.Printf("Number: %s\n", matches[3])
	}

	// Validate patterns
	patterns := []string{"test@email.com", "invalid-email", "user@domain.co.uk"}
	for _, pattern := range patterns {
		if emailRegex.MatchString(pattern) {
			fmt.Printf("'%s' is a valid email\n", pattern)
		} else {
			fmt.Printf("'%s' is not a valid email\n", pattern)
		}
	}
}

// 11. Sort package
func sortPackageExamples() {
	// Sort integers
	ints := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", ints)
	sort.Ints(ints)
	fmt.Printf("Sorted: %v\n", ints)

	// Sort strings
	strings := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Original: %v\n", strings)
	sort.Strings(strings)
	fmt.Printf("Sorted: %v\n", strings)

	// Custom sorting
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("Sorted by age: %v\n", people)

	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Printf("Sorted by name: %v\n", people)

	// Check if sorted
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Is sorted: %t\n", sort.IntsAreSorted(nums))

	// Binary search
	index := sort.SearchInts(nums, 3)
	fmt.Printf("Index of 3: %d\n", index)
}

// 12. Sync packages
func syncPackageExamples() {
	// sync.WaitGroup
	var wg sync.WaitGroup
	fmt.Println("Starting goroutines...")

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d completed\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")

	// sync.Mutex
	var mutex sync.Mutex
	var counter int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)

	// sync.Once
	var once sync.Once
	initFunction := func() {
		fmt.Println("This will only be printed once")
	}

	for i := 0; i < 3; i++ {
		once.Do(initFunction)
	}
}

// 13. Context package
func contextPackageExamples() {
	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Work completed within timeout")
	case <-ctx.Done():
		fmt.Printf("Context cancelled: %v\n", ctx.Err())
	}

	// Context with values
	ctx = context.WithValue(context.Background(), "userID", "12345")
	processRequest(ctx)

	// Context cancellation
	ctx, cancel = context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Printf("Work cancelled: %v\n", ctx.Err())
	}
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Printf("Processing request for user: %s\n", userID)
	}
}

// 14. Error handling packages
func errorPackageExamples() {
	// Creating errors
	err1 := errors.New("simple error message")
	fmt.Printf("Error 1: %v\n", err1)

	// Formatted errors
	err2 := fmt.Errorf("error processing item %d: %w", 42, err1)
	fmt.Printf("Error 2: %v\n", err2)

	// Error wrapping and unwrapping
	if errors.Is(err2, err1) {
		fmt.Println("err2 wraps err1")
	}

	// Custom error types
	customErr := &CustomError{
		Code:    404,
		Message: "Resource not found",
	}
	fmt.Printf("Custom error: %v\n", customErr)

	// Error handling patterns
	result, err := riskyOperation(true)
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	} else {
		fmt.Printf("Operation succeeded: %s\n", result)
	}
}

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func riskyOperation(shouldFail bool) (string, error) {
	if shouldFail {
		return "", &CustomError{Code: 500, Message: "Something went wrong"}
	}
	return "Success!", nil
}

// 15. Custom package creation
func customPackageExamples() {
	fmt.Println("Custom package examples:")
	fmt.Println("To create a custom package:")
	fmt.Println("1. Create a directory with your package name")
	fmt.Println("2. Create .go files with 'package packagename' declaration")
	fmt.Println("3. Export functions/types by capitalizing first letter")
	fmt.Println("4. Import using the module path")

	// Example of what a custom package might look like:
	fmt.Println("\nExample package structure:")
	fmt.Println("myproject/")
	fmt.Println("├── main.go")
	fmt.Println("├── go.mod")
	fmt.Println("└── utils/")
	fmt.Println("    ├── math.go")
	fmt.Println("    └── string.go")

	// Usage example (simulated)
	result := simulateCustomPackageUsage()
	fmt.Printf("Custom package result: %s\n", result)
}

func simulateCustomPackageUsage() string {
	// This simulates using a custom package
	return "Result from custom utility function"
}

// 16. Package initialization
func packageInitExamples() {
	fmt.Println("Package initialization:")
	fmt.Println("- init() functions run automatically when package is imported")
	fmt.Println("- Multiple init() functions can exist in a package")
	fmt.Println("- They run in the order they appear in the source")
	fmt.Println("- Package-level variables are initialized before init()")

	// Demonstrate package initialization order
	demonstrateInitOrder()
}

var packageVar = initializePackageVar()

func initializePackageVar() string {
	fmt.Println("Package variable initialized")
	return "initialized"
}

func init() {
	fmt.Println("First init function called")
}

func init() {
	fmt.Println("Second init function called")
}

func demonstrateInitOrder() {
	fmt.Printf("Package variable value: %s\n", packageVar)
}

// 17. Package visibility
func packageVisibilityExamples() {
	fmt.Println("Package visibility rules:")
	fmt.Println("- Exported: Starts with capital letter (Public)")
	fmt.Println("- Unexported: Starts with lowercase letter (Private)")

	// Examples of exported vs unexported
	fmt.Println("\nExamples:")
	fmt.Println("type PublicStruct struct { ... }  // Exported")
	fmt.Println("type privateStruct struct { ... } // Unexported")
	fmt.Println("func PublicFunction() { ... }     // Exported")
	fmt.Println("func privateFunction() { ... }    // Unexported")

	// Demonstrate with actual examples
	public := PublicStruct{Name: "Public"}
	private := privateStruct{name: "private"}

	fmt.Printf("Public struct: %+v\n", public)
	fmt.Printf("Private struct: %+v\n", private)

	PublicFunction()
	privateFunction()
}

type PublicStruct struct {
	Name string
}

type privateStruct struct {
	name string
}

func PublicFunction() {
	fmt.Println("This function is exported")
}

func privateFunction() {
	fmt.Println("This function is unexported")
}

// 18. Package documentation
func packageDocumentationExamples() {
	fmt.Println("Package documentation:")
	fmt.Println("- Use comments immediately before package declaration")
	fmt.Println("- Document exported functions, types, and variables")
	fmt.Println("- Use 'go doc' command to view documentation")
	fmt.Println("- Follow Go documentation conventions")

	fmt.Println("\nDocumentation examples:")
	fmt.Println("// Package math provides basic mathematical operations.")
	fmt.Println("package math")
	fmt.Println("")
	fmt.Println("// Add returns the sum of two integers.")
	fmt.Println("func Add(a, b int) int { ... }")

	// Example of documented function
	result := DocumentedAdd(5, 3)
	fmt.Printf("Documented function result: %d\n", result)
}

// DocumentedAdd returns the sum of two integers.
// This function demonstrates proper Go documentation.
func DocumentedAdd(a, b int) int {
	return a + b
}

// 19. Package testing
func packageTestingExamples() {
	fmt.Println("Package testing:")
	fmt.Println("- Test files end with '_test.go'")
	fmt.Println("- Test functions start with 'Test' and take *testing.T")
	fmt.Println("- Use 'go test' command to run tests")
	fmt.Println("- Benchmark functions start with 'Benchmark'")

	fmt.Println("\nExample test file structure:")
	fmt.Println("package math")
	fmt.Println("")
	fmt.Println("import \"testing\"")
	fmt.Println("")
	fmt.Println("func TestAdd(t *testing.T) {")
	fmt.Println("    result := Add(2, 3)")
	fmt.Println("    if result != 5 {")
	fmt.Println("        t.Errorf(\"Expected 5, got %d\", result)")
	fmt.Println("    }")
	fmt.Println("}")

	// Simulate running tests
	simulateTestRun()
}

func simulateTestRun() {
	fmt.Println("\nSimulated test run:")
	fmt.Println("$ go test")
	fmt.Println("PASS")
	fmt.Println("ok      mypackage    0.002s")
}

// 20. Advanced package patterns
func advancedPackagePatterns() {
	fmt.Println("Advanced package patterns:")

	// 1. Interface segregation
	fmt.Println("\n1. Interface segregation:")
	demonstrateInterfaceSegregation()

	// 2. Dependency injection
	fmt.Println("\n2. Dependency injection:")
	demonstrateDependencyInjection()

	// 3. Package organization
	fmt.Println("\n3. Package organization patterns:")
	demonstratePackageOrganization()

	// 4. Plugin architecture
	fmt.Println("\n4. Plugin architecture:")
	demonstratePluginArchitecture()

	// 5. Package versioning
	fmt.Println("\n5. Package versioning:")
	demonstratePackageVersioning()
}

// Interface segregation example
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

func demonstrateInterfaceSegregation() {
	fmt.Println("Separate interfaces for different responsibilities")
	var r Reader = &bytes.Buffer{}
	var w Writer = &bytes.Buffer{}
	var rw ReadWriter = &bytes.Buffer{}

	fmt.Printf("Reader: %T\n", r)
	fmt.Printf("Writer: %T\n", w)
	fmt.Printf("ReadWriter: %T\n", rw)
}

// Dependency injection example
type Service interface {
	Process(data string) string
}

type DatabaseService struct{}

func (d *DatabaseService) Process(data string) string {
	return "Processed by database: " + data
}

type Application struct {
	service Service
}

func NewApplication(service Service) *Application {
	return &Application{service: service}
}

func (a *Application) Run(data string) string {
	return a.service.Process(data)
}

func demonstrateDependencyInjection() {
	service := &DatabaseService{}
	app := NewApplication(service)
	result := app.Run("test data")
	fmt.Printf("Application result: %s\n", result)
}

func demonstratePackageOrganization() {
	fmt.Println("Common package organization patterns:")
	fmt.Println("- Domain-driven design")
	fmt.Println("- Layered architecture")
	fmt.Println("- Feature-based organization")
	fmt.Println("- Hexagonal architecture")

	fmt.Println("\nExample project structure:")
	fmt.Println("project/")
	fmt.Println("├── cmd/           # Main applications")
	fmt.Println("├── internal/      # Private packages")
	fmt.Println("├── pkg/           # Public packages")
	fmt.Println("├── api/           # API definitions")
	fmt.Println("├── web/           # Web assets")
	fmt.Println("└── docs/          # Documentation")
}

func demonstratePluginArchitecture() {
	fmt.Println("Plugin architecture allows extending functionality:")
	fmt.Println("- Define interfaces for plugins")
	fmt.Println("- Load plugins at runtime")
	fmt.Println("- Use Go's plugin package for dynamic loading")

	// Simulated plugin system
	plugins := []Plugin{
		&TextPlugin{},
		&NumberPlugin{},
	}

	for _, plugin := range plugins {
		result := plugin.Execute("test")
		fmt.Printf("Plugin %s result: %s\n", plugin.Name(), result)
	}
}

type Plugin interface {
	Name() string
	Execute(input string) string
}

type TextPlugin struct{}

func (p *TextPlugin) Name() string {
	return "TextPlugin"
}

func (p *TextPlugin) Execute(input string) string {
	return "Text: " + input
}

type NumberPlugin struct{}

func (p *NumberPlugin) Name() string {
	return "NumberPlugin"
}

func (p *NumberPlugin) Execute(input string) string {
	return "Number: " + input
}

func demonstratePackageVersioning() {
	fmt.Println("Package versioning with Go modules:")
	fmt.Println("- Semantic versioning (v1.2.3)")
	fmt.Println("- Major version in import path for v2+")
	fmt.Println("- Use go.mod file to manage dependencies")

	fmt.Println("\nExample go.mod file:")
	fmt.Println("module github.com/user/project")
	fmt.Println("")
	fmt.Println("go 1.19")
	fmt.Println("")
	fmt.Println("require (")
	fmt.Println("    github.com/gorilla/mux v1.8.0")
	fmt.Println("    github.com/lib/pq v1.10.7")
	fmt.Println(")")

	fmt.Println("\nVersioning commands:")
	fmt.Println("- go mod init: Initialize module")
	fmt.Println("- go get: Add/update dependencies")
	fmt.Println("- go mod tidy: Clean up dependencies")
	fmt.Println("- go mod vendor: Create vendor directory")
}
