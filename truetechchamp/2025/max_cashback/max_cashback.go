package maxcashback

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MaxCashBack(minimum, cashbackRate, total int) int {
	if total < minimum {
		return 0
	}

	return total / minimum * cashbackRate
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	minimum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	if !scanner.Scan() {
		return
	}

	cashbackRate, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	if !scanner.Scan() {
		return
	}

	total, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	if err := scanner.Err(); err != nil {
		return
	}

	result := MaxCashBack(minimum, cashbackRate, total)
	if _, err := fmt.Fprintln(os.Stdout, result); err != nil {
		return
	}
}
