package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
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
	sequence := strings.Split(input, "\r\n")

	frequency := getFrequencyUntilDoubleFound(sequence)

	fmt.Println(frequency)
}

func getFrequencyUntilDoubleFound(sequence []string) int {
	frequency := 0
	var recordedFrequencies []int

	for {
		for _, frequencyChance := range sequence {
			parsedFrequencyChance, _ := strconv.ParseInt(string(frequencyChance[1:]), 10, 32)
			if string(frequencyChance[0]) == "+" {
				frequency += int(parsedFrequencyChance)
			} else {
				frequency -= int(parsedFrequencyChance)
			}

			for _, recordedFrequency := range recordedFrequencies {
				if frequency == recordedFrequency {
					return frequency
				}
			}

			recordedFrequencies = append(recordedFrequencies, frequency)
		}
	}
}
