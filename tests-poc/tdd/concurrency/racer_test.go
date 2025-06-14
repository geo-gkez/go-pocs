package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestRacer tests the Racer function to ensure it returns the URL
// of the server that responds fastest between two given URLs
func TestRacer(t *testing.T) {

	// Create a slow test server that simulates a delayed response
	// This server sleeps for 20ms before responding
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	// Create a fast test server that responds immediately
	// This server returns a response without any delay
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Get the URLs of our test servers
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	// We expect the Racer function to return the fast server's URL
	// since it should respond first
	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	// Assert that the Racer function returned the expected URL
	if got != want {
		t.Errorf("got %q, want %q,", got, want)
	}

	// Clean up: close the test servers to free up resources
	slowServer.Close()
	fastServer.Close()
}
