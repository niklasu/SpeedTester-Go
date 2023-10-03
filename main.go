package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type DevNullWriter struct {
	receivedBytes int64
}

func (dnv *DevNullWriter) Write(p []byte) (n int, err error) {
	dnv.receivedBytes += int64(len(p))
	log.Printf("Got %d bytes\n", dnv.receivedBytes)
	if dnv.receivedBytes > 10*1000*1000 {
		return 0, errors.New("enough is enough")
	}
	return len(p), nil
}

func main() {
	log.Println("##### START ####")
	url := os.Args[2]
	log.Println(url)

	start := time.Now()
	// Get the data
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// Writer the body to file
	writer := &DevNullWriter{}
	_, _ = io.Copy(writer, resp.Body)

	elapsed := time.Since(start)
	log.Printf("Speed test took %s \n", elapsed)
	log.Printf("%v MBit/s\n", 10*8/elapsed.Seconds())

}
