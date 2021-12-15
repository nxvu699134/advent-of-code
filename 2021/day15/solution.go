// https://adventofcode.com/2021/day/15

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

	for scanner.Scan() {
		line := scanner.Text()
		tmp := make([]int, 0, len(line))
		for i := 0; i < len(line); i++ {
			tmp = append(tmp, int(line[i]-'0'))
		}
		ret = append(ret, tmp)
	}
	return ret
}

func Min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func Print(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d", grid[i][j])
		}
		fmt.Println()
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

func (this *Queue) Size() int {
	return len(this.ctn)
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

const INF = 1 << 31

func getNearPoints(grid [][]int, p *Point) []*Point {
	ret := make([]*Point, 0)
	if p.row > 0 {
		ret = append(ret, &Point{p.row - 1, p.col})
	}
	if p.col > 0 {
		ret = append(ret, &Point{p.row, p.col - 1})
	}
	if p.row+1 < len(grid) {
		ret = append(ret, &Point{p.row + 1, p.col})
	}
	if p.col+1 < len(grid[0]) {
		ret = append(ret, &Point{p.row, p.col + 1})
	}
	return ret
}

func Bfs(grid [][]int) {
	ROW, COL := len(grid), len(grid[0])
	cost := make([][]int, ROW)
	for i := 0; i < ROW; i++ {
		cost[i] = make([]int, COL)
		for j := 0; j < COL; j++ {
			cost[i][j] = INF
		}
	}
	cost[0][0] = grid[0][0]

	q := NewQueue()
	q.EnQueue(&Point{0, 0})
	for !q.IsEmpty() {
		cur := q.DeQueue()
		for _, nearPoint := range getNearPoints(grid, cur) {
			costToNear := cost[cur.row][cur.col] + grid[nearPoint.row][nearPoint.col]
			if costToNear < cost[nearPoint.row][nearPoint.col] {
				cost[nearPoint.row][nearPoint.col] = costToNear
				q.EnQueue(nearPoint)
			}
		}
	}
	fmt.Printf("Min cost: %d\n", cost[ROW-1][COL-1]-grid[0][0])
}
func partOne(grid [][]int) {
	Bfs(grid)
}

func partTwo(grid [][]int) {
	ROW, COL := len(grid), len(grid[0])
	NEW_ROW, NEW_COL := ROW*5, COL*5
	newGrid := make([][]int, NEW_ROW)
	for i := 0; i < NEW_ROW; i++ {
		newGrid[i] = make([]int, COL*5)
		for j := 0; j < NEW_COL; j++ {
			newGrid[i][j] = grid[i%ROW][j%COL] + i/ROW + j/COL
			if newGrid[i][j] > 9 {
				newGrid[i][j] %= 9
			}
		}
	}
	Bfs(newGrid)
}

func main() {
	grid := Input()
	partOne(grid)
	partTwo(grid)
}
