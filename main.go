package main

import (
	"log"
	"speedtester/configuration"
	"speedtester/downloader"
	"time"
)

func main() {
	url, err := configuration.GetUrl()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("-url is %s \n", url)
	interval, err := configuration.GetInterval()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("-interval is %d \n", interval)

	downloadSize, err := configuration.GetSizeInBytes()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("-size is %d MB\n", downloadSize/(1000*1000))

	measure(url, downloadSize)
	if interval == 0 {
		return
	}
	for range time.Tick(time.Second * time.Duration(interval)) {
		measure(url, downloadSize)
	}

}

func measure(url string, size int64) {
	start := time.Now()
	downloader.Download(url, size)
	log.Printf("%.2f MBit/s\n", float64(size/(1000000))*8/time.Since(start).Seconds())
}
