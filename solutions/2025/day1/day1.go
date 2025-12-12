package day1

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"mcbulazs/advent-of-code/registry"
)

func init() {
	registry.Register(2025, 1, loadInput, Part1, Part2)
}

func loadInput() []string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	data, err := os.ReadFile(filepath.Join(dir, "input.txt"))
	if err != nil {
		panic(err)
	}

	// Convert to string and split on newlines
	lines := strings.Split(strings.TrimSpace(strings.ReplaceAll(string(data), "\r\n", "\n")), "\n")

	return lines
}
