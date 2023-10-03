package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type NoopWriter struct {
}

func (n2 *NoopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func main() {
	url := os.Args[2]
	log.Printf("-url %s\n", url)

	start := time.Now()

	download(url)

	elapsed := time.Since(start)
	log.Printf("%.2f MBit/s\n", 10*8/elapsed.Seconds())

}

func download(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	_, _ = io.CopyN(&NoopWriter{}, resp.Body, 10*1000*1000)
}
