package main

import "encoding/hex"

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}
