package main

import (
	"bufio"
	"fmt"
	"os"
)

func countNumbers(prefix string) int {
	valid := map[rune]bool{
		'0': true,
		'2': true,
		'5': true,
		'6': true,
		'8': true,
		'9': true,
	}

	for _, char := range prefix {
		if !valid[char] {
			return 0
		}
	}

	return 24
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
