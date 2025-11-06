package waiting_goroutines

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func MultiURLTime(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		urlL := url

		go func() {
			defer wg.Done()
			URLTime(urlL)
		}()
	}

	wg.Wait()
}

func URLTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error: %q - %s", url, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("error: %q - bad request - %s", url, resp.Status)
		return
	}

	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Printf("error: %q - %s", url, err)
		return
	}

	elapsed := time.Since(start)
	log.Printf("%q - %v", url, elapsed)
}
