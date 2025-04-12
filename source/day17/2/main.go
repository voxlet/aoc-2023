package main

import (
	"fmt"
	"strconv"

	"github.com/voxlet/aoc-2023/source/day17/parse"
	"github.com/voxlet/aoc-2023/source/day17/priorityqueue"
)

type Vec2 struct {
	Row int
	Col int
}

func add(a Vec2, b Vec2) Vec2 {
	return Vec2{
		Row: a.Row + b.Row,
		Col: a.Col + b.Col,
	}
}

func neg(vec Vec2) Vec2 {
	return Vec2{
		Row: -vec.Row,
		Col: -vec.Col,
	}
}

func (vec *Vec2) inside(extent Vec2) bool {
	return vec.Row >= 0 && vec.Col >= 0 &&
		vec.Row < extent.Row && vec.Col < extent.Col
}

type Inertia struct {
	Dir   Vec2
	Count int
}

type Node struct {
	Pos     Vec2
	Inertia Inertia
}

type Dijkstra struct {
	queue     priorityqueue.PriorityQueue[Node]
	distances map[Node]int
	costs     [][]int
	extent    Vec2
}

func isImprovement(entry priorityqueue.Entry[Node], distances map[Node]int) bool {
	node := *entry.Value

	distance, ok := distances[node]
	if !ok {
		return true
	}
	return entry.Priority < distance
}

func nextNode(node Node, d Vec2, extent Vec2) (Node, bool) {
	if d == neg(node.Inertia.Dir) {
		return Node{}, false
	}

	pos := add(node.Pos, d)

	if !pos.inside(extent) {
		return Node{}, false
	}

	inertia := Inertia{
		Dir:   d,
		Count: 1,
	}

	if d == node.Inertia.Dir {
		inertia.Count = node.Inertia.Count + 1
		if inertia.Count > 10 {
			return Node{}, false
		}
	} else {
		if node.Inertia.Count < 4 {
			return Node{}, false
		}
	}

	next := Node{
		Pos:     pos,
		Inertia: inertia,
	}
	return next, true
}

func priorityFor(node Node, next Node, distances map[Node]int, costs [][]int) int {
	distance := distances[node]
	cost := costs[next.Pos.Row][next.Pos.Col]
	return distance + cost
}

func (dijk *Dijkstra) search(source Vec2, target Vec2) int {
	dijk.queue.InsertValue(&Node{Pos: source, Inertia: Inertia{Dir: Vec2{0, 0}, Count: 4}}, 0)

	for !dijk.queue.IsEmpty() {
		entry := dijk.queue.PopEntry()
		if !isImprovement(entry, dijk.distances) {
			continue
		}

		node := *entry.Value
		dijk.distances[node] = entry.Priority

		for _, d := range []Vec2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next, ok := nextNode(node, d, dijk.extent)
			if !ok {
				continue
			}

			priority := priorityFor(node, next, dijk.distances, dijk.costs)

			if next.Pos == target {
				if next.Inertia.Count < 4 {
					fmt.Println("can't stop at target")
				} else {
					return priority
				}
			}

			dijk.queue.InsertValue(&next, priority)
		}
	}

	panic("not found")
}

func toCosts(rows []string) [][]int {
	costs := make([][]int, len(rows))

	for r, row := range rows {
		for _, c := range row {
			cost, err := strconv.Atoi(string(c))
			if err != nil {
				panic(fmt.Sprintf("%v: %v", err, c))
			}
			costs[r] = append(costs[r], cost)
		}
	}

	return costs
}

func main() {
	input := parse.Parse("source/day17/input.txt")

	dijk := Dijkstra{
		queue: priorityqueue.PriorityQueue[Node]{
			IsBefore: func(priorityA, priorityB int) bool {
				return priorityA < priorityB
			},
		},
		distances: make(map[Node]int),
		costs:     toCosts(input.Rows),
		extent:    Vec2{Row: len(input.Rows), Col: input.ColSize},
	}

	min := dijk.search(Vec2{0, 0}, Vec2{len(input.Rows) - 1, input.ColSize - 1})

	fmt.Println(min)
}
