package main

type EVM struct {
	Memory *Memory
	Stack  *Stack
}

func newEVM() *EVM {
	return &EVM{
		Memory: &Memory{},
		Stack:  &Stack{},
	}
}
