// https://adventofcode.com/2021/day/17

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() []*RebootStep {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/test.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := make([]*RebootStep, 0)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		state := true
		if tmp[0] == "off" {
			state = false
		}
		tmp2 := strings.Split(tmp[1], ",")
		tmpCoord := make([][]int, 3)
		for i := 0; i < len(tmp2); i++ {
			coordStr := strings.Split(tmp2[i][2:], "..")
			from, _ := strconv.Atoi(coordStr[0])
			to, _ := strconv.Atoi(coordStr[1])
			if to < from {
				from, to = to, from
			}
			tmpCoord[i] = []int{from, to}
		}
		ret = append(ret,
			&RebootStep{
				on:   state,
				from: Coord{tmpCoord[0][0], tmpCoord[1][0], tmpCoord[2][0]},
				to:   Coord{tmpCoord[0][1], tmpCoord[1][1], tmpCoord[2][1]},
			})
	}
	standardize(ret)
	return ret
}

type Coord struct {
	x, y, z int
}

type RebootStep struct {
	on   bool
	from Coord
	to   Coord
}

func (this *Coord) toArray() []int {
	return []int{this.x, this.y, this.z}
}

func (this *Coord) shift(n int) {
	this.x += n
	this.y += n
	this.z += n
}

func hash(x, y, z int) int {
	return x*1000000000000 + y*1000000 + z
}

func standardize(steps []*RebootStep) {
	min := 1 << 32
	for i := 0; i < len(steps); i++ {
		arr := append(steps[i].from.toArray(), steps[i].to.toArray()...)
		for j := 0; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
			}
		}
	}
	if min < 0 {
		min = -min
	}
	for i := 0; i < len(steps); i++ {
		steps[i].from.shift(min)
		steps[i].to.shift(min)
	}
}

func partOne(steps []*RebootStep) {
	onMap := make(map[int]struct{})
	for i := 0; i < len(steps); i++ {
		from := steps[i].from
		to := steps[i].to
		isOn := steps[i].on
		for j := from.x; j <= to.x; j++ {
			for k := from.y; k <= to.y; k++ {
				for h := from.z; h <= to.z; h++ {
					if isOn {
						onMap[hash(j, k, h)] = struct{}{}
					} else {
						delete(onMap, hash(j, k, h))
					}
				}
			}
		}
	}
	fmt.Println(len(onMap))
}

func partTwo(steps []*RebootStep) {
	// after standardize max value all axises is ~2e5
	// Thats mean the space is 200k ^ 3 ~ 8e15
	// if a hash is a bit, space cost 8e15/1e9 = 8e6 Gigs
	// But I used int64 hash, so this solution didnt work
	onMap := make(map[int]struct{})
	for i := 0; i < len(steps); i++ {
		from := steps[i].from
		to := steps[i].to
		isOn := steps[i].on
		for j := from.x; j <= to.x; j++ {
			for k := from.y; k <= to.y; k++ {
				for h := from.z; h <= to.z; h++ {
					if isOn {
						onMap[hash(j, k, h)] = struct{}{}
					} else {
						delete(onMap, hash(j, k, h))
					}
				}
			}
		}
	}
	fmt.Println(len(onMap))
}

func main() {
	steps := Input()
	// partOne(steps[:21])
	partTwo(steps)
}
