package day

type Day struct {
	DayName string
	Puzzles []func()
}

var DayMap map[int]*Day = make(map[int]*Day)
