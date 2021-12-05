// https://adventofcode.com/2021/day/5

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	begin, end *Point
}

func NewPoint(s string) *Point {
	tmp := strings.Split(s, ",")
	x, _ := strconv.Atoi(tmp[0])
	y, _ := strconv.Atoi(tmp[1])
	return &Point{x, y}
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func Input() ([]*Line, int, int) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := make([]*Line, 0)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		// fmt.Println(strs[0], strs[len(strs)-1])
		ret = append(ret, &Line{NewPoint(strs[0]), NewPoint(strs[len(strs)-1])})
	}
	X := 0
	Y := 0
	for i := 0; i < len(ret); i++ {
		line := ret[i]
		X = max(max(line.begin.x, line.end.x), X)
		Y = max(max(line.begin.y, line.end.y), Y)
	}
	return ret, X + 1, Y + 1
}

func partOne(lines []*Line, X, Y int) {
	ocean := make([][]byte, Y)
	for i := 0; i < Y; i++ {
		ocean[i] = make([]byte, X)
	}

	const MIN_OVERLAP = byte(2)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		lengthX := line.end.x - line.begin.x
		lengthY := line.end.y - line.begin.y

		if lengthY == 0 { // horizontal
			stepX := 1
			if line.begin.x > line.end.x {
				stepX = -1
			}
			for x := line.begin.x; ; x += stepX {
				ocean[line.begin.y][x]++
				if x == line.end.x {
					break
				}
			}
		}

		if lengthX == 0 { // vertical
			stepY := 1
			if line.begin.y > line.end.y {
				stepY = -1
			}
			for y := line.begin.y; ; y += stepY {
				ocean[y][line.begin.x]++
				if y == line.end.y {
					break
				}
			}
		}
	}

	count := 0
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			if ocean[i][j] >= MIN_OVERLAP {
				count++
			}
		}
	}
	fmt.Println(count)
}

func partTwo(lines []*Line, X, Y int) {
	ocean := make([][]byte, Y)
	for i := 0; i < Y; i++ {
		ocean[i] = make([]byte, X)
	}

	const MIN_OVERLAP = byte(2)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		lengthX := line.end.x - line.begin.x
		lengthY := line.end.y - line.begin.y
		if abs(lengthX) == abs(lengthY) { // diagonal
			stepX := 1
			if line.begin.x > line.end.x {
				stepX = -1
			}
			stepY := 1
			if line.begin.y > line.end.y {
				stepY = -1
			}
			for x, y := line.begin.x, line.begin.y; ; x, y = x+stepX, y+stepY {
				ocean[y][x]++
				if x == line.end.x {
					break
				}
			}
		}

		if lengthY == 0 { // horizontal
			stepX := 1
			if line.begin.x > line.end.x {
				stepX = -1
			}
			for x := line.begin.x; ; x += stepX {
				ocean[line.begin.y][x]++
				if x == line.end.x {
					break
				}
			}
		}

		if lengthX == 0 { // vertical
			stepY := 1
			if line.begin.y > line.end.y {
				stepY = -1
			}
			for y := line.begin.y; ; y += stepY {
				ocean[y][line.begin.x]++
				if y == line.end.y {
					break
				}
			}
		}

	}

	count := 0
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			if ocean[i][j] >= MIN_OVERLAP {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	lines, X, Y := Input()
	partOne(lines, X, Y)
	partTwo(lines, X, Y)
}
