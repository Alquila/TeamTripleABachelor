package main

import (
	. "fmt"
	"math/rand"
	"time"
)

/* global variable declaration */
var r1 Register
var r2 Register
var r3 Register
var r4 Register

//sat in makeKeyStream() after init and setIndicesToOne()
var r4_after_init Register

//Lives in cipher.go Has to be manually updated
var current_frame_number int
var session_key []int

type Register struct {
	Length      int
	ArrImposter []int
	Tabs        []int
	Majs        []int
	Ært         int
}

func makeRegister(length int, tabs []int, major_idx []int, compliment_idx int) Register {
	reg := Register{
		Length:      length,
		ArrImposter: make([]int, length),
		Tabs:        tabs,
		Majs:        major_idx,
		Ært:         compliment_idx}
	return reg
}

func makeR1() Register {
	r1 = makeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14)
	return r1
}

func makeR2() Register {
	r2 = makeRegister(22, []int{21, 20}, []int{9, 13}, 16)
	return r2
}

func makeR3() Register {
	r3 = makeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13)
	return r3
}

func makeR4() Register {
	r4 := makeRegister(17, []int{16, 11}, nil, -1)
	return r4
}

//Make all the global registers
func makeRegisters() {
	r1 = makeR1()
	r2 = makeR2()
	r3 = makeR3()
	r4 = makeR4()
}

func prettyPrint(r Register) {
	Printf("%+v", r.ArrImposter)
	print("\n")
}

func printAll() {
	Printf("R1: %+v \n", r1.ArrImposter)
	Printf("R2: %+v \n", r2.ArrImposter)
	Printf("R3: %+v \n", r3.ArrImposter)
	Printf("R4: %+v \n", r4.ArrImposter)
}

// Returns the majority bit of input x, y, z
func majority(x int, y int, z int) int {
	return x*y ^ y*z ^ z*x
}

//Calls majority function for R1, R2, R3 with one inversed. Don't call on R4, it will crash
func majorityOutput(r Register) int {
	arr := r.ArrImposter

	x := arr[r.Majs[0]]
	y := arr[r.Majs[1]]
	z := arr[r.Ært] ^ 1

	return majority(x, y, z)
}

//Clock R1, R2, R3 based on R4 state
func clockingUnit(r4 Register) {
	arr := r4.ArrImposter
	maj := majority(arr[3], arr[7], arr[10])
	if maj == arr[10] {
		Clock(r1)
		// print("clock R1\n")
	}
	if maj == arr[3] {
		//clock R2
		Clock(r2)
		// print("clock R2\n")
	}
	if maj == arr[7] {
		//clock R3
		Clock(r3)
		// print("clock R3\n")
	}
}

//Move all the bits to the right, rightmost bit is discarded!!!, input bit is specified by the taps of the register
func Clock(r Register) {
	arr := r.ArrImposter

	//calculate the new bit before shifting all the numbers, using the feedback function
	newbit := calculateNewBit(r)
	//print(newbit)
	//call majorityOutput?

	//shift all the numbers to the right, start at the end, copy from index before it
	//save arr[r.Length-1] ?
	for i := r.Length - 1; i > 0; i-- { //stops after arr[1] = arr[0]
		arr[i] = arr[i-1]
	}
	//set arr[0] to the new bit
	arr[0] = newbit

}

//calculates the newbit by taking xor of all the tab indexes
func calculateNewBit(r Register) int {
	arr := r.ArrImposter

	//calculate the new bit before shifting all the numbers
	newbit := 0 //setting it to 0 should just copy the first bit when doing first xor
	//print(newbit)
	//print("\n")
	for i := range r.Tabs {
		newbit = newbit ^ arr[r.Tabs[i]]
		//Printf("tab bit %d \n",arr[r.Tabs[i]])
		//Printf("new bit %d \n",newbit)
	}
	return newbit
}

func MakeFrameNumberToBits(number int) []int {
	// frame number will always be 22 bits
	frame_bit := make([]int, 22)

	for i := 0; i < 22; i++ {
		frame_bit[i] = (number >> i) & 1 // index 0 becomes least significant bit
	}

	return frame_bit
}

//makes "random" bit array
func makeSessionKey() {
	rand.Seed(time.Now().Unix())

	key := make([]int, 64)
	for i := 0; i < 64; i++ {
		key[i] = rand.Intn(2)
	}
	session_key = key
}

//makes 0's arrays, for 64 cycles clock registers and xor with i'th key bit, for 22 cycles clock registers and xor with i'th frame bit
func initializeRegisters() { // used to have session_key and frame_number as params, but made then global variables instead
	/* do A5/2 */
	r1.ArrImposter = make([]int, r1.Length)
	r2.ArrImposter = make([]int, r2.Length)
	r3.ArrImposter = make([]int, r3.Length)
	r4.ArrImposter = make([]int, r4.Length)

	for i := 0; i < 64; i++ {
		Clock(r1)
		Clock(r2)
		Clock(r3)
		Clock(r4)

		// print("printing r1 \n")
		// prettyPrint(r1)
		// Printf("sk %d \n", session_key[i])

		r1.ArrImposter[0] = r1.ArrImposter[0] ^ session_key[i]
		r2.ArrImposter[0] = r2.ArrImposter[0] ^ session_key[i]
		r3.ArrImposter[0] = r3.ArrImposter[0] ^ session_key[i]
		r4.ArrImposter[0] = r4.ArrImposter[0] ^ session_key[i]
	}

	// prints(r1.ArrImposter, "r1")
	// prints(r2.ArrImposter, "r2")
	// prints(r3.ArrImposter, "r3")
	// prints(r4.ArrImposter, "r4")

	// makes frame_number from int -> bits in array
	frame_bits := MakeFrameNumberToBits(current_frame_number)
	//prints(r4.ArrImposter, "r4 before clocking with frame ")
	for i := 0; i < 22; i++ {
		Clock(r1)
		Clock(r2)
		Clock(r3)
		Clock(r4)

		r1.ArrImposter[0] = r1.ArrImposter[0] ^ frame_bits[i]
		r2.ArrImposter[0] = r2.ArrImposter[0] ^ frame_bits[i]
		r3.ArrImposter[0] = r3.ArrImposter[0] ^ frame_bits[i]
		r4.ArrImposter[0] = r4.ArrImposter[0] ^ frame_bits[i]
	}

	//Printf("Initialized registers 1-4:")
	// Println()
	//prettyPrint(r1)
	//prettyPrint(r2)
	//prettyPrint(r3)
	//prettyPrint(r4)
}

func setIndicesToOne() {
	r1.ArrImposter[15] = 1
	r2.ArrImposter[16] = 1
	r3.ArrImposter[18] = 1
	r4.ArrImposter[10] = 1
}

//computes the XOR of the last bits of the three registers and the results of calling majorityOutput on them.
func makeFinalXOR() int { // REVIEW: Skal tilføjes til flowdiagram
	// register R1, majs = 12, 15, ært = 14
	maj_r1 := majorityOutput(r1)
	maj_r2 := majorityOutput(r2)
	maj_r3 := majorityOutput(r3)

	last_r1 := r1.ArrImposter[r1.Length-1]
	last_r2 := r2.ArrImposter[r2.Length-1]
	last_r3 := r3.ArrImposter[r3.Length-1]

	finalXOR := maj_r1 ^ last_r1 ^ maj_r2 ^ last_r2 ^ maj_r3 ^ last_r3 // all is XOR'ed

	return finalXOR

}

/* Should we give frame number as a param ? */

/*Initializes the registers,
  sets indices to one's,
  runs for 99 clocks,
  then runs for 228 clocks and returns the key stream */
func makeKeyStream() []int {

	// all registers contain 0s
	makeRegisters()

	keyStream := make([]int, 228)

	/* Initialize internal state with K_c and frame number */
	initializeRegisters()

	/* Force bits R1[15], R2[16], R3[18], R4[10] to be 1 */
	setIndicesToOne()
	r4_after_init = makeR4()
	copy(r4_after_init.ArrImposter, r4.ArrImposter)
	prints(r1.ArrImposter, "R1 after clocking with frame")
	prints(r2.ArrImposter, "R2 after clocking with frame")
	prints(r3.ArrImposter, "R3 after clocking with frame")
	prints(r4.ArrImposter, "r4 after clocking with frame")
	// Print()

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		// do the clock thingy and ignore
		clockingUnit(r4)
		Clock(r4)
	}

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		// do the clock thingy and output
		clockingUnit(r4)
		Clock(r4)
		keyStream[i] = makeFinalXOR()
	}
	return keyStream
}

/*
#######################################################
#### THIS IS WHERE THE SIMPLE CIPHER STREAM EXISTS ####
#######################################################
*/

func InitOneRegister() Register {
	reg := makeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14)
	reg.ArrImposter = []int{1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0}
	reg.ArrImposter[15] = 1
	return reg
}

func SimpleKeyStream(r Register) []int {
	keyStream := make([]int, 228)

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		Clock(r)
	}
	// Printf("Print register efter 99 clocks: \n %d \n", r.ArrImposter)

	for i := 0; i < 228; i++ {
		Clock(r)
		keyStream[i] = r.ArrImposter[r.Length-1]
		if i == 200 {
			// Printf(" 200'th bit is %d \n", r.ArrImposter[r.Length-1])
		}
	}
	// Printf("Print register efter endnu 228 clocks: \n %d \n", r.ArrImposter)

	return keyStream
}

func SimpleKeyStreamSecondVersion(r Register) []int {
	keyStream := make([]int, 228)

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		Clock(r)
	}
	//fmt.Printf("Print register efter 99 clocks: \n %d \n", r.ArrImposter)

	for i := 0; i < 228; i++ {
		Clock(r)
		keyStream[i] = majorityOutput(r) ^ r.ArrImposter[r.Length-1]
	}
	//fmt.Printf("Print register efter endnu 228 clocks: \n %d \n", r.ArrImposter)

	return keyStream
}

// vi bruger kun register R1 som er 19 langt
func MakePlaintext() []int {
	plaintext := []int{1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0}
	return plaintext
}

// run stream-cipher, the one with actual numbers, such that the plaintext is encrypted
func EncryptSimplePlaintext(plaintext []int) []int {
	key := MakeSimpleKeyStream()
	Printf("This is the key-stream: %d \n", key)
	res := make([]int, len(plaintext))
	for i := range res {
		res[i] = key[i] ^ plaintext[i]
	}

	return res
}

func MakeSimpleKeyStream() []int {
	// init R1
	r1 = makeR1()
	r1.ArrImposter[15] = 1

	keyStream := make([]int, 50)

	for i := 0; i < 101; i++ {
		// clock R1
		Clock(r1)
		if i > 50 {
			keyStream[i-51] = r1.ArrImposter[18]
		}
	}

	return keyStream

}
