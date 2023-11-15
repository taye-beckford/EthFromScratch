package main

import (
	"fmt"
	"log"

	"github.com/holiman/uint256"
)

type Stack struct {
	data []uint256.Int
}

func newStack() *Stack {
	return &Stack{}
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

func (st *Stack) print() {
	fmt.Println(st.data)
}
