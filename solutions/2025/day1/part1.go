package day1

import (
	"log"
	"math"
	"strconv"
)

func Part1(input []string) string {
	currentRotation := 50
	counter := 0
	for _, v := range input {
		number := v[1:]
		val, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		var modifier int
		if v[0] == 'R' {
			modifier = 1
		} else {
			modifier = -1
		}
		rotation := currentRotation + (val * modifier)
		currentRotation = int(math.Mod(float64(rotation), 100))
		if currentRotation == 0 {
			counter++
		}
	}
	return strconv.Itoa(counter)
}
