package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/voxlet/aoc-2023/source/day14/parse"
)

func rotateClockwise(rows []string) []string {
	count := len(rows[0])
	res := make([]string, 0, count)

	for i := range count {
		var b strings.Builder
		for n := len(rows) - 1; n >= 0; n-- {
			b.WriteByte(rows[n][i])
		}
		res = append(res, b.String())
	}

	return res
}

func tiltedSegment(length int, rockCount int) string {
	return strings.Repeat(".", length-rockCount) + strings.Repeat("O", rockCount)
}

func tiltEast(row string) string {
	var b strings.Builder
	lastEnd := 0
	rockCount := 0

	for i, ch := range row {
		if ch == 'O' {
			rockCount++
			continue
		}

		if ch == '#' {
			b.WriteString(tiltedSegment(i-lastEnd, rockCount))
			b.WriteByte('#')

			lastEnd = i + 1
			rockCount = 0
		}
	}
	b.WriteString(tiltedSegment(len(row)-lastEnd, rockCount))

	return b.String()
}

type Cache struct {
	Keys   map[string]int
	Values [][]string
}

func makeCache() Cache {
	var cache Cache
	cache.Keys = make(map[string]int, 0)
	cache.Values = make([][]string, 0)
	return cache
}

func spin(rows []string, cache *Cache) ([]string, int) {
	key := strings.Join(rows, "")

	if order, ok := cache.Keys[key]; ok {
		return nil, order
	}

	for range 4 {
		rows = rotateClockwise(rows)

		for i := range len(rows) {
			rows[i] = tiltEast(rows[i])
		}
	}

	cache.Keys[key] = len(cache.Values)
	cache.Values = append(cache.Values, rows)

	return rows, -1
}

func main() {
	inputs := parse.Parse("source/day14/input.txt")
	cache := makeCache()
	cycleStart := -1

	for {
		inputs, cycleStart = spin(inputs, &cache)

		if cycleStart != -1 {
			break
		}
	}

	cyclePeriod := len(cache.Values) - cycleStart
	pos := (1000000000 - 1 - len(cache.Values)) % cyclePeriod
	final := cache.Values[cycleStart+pos]

	sum := 0

	for i, row := range final {
		sum += strings.Count(row, "O") * (len(row) - i)
	}

	j, _ := json.MarshalIndent(final, "", "  ")
	fmt.Println(string(j))
	fmt.Println(cycleStart, cyclePeriod, pos)
	fmt.Println(sum)
}
