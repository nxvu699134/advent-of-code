// https://adventofcode.com/2021/day/11

package main

import (
	"bufio"
	"fmt"
	"os"
)

func Input() [][]int {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([][]int, 0)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for j := 0; j < len(line); j++ {
			row = append(row, int(line[j]-'0'))
		}
		ret = append(ret, row)
	}
	return ret
}

func increaseEnergy(octos [][]int) {
	for i := 0; i < len(octos); i++ {
		for j := 0; j < len(octos); j++ {
			octos[i][j]++
		}
	}
}

type Point struct {
	row, col int
}

type Queue struct {
	ctn []*Point
}

func (this *Queue) EnQueue(v *Point) {
	this.ctn = append(this.ctn, v)
}

func (this *Queue) IsEmpty() bool {
	return len(this.ctn) == 0
}

func (this *Queue) DeQueue() *Point {
	res := this.ctn[0]
	this.ctn = this.ctn[1:len(this.ctn)]
	return res
}

func NewQueue() *Queue {
	return &Queue{
		ctn: make([]*Point, 0),
	}
}

func isValidPoint(grid [][]int, p *Point) bool {
	return p.row >= 0 && p.row < len(grid) && p.col >= 0 && p.col < len(grid[0])
}

var DIRECTIONS = []*Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}

func propagation(octos [][]int) {
	N := len(octos)
	qu := NewQueue()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if octos[i][j] > 9 {
				qu.EnQueue(&Point{i, j})
			}
		}
	}

	for !qu.IsEmpty() {
		p := qu.DeQueue()
		for i := 0; i < len(DIRECTIONS); i++ {
			adj := &Point{p.row + DIRECTIONS[i].row, p.col + DIRECTIONS[i].col}
			if !isValidPoint(octos, adj) || octos[adj.row][adj.col] > 9 {
				continue
			}
			octos[adj.row][adj.col]++
			if octos[adj.row][adj.col] > 9 {
				qu.EnQueue(adj)
			}
		}
	}
}

func reset(octos [][]int) int {
	N := len(octos)
	count := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if octos[i][j] <= 9 {
				continue
			}
			octos[i][j] = 0
			count++
		}
	}
	return count
}

func print(octos [][]int) {
	N := len(octos)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(octos[i][j])
		}
		fmt.Println()
	}
}

func partOne(grid [][]int) {
	const STEPS = 100
	octos := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		octos[i] = make([]int, len(grid))
		copy(octos[i], grid[i])
	}

	count := 0
	for i := 0; i < STEPS; i++ {
		increaseEnergy(octos)
		propagation(octos)
		count += reset(octos)
	}
	fmt.Println(count)
}

func partTwo(grid [][]int) {
	octos := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		octos[i] = make([]int, len(grid))
		copy(octos[i], grid[i])
	}

	step := 0
	THRESHOLD := len(grid) * len(grid)
	for {
		step++
		increaseEnergy(octos)
		propagation(octos)
		count := reset(octos)
		if count == THRESHOLD {
			break
		}
	}
	fmt.Println(step)
}

func main() {
	grid := Input()
	partOne(grid)
	partTwo(grid)
}
