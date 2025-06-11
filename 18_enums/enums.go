package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Go Enums Examples ===\n")

	// 1. Basic iota enum
	fmt.Println("1. Basic iota enum:")
	fmt.Printf("Red: %d\n", Red)
	fmt.Printf("Green: %d\n", Green)
	fmt.Printf("Blue: %d\n", Blue)

	// 2. Enum with custom values
	fmt.Println("\n2. Enum with custom values:")
	fmt.Printf("Small: %d\n", Small)
	fmt.Printf("Medium: %d\n", Medium)
	fmt.Printf("Large: %d\n", Large)

	// 3. String enums
	fmt.Println("\n3. String enums:")
	fmt.Printf("Monday: %s\n", Monday)
	fmt.Printf("Tuesday: %s\n", Tuesday)
	fmt.Printf("Wednesday: %s\n", Wednesday)

	// 4. Enum with methods
	fmt.Println("\n4. Enum with methods:")
	status := StatusActive
	fmt.Printf("Status: %s\n", status)
	fmt.Printf("Is valid: %t\n", status.IsValid())
	fmt.Printf("String representation: %s\n", status.String())

	// 5. Enum validation
	fmt.Println("\n5. Enum validation:")
	validStatus := StatusActive
	invalidStatus := Status(99)

	fmt.Printf("%s is valid: %t\n", validStatus, validStatus.IsValid())
	fmt.Printf("%s is valid: %t\n", invalidStatus, invalidStatus.IsValid())

	// 6. Enum parsing from string
	fmt.Println("\n6. Parsing enums from strings:")
	parsedStatus, err := ParseStatus("active")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Parsed status: %s\n", parsedStatus)
	}

	_, err = ParseStatus("invalid")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 7. Bit flag enums
	fmt.Println("\n7. Bit flag enums:")
	fmt.Printf("Read: %b (%d)\n", Read, Read)
	fmt.Printf("Write: %b (%d)\n", Write, Write)
	fmt.Printf("Execute: %b (%d)\n", Execute, Execute)

	permissions := Read | Write
	fmt.Printf("Read+Write: %b (%d)\n", permissions, permissions)
	fmt.Printf("Has read: %t\n", permissions&Read != 0)
	fmt.Printf("Has execute: %t\n", permissions&Execute != 0)

	// 8. Enum iteration
	fmt.Println("\n8. Enum iteration:")
	fmt.Println("All statuses:")
	for _, status := range AllStatuses() {
		fmt.Printf("- %s (%d)\n", status, status)
	}

	// 9. Enum switch statements
	fmt.Println("\n9. Switch with enums:")
	handleStatus(StatusActive)
	handleStatus(StatusInactive)
	handleStatus(StatusPending)

	// 10. JSON serialization/deserialization
	fmt.Println("\n10. JSON serialization:")
	user := User{
		Name:   "Alice",
		Status: StatusActive,
		Role:   RoleAdmin,
	}

	jsonData, _ := json.Marshal(user)
	fmt.Printf("JSON: %s\n", jsonData)

	var parsedUser User
	json.Unmarshal(jsonData, &parsedUser)
	fmt.Printf("Parsed: %+v\n", parsedUser)

	// 11. Enum with zero value handling
	fmt.Println("\n11. Zero value handling:")
	var defaultPriority Priority
	fmt.Printf("Default priority: %s\n", defaultPriority)
	fmt.Printf("Is unknown: %t\n", defaultPriority == PriorityUnknown)

	// 12. Enum comparison
	fmt.Println("\n12. Enum comparison:")
	p1 := PriorityHigh
	p2 := PriorityMedium
	p3 := PriorityHigh

	fmt.Printf("%s == %s: %t\n", p1, p2, p1 == p2)
	fmt.Printf("%s == %s: %t\n", p1, p3, p1 == p3)
	fmt.Printf("%s > %s: %t\n", p1, p2, p1 > p2)

	// 13. Enum maps and lookups
	fmt.Println("\n13. Enum maps:")
	statusDescriptions := map[Status]string{
		StatusActive:   "User is currently active",
		StatusInactive: "User account is disabled",
		StatusPending:  "User registration is pending",
	}

	for status, description := range statusDescriptions {
		fmt.Printf("%s: %s\n", status, description)
	}

	// 14. Enum slices
	fmt.Println("\n14. Enum slices:")
	allowedStatuses := []Status{StatusActive, StatusPending}
	testStatus := StatusActive

	fmt.Printf("Testing if %s is allowed:\n", testStatus)
	allowed := false
	for _, status := range allowedStatuses {
		if status == testStatus {
			allowed = true
			break
		}
	}
	fmt.Printf("Result: %t\n", allowed)

	// 15. Complex enum with multiple fields
	fmt.Println("\n15. Complex enums:")
	httpStatus := HTTPStatusOK
	fmt.Printf("HTTP Status: %d %s\n", httpStatus.Code(), httpStatus.Message())
	fmt.Printf("Is success: %t\n", httpStatus.IsSuccess())

	// 16. Enum state machine
	fmt.Println("\n16. State machine with enums:")
	order := &Order{State: OrderStatePending}
	fmt.Printf("Initial state: %s\n", order.State)

	order.Process()
	order.Ship()
	order.Deliver()

	// 17. Enum with custom types
	fmt.Println("\n17. Custom type enums:")
	temp := TemperatureHot
	fmt.Printf("Temperature: %s (%.1f°C)\n", temp, temp.Celsius())
	fmt.Printf("Description: %s\n", temp.Description())

	// 18. Enum validation in functions
	fmt.Println("\n18. Function validation:")
	err = processOrder(OrderStateShipped)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = processOrder(OrderState(99))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 19. Enum with interfaces
	fmt.Println("\n19. Enum implementing interfaces:")
	var logger Logger = LogLevelInfo
	logger.Log("This is an info message")

	logger = LogLevelError
	logger.Log("This is an error message")

	// 20. Advanced enum patterns
	fmt.Println("\n20. Advanced patterns:")

	// Enum factory
	vehicle := CreateVehicle(VehicleTypeCar)
	fmt.Printf("Created vehicle: %s with %d wheels\n",
		vehicle.Type(), vehicle.Wheels())

	// Enum with behavior
	fmt.Println("File operations:")
	operations := []FileOperation{FileOperationRead, FileOperationWrite, FileOperationDelete}
	filename := "test.txt"

	for _, op := range operations {
		op.Execute(filename)
	}
}

// 1. Basic iota enum
type Color int

const (
	Red Color = iota
	Green
	Blue
)

// 2. Enum with custom values
type Size int

const (
	Small  Size = 10
	Medium Size = 20
	Large  Size = 30
)

// 3. String enums
type Day string

const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"
)

// 4. Enum with methods
type Status int

const (
	StatusActive Status = iota + 1
	StatusInactive
	StatusPending
)

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "active"
	case StatusInactive:
		return "inactive"
	case StatusPending:
		return "pending"
	default:
		return "unknown"
	}
}

func (s Status) IsValid() bool {
	return s >= StatusActive && s <= StatusPending
}

// 5. Enum validation and parsing
func ParseStatus(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "active":
		return StatusActive, nil
	case "inactive":
		return StatusInactive, nil
	case "pending":
		return StatusPending, nil
	default:
		return 0, fmt.Errorf("invalid status: %s", s)
	}
}

// 6. Bit flag enums
type Permission int

const (
	Read Permission = 1 << iota
	Write
	Execute
)

// 7. Enum iteration
func AllStatuses() []Status {
	return []Status{StatusActive, StatusInactive, StatusPending}
}

// 8. Switch with enums
func handleStatus(status Status) {
	switch status {
	case StatusActive:
		fmt.Println("User is active - full access granted")
	case StatusInactive:
		fmt.Println("User is inactive - access denied")
	case StatusPending:
		fmt.Println("User is pending - limited access")
	default:
		fmt.Println("Unknown status")
	}
}

// 9. JSON serialization
type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
	RoleGuest Role = "guest"
)

type User struct {
	Name   string `json:"name"`
	Status Status `json:"status"`
	Role   Role   `json:"role"`
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	parsed, err := ParseStatus(str)
	if err != nil {
		return err
	}

	*s = parsed
	return nil
}

// 10. Priority enum with zero value
type Priority int

const (
	PriorityUnknown Priority = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
	PriorityCritical
)

func (p Priority) String() string {
	switch p {
	case PriorityUnknown:
		return "unknown"
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	case PriorityCritical:
		return "critical"
	default:
		return "invalid"
	}
}

// 11. HTTP Status enum with methods
type HTTPStatus int

const (
	HTTPStatusOK                  HTTPStatus = 200
	HTTPStatusNotFound            HTTPStatus = 404
	HTTPStatusInternalServerError HTTPStatus = 500
)

func (h HTTPStatus) Code() int {
	return int(h)
}

func (h HTTPStatus) Message() string {
	switch h {
	case HTTPStatusOK:
		return "OK"
	case HTTPStatusNotFound:
		return "Not Found"
	case HTTPStatusInternalServerError:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
}

func (h HTTPStatus) IsSuccess() bool {
	return h >= 200 && h < 300
}

// 12. State machine with enums
type OrderState int

const (
	OrderStatePending OrderState = iota
	OrderStateProcessing
	OrderStateShipped
	OrderStateDelivered
	OrderStateCancelled
)

func (os OrderState) String() string {
	switch os {
	case OrderStatePending:
		return "pending"
	case OrderStateProcessing:
		return "processing"
	case OrderStateShipped:
		return "shipped"
	case OrderStateDelivered:
		return "delivered"
	case OrderStateCancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

type Order struct {
	State OrderState
}

func (o *Order) Process() error {
	if o.State != OrderStatePending {
		return fmt.Errorf("cannot process order in state: %s", o.State)
	}
	o.State = OrderStateProcessing
	fmt.Printf("Order state changed to: %s\n", o.State)
	return nil
}

func (o *Order) Ship() error {
	if o.State != OrderStateProcessing {
		return fmt.Errorf("cannot ship order in state: %s", o.State)
	}
	o.State = OrderStateShipped
	fmt.Printf("Order state changed to: %s\n", o.State)
	return nil
}

func (o *Order) Deliver() error {
	if o.State != OrderStateShipped {
		return fmt.Errorf("cannot deliver order in state: %s", o.State)
	}
	o.State = OrderStateDelivered
	fmt.Printf("Order state changed to: %s\n", o.State)
	return nil
}

// 13. Temperature enum with values
type Temperature float64

const (
	TemperatureFreezing Temperature = 0.0
	TemperatureCold     Temperature = 10.0
	TemperatureWarm     Temperature = 20.0
	TemperatureHot      Temperature = 30.0
)

func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) String() string {
	switch t {
	case TemperatureFreezing:
		return "freezing"
	case TemperatureCold:
		return "cold"
	case TemperatureWarm:
		return "warm"
	case TemperatureHot:
		return "hot"
	default:
		return "unknown"
	}
}

func (t Temperature) Description() string {
	switch t {
	case TemperatureFreezing:
		return "Water freezes"
	case TemperatureCold:
		return "Need a jacket"
	case TemperatureWarm:
		return "Comfortable"
	case TemperatureHot:
		return "Very warm"
	default:
		return "Unknown temperature"
	}
}

// 14. Function validation with enums
func processOrder(state OrderState) error {
	if state < OrderStatePending || state > OrderStateCancelled {
		return fmt.Errorf("invalid order state: %d", state)
	}
	fmt.Printf("Processing order in state: %s\n", state)
	return nil
}

// 15. Enum implementing interfaces
type Logger interface {
	Log(message string)
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

func (l LogLevel) Log(message string) {
	prefix := l.String()
	fmt.Printf("[%s] %s\n", strings.ToUpper(prefix), message)
}

func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "debug"
	case LogLevelInfo:
		return "info"
	case LogLevelWarning:
		return "warning"
	case LogLevelError:
		return "error"
	default:
		return "unknown"
	}
}

// 16. Vehicle enum with factory pattern
type VehicleType int

const (
	VehicleTypeCar VehicleType = iota
	VehicleTypeBike
	VehicleTypeTruck
)

type Vehicle interface {
	Type() string
	Wheels() int
}

type Car struct{}

func (c Car) Type() string { return "car" }
func (c Car) Wheels() int  { return 4 }

type Bike struct{}

func (b Bike) Type() string { return "bike" }
func (b Bike) Wheels() int  { return 2 }

type Truck struct{}

func (t Truck) Type() string { return "truck" }
func (t Truck) Wheels() int  { return 6 }

func CreateVehicle(vt VehicleType) Vehicle {
	switch vt {
	case VehicleTypeCar:
		return Car{}
	case VehicleTypeBike:
		return Bike{}
	case VehicleTypeTruck:
		return Truck{}
	default:
		return Car{} // default
	}
}

// 17. File operation enum with behavior
type FileOperation int

const (
	FileOperationRead FileOperation = iota
	FileOperationWrite
	FileOperationDelete
)

func (fo FileOperation) Execute(filename string) {
	switch fo {
	case FileOperationRead:
		fmt.Printf("Reading file: %s\n", filename)
	case FileOperationWrite:
		fmt.Printf("Writing to file: %s\n", filename)
	case FileOperationDelete:
		fmt.Printf("Deleting file: %s\n", filename)
	default:
		fmt.Printf("Unknown operation on file: %s\n", filename)
	}
}

func (fo FileOperation) String() string {
	switch fo {
	case FileOperationRead:
		return "read"
	case FileOperationWrite:
		return "write"
	case FileOperationDelete:
		return "delete"
	default:
		return "unknown"
	}
}
