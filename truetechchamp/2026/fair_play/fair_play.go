package main

import (
	"bufio"
	"fmt"
	"os"
)

const modulus int64 = 998244353

type factor struct {
	prime    int64
	exponent int
}

type divisor struct {
	value int64
}

var primePowerOrders = make(map[int64]int64)
var primeTable = primesUpTo(31623)

func solve(p1, q1, p2, q2 int64) int64 {
	factorizations := factorRange(q1, q2, primeTable)
	answer := int64(0)

	for index, factors := range factorizations {
		denominator := q1 + int64(index)
		answer += sumForDenominator(p1, p2, denominator, factors)
		answer %= modulus
	}

	return answer
}

func sumForDenominator(left, right, denominator int64, factors []factor) int64 {
	divisors := makeDivisors(factors)
	values := make([]int64, len(divisors))

	indices := make(map[int64]int, len(divisors))
	for index, current := range divisors {
		indices[current.value] = index
		values[index] = periodLength(denominator/current.value, factors)
	}

	mobiusValues := append([]int64(nil), values...)

	for _, currentFactor := range factors {
		for index := len(divisors) - 1; index >= 0; index-- {
			if divisors[index].value%currentFactor.prime != 0 {
				continue
			}

			previous := indices[divisors[index].value/currentFactor.prime]
			mobiusValues[index] -= mobiusValues[previous]
		}
	}

	answer := int64(0)

	for index, current := range divisors {
		count := right/current.value - (left-1)/current.value
		term := (mobiusValues[index] % modulus) * (count % modulus) % modulus
		answer = (answer + term) % modulus
	}

	return (answer + modulus) % modulus
}

func periodLength(value int64, factors []factor) int64 {
	period := int64(1)

	for _, currentFactor := range factors {
		exponent := 0

		for value%currentFactor.prime == 0 {
			value /= currentFactor.prime
			exponent++
		}

		if currentFactor.prime == 2 || currentFactor.prime == 5 || exponent == 0 {
			continue
		}

		currentOrder := orderPrimePower(currentFactor.prime, exponent)
		period = lcm(period, currentOrder)
	}

	return period
}

func orderPrimePower(prime int64, exponent int) int64 {
	modulusValue := int64(1)
	for range exponent {
		modulusValue *= prime
	}

	if cached, ok := primePowerOrders[modulusValue]; ok {
		return cached
	}

	phi := (prime - 1) * (modulusValue / prime)

	order := phi
	for _, currentFactor := range factorNumber(phi, primeTable) {
		for order%currentFactor.prime == 0 && powMod(10, order/currentFactor.prime, modulusValue) == 1 {
			order /= currentFactor.prime
		}
	}

	primePowerOrders[modulusValue] = order

	return order
}

func makeDivisors(factors []factor) []divisor {
	divisors := []divisor{{value: 1}}
	for _, currentFactor := range factors {
		base := len(divisors)

		multiplier := int64(1)
		for exponent := 1; exponent <= currentFactor.exponent; exponent++ {
			multiplier *= currentFactor.prime
			for index := range base {
				divisors = append(divisors, divisor{value: divisors[index].value * multiplier})
			}
		}
	}

	return divisors
}

func factorRange(left, right int64, primes []int64) [][]factor {
	remaining := make([]int64, right-left+1)

	result := make([][]factor, right-left+1)
	for index := range remaining {
		remaining[index] = left + int64(index)
	}

	for _, prime := range primes {
		if prime*prime > right {
			break
		}

		first := (left + prime - 1) / prime * prime
		for value := first; value <= right; value += prime {
			index := value - left
			exponent := 0

			for remaining[index]%prime == 0 {
				remaining[index] /= prime
				exponent++
			}

			if exponent > 0 {
				result[index] = append(result[index], factor{prime: prime, exponent: exponent})
			}
		}
	}

	for index, value := range remaining {
		if value > 1 {
			result[index] = append(result[index], factor{prime: value, exponent: 1})
		}
	}

	return result
}

func factorNumber(value int64, primes []int64) []factor {
	result := make([]factor, 0)

	for _, prime := range primes {
		if prime*prime > value {
			break
		}

		if value%prime != 0 {
			continue
		}

		exponent := 0

		for value%prime == 0 {
			value /= prime
			exponent++
		}

		result = append(result, factor{prime: prime, exponent: exponent})
	}

	if value > 1 {
		result = append(result, factor{prime: value, exponent: 1})
	}

	return result
}

func primesUpTo(limit int64) []int64 {
	composite := make([]bool, limit+1)
	primes := make([]int64, 0)

	for value := int64(2); value <= limit; value++ {
		if composite[value] {
			continue
		}

		primes = append(primes, value)
		if value*value <= limit {
			for multiple := value * value; multiple <= limit; multiple += value {
				composite[multiple] = true
			}
		}
	}

	return primes
}

func powMod(base, exponent, mod int64) int64 {
	result := int64(1)
	base %= mod

	for exponent > 0 {
		if exponent%2 == 1 {
			result = result * base % mod
		}

		base = base * base % mod
		exponent /= 2
	}

	return result
}

func gcd(first, second int64) int64 {
	for second != 0 {
		first, second = second, first%second
	}

	return first
}

func lcm(first, second int64) int64 {
	return first / gcd(first, second) * second
}

func main() {
	in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	var p1, q1, p2, q2 int64
	if _, err := fmt.Fscan(in, &p1, &q1, &p2, &q2); err != nil {
		return
	}

	if _, err := fmt.Fprintln(out, solve(p1, q1, p2, q2)); err != nil {
		return
	}
}
