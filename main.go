package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"speedtester/downloader"
	"time"
)

func main() {

	url, err := getValue(os.Args, "url")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Printf("-url is %s \n", url)

	start := time.Now()

	downloader.Download(url)

	log.Printf("%.2f MBit/s\n", 10*8/time.Since(start).Seconds())

}

func getValue(args []string, s string) (string, error) {
	indexOfKey := indexOf(os.Args, "-url")
	if indexOfKey == len(args)-1 {
		return "", errors.New("found key -" + s + " but missing a value")
	}
	return os.Args[indexOfKey+1], nil

}

func indexOf(arr []string, searchTerm string) int {
	for i, s := range arr {
		if s == searchTerm {
			return i
		}
	}
	return 0
}
