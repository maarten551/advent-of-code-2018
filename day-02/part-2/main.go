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
	boxIds := strings.Split(input, "\r\n")

	for boxIdIndex, boxIdA := range boxIds {
		for _, boxIdB := range boxIds[boxIdIndex+1:] {
			differentCharacters := 0
			for i := 0; i < len(boxIdA); i++ {
				if boxIdA[i] != boxIdB[i] {
					differentCharacters++

					temp := []rune(boxIdB)
					temp[i] = ' '
					boxIdB = string(temp)
				}
			}

			if differentCharacters == 1 {
				fmt.Println(strings.ReplaceAll(boxIdB, " ", ""))

				return
			}
		}
	}
}
