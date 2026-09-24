package signalstations

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func SignalStations(signal1, signal2, signal3 int) int {
	firstSide := signal1 + signal2
	secondSide := signal3 + signal2
	thirdSide := signal3 + signal1
	sides := []int{firstSide, secondSide, thirdSide}
	sort.Ints(sides)

	if sides[0]+sides[1] <= sides[2] {
		return -1
	}

	semiperimeter := (firstSide + secondSide + thirdSide) / 2
	area := semiperimeter * (semiperimeter - firstSide) *
		(semiperimeter - secondSide) * (semiperimeter - thirdSide)

	return area
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	values := make([]int, 3)

	for index := range values {
		if !scanner.Scan() {
			return
		}

		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		values[index] = value
	}

	if err := scanner.Err(); err != nil {
		return
	}

	area := SignalStations(values[0], values[1], values[2])

	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := writer.Flush(); err != nil {
			return
		}
	}()

	if area == -1 {
		if _, err := fmt.Fprintln(writer, area); err != nil {
			return
		}
	} else {
		if _, err := fmt.Fprintln(writer, area, area); err != nil {
			return
		}
	}
}
