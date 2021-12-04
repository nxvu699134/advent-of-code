package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Input() []int {
	pwd, _ := os.Getwd()
	file, _ := os.Open(pwd + "/2021/day1/input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([]int, 0)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		ret = append(ret, val)
	}
	return ret
}

func main() {
	input := Input()
	count := 0
	curSum := input[0] + input[1] + input[2]
	for i := 1; i < len(input)-2; i++ {
		nextSum := input[i] + input[i+1] + input[i+2]
		if nextSum > curSum {
			count++
		}
		curSum = nextSum
	}
	fmt.Println(input[0:10])
	fmt.Println(count)
}
