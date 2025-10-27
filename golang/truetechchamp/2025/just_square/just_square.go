package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func isPrime(n int64, primes []int64) bool {
	if n < 2 {
		return false
	}
	for _, p := range primes {
		if p*p > n {
			break
		}
		if n%p == 0 {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Генерация простых чисел до 300000
	sieve := make([]bool, 300001)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false
	for i := 2; i*i <= 300000; i++ {
		if sieve[i] {
			for j := i * i; j <= 300000; j += i {
				sieve[j] = false
			}
		}
	}
	primes := make([]int64, 0)
	for i, isP := range sieve {
		if isP {
			primes = append(primes, int64(i))
		}
	}

	var t int
	fmt.Fscan(in, &t)

	for test := 0; test < t; test++ {
		var prefStr string
		fmt.Fscan(in, &prefStr)

		results := make([]string, 0)

		// Случай 1: A простое, B квадрат
		startK := int64(282843)
		endK := int64(299999)
		for k := startK; k <= endK; k++ {
			B := k * k
			N := B - 80000000000
			if N < 0 || N > 9999999999 {
				continue
			}
			// Проверяем префикс
			if fmt.Sprintf("%010d", N)[:3] != prefStr {
				continue
			}
			A := 70000000000 + N
			if isPrime(A, primes) {
				results = append(results, fmt.Sprintf("%010d", N))
			}
		}

		// Случай 2: B простое, A квадрат
		startM := int64(264576)
		endM := int64(282842)
		for m := startM; m <= endM; m++ {
			A := m * m
			N := A - 70000000000
			if N < 0 || N > 9999999999 {
				continue
			}
			// Проверяем префикс
			if fmt.Sprintf("%010d", N)[:3] != prefStr {
				continue
			}
			B := 80000000000 + N
			if isPrime(B, primes) {
				results = append(results, fmt.Sprintf("%010d", N))
			}
		}

		// Сортируем результаты
		sort.Strings(results)

		// Выводим
		fmt.Fprint(out, len(results))
		for _, num := range results {
			fmt.Fprint(out, " ", num)
		}
		fmt.Fprintln(out)
	}
}
