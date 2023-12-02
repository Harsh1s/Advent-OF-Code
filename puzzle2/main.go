package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Harsh1s/Advent-OF-Code/puzzle2/part1"
	"github.com/Harsh1s/Advent-OF-Code/puzzle2/part2"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	var total1, total2 int
	var s string
	defer writer.Flush()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s = scanner.Text()
		if s == "" {
			break
		}
		total1 = total1 + part1.Solve(s)
		total2 = total2 + part2.Solve(s)
	}
	if err := scanner.Err(); err != nil {
		printf("reading standard input: %v\n", err)
	}

	printf("total1: %v\ntotal2: %v\n", total1, total2)
}
