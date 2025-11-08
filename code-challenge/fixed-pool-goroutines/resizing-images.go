package fixed_pool_goroutines

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/fs"
	"os"
	"path/filepath"
)

func Center(srcFile, destFile string) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer file.Close()

	src, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	x, y := src.Bounds().Max.X, src.Bounds().Max.Y
	r := image.Rect(0, 0, x/2, y/2)
	dest := image.NewRGBA(r)
	draw.Draw(dest, dest.Bounds(), src, image.Point{x / 4, y / 4}, draw.Over)

	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, dest, nil)
}

type request struct {
	src  string
	dest string
}

func worker(ctx context.Context, in <-chan request, out chan<- error) {
	for {
		select {
		case r, ok := <-in:
			if !ok {
				return
			}
			out <- Center(r.src, r.dest)
		case <-ctx.Done():
			return
		}
	}
}

func producer(ctx context.Context, in chan<- request, srcFiles []string, destDir string) {
	defer close(in)

	for _, src := range srcFiles {
		dest := fmt.Sprintf("%s/%s", destDir, filepath.Base(src))

		select {
		case in <- request{src, dest}:
		case <-ctx.Done():
			return
		}
	}
}

func CenterDir(ctx context.Context, srcDir, destDir string, n int) error {
	if err := os.Mkdir(destDir, 0750); err != nil && !errors.Is(err, fs.ErrExist) {
		return err
	}

	matches, err := filepath.Glob(fmt.Sprintf("%s/*.jpg", srcDir))
	if err != nil {
		return err
	}

	in, out := make(chan request), make(chan error, len(matches))
	for i := 0; i < n; i++ {
		go worker(ctx, in, out)
	}

	go producer(ctx, in, matches, destDir)

	for range matches {
		select {
		case err := <-out:
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
