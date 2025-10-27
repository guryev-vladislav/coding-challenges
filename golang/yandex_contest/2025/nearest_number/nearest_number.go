package nearestnumber

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func NearestNumber(elements []int, x int) int {

	closest := elements[0]
	minDiff := math.Abs(float64(elements[0] - x))
	for _, element := range elements {
		diff := math.Abs(float64(element - x))
		if diff < minDiff {
			minDiff = diff
			closest = element
		}
	}

	return closest

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	elementsStr := strings.Split(scanner.Text(), " ")
	elements := make([]int, n)
	for i, s := range elementsStr {
		elements[i], _ = strconv.Atoi(s)
	}

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	closest := NearestNumber(elements, x)

	fmt.Println(closest)
}
