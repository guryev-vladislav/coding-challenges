package main

import (
	"bufio"
	"fmt"
	"os"
)

func countNumbers(prefix string) int {
	answer := 0
	if prefix[1:] == "22" {
		answer += 10
	}

	if prefix[2] == '2' {
		answer++
	}

	return answer
}

func main() {
	in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	var prefix string
	if _, err := fmt.Fscan(in, &prefix); err != nil {
		return
	}

	if _, err := fmt.Fprintln(out, countNumbers(prefix)); err != nil {
		return
	}
}
