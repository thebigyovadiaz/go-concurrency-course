package main

import (
	grg "github.com/thebigyovadiaz/go-concurrency-course/code-challenge/getting-results-goroutines"
	wg "github.com/thebigyovadiaz/go-concurrency-course/code-challenge/waiting-goroutines"
)

func main() {
	wg.ExecTimingHTTPCalls()
	grg.ExecValidateSigs()
}
