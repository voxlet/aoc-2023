package main

import (
	"fmt"

	"github.com/voxlet/aoc-2023/source/day8/parse"
)

type Network map[string]parse.Node

func toNetwork(nodes []parse.Node) Network {
	network := make(Network)

	for _, n := range nodes {
		network[n.Label] = n
	}

	return network
}

func main() {
	input := parse.Parse("source/day8/input.txt")

	dLen := len(input.Directions)
	network := toNetwork(input.Nodes)

	step := 0
	node := network["AAA"]

	for node.Label != "ZZZ" {
		d := input.Directions[step%dLen]
		step++

		var next string
		switch d {
		case 'L':
			next = node.L
		case 'R':
			next = node.R
		}

		fmt.Println(step, node, string(d), next)

		node = network[next]
	}

	fmt.Println(step)
}
