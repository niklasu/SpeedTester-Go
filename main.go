package main

import (
	"errors"
	"log"
	"os"
	"speedtester/downloader"
	"strconv"
	"strings"
	"time"
)

func main() {
	url, err := getValue(os.Args, "-url")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("-url is %s \n", url)
	interval, err := getValue(os.Args, "-interval")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("-interval is %s \n", interval)

	intervalAsInt, err := strconv.ParseInt(interval, 10, 64)
	if err != nil {
		return
	}

	measure(url)

	for range time.Tick(time.Second * time.Duration(intervalAsInt)) {
		measure(url)
	}

}

func measure(url string) {
	start := time.Now()
	downloader.Download(url)
	log.Printf("%.2f MBit/s\n", 10*8/time.Since(start).Seconds())
}

func getValue(args []string, key string) (string, error) {
	indexOfKey := indexOf(os.Args, key)
	if indexOfKey == len(args)-1 {
		return "", errors.New("found key -" + key + " but missing a value")
	}
	value := os.Args[indexOfKey+1]
	if strings.HasPrefix(value, "-") {
		return "", errors.New("the value for key " + key + " is " + value + " and not allowed because it starts with '-'")
	}
	return value, nil

}

func indexOf(arr []string, searchTerm string) int {
	for i, s := range arr {
		if s == searchTerm {
			return i
		}
	}
	return 0
}
