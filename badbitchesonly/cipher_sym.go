package main

import (
	. "fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/* global variable declaration */
var sr1 SymRegister
var sr2 SymRegister
var sr3 SymRegister
var sr4 SymRegister

type SymRegister struct {
	Length      int
	ArrImposter [][]int
	Tabs        []int
	Majs        []int
	Ært         int
}

var sym_session_key [][]int

func SymMakeRegister(length int, tabs []int, major_idx []int, compliment_idx int) SymRegister {
	reg := SymRegister{
		Length:      length,
		ArrImposter: make([][]int, length),
		Tabs:        tabs,
		Majs:        major_idx,
		Ært:         compliment_idx}

	for i := 0; i < reg.Length; i++ {
		reg.ArrImposter[i] = make([]int, reg.Length)
	}

	return reg
}

func SymSetRegisters() {
	sr1 = SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14)
	sr2 = SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16)
	sr3 = SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13)
	sr4 = SymMakeRegister(17, []int{16, 11}, nil, -1)

}

//Clock R1, R2, R3 based on R4 state
func SymClockingUnit(r4 SymRegister) {
	//arr := r4.ArrImposter
	//maj := majority(arr[3], arr[7], arr[10])
	//if maj == arr[10] {

	// HARDCODING all registers is clocked every fucking time
	SymClock(sr1)
	print("clock R1\n")
	//}
	//if maj == arr[3] {
	//clock R2
	SymClock(sr2)
	print("clock R2\n")
	//}
	//if maj == arr[7] {
	//clock R3
	SymClock(sr3)
	print("clock R3\n")
	//}
}

func SymClock(r SymRegister) {
	arr := r.ArrImposter

	//calculate the new bit before shifting all the numbers, using the feedback function
	newbit := SymCalculateNewBit(r)
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

//Calculate the new int slice by xor'ing the tab-slices together row-wise
func SymCalculateNewBit(r SymRegister) []int {
	slice_slice := r.ArrImposter

	newbit := make([]int, r.Length) //all 0 slice for first xor

	for i := range r.Tabs {
		tabslice := slice_slice[r.Tabs[i]] //get the slice for the tap
		//Printf("slice %d is %+v \n",r.Tabs[i], tabslice)
		for i := 0; i < r.Length; i++ { //loop through the slices and xor them index-wise
			newbit[i] = newbit[i] ^ tabslice[i]
		}
	}
	return newbit
}

func SymMakeSessionKey() {
	rand.Seed(time.Now().Unix())

	key := make([]int, 64)
	for i := 0; i < 64; i++ {
		key[i] = rand.Intn(2)
	}
	sym_session_key = make([][]int, 4) // REVIEW: der mangler noget her
}

func SymInitialiseRegisters() {
	// Reset registers
	SymSetRegisters()

	for i := 0; i < 64; i++ {
		SymClock(sr1)
		SymClock(sr2)
		SymClock(sr3)
		SymClock(sr4)

		// REVIEW: nomalt xor med sessions key - skal dette stadig gøres?
		// REVIEW: we pretend that the session key is 0 #verySafe sorry Ivan
		// session_key[i] skal XORs her
	}

	// makes frame_number from int -> bits in array
	//frame_bits := makeFrameNumberToBits(frame_number)

	for i := 0; i < 22; i++ {
		SymClock(sr1)
		SymClock(sr2)
		SymClock(sr3)
		SymClock(sr4)

		// REVIEW: xor med framenumber
		// REVIEW: we pretend that the framenumber is 0
		// frame_bits[i] skal XORs her
	}
}

// func symMakeFinalXOR() []int { // REVIEW: Skal tilføjes til flowdiagram
// 	// register R1, majs = 12, 15, ært = 14

// 	//maj_r1 := majorityOutput(r1)
// 	//maj_r2 := majorityOutput(r2)
// 	//maj_r3 := majorityOutput(r3)

// 	//last_r1 := r1.ArrImposter[r1.Length-1]
// 	//last_r2 := r2.ArrImposter[r2.Length-1]
// 	//last_r3 := r3.ArrImposter[r3.Length-1]

// 	//finalXOR := maj_r1 ^ last_r1 ^ maj_r2 ^ last_r2 ^ maj_r3 ^ last_r3 // all is XOR'ed

// 	return finalXOR

// }

/*
Makes the final xor of r1[-1] ⨁ maj(r1) ⨁ r2[-1] ⨁ maj(r2) ⨁ r3[-1] ⨁ maj(r3)
returns [vars1 | prod1 | vars2 | prod2 | vars3 | prod3]
Calls SymMajorityOutput and OverwriteXorSlice
*/
func SymMakeFinalXOR(r1 SymRegister, r2 SymRegister, r3 SymRegister) []int {
	last_r1 := r1.ArrImposter[r1.Length-1]
	last_r2 := r2.ArrImposter[r2.Length-1]
	last_r3 := r3.ArrImposter[r3.Length-1]

	maj_r1 := SymMajorityOutput(r1)
	maj_r2 := SymMajorityOutput(r2)
	maj_r3 := SymMajorityOutput(r3)

	//Xor them "locally" together first
	OverwriteXorSlice(last_r1, maj_r1) //[vars1 | prods1] = [vars1] ⨁ [vars1 | prods1]
	OverwriteXorSlice(last_r2, maj_r2)
	OverwriteXorSlice(last_r3, maj_r3)

	start := make([]int, len(maj_r1))
	copy(start, maj_r1)              //start by res = [vars1 | prod1]
	start = append(start, maj_r2...) //now [vars1 | prod1 | vars2 | prod2 ]
	start = append(start, maj_r3...) //now [vars1 | prod1 | vars2 | prod2 | vars3 | prod3]

	return start
}

/*
Symbolic majority function. Calls SymMajorityMultiply and XorSlice.
Performs xy ⨁ xz ⨁ yz ⨁ x ⨁ y on the majority tabs of the register.
Returns slice of lengt len(r)+ (len(r)*(len(r)-1))/2 with the original variables in the first len(r) entries and products in the rest
*/
func SymMajorityOutput(r SymRegister) []int {
	arr := r.ArrImposter
	x := arr[r.Majs[0]]
	y := arr[r.Majs[1]]
	z := arr[r.Ært]
	// xy ⨁ xz ⨁ yz ⨁ x ⨁ y
	xy := SymMajorityMultiply(x, y) // [vars | products ] [x1, x2, x3, ..., x12, x23]
	xz := SymMajorityMultiply(x, z)
	yz := SymMajorityMultiply(y, z)
	ee := XorSlice(xy, xz)
	long_slice := XorSlice(ee, yz) //This is a xor of normal and product variables
	short_slice := XorSlice(x, y)  //This is only a xor of the normal variables
	// xor the "normal" variables in the start of the long slice [ vars | products ] ⨁ [ vars ]
	for i := 0; i < len(short_slice); i++ {
		long_slice[i] = long_slice[i] ^ short_slice[i]
	}
	return long_slice
}

//Takes two slices and xors them indexwise together. Assumed to be of same lenght. Returns slice of size len(a)
func XorSlice(a []int, b []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] ^ b[i]
	}
	return res
}

//Takes two slices with first shorter than the second. Overwrites the first len(short) entries in long with long[i] = short[i] ^ long[i]
func OverwriteXorSlice(short []int, long []int) {
	for i := 0; i < len(short); i++ {
		long[i] = short[i] ^ long[i]
	}
}

/*
multiplies two decision vectors with result being c[i]d[j] ^ c[j]d[i] for i /= j and result = c[i]d[j] for i=j.
res slice has lenght len(c)*(len(c)-1)/2 + len(c).  c and d are assumed to be same lenght.
The original len(c) variables will be in the first len(c) indexes of result
*/
func SymMajorityMultiply(c []int, d []int) []int {
	lenc := len(c)
	leng := lenc * (lenc - 1) / 2
	res := make([]int, leng+lenc)
	acc := 0
	for i := 0; i < lenc; i++ {
		res[i] = c[i] * d[i]
		for j := i + 1; j < lenc; j++ {
			res[lenc+acc] = c[i]*d[j] ^ c[j]*d[i]
			//Printf("res[%d] = %d*%d ^ %d*%d = %d \n", lenc+acc, c[i], d[j], c[j], d[i], res[lenc+acc])
			acc++
		}
	}

	return res
}

/*
###########################################################
#### THIS IS WHERE THE SIMPLE CIPHER SYM STREAM EXISTS ####
###########################################################
*/

func InitOneSymRegister() SymRegister {
	reg := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14) // equvalent to reg1
	for i := 0; i < 19; i++ {
		// reg.ArrImposter[i] = make([]int, 19)
		reg.ArrImposter[i][i] = 1 // each entry in the diagonal set to 1 as x_i is only dependent on x_i when initialized
	}
	return reg
}

// NOT USED AS WE DO NOT CARE ABOUT THE PLAINTEXT
// func EncryptSimpleSymPlaintext(plaintext []int) [][]int {
// 	key := MakeSymPlaintext()
// 	Printf("This is the key-stream: %d \n", key)
// 	res := make([][]int, len(plaintext))
// 	for i := range res {
// 		res[i] = make([]int, 19)
// 		for j := i; i < 19; i++ {
// 			res[i][j] = key[i][i]
// 		}
// 	}
// 	return res
// }

func SimpleKeyStreamSym(r SymRegister) [][]int {

	// Init key stream array
	keyStream := make([][]int, 228)
	for i := 0; i < 228; i++ {
		keyStream[i] = make([]int, r.Length)
	}

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		SymClock(r)
	}

	// clock 228 times and save keystream
	for i := 0; i < 228; i++ {
		SymClock(r)
		keyStream[i] = r.ArrImposter[r.Length-1]
	}
	return keyStream
}

func SimpleKeyStreamSymSecondVersion(r SymRegister) [][]int {

	// Init key stream array
	keyStream := make([][]int, 228)
	for i := 0; i < 228; i++ {
		keyStream[i] = make([]int, r.Length)
	}

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		SymClock(r)
	}

	// clock 228 times and save keystream
	for i := 0; i < 228; i++ {
		SymClock(r)
		keyStream[i] = SymMajorityOutput(r)
		//Printf("Length of output from symMajorFunc %d\n", len(SymMajorityOutput(r)))
	}
	return keyStream
}

/*
func symMakeKeyStream() [][]int {

	symSetRegisters()

	frame_number++

	keyStream := make([][]int, 228)

	initialiseRgisters()

	// Run A5/2 for 99 clocks and ignore output
	for i := 0; i < 99; i++ {
		// do the clock thingy and ignore
		clockingUnit(r4)
		Clock(r4)
	}

	// Run A5/2 for 228 clocks and use outputs as key-stream
	for i := 0; i < 228; i++ {
		// do the clock thingy and output
		clockingUnit(r4)
		Clock(r4)
		keyStream[i] = makeFinalXOR()
	}
	return keyStream

}

*/

func PrettyPrint(r SymRegister) {
	rLength := r.Length
	rMatrix := r.ArrImposter

	for i := 0; i < rLength; i++ {
		accString := "x" + strconv.Itoa(i) + " = "
		for j := 0; j < rLength; j++ {
			if rMatrix[i][j] == 1 {
				str := strconv.Itoa(j)
				accString += " x" + (str) + " ⨁ "
			}
		}
		accString = strings.TrimRight(accString, "⨁ ")
		Printf("")
		println(accString)
	}
}
