package registry

type LoaderFunc func() []string

type PartFunc func(input []string) string

type Day struct {
	Loader LoaderFunc
	Part1  PartFunc
	Part2  PartFunc
}

var Solutions = map[int]map[int]Day{}

func Register(year int, day int, loader LoaderFunc, part1 PartFunc, part2 PartFunc) {
	if _, ok := Solutions[year]; !ok {
		Solutions[year] = map[int]Day{}
	}
	Solutions[year][day] = Day{Loader: loader, Part1: part1, Part2: part2}
}
