package main

import (
	"encoding/hex"
	"fmt"
)

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
