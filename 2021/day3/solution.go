package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Input() []int {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.ParseInt(line, 2, 32)
		ret = append(ret, int(n))
	}
	return ret
}

// func main() {
//   in := Input()
//   N := 12
//   bin := make([]int, N)
//   for i := 0; i < len(in); i++ {
//     for n, count := in[i], 0; n != 0; n, count = n>>1, count+1 {
//       bin[count] += n & 1
//     }
//   }
//   bound := len(in) / 2
//   gamma := 0
//   for i := len(bin) - 1; i >= 0; i-- {
//     gamma <<= 1
//     if bin[i] > bound {
//       gamma |= 1
//     }
//   }
//   mask := (1 << N) - 1
//   fmt.Println(gamma * (^gamma & mask))
// }

func getBin(nums []int, pos int, isMost bool) []int {
	mask := 1 << pos
	onesBin := make([]int, 0)
	zerosBin := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i]&mask == 0 {
			zerosBin = append(zerosBin, nums[i])
		} else {
			onesBin = append(onesBin, nums[i])
		}
	}

	if isMost {
		if len(onesBin) >= len(zerosBin) {
			return onesBin
		} else {
			return zerosBin
		}
	} else {
		if len(zerosBin) <= len(onesBin) {
			return zerosBin
		} else {
			return onesBin
		}
	}
}

func main() {
	in := Input()
	N := 12
	oxy := in
	for i := N - 1; len(oxy) > 1; i-- {
		oxy = getBin(oxy, i, true)
	}
	co := in
	for i := N - 1; len(co) > 1; i-- {
		co = getBin(co, i, false)
	}
	fmt.Println(oxy[0] * co[0])
}
