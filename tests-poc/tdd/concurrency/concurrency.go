// This example demonstrates concurrency patterns in Go using goroutines and channels.
// It's based on the "Learn Go with Tests" tutorial.
// For more information on Go concurrency patterns and best practices, see:
// /home/geo-gkez/Documents/go-projects/go-pocs/tests-poc/go-tools-and-tips.md#concurrency-in-go
package concurrency

// WebsiteChecker defines a function type that takes a URL string and returns a boolean indicating if the website is up.
type WebsiteChecker func(string) bool

// result holds the URL and the result of checking that URL.
type result struct {
	url   string
	value bool
}

// CheckWebsites runs the WebsiteChecker function concurrently for each URL in the input slice.
// It returns a map of URLs to their check results (true/false).
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Launch a goroutine for each URL to check it concurrently.
	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	// Collect the results from all goroutines.
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url] = r.value
	}

	return results
}
