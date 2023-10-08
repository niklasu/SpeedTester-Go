package downloader

import (
	"io"
	"net/http"
)

func Download(url string, size int64) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	_, _ = io.CopyN(&noopWriter{}, resp.Body, size)
}

type noopWriter struct {
}

func (n2 *noopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
