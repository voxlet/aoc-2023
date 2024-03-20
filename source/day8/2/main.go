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

func startNodes(nodes []parse.Node) []parse.Node {
	found := make([]parse.Node, 0)

	for _, n := range nodes {
		if n.Label[2] == 'A' {
			found = append(found, n)
		}
	}

	return found
}

func stepsToFinish(node parse.Node, directions string, network *Network) int {
	dLen := len(directions)
	step := 0

	for node.Label[2] != 'Z' {
		d := directions[step%dLen]
		step++

		var next string
		switch d {
		case 'L':
			next = node.L
		case 'R':
			next = node.R
		}

		node = (*network)[next]
	}

	fmt.Println(step)
	return step
}

func gcd(a int, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	input := parse.Parse("source/day8/input.txt")
	network := toNetwork(input.Nodes)
	nodes := startNodes(input.Nodes)

	steps := 1

	for _, node := range nodes {
		steps = lcm(steps, stepsToFinish(node, input.Directions, &network))
	}

	fmt.Println(steps)
}
