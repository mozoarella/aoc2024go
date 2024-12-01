package conversions

import (
	"log"
	"strconv"
)

func MAtoI(input string) (output int) {
	output, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("This number is not a number :(((")
	}
	return output
}
