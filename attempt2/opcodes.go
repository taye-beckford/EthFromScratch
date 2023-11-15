package main

type OpCode byte

const (
	PUSH1 OpCode = 0x60
	POP   OpCode = 0x50
	ADD   OpCode = 0x01
	MUL   OpCode = 0x02
	SUB   OpCode = 0x03
	DIV   OpCode = 0x04
	MOD   OpCode = 0x06
)
