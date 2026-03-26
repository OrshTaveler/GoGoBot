package utils

import (
	"os"
	"regexp"
)

func ParseSGFMoves(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := string(data)

	re := regexp.MustCompile(`;[BW]\[([a-z]{2})\]`)
	matches := re.FindAllStringSubmatch(content, -1)

	var moves []string
	for _, m := range matches {
		moves = append(moves, m[1])
	}

	return moves, nil
}
