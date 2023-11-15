package main

import (
	"encoding/hex"
	"fmt"

	"github.com/holiman/uint256"
)

func opPush1(pc *uint64, code []byte, st *Stack) {
	*pc++
	nextByte := code[*pc]
	st.push(*uint256.NewInt(uint64(nextByte)))
}

func opPop(st *Stack) {
	st.pop()
}

func opAdd(st *Stack) {
	x := st.pop()
	y := st.pop()
	st.push(*y.Add(&x, &y))
}

func main() {
	// byteCode := "60FF606050" // push and pop
	byteCode := "600a600a01"
	code, err := hex.DecodeString(byteCode)
	if err != nil {
		fmt.Println("Failed to decode hex string")
	}
	var (
		stack = &Stack{}
		pc    = uint64(0)
	)

	for {
		if pc < uint64(len(code)) {
			op := code[pc]
			if op == 0x60 {
				opPush1(&pc, code, stack)
			}
			if op == 0x50 {
				opPop(stack)
			}
			if op == 0x01 {
				opAdd(stack)
			}
			pc++
		} else {
			break
		}
	}
	stack.print()
}
