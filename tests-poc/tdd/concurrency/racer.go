// Package concurrency demonstrates concurrent programming patterns in Go,
// specifically racing HTTP requests to determine which server responds fastest.
package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

// tenSecondTimeout is the default timeout duration used by the Racer function
// to prevent indefinite waiting for slow or unresponsive servers.
var tenSecondTimeout = 10 * time.Second

// Racer races two URLs and returns the URL of the server that responds first.
// It uses a default timeout of 10 seconds to avoid hanging indefinitely.
// Returns the winning URL and any error that occurred during the race.
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer races two URLs with a custom timeout duration.
// It starts concurrent HTTP requests to both URLs and returns the URL
// of whichever server responds first. If neither responds within the
// specified timeout, it returns an error.
//
// Parameters:
//   - a, b: URLs to race against each other
//   - timeout: maximum time to wait for a response
//
// Returns:
//   - winner: URL of the server that responded first
//   - error: timeout error if neither server responds in time
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// Use select statement to handle multiple channel operations concurrently
	select {
	case <-ping(a):
		// First URL responded - return it as the winner
		return a, nil
	case <-ping(b):
		// Second URL responded - return it as the winner
		return b, nil
	case <-time.After(timeout):
		// Timeout occurred - neither server responded in time
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// ping makes an HTTP GET request to the specified URL in a goroutine
// and returns a channel that will be closed when the request completes.
// This allows the caller to wait for the HTTP request to finish without
// blocking, enabling concurrent requests and racing behavior.
//
// The channel uses struct{} as it only signals completion - no data is sent.
func ping(url string) chan struct{} {
	// Create an unbuffered channel for signaling completion
	ch := make(chan struct{})

	// Start HTTP request in a separate goroutine
	go func() {
		// Make the HTTP request (we don't care about the response or errors)
		http.Get(url)
		// Signal completion by closing the channel
		close(ch)
	}()

	// Return the channel immediately so caller can wait on it
	return ch
}

/*
Wrapping up - Key Go Concepts Demonstrated:

select:
- Helps you wait on multiple channels.
- Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

This implementation showcases:
- Racing multiple concurrent operations using select
- Using channels for signaling completion
- Implementing timeouts to prevent indefinite blocking
- Goroutines for concurrent HTTP requests
*/
