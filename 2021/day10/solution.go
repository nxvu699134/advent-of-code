// https://adventofcode.com/2021/day/9#part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Input() []string {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([]string, 0)

	for i := 0; scanner.Scan(); i++ {
		ret = append(ret, scanner.Text())
	}
	return ret
}

type Stack struct {
	ctn []byte
}

func (this *Stack) Push(v byte) {
	this.ctn = append(this.ctn, v)
}

func (this *Stack) IsEmpty() bool {
	return len(this.ctn) == 0
}

func (this *Stack) Size() int {
	return len(this.ctn)
}

func (this *Stack) Pop() byte {
	l := len(this.ctn) - 1
	res := this.ctn[l]
	this.ctn = this.ctn[:l]
	return res
}

func NewStack() *Stack {
	return &Stack{
		ctn: make([]byte, 0),
	}
}

type Checker struct {
	closeOf map[byte]byte
}

func NewChecker() *Checker {
	closeOf := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
		'<': '>',
	}
	return &Checker{closeOf}
}

func (this *Checker) isOpenTag(c byte) bool {
	_, ok := this.closeOf[c]
	return ok
}

func isValid(s string, checker *Checker) (bool, byte, string) {
	t := NewStack()
	for i := 0; i < len(s); i++ {
		if checker.isOpenTag(s[i]) {
			t.Push(s[i])
		} else {
			c := t.Pop()
			if s[i] != checker.closeOf[c] {
				return false, s[i], ""
			}
		}
	}
	incomplete := make([]byte, 0, t.Size())
	for !t.IsEmpty() {
		incomplete = append(incomplete, checker.closeOf[t.Pop()])
	}
	return true, 0, string(incomplete)
}

func partOne(lines []string, checker *Checker) {
	count := make(map[byte]int)
	for i := 0; i < len(lines); i++ {
		if ok, c, _ := isValid(lines[i], checker); !ok {
			count[c]++
		}
	}

	sum := 0
	point := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for k, v := range count {
		sum += point[k] * v
	}
	fmt.Println(sum)
}

func partTwo(lines []string, checker *Checker) {
	point := map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	scores := make([]int, 0)
	for i := 0; i < len(lines); i++ {
		if ok, _, remain := isValid(lines[i], checker); ok && len(remain) != 0 {
			sc := 0
			for j := 0; j < len(remain); j++ {
				sc = sc*5 + point[remain[j]]
			}
			scores = append(scores, sc)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func main() {
	lines := Input()
	checker := NewChecker()
	partOne(lines, checker)
	partTwo(lines, checker)
}
