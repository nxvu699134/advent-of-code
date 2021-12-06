// https://adventofcode.com/2021/day/6#part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const COOLDOWN = 7

// I see animal can "birth and undying",  I remember rabbit and mouse
// In Fibonacci seq
// 80th is 14472334024676221
// An array with every single element is "byte" (8 bits, so I can circular shift to mark internal)
// Thats cost 14472334024676221 bytes = 14472334 Gigs
// I see input with a tons of number so
// My computer may become a super NTN bomb
func Input() []int {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := make([]int, COOLDOWN+2)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), ",")
	for i := 0; i < len(strs); i++ { // How how many fish has internal couter = internal
		internal, _ := strconv.Atoi(strs[i])
		ret[internal]++
	}
	return ret
}

func internalAfter(internal []int, nDay int) []int {
	for i := 0; i < nDay; i++ {
		tmp := internal[0]
		for j := 0; j < len(internal)-1; j++ {
			internal[j] = internal[j+1]
		}
		// all zeroes becom 6 and birth 8
		// maybe can form a sequence here
		// something like seq(n) = seq(n-6) + seq(n-8)
		internal[COOLDOWN-1] += tmp
		internal[COOLDOWN+1] = tmp
	}
	return internal
}

func partOne(internal []int) {
	// atDay18 := internalAfter(internal, 18)
	// count := 0
	// for i := 0; i < len(atDay18); i++ {
	//   count += atDay18[i]
	// }
	atDay80 := internalAfter(internal, 80)
	count := 0
	for i := 0; i < len(atDay80); i++ {
		count += atDay80[i]
	}
	fmt.Println(count)
}

func partTwo(internal []int) {
	atDay256 := internalAfter(internal, 256)
	count := 0
	for i := 0; i < len(atDay256); i++ {
		count += atDay256[i]
	}
	fmt.Println(count)
}

func main() {
	internal := Input()
	partOne(internal)
	partTwo(internal)
}
