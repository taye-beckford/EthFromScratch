package main

import (
	"fmt"
	"log"

	"github.com/holiman/uint256"
)

type Stack struct {
	data []uint256.Int
}

func (st *Stack) push(value uint256.Int) {
	if len(st.data) == 1024 {
		log.Fatal("Exceeded max stack size")
	}
	st.data = append(st.data, value)
}

func (st *Stack) pop() uint256.Int {
	if len(st.data) == 0 {
		log.Fatal("The stack is empty")
	}
	end := len(st.data) - 1
	lastElement := (st.data)[end]
	st.data = (st.data)[:end]
	return lastElement
}

func (st *Stack) peek() *uint256.Int {
	return &st.data[len(st.data)-1]
}

func (st *Stack) print() {
	for i, v := range st.data {
		fmt.Println(i, v)
	}
}

func (st *Stack) len() int {
	return len(st.data)
}

func (st *Stack) swap(n int) {
	st.data[st.len()-n], st.data[st.len()-1] = st.data[st.len()-1], st.data[st.len()-n]
}
