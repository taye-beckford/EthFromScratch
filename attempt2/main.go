package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	stack := run("600a600a01")
	stack.print()
}

func run(byteCode string) *Stack {
	code, err := hex.DecodeString(byteCode)
	if err != nil {
		fmt.Println("Failed to decode hex string")
	}
	var (
		stack = newStack()
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
	return stack
}
