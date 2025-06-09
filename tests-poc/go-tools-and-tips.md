# Go Tools and Tips

This document is a compilation of useful Go tools, tips, and best practices derived from the examples in the "Learn Go with Tests" tutorial.

## Table of Contents
- [Testing Tools](#testing-tools)
- [Code Quality Tools](#code-quality-tools)
- [Documentation Tools](#documentation-tools)
- [Go Commands](#go-commands)
- [TDD Tips](#tdd-tips)
- [Useful Go Concepts](#useful-go-concepts)

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
