package getting_results_goroutines

import (
	"log"
	"time"
)

func ExecValidateSigs() {
	start := time.Now()

	files := []File{
		{Name: "srv-09.log", Content: []byte("srv-09.log"), Signature: "srv09log"},
		{Name: "srv-10.log", Content: []byte("srv-10.log"), Signature: "srv10log"},
		{Name: "srv-01.log", Content: []byte("srv-01.log"), Signature: "srv01log"},
		{Name: "srv-02.log", Content: []byte("srv-02.log"), Signature: "srv02log"},
	}

	ok, bad, err := ValidateSigs(files)
	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(start)
	log.Printf("info: %d files in %v\n", len(ok)+len(bad), duration)
	log.Printf("ok: %v", ok)
	log.Printf("bad: %v", bad)
}
