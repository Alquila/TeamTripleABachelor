package main

import (
	"math/rand"
	"time"
)

/* global variable declaration */
var r1 Register
var r2 Register
var r3 Register
var r4 Register

// Sat in MakeKeyStream() after init and SetIndicesToOne()
var r4_after_init Register

// Lives in cipher.go - Has to be manually updated
var currentFrameNumber int
var sessionKey []int

type Register struct {
	Length   int   // Length of the Register
	RegSlice []int // RegSlice consists of the registers content
	Taps     []int // Taps stores the indices used for the Feedback Function
	MajsTaps []int // MajsTaps stores the indices used for the Majority Function
	NegTap   int   // NegTap stores the index for the tap having to be flip before used in the Majority Function
}

func MakeRegister(length int, taps []int, major_idx []int, compliment_idx int) Register {
	reg := Register{
		Length:   length,
		RegSlice: make([]int, length),
		Taps:     taps,
		MajsTaps: major_idx,
		NegTap:   compliment_idx}
	return reg
}

func MakeR1() Register {
	r1 = MakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14)
	return r1
}

func MakeR2() Register {
	r2 = MakeRegister(22, []int{21, 20}, []int{9, 13}, 16)
	return r2
}

func MakeR3() Register {
	r3 = MakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13)
	return r3
}

func MakeR4() Register {
	r4 := MakeRegister(17, []int{16, 11}, nil, -1)
	return r4
}

// MakeRegisters initialises all 4 global registers with their correct length and so on
func MakeRegisters() {
	r1 = MakeR1()
	r2 = MakeR2()
	r3 = MakeR3()
	r4 = MakeR4()
}

// Majority
// returns the Majority bit of input x, y, z
func Majority(x int, y int, z int) int {
	return x*y ^ y*z ^ z*x
}

// MajorityOutput
// calls Majority function for R1, R2, R3 with one index negated.
// Important! Don't call on R4, it will crash :)
func MajorityOutput(r Register) int {
	arr := r.RegSlice

	x := arr[r.MajsTaps[0]]
	y := arr[r.MajsTaps[1]]
	z := arr[r.NegTap] ^ 1

	return Majority(x, y, z)
}

// Clock
// Moves all the bits to the right.
// Rightmost bit is discarded!
// Input bit is specified by the taps of the register
func Clock(r Register) {
	arr := r.RegSlice

	// calculate the new bit before shifting all the numbers, using the feedback function
	newbit := FeedbackFunction(r)

	// shift all the numbers to the right, start at the end, copy from index before it
	for i := r.Length - 1; i > 0; i-- { //stops after arr[1] = arr[0]
		arr[i] = arr[i-1]
	}

	//set arr[0] to the new bit
	arr[0] = newbit
}

// ClockingUnit
// Clock R1, R2, R3 based on R4 state
func ClockingUnit(r4 Register) {
	arr := r4.RegSlice
	maj := Majority(arr[3], arr[7], arr[10])

	if maj == arr[10] {
		Clock(r1)
		// print("clock R1\n")
	}
	if maj == arr[3] {
		Clock(r2)
		// print("clock R2\n")
	}
	if maj == arr[7] {
		Clock(r3)
		// print("clock R3\n")
	}
}

// FeedbackFunction
// calculates the newbit by XOR'ing Register.Taps
func FeedbackFunction(r Register) int {
	arr := r.RegSlice

	// calculate the new bit before shifting all the numbers
	newbit := 0 // setting it to 0 should just copy the first bit when doing first xor

	for i := range r.Taps {
		newbit = newbit ^ arr[r.Taps[i]]
	}

	return newbit
}

// MakeFrameNumberToBits
// translates the given frame number to the equivalent bit string return as a slice
func MakeFrameNumberToBits(number int) []int {
	// frame number will always be 22 bits
	frame_bit := make([]int, 22)

	for i := 0; i < 22; i++ {
		frame_bit[i] = (number >> i) & 1 // index 0 becomes least significant bit
	}

	return frame_bit
}

// MakeSessionKey
// makes "random" bit slice of size 64
func MakeSessionKey() {
	rand.Seed(time.Now().Unix())

	key := make([]int, 64)
	for i := 0; i < 64; i++ {
		key[i] = rand.Intn(2)
	}
	sessionKey = key
}

// InitializeRegisters
// makes 0's arrays, for 64 cycles clock registers and XOR with i'th key bit, for 22 cycles clock registers and XOR with i'th frame bit
func InitializeRegisters() {
	/* do A5/2 */
	r1.RegSlice = make([]int, r1.Length)
	r2.RegSlice = make([]int, r2.Length)
	r3.RegSlice = make([]int, r3.Length)
	r4.RegSlice = make([]int, r4.Length)

	// Clock all registers 64 times and XOR with the session key
	for i := 0; i < 64; i++ {
		Clock(r1)
		Clock(r2)
		Clock(r3)
		Clock(r4)

		r1.RegSlice[0] = r1.RegSlice[0] ^ sessionKey[i]
		r2.RegSlice[0] = r2.RegSlice[0] ^ sessionKey[i]
		r3.RegSlice[0] = r3.RegSlice[0] ^ sessionKey[i]
		r4.RegSlice[0] = r4.RegSlice[0] ^ sessionKey[i]
	}

	// makes frame_number from int -> bits in slice
	frame_bits := MakeFrameNumberToBits(currentFrameNumber)

	// Clock all registers 22 times and XOR with frame number
	for i := 0; i < 22; i++ {
		Clock(r1)
		Clock(r2)
		Clock(r3)
		Clock(r4)

		r1.RegSlice[0] = r1.RegSlice[0] ^ frame_bits[i]
		r2.RegSlice[0] = r2.RegSlice[0] ^ frame_bits[i]
		r3.RegSlice[0] = r3.RegSlice[0] ^ frame_bits[i]
		r4.RegSlice[0] = r4.RegSlice[0] ^ frame_bits[i]
	}

}

// SetIndicesToOne
// Forces a specific bit in each register to be 1
func SetIndicesToOne() {
	r1.RegSlice[15] = 1
	r2.RegSlice[16] = 1
	r3.RegSlice[18] = 1
	r4.RegSlice[10] = 1
}

// MakeFinalXOR
// computes the XOR of the last bits of the three registers and the results of calling MajorityOutput on each of them.
func MakeFinalXOR() int {
	maj_r1 := MajorityOutput(r1)
	maj_r2 := MajorityOutput(r2)
	maj_r3 := MajorityOutput(r3)

	last_r1 := r1.RegSlice[r1.Length-1]
	last_r2 := r2.RegSlice[r2.Length-1]
	last_r3 := r3.RegSlice[r3.Length-1]

	finalXOR := maj_r1 ^ last_r1 ^ maj_r2 ^ last_r2 ^ maj_r3 ^ last_r3 // all is XOR'ed

	return finalXOR

}

// MakeKeyStream
// initializes the registers, sets indices to one's,
// runs for 99 clocks, then runs for 228 clocks and returns the key stream
func MakeKeyStream() []int {

	// all registers Contains 0s
	MakeRegisters()

	keyStream := make([]int, 228)

	// Initialize internal state with K_c and frame number
	InitializeRegisters()

	// Force bits R1[15], R2[16], R3[18], R4[10] to be 1
	SetIndicesToOne()

	r4_after_init = MakeR4()
	copy(r4_after_init.RegSlice, r4.RegSlice)
	Prints(r1.RegSlice, "R1 after clocking with frame")
	Prints(r2.RegSlice, "R2 after clocking with frame")
	Prints(r3.RegSlice, "R3 after clocking with frame")
	Prints(r4.RegSlice, "r4 after clocking with frame")

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		ClockingUnit(r4)
		Clock(r4)
	}

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		ClockingUnit(r4)
		Clock(r4)
		keyStream[i] = MakeFinalXOR()
	}
	return keyStream
}
