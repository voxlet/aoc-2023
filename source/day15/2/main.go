package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/voxlet/aoc-2023/source/day15/parse"
)

func hash(s string) int {
	code := 0

	for _, c := range []byte(s) {
		code += int(c)
		code *= 17
		code %= 256
	}

	return code
}

type Box struct {
	lensIndexes map[string]int
	lenses      []int
}

func NewBox() Box {
	return Box{
		lensIndexes: make(map[string]int),
		lenses:      make([]int, 0),
	}
}

type Boxes *[256]Box

func NewBoxes() Boxes {
	boxes := [256]Box{}

	for i := range len(boxes) {
		boxes[i] = NewBox()
	}

	return &boxes
}

func removeLens(step string, boxes Boxes) {
	label := strings.TrimSuffix(step, "-")
	boxIndex := hash(label)
	box := &boxes[boxIndex]

	lensIndex, ok := box.lensIndexes[label]
	if !ok {
		fmt.Println("remove not found", label, boxIndex, lensIndex, *box)
		return
	}

	fmt.Println("remove", label, boxIndex, lensIndex, *box)
	delete(box.lensIndexes, label)
	for l, i := range box.lensIndexes {
		if i > lensIndex {
			box.lensIndexes[l] = i - 1
		}
	}
	box.lenses = slices.Delete(box.lenses, lensIndex, lensIndex+1)
}

func upsertLens(step string, boxes Boxes) {
	tokens := strings.Split(step, "=")
	if len(tokens) != 2 {
		panic(fmt.Sprintf("bad upsert step: %v", step))
	}

	label, focalLengthToken := tokens[0], tokens[1]
	focalLength, err := strconv.Atoi(focalLengthToken)
	if err != nil {
		panic(fmt.Sprintf("bad focal length: %v", step))
	}

	boxIndex := hash(label)
	box := &boxes[boxIndex]

	if lensIndex, ok := box.lensIndexes[label]; ok {
		fmt.Println("update", label, boxIndex, lensIndex, focalLength, *box)
		box.lenses[lensIndex] = focalLength
	} else {
		fmt.Println("insert", label, boxIndex, lensIndex, focalLength, *box)
		box.lensIndexes[label] = len(box.lenses)
		box.lenses = append(box.lenses, focalLength)
	}
}

func applyStep(step string, boxes Boxes) {
	if strings.HasSuffix(step, "-") {
		removeLens(step, boxes)
	} else {
		upsertLens(step, boxes)
	}
}

func main() {
	input := parse.Parse("source/day15/input.txt")

	boxes := NewBoxes()

	for _, step := range input.Steps {
		applyStep(step, boxes)
	}

	sum := 0
	for boxIndex, box := range boxes {
		for lensIndex, focalLength := range box.lenses {
			sum += (boxIndex + 1) * (lensIndex + 1) * focalLength
		}
	}

	println(sum)
}
