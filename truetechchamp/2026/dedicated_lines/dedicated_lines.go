package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxCapacity = 100

type expansion struct {
	capacityLoss int
	revenue      int64
}

func maxRevenue(upgrade int, expansions []expansion) int64 {
	dp := make([]int64, maxCapacity+1)
	for capacity := range dp {
		dp[capacity] = -1
	}

	dp[maxCapacity] = 0

	for _, option := range expansions {
		next := make([]int64, maxCapacity+1)
		for capacity := range next {
			next[capacity] = -1
		}

		for capacity := 1; capacity <= maxCapacity; capacity++ {
			if dp[capacity] < 0 {
				continue
			}

			applyTransitions(next, capacity, upgrade, option, dp[capacity])
		}

		dp = next
	}

	var answer int64
	for capacity := 1; capacity <= maxCapacity; capacity++ {
		if dp[capacity] > answer {
			answer = dp[capacity]
		}
	}

	return answer
}

func applyTransitions(next []int64, capacity, upgrade int, option expansion, revenue int64) {
	upgradedCapacity := min(maxCapacity, capacity+upgrade)
	next[upgradedCapacity] = max(next[upgradedCapacity], revenue)

	remainingCapacity := capacity - option.capacityLoss
	if remainingCapacity > 0 {
		next[remainingCapacity] = max(next[remainingCapacity], revenue+option.revenue)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	var n, upgrade int
	if _, err := fmt.Fscan(in, &n, &upgrade); err != nil {
		return
	}

	expansions := make([]expansion, n)
	for index := range expansions {
		if _, err := fmt.Fscan(in, &expansions[index].capacityLoss, &expansions[index].revenue); err != nil {
			return
		}
	}

	if _, err := fmt.Fprintln(out, maxRevenue(upgrade, expansions)); err != nil {
		return
	}
}
