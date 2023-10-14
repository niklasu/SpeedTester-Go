package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Url         string
	Interval    int64
	SizeInBytes int64
}

func LoadConfig(args []string) (Config, error) {
	url, found := os.LookupEnv("URL")
	if !found {
		urlFromParams, err := getValue(args, "-url")
		if err != nil {
			log.Println(err.Error())
			return Config{}, err
		}
		url = urlFromParams
	}
	log.Printf("url is %s \n", url)

	intervalAsString, found := os.LookupEnv("INTERVAL")
	if !found {
		intervalAsStringFromParams, err := getValue(args, "-interval")
		if err != nil {
			return Config{}, err
		}
		intervalAsString = intervalAsStringFromParams
	}
	interValAsInt, err := strconv.ParseInt(intervalAsString, 10, 64)
	if err != nil {
		return Config{}, err
	}
	log.Printf("intervalAsStringFromEnv is %d \n", interValAsInt)

	downloadSizeAsString, found := os.LookupEnv("SIZE")
	if !found {
		downloadSizeAsStringFromParams, err := getValue(args, "-size")
		if err != nil {
			return Config{}, err
		}
		downloadSizeAsString = downloadSizeAsStringFromParams
	}
	downloadSizeAsIntInMb, err := strconv.ParseInt(downloadSizeAsString, 10, 64)
	if err != nil {
		return Config{}, err
	}
	log.Printf("size is %d \n", downloadSizeAsIntInMb)

	return Config{
		Url:         url,
		Interval:    interValAsInt,
		SizeInBytes: downloadSizeAsIntInMb * 1000 * 1000,
	}, nil
}

func getValue(args []string, key string) (string, error) {
	indexOfKey := indexOf(os.Args, key)
	if indexOfKey == -1 {
		return "", errors.New("key " + key + " was not found in the arguments")
	}
	if indexOfKey == len(args)-1 {
		return "", errors.New("found key " + key + " but missing a value")
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
	return -1
}
