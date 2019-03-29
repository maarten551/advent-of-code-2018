package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	input := strings.Split(getFileContent("/../../input/input.txt"), "\r\n")
	claims := make([]*FabricCut, len(input))

	for i, unparsedClaim := range input {
		claims[i] = parseInputLine(unparsedClaim)
	}

	claimedFabricParts := [1000][1000]int8{}

	for _, claim := range claims {
		for y := claim.startPosition.y; y < claim.startPosition.y+claim.size.y; y++ {
			for x := claim.startPosition.x; x < claim.startPosition.x+claim.size.x; x++ {
				claimedFabricParts[y][x]++
			}
		}
	}

	for _, claim := range claims {
		for y := claim.startPosition.y; y < claim.startPosition.y+claim.size.y; y++ {
			for x := claim.startPosition.x; x < claim.startPosition.x+claim.size.x; x++ {
				if claimedFabricParts[y][x] >= 2 {
					goto BreakFromOverlap
				}
			}
		}

		fmt.Println(claim.id)

		break
	BreakFromOverlap:
	}
}

func parseInputLine(inputLine string) *FabricCut {
	parseRegex, _ := regexp.Compile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	matchResults := parseRegex.FindStringSubmatch(inputLine)

	fabricCut := newFabricCut()
	fabricCut.id, _ = strconv.ParseInt(matchResults[1], 10, 64)
	fabricCut.startPosition.x, _ = strconv.ParseInt(matchResults[2], 10, 64)
	fabricCut.startPosition.y, _ = strconv.ParseInt(matchResults[3], 10, 64)
	fabricCut.size.x, _ = strconv.ParseInt(matchResults[4], 10, 64)
	fabricCut.size.y, _ = strconv.ParseInt(matchResults[5], 10, 64)

	return fabricCut
}
