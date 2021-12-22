// https://adventofcode.com/2021/day/17

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() []*Cuboid {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := make([]*Cuboid, 0)
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
			&Cuboid{
				on:   state,
				from: Coord{tmpCoord[0][0], tmpCoord[1][0], tmpCoord[2][0]},
				to:   Coord{tmpCoord[0][1], tmpCoord[1][1], tmpCoord[2][1]},
			})
	}
	return ret
}

type Coord struct {
	x, y, z int
}

type Cuboid struct {
	on   bool
	from Coord
	to   Coord
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func (this *Cuboid) intersectWith(cuboid *Cuboid) *Cuboid {
	//    -------------------------->
	//    |       ________
	//    |      |    S2  |
	//    |  ____|___     |
	//    | |    | S3|    |
	//    | |    |___|____|
	//    | | S1     |
	//    | |________|
	//    |
	//    v
	// Determine S : top left point, bottom right point
	//S1 : (x,y), (x',y')   x < x', y < y'
	//S2: (a,b), (a',b')    a < a', b < b'
	//S3: (max(x,a), max(y,b)), (min(x',a'), min(y',b'))
	// S3 must have: max(x,a) < min(x',a') and max(y,b) < min(y',b')
	// I think in 3D space, rule for z-axis is same with 2D
	c1 := Coord{
		max(this.from.x, cuboid.from.x),
		max(this.from.y, cuboid.from.y),
		max(this.from.z, cuboid.from.z),
	}
	c2 := Coord{
		min(this.to.x, cuboid.to.x),
		min(this.to.y, cuboid.to.y),
		min(this.to.z, cuboid.to.z),
	}
	if c1.x <= c2.x && c1.y <= c2.y && c1.z <= c2.z {
		return &Cuboid{
			on:   false,
			from: c1,
			to:   c2,
		}
	}
	return nil
}

func (this *Cuboid) getVolumn() int {
	v := (this.to.x - this.from.x + 1) * (this.to.y - this.from.y + 1) * (this.to.z - this.from.z + 1)
	if this.on {
		return v
	}
	return -v
}

func spaceVolumn(cuboids []*Cuboid) int {
	space := make([]*Cuboid, 0)
	for _, cuboid := range cuboids {
		for _, spaceCuboid := range space {
			intersectCuboid := cuboid.intersectWith(spaceCuboid)
			if intersectCuboid != nil {
				intersectCuboid.on = !spaceCuboid.on
				// Why have this line
				// space list we have 3 squares: S1 + S2 - S3
				// Did u notice S3 counted 2 times(in S1 and S2)
				// so we should "compensate" thats region
				// Same in reverse case
				space = append(space, intersectCuboid)
			}
		}
		if cuboid.on {
			space = append(space, cuboid)
		}
	}

	ret := 0
	for i := 0; i < len(space); i++ {
		ret += space[i].getVolumn()
	}
	return ret
}

func partOne(cuboids []*Cuboid) {
	fmt.Println(spaceVolumn(cuboids[:20]))
}

func partTwo(cuboids []*Cuboid) {
	fmt.Println(spaceVolumn(cuboids))
}

func main() {
	cuboids := Input()
	partOne(cuboids)
	partTwo(cuboids)
}
