package parse

import (
	"os"
	"strings"
)

type Input struct {
	Steps []string
}

func Parse(path string) Input {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(path)
	}

	return Input{
		strings.Split(strings.TrimSpace(string(bytes)), ","),
	}
}
