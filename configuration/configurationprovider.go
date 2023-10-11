package configuration

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
	url, err := getUrl(args)
	if err != nil {
		log.Println(err.Error())
		return Config{}, err
	}
	log.Printf("-url is %s \n", url)
	interval, err := getInterval(args)
	if err != nil {
		log.Println(err.Error())
		return Config{}, err
	}
	log.Printf("-interval is %d \n", interval)

	downloadSize, err := getSizeInBytes(args)
	if err != nil {
		log.Println(err.Error())
		return Config{}, err
	}
	log.Printf("-size is %d MB\n", downloadSize/(1000*1000))
	return Config{
		Url:         url,
		Interval:    interval,
		SizeInBytes: downloadSize,
	}, nil
}

func getUrl(args []string) (string, error) {
	return getValue(args, "-url")
}

func getInterval(args []string) (int64, error) {
	intervalAsString, err := getValue(args, "-interval")
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(intervalAsString, 10, 64)
}

func getSizeInBytes(args []string) (int64, error) {
	intervalAsString, err := getValue(args, "-size")
	if err != nil {
		return 0, err
	}
	sizeInMb, err := strconv.ParseInt(intervalAsString, 10, 64)
	if err != nil {
		return sizeInMb, err
	}
	return sizeInMb * 1000 * 1000, err
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
