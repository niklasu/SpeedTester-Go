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
	log.Printf("-interval is %s \n", interval)

	measure(url)
	if interval == 0 {
		return
	}
	for range time.Tick(time.Second * time.Duration(interval)) {
		measure(url)
	}

}

func measure(url string) {
	start := time.Now()
	downloader.Download(url)
	log.Printf("%.2f MBit/s\n", 10*8/time.Since(start).Seconds())
}
