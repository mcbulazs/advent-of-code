package day1

import (
	"math"
	"strconv"
)

func Part2(input []string) string {
	currentRotation := 50
	counter := 0
	for _, v := range input {
		val, _ := strconv.Atoi(v[1:])

		var modifier int = 1
		if v[0] == 'L' {
			modifier = -1
		}

		var startfromZero bool = currentRotation == 0
		currentRotation = currentRotation + (val * modifier)

		if !startfromZero && currentRotation <= 0 {
			counter++
		}
		counter += int(math.Abs(float64(currentRotation)) / 100)

		currentRotation = (currentRotation%100 + 100) % 100
	}
	return strconv.Itoa(counter)
}
