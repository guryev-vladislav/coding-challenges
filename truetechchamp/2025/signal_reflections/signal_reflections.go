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

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var p, q int64
		fmt.Fscan(in, &p, &q)

		g := gcd(p, 180)
		n := (180 * q) / g
		fmt.Fprintln(out, n-1)
	}
}
