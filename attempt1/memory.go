package main

import (
	"fmt"

	"github.com/holiman/uint256"
)

type Memory struct {
	store []byte
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Len() int {
	return len(m.store)
}

func (m *Memory) Data() []byte {
	return m.store
}

// Set32 sets the 32 bytes starting at offset to the value of val, left-padded with zeroes to
// 32 bytes.
func (m *Memory) Set32(offset uint64, val *uint256.Int) {
	// length of store may never be less than offset + size.
	// The store should be resized PRIOR to setting the memory
	if offset+32 > uint64(len(m.store)) {
		panic("invalid memory: store empty")
	}
	// Fill in relevant bits
	b32 := val.Bytes32()
	copy(m.store[offset:], b32[:])
}

func (m *Memory) Resize(size uint64) {
	if uint64(m.Len()) < size {
		m.store = append(m.store, make([]byte, size-uint64(m.Len()))...)
	}
}

func (m *Memory) Print() {
	lineWidth := 32 // Number of bytes per line

	for i, b := range m.store {
		fmt.Printf("%02x", b)
		if (i+1)%lineWidth == 0 {
			fmt.Println()
		}
	}

	// Check if we need an extra newline at the end
	if len(m.store)%lineWidth != 0 {
		fmt.Println()
	}
}
