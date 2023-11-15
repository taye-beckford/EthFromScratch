package main

type OpCode byte

const (
	PUSH1 OpCode = 0x60
	POP   OpCode = 0x50
	ADD   OpCode = 0x01
)
