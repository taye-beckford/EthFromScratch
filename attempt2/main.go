package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/holiman/uint256"
)

type Stack struct {
	data []uint256.Int
}

func (st *Stack) push(value uint256.Int) {
	st.data = append(st.data, value)
}

func (st *Stack) pop() uint256.Int {
	if len(st.data) == 0 {
		log.Fatal("The Stack is empty cannot pop")
	}
	end := len(st.data) - 1
	lastElement := st.data[end]
	st.data = st.data[:end]
	return lastElement
}

func opPush1() {

}

func main() {
	// code, err := hex.DecodeString("60FF")
	// if err != nil {
	// 	fmt.Println("Failed to decode hex string")
	// }

	// for i, op := range code {
	// 	if op == 0x60 {
	// 		opPush1()
	// 	}
	// }

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
