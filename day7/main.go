package main

import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		strings.Split(line, " contain")
	}
}
