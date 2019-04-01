package main

import "sort"

type Shift struct {
	year            int
	month           int
	day             int
	minutesAsleep   [60]bool
	sleepAtMinutes  []int
	wakeUpAtMinutes []int
}

func CreateShift(year int, month int, day int) *Shift {
	instance := new(Shift)
	instance.year = year
	instance.month = month
	instance.day = day

	instance.sleepAtMinutes = make([]int, 0)
	instance.wakeUpAtMinutes = make([]int, 0)

	return instance
}

func (s *Shift) AddSleepAtMinute(minuteAtSleep int) {
	s.sleepAtMinutes = append(s.sleepAtMinutes, minuteAtSleep)
	sort.Ints(s.sleepAtMinutes)
}

func (s *Shift) WakeUpAtMinute(wakeUpAtMinute int) {
	s.wakeUpAtMinutes = append(s.wakeUpAtMinutes, wakeUpAtMinute)
	sort.Ints(s.wakeUpAtMinutes)
}

func (s *Shift) processSleep() {
	amountOfSleepBreaks := len(s.sleepAtMinutes)
	for i := 0; i < amountOfSleepBreaks; i++ {
		var sleepAtMinute, wakeAtMinute int
		sleepAtMinute, s.sleepAtMinutes = s.sleepAtMinutes[0], s.sleepAtMinutes[1:]
		wakeAtMinute, s.wakeUpAtMinutes = s.wakeUpAtMinutes[0], s.wakeUpAtMinutes[1:]

		for x := sleepAtMinute; x < (wakeAtMinute - sleepAtMinute); x++ {
			s.minutesAsleep[x] = true
		}
	}
}
