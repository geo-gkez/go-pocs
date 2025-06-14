package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestRacer tests the Racer function with multiple scenarios using subtests
func TestRacer(t *testing.T) {

	// Test case 1: Normal operation - comparing server speeds
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		// Create servers with different response delays
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// Ensure cleanup of test servers
		defer slowServer.Close()
		defer fastServer.Close()

		// Get URLs for testing
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		// Expected result should be the fast server's URL
		want := fastURL
		got, err := Racer(slowURL, fastURL)

		// Verify no error occurred during the race
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		// Verify the fastest server's URL was returned
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	// Test case 2: Timeout scenario - server takes too long to respond
	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		// Create a server that will exceed the timeout (25ms delay vs 20ms timeout)
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		// Test with a timeout shorter than the server's response time
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		// Verify that an error was returned due to timeout
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

// makeDelayedServer is a helper function that creates a test HTTP server
// with a configurable delay before responding. This allows us to simulate
// servers with different response times for testing race conditions.
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Introduce artificial delay to simulate slow/fast servers
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

/*
Wrapping up - Key Go Testing Concepts Demonstrated:

httptest:
- A convenient way of creating test servers so you can have reliable and controllable tests.
- Uses the same interfaces as the "real" net/http servers which is consistent and less for you to learn.

select (tested indirectly):
- Helps you wait on multiple channels.
- Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

This test suite showcases:
- Creating mock HTTP servers with configurable delays
- Testing concurrent operations and race conditions
- Testing timeout scenarios with time-based assertions
- Using subtests to organize different test scenarios
- Proper cleanup with defer statements
*/
