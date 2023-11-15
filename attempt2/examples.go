package main

import (
	"encoding/hex"
	"fmt"

	"github.com/holiman/uint256"
)

func basicStack() {
	stack := Stack{}
	bytes, _ := hex.DecodeString("606060606060606060")
	var integer = new(uint256.Int)
	integer.SetBytes(bytes)

	stack.push(*integer)
	fmt.Println(stack.data)
	stack.push(*uint256.NewInt(100))
	fmt.Println(stack.data)
	stack.pop()
	fmt.Println(stack.data)
}
