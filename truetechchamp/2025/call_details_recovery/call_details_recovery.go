package calldetailsrecovery

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func recoverDuration(records []int) int {
	recordCount := len(records)
	positives := make([]int, 0)
	negAbs := make([]int, 0)

	for _, record := range records {
		if record > 0 {
			positives = append(positives, record)
		} else {
			negAbs = append(negAbs, -record)
		}
	}

	if recordCount%2 != 0 || len(positives) != len(negAbs) {
		return -1
	}

	sort.Ints(positives)
	sort.Ints(negAbs)

	if !hasValidOrder(positives, negAbs) {
		return -1
	}

	return duration(positives, negAbs)
}

func hasValidOrder(positives, negAbs []int) bool {
	all := mergeValues(positives, negAbs)
	for index := 0; index < len(all); index += 2 {
		if all[index] != positives[index/2] {
			return false
		}
	}

	for index := 1; index < len(all); index += 2 {
		if all[index] != negAbs[index/2] {
			return false
		}
	}

	return true
}

func mergeValues(positives, negAbs []int) []int {
	all := make([]int, 0, len(positives)+len(negAbs))
	positiveIndex, negativeIndex := 0, 0

	for positiveIndex < len(positives) && negativeIndex < len(negAbs) {
		if positives[positiveIndex] < negAbs[negativeIndex] {
			all = append(all, positives[positiveIndex])
			positiveIndex++
		} else {
			all = append(all, negAbs[negativeIndex])
			negativeIndex++
		}
	}

	for positiveIndex < len(positives) {
		all = append(all, positives[positiveIndex])
		positiveIndex++
	}

	for negativeIndex < len(negAbs) {
		all = append(all, negAbs[negativeIndex])
		negativeIndex++
	}

	return all
}

func duration(positives, negAbs []int) int {
	sumPos, sumNeg := 0, 0
	for _, positive := range positives {
		sumPos += positive
	}

	for _, negativeAbs := range negAbs {
		sumNeg += negativeAbs
	}

	return sumNeg - sumPos
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	var testCount int
	if _, err := fmt.Fscan(in, &testCount); err != nil {
		return
	}

	for range testCount {
		var recordCount int
		if _, err := fmt.Fscan(in, &recordCount); err != nil {
			return
		}

		records := make([]int, recordCount)
		for index := range records {
			if _, err := fmt.Fscan(in, &records[index]); err != nil {
				return
			}
		}

		if _, err := fmt.Fprintln(out, recoverDuration(records)); err != nil {
			return
		}
	}
}
