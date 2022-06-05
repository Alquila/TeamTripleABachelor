package main

import (
	. "fmt"
)

// InitOneRegister
// initialises a single Register, specifically R1.
func InitOneRegister() Register {
	reg := MakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14)
	reg.RegSlice = []int{1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0}
	reg.RegSlice[15] = 1
	return reg
}

// SimpleKeyStream
// this function is used for testing. Creates a keystream from a single Register.
func SimpleKeyStream(r Register) []int {
	keyStream := make([]int, 228)

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		Clock(r)
	}

	// Clock the register 228 times and save the last bit in keystream
	for i := 0; i < 228; i++ {
		Clock(r)
		keyStream[i] = r.RegSlice[r.Length-1]
	}

	return keyStream
}

// SimpleKeyStreamWithMajorityFunc
// this function is used for testing.
// Creates a keystream from a single Register using the Majority function.
func SimpleKeyStreamWithMajorityFunc(r Register) []int {
	keyStream := make([]int, 228)

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		Clock(r)
	}

	// Clock the register 228 times and save the last bit XOR'ed with the Majority output in keystream
	for i := 0; i < 228; i++ {
		Clock(r)
		keyStream[i] = MajorityOutput(r) ^ r.RegSlice[r.Length-1]
	}

	return keyStream
}

// MakePlaintext
// This function used to be used for testing.
func MakePlaintext() []int {
	plaintext := []int{1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0}
	return plaintext
}

// EncryptSimplePlaintext
// used to be used for testing.
// Runs streamcipher, the one with actual numbers, such that the plaintext is encrypted
func EncryptSimplePlaintext(plaintext []int) []int {
	key := MakeSimpleKeyStream()
	Printf("This is the key-stream: %d \n", key)
	res := make([]int, len(plaintext))
	for i := range res {
		res[i] = key[i] ^ plaintext[i]
	}

	return res
}

// MakeSimpleKeyStream
// used for testing. Returns a simple keystream using only Register R1.
func MakeSimpleKeyStream() []int {
	// init R1
	r1 = MakeR1()
	r1.RegSlice[15] = 1

	keyStream := make([]int, 50)

	for i := 0; i < 101; i++ {
		Clock(r1)
		if i > 50 {
			keyStream[i-51] = r1.RegSlice[18]
		}
	}

	return keyStream

}
