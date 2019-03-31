package main

type Guard struct {
	id     int
	shifts []Shift
}

func CreateGuard(id int) *Guard {
	instance := new(Guard)
	instance.id = id
	instance.shifts = make([]Shift, 0)

	return instance
}
