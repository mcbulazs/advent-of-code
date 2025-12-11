package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"mcbulazs/advent-of-code/registry" // blank import makes init() run automatically
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: go run . <year> <day> <part>")
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid year %s", os.Args[1])
	}
	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid day %s", os.Args[2])
	}
	part, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Invalid part %s", os.Args[3])
	}

	d, ok := registry.Solutions[year][day]
	if !ok {
		log.Fatalf("No solution for %d day %d", year, day)
	}

	if part == 1 {
		fmt.Println(d.Part1(d.Loader()))
	} else if part == 2 {
		fmt.Println(d.Part2(d.Loader()))
	} else {
		log.Fatalf("Invalid part %d", part)
	}
}
