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

func opPop(pc *uint64, code []byte, st *Stack) {
	st.pop()
}

func opAdd(pc *uint64, code []byte, st *Stack) {
	x := st.pop()
	y := st.pop()
	st.push(*y.Add(&x, &y))
}

type OpCode byte

type JumpTable [256]*operation

type executionFunc func(pc *uint64, code []byte, stack *Stack)

type operation struct {
	execute executionFunc
}

func newJumpTable() JumpTable {
	tbl := JumpTable{
		0x60: {
			opPush1,
		},
		0x50: {
			opPop,
		},
		0x01: {
			opAdd,
		},
	}
	return tbl
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
		jt    = newJumpTable()
	)

	for {
		if pc < uint64(len(code)) {
			op := code[pc]
			jt[op].execute(&pc, code, stack)
			pc++
		} else {
			break
		}
	}
	stack.print()
}
