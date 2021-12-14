// https://adventofcode.com/2021/day/14

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() (string, map[string]byte) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	templateStr := scanner.Text()
	scanner.Scan()

	rules := make(map[string]byte)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " -> ")
		rules[tmp[0]] = tmp[1][0]
	}
	return templateStr, rules
}

func fromStringToMap(s string) (map[string]int, map[byte]int) {
	template := make(map[string]int)
	charCount := make(map[byte]int)
	for i := 0; i < len(s)-1; i++ {
		token := s[i : i+2]
		template[token]++
		charCount[s[i]]++
	}
	charCount[s[len(s)-1]]++
	return template, charCount
}

func process(template map[string]int, rules map[string]byte, charCount map[byte]int) map[string]int {
	newTemplate := make(map[string]int)
	for token, freq := range template {
		newToken1 := string([]byte{token[0], rules[token]})
		newToken2 := string([]byte{rules[token], token[1]})
		newTemplate[newToken1] += freq
		newTemplate[newToken2] += freq
		charCount[rules[token]] += freq
	}
	return newTemplate
}

func printMap(m map[byte]int) {
	for k, v := range m {
		fmt.Printf("%c:%d ", k, v)
	}
	fmt.Println()
}

func partOne(templateStr string, rules map[string]byte) {
	const STEP = 10
	template, charCount := fromStringToMap(templateStr)
	for i := 0; i < STEP; i++ {
		template = process(template, rules, charCount)
	}

	// I was fucked up cuz didnt count char step by step
	// If count char at the end, my result alway approximate x2 the answer
	// Because of the way we tokenize (last char of prev is begin of the next)

	min, max := 1<<31, -1
	for _, v := range charCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}

func partTwo(templateStr string, rules map[string]byte) {
	const STEP = 40
	template, charCount := fromStringToMap(templateStr)
	for i := 0; i < STEP; i++ {
		template = process(template, rules, charCount)
	}

	min, max := 1<<60, -1
	for _, v := range charCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}

func main() {
	template, rules := Input()
	partOne(template, rules)
	partTwo(template, rules)
}
