package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
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
	guards := parseGuardsWithShiftsFromInput(input)

	for _, guard := range guards {
		fmt.Println(guard.id, len(guard.shifts))
	}

	fmt.Println(guards)
}

func parseGuardsWithShiftsFromInput(inputLines []string) map[int]*Guard {
	parseRegex, _ := regexp.Compile("\\[(\\d+)-(\\d+)-(\\d+) (\\d+):(\\d+)] Guard #(\\d+) begins shift")
	guards := make(map[int]*Guard)

	for _, inputLine := range inputLines {
		if parseRegex.MatchString(inputLine) {
			matchResults := parseRegex.FindStringSubmatch(inputLine)

			guardId, _ := strconv.ParseInt(matchResults[6], 10, 32)
			guard, guardAlreadyExists := guards[int(guardId)]

			if !guardAlreadyExists {
				guard = CreateGuard(int(guardId))
				guards[int(guardId)] = guard
			}

			shiftYear, _ := strconv.ParseInt(matchResults[1], 10, 32)
			shiftMonth, _ := strconv.ParseInt(matchResults[2], 10, 32)
			shiftDay, _ := strconv.ParseInt(matchResults[3], 10, 32)
			shiftHour, _ := strconv.ParseInt(matchResults[4], 10, 32)

			t := fmt.Sprintf("%d-%02d-%02d", shiftYear, shiftMonth, shiftDay)

			shiftTime, _ := time.Parse("2006-01-02", t)

			if shiftHour == 23 {
				shiftTime.Add(time.Hour * 24)
			}

			shift := CreateShift(shiftTime.Year(), int(shiftTime.Month()), shiftTime.Day())
			guard.shifts = append(guard.shifts, *shift)
		}
	}

	return guards
}
