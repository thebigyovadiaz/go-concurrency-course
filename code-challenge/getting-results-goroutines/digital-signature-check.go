package getting_results_goroutines

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
)

// sha1sig return SHA1 signature in the format "35aabcd5a32e01d18a5ef688111624f3c547e13b"
func sha1Sig(data []byte) (string, error) {
	w := sha1.New()
	r := bytes.NewReader(data)
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := fmt.Sprintf("%x", w.Sum(nil))
	return sig, nil
}

type File struct {
	Name      string
	Content   []byte
	Signature string
}

type Reply struct {
	fileName string
	match    bool
	err      error
}

func workerSigs(file File, chR chan<- Reply) {
	sig, err := sha1Sig(file.Content)
	r := Reply{
		fileName: file.Name,
		match:    file.Signature == sig,
		err:      err,
	}
	chR <- r
}

// ValidateSigs return slice of OK files and slice of mismatched files
func ValidateSigs(files []File) ([]string, []string, error) {
	var okFiles []string
	var badFiles []string
	chR := make(chan Reply)

	for _, file := range files {
		go workerSigs(file, chR)
	}

	for range files {
		r := <-chR
		if !r.match || r.err != nil {
			badFiles = append(badFiles, r.fileName)
		} else {
			okFiles = append(okFiles, r.fileName)
		}
	}

	close(chR)

	return okFiles, badFiles, nil
}
