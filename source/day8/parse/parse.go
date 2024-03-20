package parse

import (
	"bufio"
	"os"
	"regexp"
)

type Node struct {
	Label string
	L     string
	R     string
}

type Input struct {
	Directions string
	Nodes      []Node
}

func Parse(path string) Input {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	lines.Scan()
	directions := lines.Text()
	lines.Scan()

	nodes := make([]Node, 0)
	re := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

	for lines.Scan() {
		matches := re.FindStringSubmatch(lines.Text())
		nodes = append(nodes, Node{matches[1], matches[2], matches[3]})
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	return Input{directions, nodes}
}
