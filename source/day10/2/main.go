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

type Tile int

const (
	Unknown Tile = iota
	Loop
	In
	Out
	Visiting
)

type Tiles [][]Tile

func (tiles Tiles) at(pos Pos) Tile {
	return tiles[pos.row][pos.col]
}

func (tiles Tiles) mark(pos Pos, tile Tile) {
	tiles[pos.row][pos.col] = tile
}

func makeTiles(inputs []string) Tiles {
	tiles := make(Tiles, len(inputs))
	for row := range len(tiles) {
		tiles[row] = make([]Tile, len(inputs[row]))
	}
	return tiles
}

func markLoop(inputs []string, tiles Tiles) {
	pos := findStart(inputs)
	tiles.mark(pos, Loop)

	dir := firstStep(inputs, pos)
	pos = move(pos, dir)

	for at(inputs, pos) != 'S' {
		println(fmt.Sprintf("Loop: %s @%v:%v", string(at(inputs, pos)), pos, dir))
		tiles.mark(pos, Loop)

		dir = nextDir(at(inputs, pos), dir)
		pos = move(pos, dir)
	}
}

func outOfBounds(pos Pos, tiles Tiles) bool {
	return pos.row < 0 || pos.col < 0 || pos.row >= len(tiles) || pos.col >= len(tiles[pos.row])
}

func leftConnected(left, right byte) bool {
	switch right {
	case '|':
	case 'F':
	case 'L':
		return false
	case 'J':
	case '-':
	case '7':
		switch left {
		case 'L':
		case '-':
		case 'F':
			return true
		}
	}
	return false
}

func verfiyInside(pos Pos, tiles Tiles, inputs []string) bool {
	fmt.Printf("verify: %s of %d @ %#v\n", string(at(inputs, pos)), tiles.at(pos), pos)

	loopCount := 0
	prev := byte(0)

	for i := range pos.col {
		col := pos.col - i - 1

		if tiles[pos.row][col] != Loop {
			prev = 0
			continue
		}

		c := inputs[pos.row][col]

		if prev == 0 {
			if c == '|' || c == 'S' {
				println("++", string(c))
				loopCount++
				continue
			}

			if c != 'J' && c != '7' {
				panic("bad state: " + string(c))
			}
			prev = c
			fmt.Print(string(c))
			continue
		}

		if c == '-' {
			fmt.Print(string(c))
			continue
		}

		if !((c == 'L' && prev == 'J') || (c == 'F' && prev == '7')) {
			loopCount++
			fmt.Print(" ++")
		}
		fmt.Println(string(c))

		prev = 0
	}

	return loopCount%2 == 1
}

func dfsInOut(tiles Tiles, inputs []string, pos Pos, inCount *uint) bool {
	if outOfBounds(pos, tiles) {
		return false
	}

	tile := tiles.at(pos)

	if tile == Out {
		return false
	}
	if tile == In || tile == Loop || tile == Visiting {
		return true
	}

	tiles.mark(pos, Visiting)

	inside := true
	for _, next := range []Dir{Up, Down, Left, Right} {
		inside = inside && dfsInOut(tiles, inputs, move(pos, next), inCount)
	}

	if inside && verfiyInside(pos, tiles, inputs) {
		tiles.mark(pos, In)
		(*inCount)++
	} else {
		tiles.mark(pos, Out)
	}

	return inside
}

func markInOut(tiles Tiles, inputs []string) uint {
	inCount := uint(0)

	for row := range len(tiles) {
		for col := range len(tiles[row]) {
			dfsInOut(tiles, inputs, Pos{row, col}, &inCount)
		}
	}

	return inCount
}

func main() {
	inputs := parse.Parse("source/day10/input.txt")
	tiles := makeTiles(inputs)

	markLoop(inputs, tiles)
	inCount := markInOut(tiles, inputs)

	for _, row := range tiles {
		for _, b := range row {
			fmt.Print(b)
		}
		fmt.Println()
	}

	fmt.Println(inCount)
}
