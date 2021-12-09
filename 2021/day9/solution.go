// https://adventofcode.com/2021/day/9#part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func isLowerPoint(grid [][]int, row, col int) bool {
	if row > 0 && grid[row][col] >= grid[row-1][col] {
		return false
	}
	if col+1 < len(grid[row]) && grid[row][col] >= grid[row][col+1] {
		return false
	}
	if row+1 < len(grid) && grid[row][col] >= grid[row+1][col] {
		return false
	}
	if col > 0 && grid[row][col] >= grid[row][col-1] {
		return false
	}
	return true
}

func partOne(grid [][]int) {
	low := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isLowerPoint(grid, i, j) {
				low = append(low, grid[i][j])
			}
		}
	}

	sum := 0
	for i := 0; i < len(low); i++ {
		sum += low[i] + 1
	}
	fmt.Println(sum)
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

func checkBasinFrom(grid [][]int, checked [][]bool) func(low, high *Point) bool {
	return func(low *Point, high *Point) bool {
		return !checked[high.row][high.col] && grid[low.row][low.col] < grid[high.row][high.col] && grid[high.row][high.col] != 9
	}
}

func floodFill(grid [][]int, start *Point) int {
	ROW := len(grid)
	COL := len(grid[0])
	q := NewQueue()
	q.EnQueue(start)
	checked := make([][]bool, ROW)
	for i := 0; i < len(checked); i++ {
		checked[i] = make([]bool, COL)
	}
	count := 0
	isBasin := checkBasinFrom(grid, checked)
	for !q.IsEmpty() {
		p := q.DeQueue()
		if checked[p.row][p.col] {
			continue
		}
		checked[p.row][p.col] = true
		count++

		if top := (&Point{p.row - 1, p.col}); top.row >= 0 && isBasin(p, top) {
			q.EnQueue(top)
		}
		if right := (&Point{p.row, p.col + 1}); right.col < COL && isBasin(p, right) {
			q.EnQueue(right)
		}
		if bot := (&Point{p.row + 1, p.col}); bot.row < ROW && isBasin(p, bot) {
			q.EnQueue(bot)
		}
		if left := (&Point{p.row, p.col - 1}); left.col >= 0 && isBasin(p, left) {
			q.EnQueue(left)
		}
	}

	// for i := 0; i < ROW; i++ {
	//   for j := 0; j < COL; j++ {
	//     fmt.Print(grid[i][j])
	//     if checked[i][j] {
	//       fmt.Print("*")
	//     } else {
	//       fmt.Print(" ")
	//     }
	//   }
	//   fmt.Println()
	// }
	// fmt.Println()
	return count
}

func partTwo(grid [][]int) {
	low := make([]*Point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isLowerPoint(grid, i, j) {
				low = append(low, &Point{i, j})
			}
		}
	}

	basinSizes := make([]int, 0)
	for i := 0; i < len(low); i++ {
		basinSizes = append(basinSizes, floodFill(grid, low[i]))
	}

	sort.Slice(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}

func main() {
	grid := Input()
	partOne(grid)
	partTwo(grid)
}
