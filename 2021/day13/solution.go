// https://adventofcode.com/2021/day/13

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() ([][]bool, []Instruction) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxX, maxY := 0, 0
	tmp := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		cord := strings.Split(line, ",")
		x, _ := strconv.Atoi(cord[0])
		y, _ := strconv.Atoi(cord[1])
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		tmp = append(tmp, []int{x, y})
	}

	ins := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line[13:])
		if line[11] == 'x' {
			ins = append(ins, FoldX{n})
		} else {
			ins = append(ins, FoldY{n})
		}
	}

	grid := make([][]bool, maxY+1)
	for i := 0; i <= maxY; i++ {
		grid[i] = make([]bool, maxX+1)
	}
	for i := 0; i < len(tmp); i++ {
		grid[tmp[i][1]][tmp[i][0]] = true
	}
	return grid, ins
}

type FoldX struct {
	at int
}

type FoldY struct {
	at int
}

type Instruction interface {
	fold(grid [][]bool) [][]bool
}

func (this FoldX) fold(grid [][]bool) [][]bool {
	newGrid := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 1; this.at+j < len(grid[0]); j++ {
			grid[i][this.at-j] = grid[i][this.at-j] || grid[i][this.at+j]
		}
		newGrid[i] = grid[i][:this.at]
	}
	return newGrid
}

func (this FoldY) fold(grid [][]bool) [][]bool {
	for i := 1; this.at+i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[this.at-i][j] = grid[this.at-i][j] || grid[this.at+i][j]
		}
	}
	return grid[:this.at]
}

func printGrid(grid [][]bool) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println("==================================")
}

func partOne(grid [][]bool, in Instruction) {
	newGrid := in.fold(grid)
	count := 0
	for i := 0; i < len(newGrid); i++ {
		for j := 0; j < len(newGrid[i]); j++ {
			if newGrid[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func partTwo(grid [][]bool, ins []Instruction) {
	for i := 0; i < len(ins); i++ {
		grid = ins[i].fold(grid)
	}
	printGrid(grid)
}

func main() {
	grid, ins := Input()
	// printGrid(grid)

	partOne(grid, ins[0])
	partTwo(grid, ins)
}
