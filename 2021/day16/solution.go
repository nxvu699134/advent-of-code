// https://adventofcode.com/2021/day/16

package main

import (
	"bufio"
	"fmt"
	"os"
)

func Input() *BitStream {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/input.txt")
	fmt.Println(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	hexStr := scanner.Text()
	bt := NewBitStream(hexStr)
	// fmt.Printf("%08b", bt.cnt)
	// fmt.Println()
	return bt
}

func hexToDec(c byte) byte {
	if '0' <= c && c <= '9' {
		return c - '0'
	}
	if 'A' <= c && c <= 'F' {
		return c - 'A' + 10
	}
	return 1<<8 - 1
}

type BitStream struct {
	cnt     []byte
	idxByte int
	mask    byte
}

type Packet interface {
	getValue() int
	isOperator() bool
}

type OperatorPacket struct {
	version  int
	pType    int
	packages []Packet
}

type LiteralPacket struct {
	version int
	pType   int
	content []int
}

var PacketType = struct{ Sum, Product, Minimum, Maximum, Literal, GreaterThan, LessThan, EqualTo int }{0, 1, 2, 3, 4, 5, 6, 7}

func NewBitStream(s string) *BitStream {
	nLoop := len(s) / 2
	cnt := make([]byte, 0, nLoop+(len(s)%2)&1)
	i := 0
	for j := 0; j < nLoop; i, j = i+2, j+1 {
		tmp := (hexToDec(s[i]) << 4) | hexToDec(s[i+1])
		cnt = append(cnt, tmp)
	}
	if len(s)%2 != 0 {
		cnt = append(cnt, hexToDec(s[i])<<4)
	}
	return &BitStream{cnt: cnt, idxByte: 0, mask: 1 << 7}
}

func (this *BitStream) nBits() int {
	return len(this.cnt) * 8
}

func (this *BitStream) reset() {
	this.idxByte = 0
	this.mask = 1 << 7
}

// TODO: should return number of bit proceseed here
func (this *BitStream) getBits(nBits int) (int, bool) {
	if this.idxByte >= len(this.cnt) {
		return 0, false
	}
	ret := 0
	for i := 0; i < nBits; i++ {
		if this.idxByte >= len(this.cnt) {
			return 0, false
		}
		ret <<= 1
		if this.mask&this.cnt[this.idxByte] != 0 {
			ret |= 1
		}
		this.mask >>= 1
		if this.mask == 0 {
			this.idxByte++
			this.mask = 1 << 7
		}
	}
	return ret, true
}

func (this *LiteralPacket) getValue() int {
	ret := 0
	for i := 0; i < len(this.content); i++ {
		ret <<= 4
		ret |= int(this.content[i])
	}
	return ret
}

func (this *LiteralPacket) isOperator() bool { return false }

func (this *OperatorPacket) getValue() int {
	switch this.pType {
	case PacketType.Sum:
		sum := 0
		for i := 0; i < len(this.packages); i++ {
			sum += this.packages[i].getValue()
		}
		return sum
	case PacketType.Product:
		product := 1
		for i := 0; i < len(this.packages); i++ {
			product *= this.packages[i].getValue()
		}
		return product
	case PacketType.Minimum:
		min := this.packages[0].getValue()
		for i := 1; i < len(this.packages); i++ {
			if v := this.packages[i].getValue(); v < min {
				min = v
			}
		}
		return min
	case PacketType.Maximum:
		max := this.packages[0].getValue()
		for i := 1; i < len(this.packages); i++ {
			if v := this.packages[i].getValue(); v > max {
				max = v
			}
		}
		return max
	case PacketType.GreaterThan:
		if this.packages[0].getValue() > this.packages[1].getValue() {
			return 1
		}
		return 0
	case PacketType.LessThan:
		if this.packages[0].getValue() < this.packages[1].getValue() {
			return 1
		}
		return 0
	case PacketType.EqualTo:
		if this.packages[0].getValue() == this.packages[1].getValue() {
			return 1
		}
		return 0
	}
	return 0
}

func (this *OperatorPacket) isOperator() bool { return true }

func Parse(bitstream *BitStream, nBits int) ([]Packet, int) {
	ret := make([]Packet, 0)
	nProcessedBits := 0
	//Why have this cheat
	// In case parse operator package
	// we have 2 ways, use size of package and number of sub-package contained
	// I want bypass check by set nbits to -1 and the init processed bits to any number < -1
	// Thats guarantee this function alway do 1 time
	// I was fucked here because I forgot to re +2 to nProcessedBits and thats cause much problem
	if nBits == -1 {
		nProcessedBits = -2
	}
	for nProcessedBits < nBits {
		version, ok := bitstream.getBits(3)
		if !ok {
			fmt.Println("version")
			break
		}
		nProcessedBits += 3
		pType, ok := bitstream.getBits(3)
		if !ok {
			//TODO: revert processed bits
			break
		}
		nProcessedBits += 3
		if pType == PacketType.Literal {
			content := make([]int, 0)
			parity := 1 << 4
			for {
				chunk, ok := bitstream.getBits(5)
				if !ok {
					//TODO: revert processed bits
				}
				nProcessedBits += 5
				content = append(content, chunk&(parity-1))
				if parity&chunk == 0 {
					break
				}
			}
			ret = append(ret, &LiteralPacket{version, pType, content})
		} else {
			label, ok := bitstream.getBits(1)
			if !ok {
				//TODO: revert processed bits
			}
			nProcessedBits += 1
			if label == 1 {
				nSub, ok := bitstream.getBits(11)
				if !ok {
					//TODO: revert processed bits
				}
				nProcessedBits += 11
				ps := make([]Packet, 0)
				for i := 0; i < nSub; i++ {
					pack, nParsedBits := Parse(bitstream, -1)
					nProcessedBits += nParsedBits
					ps = append(ps, pack...)
				}
				ret = append(ret, &OperatorPacket{version, pType, ps})
			} else {
				nNextBits, ok := bitstream.getBits(15)
				if !ok {
					//TODO: revert processed bits
				}
				nProcessedBits += 15
				ps, nParsedBits := Parse(bitstream, nNextBits)
				nProcessedBits += nParsedBits
				ret = append(ret, &OperatorPacket{version, pType, ps})
			}
		}
		if nBits < 0 {
			nProcessedBits += 2 // This line save my life, fuck that
			break
		}
	}
	return ret, nProcessedBits
}

func getSumVersion(packets []Packet) int {
	sum := 0
	for i := 0; i < len(packets); i++ {
		if packets[i].isOperator() {
			op, _ := packets[i].(*OperatorPacket)
			sum += op.version
			sum += getSumVersion(op.packages)
		} else {
			lp, _ := packets[i].(*LiteralPacket)
			sum += lp.version
		}
	}
	return sum
}

func partOne(bitstream *BitStream) {
	ps, _ := Parse(bitstream, bitstream.nBits())
	fmt.Println(getSumVersion(ps))
}

func partTwo(bitstream *BitStream) {
	bitstream.reset()
	ps, nProcessedBits := Parse(bitstream, bitstream.nBits())
	fmt.Println(nProcessedBits, bitstream.nBits())
	fmt.Println(ps[0].getValue())
}

func main() {
	bitstream := Input()
	partOne(bitstream)
	partTwo(bitstream)
}
