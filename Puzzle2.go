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
	var total int
	var s string
	defer writer.Flush()

	numbers := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s = scanner.Text()
		if s == "" {
			break
		}

		calVal := 0
	OuterLoop:
		for i := 0; i < len(s); i++ {
			if s[i] <= '9' && s[i] >= '0' {
				calVal += int(s[i]-'0') * 10
				break OuterLoop
			}
			windows := []string{}
			if i+3 <= len(s) {
				windows = append(windows, s[i:i+3])
			}
			if i+4 <= len(s) {
				windows = append(windows, s[i:i+4])
			}
			if i+5 <= len(s) {
				windows = append(windows, s[i:i+5])
			}
			for _, window := range windows {
				if val, ok := numbers[window]; ok {
					calVal += val * 10
					break OuterLoop
				}
			}
		}
	OuterLoopRev:
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] <= '9' && s[i] >= '0' {
				calVal += int(s[i] - '0')
				break OuterLoopRev
			}
			windows := []string{}
			if i-2 >= 0 {
				windows = append(windows, s[i-2:i+1])
			}
			if i-3 >= 0 {
				windows = append(windows, s[i-3:i+1])
			}
			if i-4 >= 0 {
				windows = append(windows, s[i-4:i+1])
			}
			for _, window := range windows {
				if val, ok := numbers[window]; ok {
					calVal += val
					break OuterLoopRev
				}
			}
		}

		total += calVal
	}

	if err := scanner.Err(); err != nil {
		printf("reading standard input: %v\n", err)
	}

	printf("%d\n", total)
}
