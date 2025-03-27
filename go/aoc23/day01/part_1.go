package day01

import (
	"bufio"
	"fmt"
	"io"
)

func ParseLine(line []byte) int {
	var curVal, lastVal uint8 = 20, 20
	for i := range line {
		if (line)[i] >= '0' && (line)[i] <= '9' {
			if curVal >= 10 {
				curVal = (line)[i] - '0'
				lastVal = (line)[i] - '0'
			} else {
				lastVal = (line)[i] - '0'
			}
		}
	}
	return int(curVal*10 + lastVal)
}
func ParseFile(reader io.Reader) int {
	s := bufio.NewScanner(reader)
	summe := 0
	for s.Scan() {
		summe += ParseLine(s.Bytes())
	}
	return summe
}
func Part1(in io.Reader) {
	fmt.Printf("Result: %d\n", ParseFile(in))
}
