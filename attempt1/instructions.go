package main

import "github.com/holiman/uint256"

func (e *EVM) opPush1(value byte) {
	integer := new(uint256.Int)
	e.Stack.push(*integer.SetUint64(uint64(value)))
}

func (e *EVM) opPop() {
	e.Stack.pop()
}

func (e *EVM) opAdd() {
	x, y := e.Stack.pop(), e.Stack.peek()
	y.Add(&x, y)
}

func (e *EVM) opSub() {
	x, y := e.Stack.pop(), e.Stack.peek()
	y.Sub(&x, y)
}

func (e *EVM) makeSwap(size int) {
	e.Stack.swap(size + 1)
}

func (e *EVM) opMstore() {
	mStart, val := e.Stack.pop(), e.Stack.pop()
	e.Memory.Set32(mStart.Uint64(), &val)
}

func (e *EVM) opMstore8() {
	off, val := e.Stack.pop(), e.Stack.pop()
	e.Memory.store[off.Uint64()] = byte(val.Uint64())
}
