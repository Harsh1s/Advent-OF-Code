package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	var total, calVal int
	var s string
	defer writer.Flush()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s = scanner.Text()
		if s == "" {
			break
		}
		for _, c := range s {
			if c >= '0' && c <= '9' {
				calVal += int(c-'0') * 10
				break
			}
		}
		for i := range s {
			c := s[len(s)-i-1]
			if c >= '0' && c <= '9' {
				calVal += int(c - '0')
				break
			}
		}

		total += calVal
		calVal = 0
	}

	if err := scanner.Err(); err != nil {
		printf("reading standard input: %v\n", err)
	}

	printf("%d\n", total)
}
