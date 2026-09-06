package maxcashback

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MaxCashBack(min, casback, sum int) int {
	if sum < min {
		return 0
	}
	max := sum / min * casback

	return max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	min, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	cashback, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	sum, _ := strconv.Atoi(scanner.Text())

	result := MaxCashBack(min, cashback, sum)
	fmt.Println(result)
}
