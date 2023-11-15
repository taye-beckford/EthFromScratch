package main

import "github.com/holiman/uint256"

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

func opMul(pc *uint64, code []byte, st *Stack) {
	x := st.pop()
	y := st.pop()
	st.push(*y.Mul(&x, &y))
}

func opSub(pc *uint64, code []byte, st *Stack) {
	x := st.pop()
	y := st.pop()
	st.push(*y.Sub(&x, &y))
}

func opDiv(pc *uint64, code []byte, st *Stack) {
	x := st.pop()
	y := st.pop()
	st.push(*y.Div(&x, &y))
}
