package main

import (
	"bufio"
	"fmt"
	"os"
)

func chooseOption(distanceToOffice, distanceToScooter, charge int64) int {
	if distanceToScooter >= distanceToOffice {
		return 1
	}

	distanceSquared := distanceToOffice*distanceToOffice + distanceToScooter*distanceToScooter
	if distanceSquared <= charge*charge {
		return 2
	}

	threshold := distanceToOffice - distanceToScooter + charge
	if threshold > 0 && distanceSquared < threshold*threshold {
		return 2
	}

	return 1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	var x, y, charge int64
	if _, err := fmt.Fscan(in, &x, &y, &charge); err != nil {
		return
	}

	if _, err := fmt.Fprintln(out, chooseOption(x, y, charge)); err != nil {
		return
	}
}
