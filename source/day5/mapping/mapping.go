package mapping

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Entry struct {
	Dest int
	Src  int
	Size int
}

func parseSeeds(input *bufio.Reader) []int {
	seeds := make([]int, 0)

	input.Discard(len("seeds: "))

	seedsLineString, isPrefix, err := input.ReadLine()
	if err != nil || isPrefix {
		panic(err)
	}

	seedsLine := strings.NewReader(string(seedsLineString))

	for {
		var seed int
		_, err := fmt.Fscan(seedsLine, &seed)

		if err != nil {
			break
		}

		seeds = append(seeds, seed)
	}

	return seeds
}

func parseMappingEntry(input *bufio.Reader) (Entry, error) {
	var entry Entry
	_, err := fmt.Fscan(input, &entry.Dest, &entry.Src, &entry.Size)
	return entry, err
}

func parseMapping(input *bufio.Reader) ([]Entry, error) {
	mapping := make([]Entry, 0)

	entry, err := parseMappingEntry(input)
	if err == io.EOF {
		panic("EOF at preamble")
	}

	for err != nil {
		input.ReadLine()
		entry, err = parseMappingEntry(input)
	}
	mapping = append(mapping, entry)

	for {
		entry, err = parseMappingEntry(input)

		if err == nil {
			mapping = append(mapping, entry)
			continue
		}

		return mapping, err
	}
}

func parseMappings(input *bufio.Reader) [][]Entry {
	mappings := make([][]Entry, 0)

	for {
		mapping, err := parseMapping(input)
		if len(mapping) > 0 {
			mappings = append(mappings, mapping)
		}
		if err == io.EOF {
			return mappings
		}
	}
}

func Parse(inputPath string) ([]int, [][]Entry) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := bufio.NewReader(file)

	seeds := parseSeeds(input)
	mappings := parseMappings(input)

	return seeds, mappings
}

func Apply(v int, mapping []Entry) int {
	for _, entry := range mapping {
		if v < entry.Src || v >= entry.Src+entry.Size {
			continue
		}
		offset := entry.Dest - entry.Src
		return v + offset
	}
	return v
}
