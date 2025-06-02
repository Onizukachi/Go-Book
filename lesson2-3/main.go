package main

import (
	"fmt"
	"math/bits"
)

var pc [256]byte

func main() {
	var x uint64 = 1844674407370955312

	fmt.Println(PopCount(x))
	fmt.Println(MyCount(x))
	fmt.Println(bits.OnesCount64(x))
}

func onceCount(b byte) int {
	count := 0

	for i := 0; i < 8; i++ {
		if b&1 == 1 {
			count++
		}
		b >>= 1
	}

	return count
}

func MyCount(x uint64) int {
	count := 0

	for i := 0; i < 8; i++ {
		count += onceCount(byte(x >> (i * 8)))
	}

	return count
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
