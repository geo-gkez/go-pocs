# Go Tools and Tips

This document is a compilation of useful Go tools, tips, and best practices derived from the examples in the "Learn Go with Tests" tutorial.

## Resources

- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) - A great resource for learning Go through test-driven development
- [Learn Go with Tests GitHub Repository](https://github.com/quii/learn-go-with-tests) - Source code for all examples

## Table of Contents
- [Testing Tools](#testing-tools)
- [Code Quality Tools](#code-quality-tools)
- [Documentation Tools](#documentation-tools)
- [Go Commands](#go-commands)
- [TDD Tips](#tdd-tips)
- [Useful Go Concepts](#useful-go-concepts)
- [Concurrency in Go](#concurrency-in-go)

## Testing Tools

### `go test`
The primary command for running tests in Go.

```bash
# Run all tests in the current package
go test

# Run with verbose output
go test -v

# Run tests in a specific package
go test ./path/to/package

# Run a specific test
go test -run TestFunctionName

# Run a specific subtest
go test -run TestFunctionName/SubtestName
```

### Test Coverage
Go has a built-in test coverage tool to identify areas of your code not covered by tests.

```bash
# Run tests with coverage
go test -cover

# Generate a coverage profile
go test -coverprofile=coverage.out

# View coverage in the browser
go tool cover -html=coverage.out
```

### Table-Driven Tests
A pattern for organizing multiple test cases within a single test function.

```go
func TestSomething(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected int
    }{
        {"case 1", 1, 2},
        {"case 2", 2, 4},
        {"case 3", 3, 6},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            got := functionUnderTest(tc.input)
            if got != tc.expected {
                t.Errorf("got %d want %d", got, tc.expected)
            }
        })
    }
}
```

### `t.Helper()`
Mark a function as a test helper function. When a helper function fails the test, the line reported is the line in the test function that called the helper.

```go
func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper() // Mark this as a helper function
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## Code Quality Tools

### `errcheck`
A tool to check that you actually check error returns.

```bash
# Install errcheck
go install github.com/kisielk/errcheck@latest

# Run errcheck on your code
errcheck ./...
```

### `go vet`
Examines Go source code and reports suspicious constructs.

```bash
# Run go vet on current package
go vet

# Run on all packages in module
go vet ./...
```

### `golint`
Lints Go source code and reports style mistakes.

```bash
# Install golint
go install golang.org/x/lint/golint@latest

# Run golint
golint ./...
```

### `staticcheck`
A state-of-the-art linter for Go.

```bash
# Install staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest

# Run staticcheck
staticcheck ./...
```

### `go mod tidy`
Adds missing and removes unused modules from go.mod.

```bash
go mod tidy
```

## Documentation Tools

### `pkgsite`
Serves documentation for your module locally.

```bash
# Install pkgsite
go install golang.org/x/pkgsite/cmd/pkgsite@latest

# Run pkgsite on your module
pkgsite
```

### `godoc`
Documentation generator and viewer.

```bash
# Install godoc
go install golang.org/x/tools/cmd/godoc@latest

# Run godoc server
godoc -http=:6060
```

### `go doc`
Command line documentation tool.

```bash
# View documentation for a package
go doc fmt

# View documentation for a function
go doc fmt.Println
```

## Go Commands

### `go build`
Compiles packages and dependencies.

```bash
# Build the current package
go build

# Build and specify output name
go build -o myapp

# Build for a different OS/architecture
GOOS=linux GOARCH=amd64 go build
```

### `go run`
Compiles and runs a program.

```bash
go run main.go
```

### `go fmt`
Formats Go code according to the Go standard.

```bash
# Format current package
go fmt

# Format all packages in module
go fmt ./...
```

### `go mod`
Module management.

```bash
# Initialize a new module
go mod init modulename

# Add missing and remove unused modules
go mod tidy

# Verify dependencies
go mod verify

# Show module dependencies
go mod graph
```

## TDD Tips

1. **Write a failing test first**: This ensures you understand the requirements.
2. **Write the minimal code to make the test pass**: Don't implement more than needed.
3. **Refactor**: Clean up your code while keeping the tests passing.
4. **Small steps**: Make incremental changes, run tests frequently.
5. **Red, Green, Refactor**: This is the TDD cycle - write a failing test (red), make it pass (green), then refactor.

## Useful Go Concepts

### Interfaces
Interfaces in Go are implemented implicitly. If a type has all the methods declared in an interface, it implements that interface.

```go
type Shape interface {
    Area() float64
}

// Rectangle implements Shape because it has an Area() method
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### Pointers
Use pointers when you need to mutate the value of a variable:

```go
func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}
```

### Error Handling
Go's standard error handling pattern:

```go
func DoSomething() (Result, error) {
    if somethingWentWrong {
        return Result{}, errors.New("something went wrong")
    }
    return result, nil
}

// When using:
result, err := DoSomething()
if err != nil {
    // Handle error
}
```

### Slices
Slices are more flexible than arrays:

```go
// Create a slice
numbers := []int{1, 2, 3, 4, 5}

// Slice of a slice
twoToFour := numbers[1:4] // [2, 3, 4]

// Append to a slice
numbers = append(numbers, 6)

// Create a slice with make
slice := make([]int, 3, 5) // len=3, cap=5
```

### Variadic Functions
Functions that accept a variable number of arguments:

```go
func Sum(numbers ...int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

// Call with multiple arguments
Sum(1, 2, 3)

// Or with a slice using the spread operator
nums := []int{1, 2, 3}
Sum(nums...)
```

### Benchmarks
Go has built-in support for benchmarking:

```go
func BenchmarkSomething(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Code to benchmark
    }
}
```

Run with:
```bash
go test -bench=.
```

### Examples
Testable examples in Go documentation:

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```

These examples serve as both documentation and tests.

### Maps

Maps in Go are powerful built-in data structures that associate keys with values. Here are some important concepts and best practices:

#### Maps as Reference Types

Maps in Go are reference types (similar to slices), which means:

- When you pass a map to a function, you're passing a reference to the map
- Changes to the map inside a function are visible outside the function
- You don't need to use pointers to maps to modify them in functions or methods

```go
// This works fine without a pointer receiver
func (d Dictionary) Add(word, definition string) {
    d[word] = definition  // Modifies the map directly
}
```

#### Map Operations

```go
// Creating a map
m := make(map[string]string)
// Or with initial values
m := map[string]string{
    "key1": "value1",
    "key2": "value2",
}

// Adding or updating a key
m["key"] = "value"

// Retrieving a value
value := m["key"]

// Checking if a key exists (comma ok idiom)
value, ok := m["key"]
if ok {
    // Key exists, use value
}

// Deleting a key
delete(m, "key")

// Getting the number of items
length := len(m)
```

#### Common Map Pitfalls

1. **Accessing a nil map**: Attempting to write to a nil map will cause a panic
   ```go
   var m map[string]string  // nil map
   m["key"] = "value"       // panic: assignment to entry in nil map
   ```
   Always initialize a map before use: `m := make(map[string]string)`

2. **Using pointer receivers with maps unnecessarily**:
   ```go
   // Unnecessarily complex - maps are already references
   func (d *Dictionary) Add(word, definition string) {
       (*d)[word] = definition  // Need to dereference to access the map
   }
   ```

3. **Not checking if a key exists before accessing**:
   ```go
   // Safe way to access a map
   value, exists := myMap[key]
   if exists {
       // Use value
   }
   ```

4. **Concurrent map access**: Maps are not safe for concurrent use. Use sync.Mutex or sync.RWMutex for concurrent access.

#### Custom Map Types

You can create custom types based on maps to add methods, as shown in the dictionary example:

```go
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", errors.New("word not found")
    }
    return definition, nil
}
```

## Concurrency in Go

Go provides powerful concurrency primitives: goroutines and channels. Here are some tips and best practices for writing safe and efficient concurrent code.

### Goroutines

A goroutine is a lightweight thread managed by the Go runtime. You start a goroutine by prefixing a function call with the `go` keyword:

```go
go func() {
    // concurrent code here
}()
```

### Channels

Channels are used to communicate between goroutines and synchronize execution.

```go
ch := make(chan int)

// Send to channel
ch <- 42

// Receive from channel
value := <-ch
```

### Common Concurrency Pitfalls

#### 1. Loop Variable Capture in Goroutines

When launching goroutines inside a loop, always pass the loop variable as an argument to the goroutine function to avoid unexpected behavior:

```go
for _, url := range urls {
    go func(url string) {
        // use url safely here
    }(url)
}
```

If you use the loop variable directly, all goroutines may capture the same (last) value.

#### 2. Race Conditions with Shared Data

Maps are not safe for concurrent writes. If multiple goroutines write to a map, use a `sync.Mutex` or a channel to synchronize access.

```go
var mu sync.Mutex
results := make(map[string]bool)

for _, url := range urls {
    go func(url string) {
        mu.Lock()
        results[url] = check(url)
        mu.Unlock()
    }(url)
}
```

#### 3. Deadlocks

Always ensure that every send on a channel has a corresponding receive, and vice versa. Buffered channels or closing channels can help avoid deadlocks in some cases.

### Example: Concurrent Website Checker

```go
type result struct {
    url   string
    value bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    resultChannel := make(chan result)

    for _, url := range urls {
        go func(url string) {
            resultChannel <- result{url, wc(url)}
        }(url)
    }

    for i := 0; i < len(urls); i++ {
        r := <-resultChannel
        results[r.url] = r.value
    }

    return results
}
```

### Tools for Concurrency

- `go test -race`: Detects race conditions in your code.
- `sync.WaitGroup`: Waits for a collection of goroutines to finish.
- `sync.Mutex`: Provides mutual exclusion for shared data.

### Further Reading

- [Go by Example: Goroutines](https://gobyexample.com/goroutines)
- [Go by Example: Channels](https://gobyexample.com/channels)
- [Go Blog: Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
- [Go Blog: Share Memory by Communicating](https://blog.golang.org/share-memory-by-communicating)
