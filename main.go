package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func ping(client *http.Client, url string) (int, time.Duration, error) {
	start := time.Now()
	resp, err := client.Get(url)
	dur := time.Since(start).Round(time.Millisecond)

	if err != nil {
		return 0, dur, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, dur, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: beacon <url>...")
		os.Exit(1)
	}

	client := &http.Client{Timeout: 5 * time.Second}

	for {
		var wg sync.WaitGroup
		for _, url := range os.Args[1:] {
			if !strings.Contains(url, "://") {
				url = "https://" + url
			}

			wg.Add(1)
			go func() {
				defer wg.Done()

				code, dur, err := ping(client, url)
				switch {
				case code >= 200 && code < 300:
					fmt.Printf("PASS  %-6s  %s\n", dur, url)
				case err == nil:
					fmt.Printf("WARN  %-6s  %s (%d)\n", dur, url, code)
				case err != nil:
					fmt.Printf("FAIL  %-6s  %s\n", dur, url)
				}
			}()
		}
		wg.Wait()

		fmt.Println()
		time.Sleep(30 * time.Second)
	}
}
