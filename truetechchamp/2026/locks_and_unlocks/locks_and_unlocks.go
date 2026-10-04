package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

type segment struct {
	start point
	end   point
}

var points = [10]point{
	1: {x: 0, y: 2},
	2: {x: 1, y: 2},
	3: {x: 2, y: 2},
	4: {x: 0, y: 1},
	5: {x: 1, y: 1},
	6: {x: 2, y: 1},
	7: {x: 0, y: 0},
	8: {x: 1, y: 0},
	9: {x: 2, y: 0},
}

func solveKey(code string) int {
	if len(code) == 0 || !validDigit(code[0]) {
		return -1
	}

	red := [10]bool{}
	red[code[0]-'0'] = true
	segments := make([]segment, 0, len(code)-1)

	for index := 1; index < len(code); index++ {
		if !validDigit(code[index]) {
			return -1
		}

		from := code[index-1] - '0'

		to := code[index] - '0'
		if red[to] {
			return -1
		}

		currentSegment := segment{start: points[from], end: points[to]}
		markGreenNodes(red[:], currentSegment)

		segments = append(segments, currentSegment)
	}

	answer := 0

	for first := range segments {
		for second := first + 1; second < len(segments); second++ {
			answer += overlapLengthSquared(segments[first], segments[second])
		}
	}

	return answer
}

func validDigit(digit byte) bool {
	return digit >= '1' && digit <= '9'
}

func markGreenNodes(red []bool, currentSegment segment) {
	for digit := byte(1); digit <= 9; digit++ {
		if !red[digit] && onSegment(points[digit], currentSegment) {
			red[digit] = true
		}
	}
}

func onSegment(candidate point, current segment) bool {
	if cross(current.start, current.end, candidate) != 0 {
		return false
	}

	return candidate.x >= min(current.start.x, current.end.x) &&
		candidate.x <= max(current.start.x, current.end.x) &&
		candidate.y >= min(current.start.y, current.end.y) &&
		candidate.y <= max(current.start.y, current.end.y)
}

func cross(start, end, candidate point) int {
	return (end.x-start.x)*(candidate.y-start.y) - (end.y-start.y)*(candidate.x-start.x)
}

func overlapLengthSquared(first, second segment) int {
	if cross(first.start, first.end, second.start) != 0 ||
		cross(first.start, first.end, second.end) != 0 {
		return 0
	}

	dx := first.end.x - first.start.x

	dy := first.end.y - first.start.y
	if dx == 0 {
		left := max(min(first.start.y, first.end.y), min(second.start.y, second.end.y))

		right := min(max(first.start.y, first.end.y), max(second.start.y, second.end.y))
		if right <= left {
			return 0
		}

		length := right - left

		return length * length * (dx*dx + dy*dy) / (dy * dy)
	}

	left := max(min(first.start.x, first.end.x), min(second.start.x, second.end.x))

	right := min(max(first.start.x, first.end.x), max(second.start.x, second.end.x))
	if right <= left {
		return 0
	}

	length := right - left

	return length * length * (dx*dx + dy*dy) / (dx * dx)
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
		var code string
		if _, err := fmt.Fscan(in, &code); err != nil {
			return
		}

		if _, err := fmt.Fprintln(out, solveKey(code)); err != nil {
			return
		}
	}
}
