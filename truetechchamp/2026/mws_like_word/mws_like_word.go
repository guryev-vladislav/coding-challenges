package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildWord(n int) string {
	pairs := n / 2

	answer := strings.Repeat("M", pairs)
	if n%2 != 0 {
		answer += "S"
	}

	answer += strings.Repeat("W", pairs)

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

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	if _, err := fmt.Fprintln(out, buildWord(n)); err != nil {
		return
	}
}
