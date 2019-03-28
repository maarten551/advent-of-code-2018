package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
)

func checkErr(error error) {
	if error != nil {
		panic(error)
	}
}

func getFileContent(path string) string {
	_, currentPath, _, _ := runtime.Caller(1)

	dat, err := ioutil.ReadFile(currentPath + path)
	checkErr(err)

	return string(dat[:])
}

func main() {
	input := getFileContent("/../../input/input.txt")
	checksumPartsCount := [2]int{}

	for _, boxId := range strings.Split(input, "\r\n") {
		containsTwoLetters, containsThreeLetters := retrieveChecksumPartFromId(boxId)

		if containsTwoLetters {
			checksumPartsCount[0]++
		}

		if containsThreeLetters {
			checksumPartsCount[1]++
		}
	}

	fmt.Println(checksumPartsCount[0] * checksumPartsCount[1])
}

func retrieveChecksumPartFromId(boxId string) (containsTwoLetters bool, containsThreeLetters bool) {
	letterCount := make(map[string]int)
	letterCount["a"] = 0

	for _, char := range boxId {
		if _, ok := letterCount[string(char)]; ok {
			letterCount[string(char)]++
		} else {
			letterCount[string(char)] = 1
		}
	}

	for _, count := range letterCount {
		switch count {
		case 2:
			containsTwoLetters = true
		case 3:
			containsThreeLetters = true
		}
	}

	return
}
