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
	guardsById, guardsByDate := parseGuardsWithShiftsFromInput(input)
	parseSleepEvents(guardsByDate, &input)

	guardWithMostAverageSleep, mostAverageSleep := 0, 0.
	for _, guard := range guardsById {
		averageSleepPerShift, _ := guard.calculateSleepStatistics()

		if averageSleepPerShift > mostAverageSleep {
			guardWithMostAverageSleep = guard.Id
			mostAverageSleep = averageSleepPerShift
		}
	}

	_, sleepsMostAtMinute := guardsById[guardWithMostAverageSleep].calculateSleepStatistics()
	fmt.Println(guardWithMostAverageSleep * sleepsMostAtMinute)
}

func parseGuardsWithShiftsFromInput(inputLines []string) (guardsById map[int]*Guard, guardsByDate map[string]*Guard) {
	parseRegex, _ := regexp.Compile("\\[(\\d+)-(\\d+)-(\\d+) (\\d+):(\\d+)] Guard #(\\d+) begins shift")
	guardsById = make(map[int]*Guard)
	guardsByDate = make(map[string]*Guard)

	for _, inputLine := range inputLines {
		if parseRegex.MatchString(inputLine) {
			matchResults := parseRegex.FindStringSubmatch(inputLine)

			guardId, _ := strconv.ParseInt(matchResults[6], 10, 32)
			guard, guardAlreadyExists := guardsById[int(guardId)]

			if !guardAlreadyExists {
				guard = CreateGuard(int(guardId))
				guardsById[int(guardId)] = guard
			}

			shiftYear, _ := strconv.ParseInt(matchResults[1], 10, 32)
			shiftMonth, _ := strconv.ParseInt(matchResults[2], 10, 32)
			shiftDay, _ := strconv.ParseInt(matchResults[3], 10, 32)
			shiftHour, _ := strconv.ParseInt(matchResults[4], 10, 32)

			t := fmt.Sprintf("%d-%02d-%02d", shiftYear, shiftMonth, shiftDay)

			shiftTime, _ := time.Parse("2006-01-02", t)

			if shiftHour == 23 {
				shiftTime = shiftTime.Add(time.Hour * 24)
			}

			shift := CreateShift(shiftTime.Year(), int(shiftTime.Month()), shiftTime.Day())

			guard.Shifts[shiftTime.Format("2006-01-02")] = shift
			guardsByDate[shiftTime.Format("2006-01-02")] = guard
		}
	}

	return
}

func parseSleepEvents(guardsByDate map[string]*Guard, inputLines *[]string) {
	parseRegex, _ := regexp.Compile("\\[(\\d+)-(\\d+)-(\\d+) (\\d+):(\\d+)] ((wakes up)|(falls asleep))")

	for _, inputLine := range *inputLines {
		if parseRegex.MatchString(inputLine) {
			matchResults := parseRegex.FindStringSubmatch(inputLine)

			shiftYear, _ := strconv.ParseInt(matchResults[1], 10, 32)
			shiftMonth, _ := strconv.ParseInt(matchResults[2], 10, 32)
			shiftDay, _ := strconv.ParseInt(matchResults[3], 10, 32)
			shiftMinute, _ := strconv.ParseInt(matchResults[5], 10, 32)

			t := fmt.Sprintf("%d-%02d-%02d", shiftYear, shiftMonth, shiftDay)
			shift := guardsByDate[t].Shifts[t]

			if matchResults[6] == "wakes up" {
				shift.WakeUpAtMinute(int(shiftMinute))
			} else {
				shift.AddSleepAtMinute(int(shiftMinute))
			}
		}
	}
}
