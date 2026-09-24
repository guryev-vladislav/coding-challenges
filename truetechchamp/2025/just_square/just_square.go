package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func isPrime(number int64, primes []int64) bool {
	if number < 2 {
		return false
	}

	for _, prime := range primes {
		if prime*prime > number {
			break
		}

		if number%prime == 0 {
			return false
		}
	}

	return true
}

func generatePrimes() []int64 {
	const limit = 300000

	sieve := make([]bool, limit+1)
	for index := range sieve {
		sieve[index] = true
	}

	sieve[0], sieve[1] = false, false

	for index := 2; index*index <= limit; index++ {
		if sieve[index] {
			for multiple := index * index; multiple <= limit; multiple += index {
				sieve[multiple] = false
			}
		}
	}

	primes := make([]int64, 0)

	for index, prime := range sieve {
		if prime {
			primes = append(primes, int64(index))
		}
	}

	return primes
}

func findCandidates(prefix string, primes []int64, start, end, squareBase, primeBase int64) []string {
	results := make([]string, 0)

	for value := start; value <= end; value++ {
		square := value * value

		number := square - squareBase
		if number < 0 || number > 9999999999 {
			continue
		}

		formatted := fmt.Sprintf("%010d", number)
		if formatted[:3] != prefix {
			continue
		}

		if isPrime(primeBase+number, primes) {
			results = append(results, formatted)
		}
	}

	return results
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	primes := generatePrimes()

	var testCount int
	if _, err := fmt.Fscan(in, &testCount); err != nil {
		return
	}

	for range testCount {
		var prefix string
		if _, err := fmt.Fscan(in, &prefix); err != nil {
			return
		}

		results := findCandidates(prefix, primes, 282843, 299999, 80000000000, 70000000000)
		results = append(results, findCandidates(prefix, primes, 264576, 282842, 70000000000, 80000000000)...)
		sort.Strings(results)

		if _, err := fmt.Fprint(out, len(results)); err != nil {
			return
		}

		for _, number := range results {
			if _, err := fmt.Fprint(out, " ", number); err != nil {
				return
			}
		}

		if _, err := fmt.Fprintln(out); err != nil {
			return
		}
	}
}
