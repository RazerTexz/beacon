package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

func ping(url string) {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}

	start := time.Now()
	resp, err := client.Get(url)
	dur := time.Since(start).Round(time.Millisecond)

	if err != nil {
		fmt.Printf("FAIL  %-6s  %s\n", dur, url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("PASS  %-6s  %s\n", dur, url)
	} else {
		fmt.Printf("WARN  %-6s  %s (%d)\n", dur, url, resp.StatusCode)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: beacon <url>...")
		os.Exit(1)
	}

	urls := os.Args[1:]

	for {
		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ping(url)
			}()
		}
		wg.Wait()

		fmt.Println()
		time.Sleep(30 * time.Second)
	}
}
