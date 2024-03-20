package main

import (
	"fmt"
	"strings"

	"github.com/voxlet/aoc-2023/source/day10/parse"
)

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

type Pos struct {
	row int
	col int
}

func move(pos Pos, dir Dir) Pos {
	row := pos.row
	col := pos.col

	switch dir {
	case Up:
		row--
	case Down:
		row++
	case Left:
		col--
	case Right:
		col++
	}

	return Pos{row, col}
}

func at(inputs []string, pos Pos) byte {
	return inputs[pos.row][pos.col]
}

func findStart(inputs []string) Pos {
	var pos Pos

	for pos.row = range len(inputs) {
		pos.col = strings.IndexByte(inputs[pos.row], 'S')
		if pos.col >= 0 {
			break
		}
	}

	return pos
}

func firstStep(inputs []string, pos Pos) Dir {
	if next := move(pos, Up); next.row >= 0 && strings.IndexByte("7|F", at(inputs, next)) >= 0 {
		return Up
	}
	if next := move(pos, Down); next.row < len(inputs) && strings.IndexByte("J|L", at(inputs, next)) >= 0 {
		return Down
	}
	if next := move(pos, Left); next.col >= 0 && strings.IndexByte("F-L", at(inputs, next)) >= 0 {
		return Left
	}
	if next := move(pos, Right); next.col < len(inputs[pos.row]) && strings.IndexByte("J-7", at(inputs, next)) >= 0 {
		return Right
	}
	panic(fmt.Sprintf("no connection: %#v", pos))
}

func nextDir(ch byte, dir Dir) Dir {
	switch ch {
	case '|':
		switch dir {
		case Up:
			return Up
		case Down:
			return Down
		}
	case '-':
		switch dir {
		case Left:
			return Left
		case Right:
			return Right
		}
	case '7':
		switch dir {
		case Right:
			return Down
		case Up:
			return Left
		}
	case 'F':
		switch dir {
		case Up:
			return Right
		case Left:
			return Down
		}
	case 'J':
		switch dir {
		case Right:
			return Up
		case Down:
			return Left
		}
	case 'L':
		switch dir {
		case Down:
			return Right
		case Left:
			return Up
		}
	}

	panic(fmt.Sprintf("bad dir: %v %v", ch, dir))
}

func main() {
	inputs := parse.Parse("source/day10/input.txt")

	pos := findStart(inputs)
	dir := firstStep(inputs, pos)

	pos = move(pos, dir)
	steps := 1

	for at(inputs, pos) != 'S' {
		println(fmt.Sprintf("%d: %s @%v:%v", steps, string(at(inputs, pos)), pos, dir))

		dir = nextDir(at(inputs, pos), dir)
		pos = move(pos, dir)
		steps++
	}

	fmt.Println(steps / 2)
}
