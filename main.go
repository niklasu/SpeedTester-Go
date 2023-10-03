package main

import (
	"log"
	"os"
	"speedtester/downloader"
	"time"
)

func main() {
	url := os.Args[2]
	log.Printf("-url %s\n", url)

	start := time.Now()

	downloader.Download(url)

	log.Printf("%.2f MBit/s\n", 10*8/time.Since(start).Seconds())

}
