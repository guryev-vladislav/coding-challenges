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
	if !scanner.Scan() {
		return
	}

	elementCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	if !scanner.Scan() {
		return
	}

	elementsStr := strings.Split(scanner.Text(), " ")

	elements := make([]int, elementCount)
	for i, s := range elementsStr {
		elements[i], err = strconv.Atoi(s)
		if err != nil {
			return
		}
	}

	if !scanner.Scan() {
		return
	}

	target, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	if err := scanner.Err(); err != nil {
		return
	}

	closest := NearestNumber(elements, target)

	if _, err := fmt.Fprintln(os.Stdout, closest); err != nil {
		return
	}
}
