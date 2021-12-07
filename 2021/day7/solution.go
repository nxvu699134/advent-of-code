// https://adventofcode.com/2021/day/7#part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	forward int
	down    int
}

func Input() []int {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	strs := strings.Split(scanner.Text(), ",")
	ret := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		n, _ := strconv.Atoi(strs[i])
		ret[i] = n
	}
	return ret
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func partOne(nums []int) {
	min := 1 << 31
	for i := 0; i < len(nums); i++ {
		diff := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			diff += abs(nums[j] - nums[i])
		}
		if diff < min {
			min = diff
		}
	}
	fmt.Println(min)
}

func partTwo(nums []int) {
	min := 1 << 31
	for i := 0; i < len(nums); i++ {
		diff := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			d := abs(nums[j] - nums[i])
			diff += (d + 1) * d / 2
		}
		if diff < min {
			min = diff
		}
	}
	fmt.Println(min)
}

func main() {
	nums := Input()
	partOne(nums)
	partTwo(nums)
}
