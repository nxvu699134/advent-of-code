// https://adventofcode.com/2021/day/21

package main

import (
	"bufio"
	"fmt"
	"os"
)

func Input() map[bool]*Player {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	l1 := scanner.Text()
	scanner.Scan()
	l2 := scanner.Text()
	return map[bool]*Player{
		true:  {int(l1[len(l1)-1] - '0'), 0},
		false: {int(l2[len(l2)-1] - '0'), 0},
	}
}

type Dice struct {
	current int
}

type Player struct {
	pos   int
	score int
}

func NewDice() *Dice {
	return &Dice{0}
}

func (this *Dice) Next() int {
	this.current++
	if this.current > 100 {
		this.current = 1
	}
	return this.current
}

func (this *Player) Roll(dice *Dice) bool {
	p := (dice.Next() + dice.Next() + dice.Next()) % 10
	this.pos += p
	if this.pos > 10 {
		this.pos = this.pos % 10
	}
	this.score += this.pos
	return this.score >= 1000
}

func partOne(players map[bool]*Player) {
	die := NewDice()
	turn := true
	numTurn := 0
	var loser *Player
	for {
		numTurn++
		if players[turn].Roll(die) {
			loser = players[!turn]
			break
		}
		turn = !turn
	}
	fmt.Println(loser.score * numTurn * 3)
}

func calcNewPlayer(p Player, rolled int) Player {
	p.pos += rolled
	if p.pos > 10 {
		p.pos = p.pos % 10
	}
	return Player{pos: p.pos, score: p.score + p.pos}
}

const WIN_SCORE = 21

func Dfs(p1, p2 Player, isP1 bool, freq map[int]int) (int, int) {
	if p1.score >= WIN_SCORE {
		return 1, 0
	}

	if p2.score >= WIN_SCORE {
		return 0, 1
	}

	countP1, countP2 := 0, 0
	for rolled, fz := range freq {
		if isP1 {
			p1Win, p2Win := Dfs(calcNewPlayer(p1, rolled), p2, !isP1, freq)
			countP1 += fz * p1Win
			countP2 += fz * p2Win
		} else {
			p1Win, p2Win := Dfs(p1, calcNewPlayer(p2, rolled), !isP1, freq)
			countP1 += fz * p1Win
			countP2 += fz * p2Win
		}
	}
	return countP1, countP2
}

func partTwo(players map[bool]*Player) {
	freq := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				freq[i+j+k]++
			}
		}
	}
	c1, c2 := Dfs(*players[true], *players[false], true, freq)
	fmt.Println(c1)
	fmt.Println(c2)
}

func main() {
	ps := Input()
	// partOne(ps)
	partTwo(ps)
}
