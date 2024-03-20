package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/voxlet/aoc-2023/source/day5/mapping"
)

type Span = struct {
	Start int
	Size  int
}

func toSpans(seeds []int) []Span {
	if len(seeds)%2 != 0 {
		panic(len(seeds))
	}

	count := len(seeds) / 2
	seedSpans := make([]Span, 0, count)

	for i := range count {
		seedSpans = append(seedSpans, Span{Start: seeds[2*i], Size: seeds[2*i+1]})
	}

	return seedSpans
}

func intersect(a Span, b Span) (Span, []Span, bool) {
	left, right := a, b
	if right.Start < left.Start {
		left, right = right, left
	}

	leftEnd := left.Start + left.Size

	if leftEnd <= right.Start {
		return Span{}, nil, false
	}

	end := min(leftEnd, right.Start+right.Size)
	intersection := Span{right.Start, end - right.Start}

	diff := make([]Span, 0)

	if a.Start < intersection.Start {
		diff = append(diff, Span{a.Start, intersection.Start - a.Start})
	}

	aEnd := a.Start + a.Size
	if end < aEnd {
		diff = append(diff, Span{end, aEnd - end})
	}

	return intersection, diff, true
}

func findIntersection(m []mapping.Entry, span Span) (*mapping.Entry, Span, []Span) {
	for _, entry := range m {
		intersection, diff, intersects := intersect(span, Span{entry.Src, entry.Size})

		if intersects {
			return &entry, intersection, diff
		}
	}

	return nil, span, nil
}

func apply(span Span, mappings [][]mapping.Entry) []Span {
	spans := []Span{span}

	for _, m := range mappings {
		next := make([]Span, 0, len(spans))

		for len(spans) > 0 {
			s := spans[len(spans)-1]
			spans = spans[:len(spans)-1]

			entry, intersection, diff := findIntersection(m, s)

			if entry == nil {
				next = append(next, s)
				continue
			}

			intersection.Start += entry.Dest - entry.Src
			next = append(next, intersection)

			spans = slices.Concat(spans, diff)
		}

		spans = next
	}

	return spans
}

func mapToLocations(spans []Span, mappings [][]mapping.Entry) []Span {
	locations := make([]Span, 0, len(spans))

	for _, span := range spans {
		locations = slices.Concat(locations, apply(span, mappings))
	}

	return locations
}

func main() {
	seeds, mappings := mapping.Parse("source/day5/input.txt")
	spans := toSpans(seeds)

	locations := mapToLocations(spans, mappings)

	slices.SortFunc(locations, func(a Span, b Span) int { return cmp.Compare(a.Start, b.Start) })

	s, _ := json.MarshalIndent(locations, "", "\t")
	fmt.Println(string(s))
}
