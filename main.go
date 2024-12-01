package main

import (
	"flag"
	"fmt"

	day01 "github.com/mozoarella/aoc2024go/days/01"
	"github.com/mozoarella/aoc2024go/lib/day"
)

var (
	dayFlag    *int = flag.Int("d", -1, "Pick the day for which you would like to run the code")
	puzzleFlag *int = flag.Int("p", -1, "Pick the part for which you would like to run the code")
)

func init() {
	flag.Parse()
	day01.Init()
}

func main() {

	if *dayFlag == -1 && *puzzleFlag == -1 {
		listDays()
	} else if dayVal, dayOk := day.DayMap[*dayFlag]; dayOk {
		if *puzzleFlag >= 1 && *puzzleFlag < len(dayVal.Puzzles)+1 {
			dayVal.Puzzles[*puzzleFlag-1]()
		} else {
			fmt.Println("No or invalid puzzle specified, displaying all results for the day")
			for i, v := range dayVal.Puzzles {
				fmt.Printf("Puzzle %d\n", i+1)
				v()
			}
		}
	} else {
		fmt.Println("Invalid day specified")
		listDays()
	}

}

func listDays() {
	for k, v := range day.DayMap {
		fmt.Printf("%d: %s\n", k, v.DayName)
		for k := range v.Puzzles {
			fmt.Printf("\t%d: Puzzle %d\n", k+1, k+1)
		}
	}
}
