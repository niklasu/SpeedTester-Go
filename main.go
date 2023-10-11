package main

import (
	"log"
	"os"
	"speedtester/configuration"
	"speedtester/downloader"
	"time"
)

func main() {
	conf, err := configuration.LoadConfig(os.Args)
	if err != nil {
		log.Println(err.Error())
		return
	}
	measure(conf.Url, conf.SizeInBytes)
	if conf.Interval == 0 {
		return
	}
	for range time.Tick(time.Second * time.Duration(conf.Interval)) {
		measure(conf.Url, conf.SizeInBytes)
	}

}

func measure(url string, size int64) {
	start := time.Now()
	downloader.Download(url, size)
	log.Printf("%.2f MBit/s\n", float64(size/(1000000))*8/time.Since(start).Seconds())
}
