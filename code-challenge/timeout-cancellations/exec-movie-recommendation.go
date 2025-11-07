package timeout_cancellations

import (
	"context"
	"log"
)

func ExecNextMovie() {
	log.Printf("info: checking finish in time")
	ctx, cancel := context.WithTimeout(context.Background(), bmvTime*2)
	defer cancel()

	mOK := NextMovie(ctx, "ridley")
	log.Printf("info: got %+v", mOK)

	log.Printf("info: checking timeout")
	ctx, cancel = context.WithTimeout(context.Background(), bmvTime/2)
	defer cancel()

	mTimeout := NextMovie(ctx, "ridley")
	log.Printf("info: got %+v", mTimeout)
}
