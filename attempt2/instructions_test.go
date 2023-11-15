package main

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestAdd(t *testing.T) {
	testCode := "600a600a01" // PUSH1 10; PUSH 10; ADD;
	stack := run(testCode)
	expected := uint256.NewInt(20)
	actual := stack.pop()
	if actual.Cmp(expected) != 0 {
		t.Errorf("Expected %x, got %x", expected, actual)
	}
}

func BenchmarkAdd(b *testing.B) {
	stack := newStack()
	pc := uint64(0)
	stack.push(*uint256.NewInt(100))
	stack.push(*uint256.NewInt(100))
	opAdd(&pc, []byte{}, stack)
	stack.pop()
}
