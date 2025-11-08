package fixed_pool_goroutines

import (
	"context"
	"log"
	"runtime"
	"time"
)

func ExecCenterDiv() {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	n := runtime.GOMAXPROCS(0) // number of cores
	err := CenterDir(ctx, "srcDir", "destDir", n)

	duration := time.Since(start)
	log.Printf("info: finished in %v (err=%v)", duration, err)
}
