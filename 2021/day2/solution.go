package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Command struct {
	forward int
	down    int
}

func Input() []Command {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([]Command, 0)
	for scanner.Scan() {
		line := scanner.Text()
		spaceIdx := len(line) - 1
		for ; spaceIdx >= 0; spaceIdx-- {
			if line[spaceIdx] == ' ' {
				break
			}
		}
		typeCommand := line[:spaceIdx]
		unit, _ := strconv.Atoi(line[spaceIdx+1:])
		if typeCommand == "forward" {
			ret = append(ret, Command{unit, 0})
		} else {
			sign := 1
			if typeCommand == "up" {
				sign = -1
			}
			ret = append(ret, Command{0, unit * sign})
		}
	}
	return ret
}

func main() {
	input := Input()
	h, d, a := 0, 0, 0
	for i := 0; i < len(input); i++ {
		h += input[i].forward
		a += input[i].down
		d += input[i].forward * a
	}
	fmt.Println(h * d)
}
