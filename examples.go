package main

func basicStack() {
	evm := newEVM()
	evm.opPush1(10)
	evm.opPush1(20)
	evm.opAdd() // 30
	evm.opPush1(5)
	evm.makeSwap(1)
	evm.opSub() // 25
	evm.Stack.print()
}

func basicMemory() {
	evm := newEVM()
	evm.Memory.Resize(64)
	evm.opPush1(200)
	evm.opPush1(0)
	evm.opMstore()
	evm.Memory.Print()
}
