package main

type JumpTable [256]*operation

type executionFunc func(pc *uint64, code []byte, stack *Stack)

type operation struct {
	execute executionFunc
}

func newJumpTable() JumpTable {
	tbl := JumpTable{
		PUSH1: {
			opPush1,
		},
		POP: {
			opPop,
		},
		ADD: {
			opAdd,
		},
	}
	return tbl
}
