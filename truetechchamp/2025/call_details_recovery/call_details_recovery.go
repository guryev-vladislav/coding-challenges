package calldetailsrecovery

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		positives := make([]int, 0)
		negAbs := make([]int, 0)

		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(in, &x)
			if x > 0 {
				positives = append(positives, x)
			} else {
				negAbs = append(negAbs, -x)
			}
		}

		if n%2 != 0 || len(positives) != len(negAbs) {
			fmt.Fprintln(out, -1)
			continue
		}

		sort.Ints(positives)
		sort.Ints(negAbs)

		// Создаём объединённый отсортированный список всех абсолютных значений
		all := make([]int, 0, n)
		i, j := 0, 0
		for i < len(positives) && j < len(negAbs) {
			if positives[i] < negAbs[j] {
				all = append(all, positives[i])
				i++
			} else {
				all = append(all, negAbs[j])
				j++
			}
		}
		for i < len(positives) {
			all = append(all, positives[i])
			i++
		}
		for j < len(negAbs) {
			all = append(all, negAbs[j])
			j++
		}

		// Проверяем чередование
		valid := true
		for k := 0; k < n; k += 2 {
			if all[k] != positives[k/2] {
				valid = false
				break
			}
		}
		for k := 1; k < n; k += 2 {
			if all[k] != negAbs[k/2] {
				valid = false
				break
			}
		}

		if !valid {
			fmt.Fprintln(out, -1)
			continue
		}

		// Считаем сумму
		sumPos, sumNeg := 0, 0
		for _, p := range positives {
			sumPos += p
		}
		for _, na := range negAbs {
			sumNeg += na
		}
		fmt.Fprintln(out, sumNeg-sumPos)
	}
}
