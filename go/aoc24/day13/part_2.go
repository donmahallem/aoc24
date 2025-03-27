package day13

import (
	"fmt"
	"io"
)

func Part2(in io.Reader) {
	data := LoadFile(in)
	totalSum := 0
	for _, inp := range data {
		inp.target[0] += 10000000000000
		inp.target[1] += 10000000000000
		a, b, ok := Calculate(&inp)
		if !ok {
			continue
		}
		totalSum += a*3 + b
	}
	fmt.Printf("%d\n", totalSum)
}
