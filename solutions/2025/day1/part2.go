package day1

import (
	"log"
	"math"
	"strconv"
)

func Part2(input []string) string {
	currentRotation := 50
	counter := 0
	for _, v := range input {
		val, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}

		var modifier int = 1
		if v[0] == 'L' {
			modifier = -1
		}

		currentRotation = currentRotation + (val * modifier)
		if currentRotation >= 100 {
			counter += (currentRotation / 100)
		} else if currentRotation <= 0 {
			counter -= (currentRotation / 100) - 1
		}

		currentRotation = int(math.Mod(float64(currentRotation), 100))
	}
	return strconv.Itoa(counter)
}
