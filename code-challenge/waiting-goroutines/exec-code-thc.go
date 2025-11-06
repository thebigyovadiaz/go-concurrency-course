package waiting_goroutines

import (
	"log"
	"time"
)

func ExecTimingHTTPCalls() {
	start := time.Now()

	urls := []string{
		"http://localhost:8080/200",
		"http://localhost:8080/100",
		"http://localhost:8080/50",
	}

	MultiURLTime(urls)

	duration := time.Since(start)
	log.Printf("%d URLs in %v", len(urls), duration)
}
