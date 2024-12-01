package day01

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/mozoarella/aoc2024go/lib/calculations"
	"github.com/mozoarella/aoc2024go/lib/conversions"
	"github.com/mozoarella/aoc2024go/lib/day"
)

func Init() {
	day.DayMap[1] = &day.Day{
		DayName: "Historian Hysteria",
	}
	day.DayMap[1].Puzzles = append(day.DayMap[1].Puzzles, RunDayOnePuzzleOne)
	day.DayMap[1].Puzzles = append(day.DayMap[1].Puzzles, RunDayOnePuzzleTwo)
}

func generateLists(inputFile string) ([]int, []int) {
	var listOne []int
	var listTwo []int
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
		listOne = append(listOne, conversions.MAtoI(numbers[0]))
		listTwo = append(listTwo, conversions.MAtoI(numbers[1]))
	}

	slices.SortFunc(listOne, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	slices.SortFunc(listTwo, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	return listOne, listTwo
}

func RunDayOnePuzzleOne() {
	listOne, listTwo := generateLists("./days/01/input.txt")
	var differences []int

	for i := range listOne {
		diff := calculations.DiffBetweenInts(listOne[i], listTwo[i])
		differences = append(differences, diff)
	}
	fmt.Println(differences)

	differenceTotal := calculations.AddNumbersInSlice(differences)
	fmt.Println(differenceTotal)
}

func RunDayOnePuzzleTwo() {
	listOne, listTwo := generateLists("./days/01/input.txt")
	var similarityCounters []int

	for _, v1 := range listOne {
		samesies := 0
		for _, v2 := range listTwo {
			if v1 == v2 {
				samesies++
			}
		}
		similarityCounters = append(similarityCounters, v1*samesies)
	}
	fmt.Println(similarityCounters)
	fmt.Println(calculations.AddNumbersInSlice(similarityCounters))
}
