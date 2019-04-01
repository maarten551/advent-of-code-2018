package main

type Guard struct {
	Id     int
	Shifts map[string]*Shift
}

func CreateGuard(id int) *Guard {
	instance := new(Guard)
	instance.Id = id
	instance.Shifts = make(map[string]*Shift, 0)

	return instance
}

func (g Guard) calculateSleepStatistics() (maxAmountOfMinutes int, sleepsMostAtMinute int) {
	sleepInMinutesSum := 0.
	sleepCountAtMinute := [60]int{}

	for _, shift := range g.Shifts {
		shift.processSleep()

		for minute, isAsleep := range shift.minutesAsleep {
			if isAsleep {
				sleepCountAtMinute[minute]++
				sleepInMinutesSum++

				if sleepCountAtMinute[minute] > sleepCountAtMinute[sleepsMostAtMinute] {
					sleepsMostAtMinute = minute
				}
			}
		}
	}

	maxAmountOfMinutes = sleepCountAtMinute[sleepsMostAtMinute]

	return
}
