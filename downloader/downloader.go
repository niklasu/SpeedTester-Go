package downloader

import (
	"io"
	"net/http"
)

func Download(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	_, _ = io.CopyN(&NoopWriter{}, resp.Body, 10*1000*1000)
}

type NoopWriter struct {
}

func (n2 *NoopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
