package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Entry struct {
	patterns []string
	output   []string
}

type Decoder struct {
	config map[byte]byte
	code   map[byte]byte
}

func NewEntry(s string) *Entry {
	tmp := strings.Split(s, " | ")
	return &Entry{
		patterns: strings.Fields(tmp[0]),
		output:   strings.Fields(tmp[1]),
	}
}

func Input() []*Entry {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([]*Entry, 0)
	for scanner.Scan() {
		entry := NewEntry(scanner.Text())
		ret = append(ret, entry)
	}
	return ret
}

func partOne(entries []*Entry) {
	count := 0
	for _, entry := range entries {
		for i := 0; i < len(entry.output); i++ {
			l := len(entry.output[i])
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func makeDecoder(config string) *Decoder {
	m := make(map[byte]byte)
	//direction of map: top -> down, left -> right
	//  0000
	// 1    2
	// 1    2
	//  3333
	// 4    5
	// 4    5
	//  6666
	for i := 0; i < len(config); i++ {
		m[config[i]] = byte(i)
	}
	code := make(map[byte]byte)
	code[byte(0b01110111)] = 0
	code[byte(0b00100100)] = 1
	code[byte(0b01011101)] = 2
	code[byte(0b01101101)] = 3
	code[byte(0b00101110)] = 4
	code[byte(0b01101011)] = 5
	code[byte(0b01111011)] = 6
	code[byte(0b00100101)] = 7
	code[byte(0b01111111)] = 8
	code[byte(0b01101111)] = 9

	return &Decoder{
		config: m,
		code:   code,
	}
}

func fromPatternToBits(s string) (int, int) {
	// return bits and len
	ret := 0
	for i := 0; i < len(s); i++ {
		ret |= 1 << (s[i] - 'a')
	}
	return ret, len(s)
}

func makeConfig(patterns []string) string {
	// consider a string in patterns
	// each char is position in mask
	// for example: 'acd' -> 1101
	// 0 1 2 3 4 5 6 7 8 9     1st line
	// 6 2 5 5 4 5 6 3 7 6     2nd line
	//   -     -     - -
	// some observations, with each observation is a exp: left = right
	// left = number of led is on (thats mean the 2nd line)
	// right = which leds are on (based on direction map)

	// I use xor to get exclution led
	// for example: 2 ^ 3 means (1(25) ^ 7(025) -> remaining 0)
	// 2 ^ 3 = 0                  #(0)
	// 2 ^ 4 = 13
	// 2 ^ 7 = 01346
	// 3 ^ 4 = 013
	// 3 ^ 7 = 1346
	// 4 ^ 7 = 046
	// 6 ^ 7 = 2,3,5 (we have 3 number has 6 leds on)
	// 5 ^ 7 = 15, 14, 24
	// 2 ^ 4 & any(6 ^ 7) = 3    #(1)
	// remove #(1) from 2 ^4 = 1 #(2)
	// 5 ^7 remove #(2) = 5,4,24

	bin := make(map[int][]int)
	for _, patt := range patterns {
		bits, l := fromPatternToBits(patt)
		if _, ok := bin[l]; !ok {
			bin[l] = make([]int, 0)
		}
		bin[l] = append(bin[l], bits)
	}

	led := make([]int, 7)
	led[0] = bin[2][0] ^ bin[3][0]

	b2xor4 := bin[2][0] ^ bin[4][0]
	for i := 0; i < len(bin[6]); i++ {
		b6xor7 := bin[6][i] ^ bin[7][0]
		if l3 := b6xor7 & b2xor4; l3 != 0 {
			led[3] = l3
			break
		}
	}

	led[1] = b2xor4 & (^led[3])

	b5xor7RemoveLed1 := make([]int, 0)
	for i := 0; i < len(bin[5]); i++ {
		b5xor7RemoveLed1 = append(b5xor7RemoveLed1, (bin[5][i]^bin[7][0])&(^led[1]))
	}
	for i := 0; i < len(b5xor7RemoveLed1); i++ {
		for j := i + 1; j < len(b5xor7RemoveLed1); j++ {
			if l4 := b5xor7RemoveLed1[i] & b5xor7RemoveLed1[j]; l4 != 0 {
				led[4] = l4
				break
			}
		}
	}
	for i := 0; i < len(b5xor7RemoveLed1); i++ {
		if b5xor7RemoveLed1[i]&led[4] == 0 {
			led[5] = b5xor7RemoveLed1[i]
		} else if b5xor7RemoveLed1[i] != led[4] {
			led[2] = b5xor7RemoveLed1[i] & (^led[4])
		}
	}

	led[6] = bin[4][0] ^ bin[7][0]&(^led[0])&(^led[4])

	ret := make([]byte, 0, len(led))
	for i := 0; i < len(led); i++ {
		c := byte(math.Log2(float64(led[i])))
		ret = append(ret, 'a'+c)
	}
	return string(ret)
}

func makeTranslaterFrom(decoder *Decoder) func(string) byte {
	return func(s string) byte {
		code := byte(0)
		for i := 0; i < len(s); i++ {
			code |= 1 << decoder.config[s[i]]
		}
		return decoder.code[code]
	}
}

func partTwo(entries []*Entry) {
	sum := 0
	for i := 0; i < len(entries); i++ {
		config := makeConfig(entries[i].patterns)
		decoder := makeDecoder(config)
		fromPatternToDigit := makeTranslaterFrom(decoder)
		n := 0
		for j := 0; j < len(entries[i].output); j++ {
			n = n*10 + int(fromPatternToDigit(entries[i].output[j]))
		}
		sum += n
	}
	fmt.Println(sum)
}

func main() {
	entries := Input()
	partOne(entries)
	partTwo(entries)
}
