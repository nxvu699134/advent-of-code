// https://adventofcode.com/2021/day/17

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Input() []*Point {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/test.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.Split(scanner.Text()[13:], ", ")
	xs := strings.Split(line[0][2:], "..")
	ys := strings.Split(line[1][2:], "..")
	fromX, _ := strconv.Atoi(xs[0])
	toX, _ := strconv.Atoi(xs[1])
	fromY, _ := strconv.Atoi(ys[1])
	toY, _ := strconv.Atoi(ys[0])
	return []*Point{
		{fromX, fromY},
		{toX, toY},
	}

}

type Point struct {
	x, y int
}

func visualize(initVelocity *Point, basket []*Point, step int) {
	N := 100
	grid := make([][]byte, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]byte, N)
	}
	s := &Point{0, N / 2}
	for i := s.y - basket[1].y; i >= s.y-basket[0].y; i-- {
		for j := s.x + basket[0].x; j <= s.x+basket[1].x; j++ {
			grid[i][j] = 2
		}
	}

	grid[s.y][s.x] = 1
	vel := initVelocity
	for i := 0; i < step; i++ {
		s.x += vel.x
		s.y -= vel.y
		if vel.x > 0 {
			vel.x--
		}
		vel.y--
		grid[s.y][s.x] = 3
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if v := grid[i][j]; v == 0 {
				fmt.Print(". ")
			} else if v == 1 {
				fmt.Print("s ")
			} else if v == 2 {
				fmt.Print("T ")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
	fmt.Println("=============================")
}

func partOne(ps []*Point) {
	// Consider only x-axis
	// after any step velocity slowdown by 1 until 0
	// so valid velocity is fromX <= 0+1+2+3+....n <= toX
	// n^2 + n                  >= 2*fromX
	// n^2 + 2*1/2*n + 1/4 - 1/4 >= 2*fromX
	// (n + 1/2)^2              >= 2*fromX + 1/4
	// n  = sqrt(2*fromX + 1/4) - 1/2 (eliminate -n)
	//
	// About y-axis
	//              #
	//            # |
	//          #   |
	//       #      | height = S
	//    #         |
	// #            |
	// s-------------------------------
	//                  |
	//                  |
	//                  | height = S1
	//            TTTTT |
	//            TTTTT |
	//            TTTTT |<-- lowest point
	//
	// S is an arithmetic sequence 0+1+2+...+n
	// To maximize height S and bot drop into target area
	// S1 must be the next number in sequence n+1

	vX := int(math.Ceil(math.Sqrt(float64(2*ps[0].x)+0.25) - 0.5))
	vY := -ps[1].y - 1
	vel := &Point{vX, vY}
	fmt.Println(vel)
	height := (vY + 1) * vY / 2
	fmt.Println(height)
}

func partTwo(ps []*Point) {
	min_vX := int(math.Ceil(math.Sqrt(float64(2*ps[0].x)+0.25) - 0.5))
	// max_vX := ps[1].x
	_ = min_vX
	// all valid x value is sum of last n number of sequence
	// n reach maximum when last n number is full sequence
	// fmt.Println(validX)
	// Do it latter
	// ......
}

func main() {
	ps := Input()
	// partOne(ps)
	partTwo(ps)
}
