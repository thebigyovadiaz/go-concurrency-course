package main

import (
	fpg "github.com/thebigyovadiaz/go-concurrency-course/code-challenge/fixed-pool-goroutines"
	grg "github.com/thebigyovadiaz/go-concurrency-course/code-challenge/getting-results-goroutines"
	tc "github.com/thebigyovadiaz/go-concurrency-course/code-challenge/timeout-cancellations"
	wg "github.com/thebigyovadiaz/go-concurrency-course/code-challenge/waiting-goroutines"
)

func main() {
	wg.ExecTimingHTTPCalls()
	grg.ExecValidateSigs()
	tc.ExecNextMovie()
	fpg.ExecCenterDiv()
}
