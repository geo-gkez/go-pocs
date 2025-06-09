package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// mockWebsiteChecker is a test double that returns false for a specific invalid URL,
// and true for all others. Used to simulate website checking logic in tests.
func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

// TestCheckWebsites verifies that CheckWebsites correctly maps URLs to their check results
// using the provided WebsiteChecker function.
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

// slowStubWebsiteChecker simulates a slow website check by sleeping for 20ms before returning true.
// Used to benchmark the performance of CheckWebsites.
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// BenchmarkCheckWebsites measures the performance of CheckWebsites when checking 100 URLs
// using a slow checker function.
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
