package main

import (
	"math"
	"strconv"
	"strings"
)

const input = `Time:        35     93     73     66
Distance:   212   2060   1201   1044`

func toInts(s string) []int {
	ints := make([]int, 0)

	tokens := strings.Fields(s)
	for _, t := range tokens {
		n, err := strconv.Atoi(t)

		if err != nil {
			continue
		}

		ints = append(ints, n)
	}

	return ints
}

func parse(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")

	return toInts(lines[0]), toInts(lines[1])
}

// f(t, x) = (t-x)x = tx - x^2
// f(t, x) > d = -x^2 + tx - d > 0

func solveQuad(a, b, c float64) (float64, float64) {
	axis := -b / (2 * a)
	halfBase := math.Sqrt(b*b-4*a*c) / (2 * a)

	min := axis - halfBase
	max := axis + halfBase

	if max < min {
		min, max = max, min
	}

	return min, max
}

func winMinMax(time int, distance int) (int, int) {
	min, max := solveQuad(-1, float64(time), -float64(distance))
	return int(math.Ceil(min)), int(math.Floor(max))
}

func main() {
	times, distances := parse(input)

	margin := 1

	for i := range times {
		min, max := winMinMax(times[i], distances[i])
		margin *= max - min + 1
	}

	println(margin)
}
