package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func preprocessURL(url string) string {

	url = strings.TrimSpace(url)
	url = strings.ToLower(url)
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		return url
	}
	return "https://" + url
}

func checkURL(url string, maxLen int, client *http.Client) {

	url = preprocessURL(url)
	start := time.Now()
	resp, err := client.Get(url)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	maxLen += 4 // add a natural spacing
	format := fmt.Sprintf("%%-%ds Status: %%d | Time: %%vms\n", maxLen)

	fmt.Printf(format, url, resp.StatusCode, elapsed.Milliseconds())
}

func main() {
	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "Request timeout in seconds")
	flag.IntVar(&timeout, "t", 10, "Request timeout in seconds (shorthand)")

	var help bool
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.BoolVar(&help, "h", false, "Show help message (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "URL Status Checker\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <url1> <url2> ... \n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, " %s google.com github.com\n", os.Args[0])
		fmt.Fprintf(os.Stderr, " %s -timeout 5 google.com\n", os.Args[0])
	}

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("Error: No URLs provided!\n")
		flag.Usage()
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	maxLen := 0
	for _, url := range urls {
		processedURL := preprocessURL(url)
		if len(processedURL) > maxLen {
			maxLen = len(processedURL)
		}
	}

	fmt.Printf("Checking %d URL(s)...\n", len(urls))

	var wg sync.WaitGroup
	start := time.Now()
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			checkURL(u, maxLen, client)
		}(url)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("\nAll checks complete!")
	fmt.Printf("Total time: %vms\n", elapsed.Milliseconds())
}
