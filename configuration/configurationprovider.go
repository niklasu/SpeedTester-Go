package configuration

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func GetUrl() (string, error) {
	return getValue(os.Args, "-url")
}

func GetInterval() (int64, error) {
	intervalAsString, err := getValue(os.Args, "-interval")
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(intervalAsString, 10, 64)
}

func GetSizeInBytes() (int64, error) {
	intervalAsString, err := getValue(os.Args, "-size")
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
