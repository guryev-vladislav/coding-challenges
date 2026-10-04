package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const negativeInfinity = -1 << 30

type matrix [9]int

type fastScanner struct {
	data []byte
	pos  int
}

func identityMatrix() matrix {
	result := matrix{}
	for state := range 3 {
		result[state*3+state] = 0
	}

	for index := range result {
		if result[index] == 0 && index%4 == 0 {
			continue
		}

		if result[index] == 0 {
			result[index] = negativeInfinity
		}
	}

	return result
}

func leafMatrix(character byte) matrix {
	result := identityMatrix()
	from, to := -1, -1
	score := 0

	switch character {
	case 'M':
		from, to = 0, 1
	case 'W':
		from, to = 1, 2
	case 'S':
		from, to = 2, 0
		score = 1
	}

	if from >= 0 {
		result[from*3+to] = score
	}

	return result
}

func combine(left, right matrix) matrix {
	result := matrix{}
	for index := range result {
		result[index] = negativeInfinity
	}

	for from := range 3 {
		for middle := range 3 {
			leftValue := left[from*3+middle]
			if leftValue == negativeInfinity {
				continue
			}

			for to := range 3 {
				rightValue := right[middle*3+to]
				if rightValue == negativeInfinity {
					continue
				}

				index := from*3 + to
				if candidate := leftValue + rightValue; candidate > result[index] {
					result[index] = candidate
				}
			}
		}
	}

	return result
}

type segmentTree struct {
	size int
	tree []matrix
}

func newSegmentTree(value string) *segmentTree {
	size := 1
	for size < len(value) {
		size *= 2
	}

	identity := identityMatrix()

	tree := make([]matrix, size*2)
	for index := range tree {
		tree[index] = identity
	}

	for index := range value {
		tree[size+index] = leafMatrix(value[index])
	}

	for index := size - 1; index > 0; index-- {
		tree[index] = combine(tree[index*2], tree[index*2+1])
	}

	return &segmentTree{size: size, tree: tree}
}

func (tree *segmentTree) update(position int, character byte) {
	index := tree.size + position

	tree.tree[index] = leafMatrix(character)
	for index /= 2; index > 0; index /= 2 {
		tree.tree[index] = combine(tree.tree[index*2], tree.tree[index*2+1])
	}
}

func (tree *segmentTree) query(left, right int) int {
	left += tree.size
	right += tree.size
	leftResult := identityMatrix()
	rightResult := identityMatrix()

	for left < right {
		if left%2 == 1 {
			leftResult = combine(leftResult, tree.tree[left])
			left++
		}

		if right%2 == 1 {
			right--
			rightResult = combine(tree.tree[right], rightResult)
		}

		left /= 2
		right /= 2
	}

	return combine(leftResult, rightResult)[0]
}

func (scanner *fastScanner) nextInt() int {
	for scanner.pos < len(scanner.data) && (scanner.data[scanner.pos] < '0' || scanner.data[scanner.pos] > '9') {
		scanner.pos++
	}

	result := 0
	for scanner.pos < len(scanner.data) && scanner.data[scanner.pos] >= '0' && scanner.data[scanner.pos] <= '9' {
		result = result*10 + int(scanner.data[scanner.pos]-'0')
		scanner.pos++
	}

	return result
}

func (scanner *fastScanner) nextString() string {
	for scanner.pos < len(scanner.data) && scanner.data[scanner.pos] <= ' ' {
		scanner.pos++
	}

	start := scanner.pos
	for scanner.pos < len(scanner.data) && scanner.data[scanner.pos] > ' ' {
		scanner.pos++
	}

	return string(scanner.data[start:scanner.pos])
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	scanner := fastScanner{data: data}
	scanner.nextInt()
	queryCount := scanner.nextInt()
	value := scanner.nextString()
	tree := newSegmentTree(value)

	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			return
		}
	}()

	for range queryCount {
		typeID := scanner.nextInt()
		if typeID == 1 {
			position := scanner.nextInt() - 1
			character := scanner.nextString()[0]
			tree.update(position, character)
		} else {
			left := scanner.nextInt() - 1

			right := scanner.nextInt()
			if _, err := fmt.Fprintln(out, tree.query(left, right)); err != nil {
				return
			}
		}
	}
}
