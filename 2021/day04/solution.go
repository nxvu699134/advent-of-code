// https://adventofcode.com/2021/day/4

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BOARD_SIZE = 5

type Board struct {
	state int
	board [][]int
}

type Winner struct {
	id       int
	bingoNum int
	score    int
}

type Game struct {
	boards    []*Board
	criteria  []int
	winners   []*Winner
	winnerMap map[int]struct{}
}

func ConstructBoard(nums []int) *Board {
	board := make([][]int, BOARD_SIZE)
	for i := 0; i < BOARD_SIZE; i++ {
		board[i] = make([]int, BOARD_SIZE)
		for j := 0; j < BOARD_SIZE; j++ {
			board[i][j] = nums[i*BOARD_SIZE+j]
		}
	}
	return &Board{0, board}
}

func (this *Board) mark(num int) bool {
	shift := -1
	for i := 0; i < len(this.board); i++ {
		for j := 0; j < len(this.board); j++ {
			if this.board[i][j] == num {
				shift = i*BOARD_SIZE + j
				break
			}
			if shift != -1 {
				break
			}
		}
	}
	if shift != -1 {
		this.state |= 1 << shift
		return true
	}
	return false
}

func (this *Board) checkWin(criteria []int) bool {
	for i := 0; i < len(criteria); i++ {
		if this.state&criteria[i] == criteria[i] {
			return true
		}
	}
	return false
}

func (this *Board) getScore() int {
	score := 0
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			mask := 1 << (i*BOARD_SIZE + j)
			if mask&this.state == 0 {
				score += this.board[i][j]
			}
		}
	}
	return score
}

func ConstructGame() *Game {
	criteria := make([]int, 0, BOARD_SIZE*BOARD_SIZE)

	//check row
	criteria = append(criteria, (1<<BOARD_SIZE)-1)
	for i := 1; i < BOARD_SIZE; i++ {
		criteria = append(criteria, criteria[len(criteria)-1]<<BOARD_SIZE)
	}

	//check col
	// criteria = append(criteria, int(0b00000000000100001000010000100001))
	colCre := 1
	for i := 1; i < BOARD_SIZE; i++ {
		colCre <<= BOARD_SIZE
		colCre |= 1
	}
	criteria = append(criteria, colCre)
	for i := 1; i < BOARD_SIZE; i++ {
		criteria = append(criteria, criteria[len(criteria)-1]<<1)
	}

	return &Game{
		boards:    make([]*Board, 0),
		criteria:  criteria,
		winners:   make([]*Winner, 0),
		winnerMap: make(map[int]struct{}),
	}
}

func (this *Game) addBoard(board *Board) {
	this.boards = append(this.boards, board)
}

func (this *Game) draw(num int) {
	for i := 0; i < len(this.boards); i++ {
		_, isWinner := this.winnerMap[i]
		if !isWinner && this.boards[i].mark(num) && this.boards[i].checkWin(this.criteria) {
			this.winners = append(this.winners, &Winner{i, num, this.boards[i].getScore()})
			this.winnerMap[i] = struct{}{}
		}
	}
}

func Input() ([]int, *Game) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	guessStr := strings.Split(scanner.Text(), ",")
	guessNums := make([]int, 0, len(guessStr))

	for i := 0; i < len(guessStr); i++ {
		n, _ := strconv.Atoi(guessStr[i])
		guessNums = append(guessNums, n)
	}

	game := ConstructGame()
	for scanner.Scan() {
		nums := make([]int, 0, BOARD_SIZE*BOARD_SIZE)
		for i := 0; i < BOARD_SIZE; i++ {
			scanner.Scan()
			line := scanner.Text()
			strs := strings.Fields(line)
			for j := 0; j < len(strs); j++ {
				n, _ := strconv.Atoi(strs[j])
				nums = append(nums, n)
			}
		}
		game.addBoard(ConstructBoard(nums))
	}
	return guessNums, game
}

func partOne(game *Game, guess []int) {
	for i := 0; i < len(guess); i++ {
		game.draw(guess[i])
		if len(game.winners) != 0 {
			first := game.winners[0]
			fmt.Printf("1. Board %2d, bingo number %2d, score %d\n", first.id, first.bingoNum, first.score*first.bingoNum)
			break
		}
	}
}
func partTwo(game *Game, guess []int) {
	for i := 0; i < len(guess); i++ {
		game.draw(guess[i])
	}
	last := game.winners[len(game.winners)-1]
	fmt.Printf("%2d. Board %2d, bingo number %2d, score %d\n", len(game.winners)+1, last.id, last.bingoNum, last.score*last.bingoNum)
}

func main() {
	guess, game := Input()
	partOne(game, guess)
	partTwo(game, guess)
}
