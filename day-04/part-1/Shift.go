package main

type Shift struct {
	year          int
	month         int
	day           int
	minutesAsleep [60]bool
}

func CreateShift(year int, month int, day int) *Shift {
	instance := new(Shift)
	instance.year = year
	instance.month = month
	instance.day = day

	return instance
}
