// https://adventofcode.com/2021/day/12

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() *Passage {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := NewPassage()

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		ids := strings.Split(line, "-")
		ret.AddEdge(ids[0], ids[1])
	}
	return ret
}

type Node struct {
	id   string
	Next []string
}

type Passage struct {
	nodes map[string]*Node
}

func NewPassage() *Passage {
	return &Passage{
		nodes: make(map[string]*Node),
	}
}

func NewNode(id string) *Node {
	return &Node{
		id:   id,
		Next: make([]string, 0),
	}
}

func (this *Node) Connect(to string) {
	this.Next = append(this.Next, to)
}

func (this *Node) IsSmall() bool {
	return this.id != "start" && this.id != "end" && this.id[0] >= 'a' && this.id[0] <= 'z'
}

func (this *Passage) AddEdge(n1 string, n2 string) {
	if _, ok := this.nodes[n1]; !ok {
		this.nodes[n1] = NewNode(n1)
	}
	if _, ok := this.nodes[n2]; !ok {
		this.nodes[n2] = NewNode(n2)
	}
	this.nodes[n1].Connect(n2)
	this.nodes[n2].Connect(n1)
}

func (this *Passage) Print() {
	for k, v := range this.nodes {
		fmt.Printf("%5s ", k)
		fmt.Println(v.Next)
	}
	fmt.Println("=================================")
}

func Dfs(passage *Passage, id string, visitedSmall map[string]struct{}, path string, byPass bool) int {
	if id == "end" {
		// fmt.Println(path + ",end")
		return 1
	}
	check := false // this flag for backtracking
	if _, ok := visitedSmall[id]; ok {
		if byPass {
			byPass = false
			check = true
		} else {
			return 0
		}
	}
	n := passage.nodes[id]
	if n.IsSmall() {
		visitedSmall[id] = struct{}{}
	}
	path += "," + id
	count := 0
	for i := 0; i < len(n.Next); i++ {
		if n.Next[i] == "start" {
			continue
		}
		count += Dfs(passage, n.Next[i], visitedSmall, path, byPass)
	}
	if !check {
		delete(visitedSmall, id)
	}
	return count
}

func partOne(passage *Passage) {
	visitedSmall := make(map[string]struct{})
	numOfPath := Dfs(passage, "start", visitedSmall, "", false)
	fmt.Println(numOfPath)
}

func partTwo(passage *Passage) {
	visitedSmall := make(map[string]struct{})
	numOfPath := Dfs(passage, "start", visitedSmall, "", true)
	fmt.Println(numOfPath)
}

func main() {
	passage := Input()
	partOne(passage)
	partTwo(passage)
}
