package day1

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Part2(input []string) string {
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
		currentRotation = currentRotation + (val * modifier)
		fmt.Println(currentRotation)
		if currentRotation >= 100 || currentRotation <= 0 {
			counter++
		}
		currentRotation = int(math.Mod(float64(currentRotation), 100))

		fmt.Println(currentRotation)
	}
	return strconv.Itoa(counter)
}
