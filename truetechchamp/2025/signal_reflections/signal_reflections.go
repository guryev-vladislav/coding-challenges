package signalreflections

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func reflections(p, q int64) int64 {
	return (180*q)/gcd(p, 180) - 1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for range t {
		var p, q int64
		if _, err := fmt.Fscan(in, &p, &q); err != nil {
			return
		}

		if _, err := fmt.Fprintln(out, reflections(p, q)); err != nil {
			return
		}
	}
}
