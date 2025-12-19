package day2

import (
	"strconv"
	"strings"
)

type myrange struct {
	min int64
	max int64
}

func Part1(input []string) string {
	//rossz!
	var ranges []myrange
	for _, v := range strings.Split(input[0], ",") {
		spl := strings.Split(v, "-")
		min, _ := strconv.ParseInt(spl[0], 10, 64)
		max, _ := strconv.ParseInt(spl[1], 10, 64)
		ranges = append(ranges, myrange{min: min, max: max})
	}
	counter := 0
	for _, v := range ranges {
		length := len(strconv.FormatUint(uint64(v.min), 10))
		start := v.min >> (int64(length) / 2)

		for i := start; i < v.max>>(int64(length)/2); i++ {
			if i<<(int64(length)/2)+i >= v.min && i<<(int64(length)/2)+i <= v.max {
				counter++
			}
		}
	}
	return strconv.Itoa(counter)
}
