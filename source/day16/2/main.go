package main

import (
	"github.com/voxlet/aoc-2023/source/day16/parse"
)

type Vec2 struct {
	Row int
	Col int
}

func (vec *Vec2) add(other Vec2) {
	vec.Row += other.Row
	vec.Col += other.Col
}

func (vec *Vec2) inside(extent Vec2) bool {
	return vec.Row >= 0 && vec.Col >= 0 &&
		vec.Row < extent.Row && vec.Col < extent.Col
}

func (vec *Vec2) neg() {
	vec.Row = -vec.Row
	vec.Col = -vec.Col
}

func (vec *Vec2) swizzle() {
	vec.Row, vec.Col = vec.Col, vec.Row
}

type Ray struct {
	Pos Vec2
	Dir Vec2
}

func (ray *Ray) move() {
	ray.Pos.add(ray.Dir)
}

type Energized = map[Vec2]map[Vec2]struct{}

func energizeAndPrune(rays []Ray, energized Energized) []Ray {
	pruned := make([]Ray, 0, len(rays))

	for i := range len(rays) {
		ray := rays[i]
		if dirs, posOk := energized[ray.Pos]; posOk {
			if _, dirOk := dirs[ray.Dir]; dirOk {
				continue
			}
		} else {
			energized[ray.Pos] = make(map[Vec2]struct{})
		}

		energized[ray.Pos][ray.Dir] = struct{}{}
		pruned = append(pruned, ray)
	}

	return pruned
}

func collide(rays []Ray, input parse.Input) []Ray {
	count := len(rays)

	for i := range count {
		ray := &rays[i]
		tile := input.Rows[ray.Pos.Row][ray.Pos.Col]

		if tile == '.' {
			continue
		}

		if tile == '/' || tile == '\\' {
			if tile == '/' {
				ray.Dir.neg()
			}
			ray.Dir.swizzle()
			continue
		}

		if (ray.Dir.Col == 0 && tile == '-') ||
			(ray.Dir.Row == 0 && tile == '|') {
			ray.Dir.swizzle()

			split := *ray
			split.Dir.neg()
			rays = append(rays, split)
		}
	}

	return rays
}

func move(rays []Ray, extent Vec2) []Ray {
	moved := make([]Ray, 0, len(rays))
	for _, ray := range rays {
		ray.move()
		if ray.Pos.inside(extent) {
			moved = append(moved, ray)
		}
	}
	return moved
}

func analyze(input parse.Input, initial Ray) int {
	extent := Vec2{Col: input.ColCount, Row: len(input.Rows)}
	energized := Energized{}
	rays := []Ray{initial}

	for len(rays) > 0 {
		rays = energizeAndPrune(rays, energized)
		rays = collide(rays, input)
		rays = move(rays, extent)
	}

	return len(energized)
}

func main() {
	input := parse.Parse("source/day16/input.txt")
	maxEnergized := 0

	for row := range len(input.Rows) {
		maxEnergized = max(maxEnergized, analyze(input, Ray{Vec2{row, 0}, Vec2{0, 1}}))
		maxEnergized = max(maxEnergized, analyze(input, Ray{Vec2{row, input.ColCount - 1}, Vec2{0, -1}}))
	}

	for col := range input.ColCount {
		maxEnergized = max(maxEnergized, analyze(input, Ray{Vec2{0, col}, Vec2{1, 0}}))
		maxEnergized = max(maxEnergized, analyze(input, Ray{Vec2{len(input.Rows) - 1, 0}, Vec2{-1, 0}}))
	}

	println(maxEnergized)
}
