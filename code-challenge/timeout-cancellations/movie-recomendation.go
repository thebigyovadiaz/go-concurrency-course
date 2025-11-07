package timeout_cancellations

import (
	"context"
	"time"
)

var (
	defaultMovie = Movie{
		ID:    "tt0093779",
		Title: "The Princess Bride",
	}
	bmvTime = 50 * time.Millisecond
)

type Movie struct {
	ID    string
	Title string
}

func BestNextMovie(user string) Movie {
	time.Sleep(bmvTime)

	return Movie{
		ID:    "tt0083658",
		Title: "Blade Runner",
	}
}

func NextMovie(ctx context.Context, user string) Movie {
	ch := make(chan Movie)

	go func() {
		m := BestNextMovie(user)
		ch <- m
	}()

	select {
	case r := <-ch:
		return r
	case <-ctx.Done():
		return defaultMovie
	}
}
