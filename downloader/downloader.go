package downloader

import (
	"io"
	"net/http"
)

func Download(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	_, _ = io.CopyN(&noopWriter{}, resp.Body, 50*1000*1000)
}

type noopWriter struct {
}

func (n2 *noopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
