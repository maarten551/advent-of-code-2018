package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
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
	var input = getFileContent("/../../input/input.txt")

	fmt.Println(input)
}
