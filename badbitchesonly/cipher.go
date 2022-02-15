package badbitchesonly

import (
	_ "fmt"
)

/* global variable declaration */
var r1 Register
var r2 Register
var r3 Register
var r4 Register

type Register struct {
	Number      int
	Length      int
	ArrImposter []int
	Tabs        []int
	Majs        []int
	Ært         int
}

func makeR1() Register {
	r1 = Register{Number: 1,
		Length:      19,
		ArrImposter: make([]int, 19),
		Tabs:        []int{18, 17, 16, 13},
		Majs:        []int{12, 15},
		Ært:         14}
	return r1
}

func makeR2() Register {
	r2 = Register{Number: 2,
		Length:      22,
		ArrImposter: make([]int, 22),
		Tabs:        []int{21, 20},
		Majs:        []int{9, 13},
		Ært:         16}
	return r2
}

func makeR3() Register {
	r3 = Register{Number: 3,
		Length:      23,
		ArrImposter: make([]int, 23),
		Tabs:        []int{22, 21, 20, 7},
		Majs:        []int{16, 18},
		Ært:         13}
	return r3
}

func makeR4() Register {
	r4 = Register{Number: 4,
		Length:      17,
		ArrImposter: make([]int, 17),
		Tabs:        []int{16, 11}}
	return r4
}

// Returns the majority bit of input x, y, z
func majority(x int, y int, z int) int {
	if x+y+z >= 2 {
		return 1
	} else {
		return 0
	}
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
	if maj == arr[3] {
		//clock R2
		print("clock R2")
	}
	if maj == arr[7] {
		//clock R3
		print("clock R3")
	}
	if maj == arr[10] {
		print("clock R1")
	}
}

//TODO handle output bit
//Move all the bits to the right, rightmost bit is dicarded!!!, input bit is specified by the taps of the register
func Clock(r Register) {
	arr := r.ArrImposter

	//calculate the new bit before shifting all the numbers, using the feedback function
	newbit := calculateNewBit(r)

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
	for i := range r.Tabs {
		newbit = ^arr[r.Tabs[i]]
	}
	return newbit
}

func initializeRegisters(session_key []int, frame_number []int) {
	/* do A5/2 */

	for i := 0; i < 63; i++ {
		/* Clock all registers */
		r1.ArrImposter[0] = ^session_key[i]
		r2.ArrImposter[0] = ^session_key[i]
		r3.ArrImposter[0] = ^session_key[i]
		r4.ArrImposter[0] = ^session_key[i]
	}

	for i := 0; i < 21; i++ {
		/* Clock all registers */
		r1.ArrImposter[0] = ^frame_number[i]
		r2.ArrImposter[0] = ^frame_number[i]
		r3.ArrImposter[0] = ^frame_number[i]
		r4.ArrImposter[0] = ^frame_number[i]
	}
}

/* Should we give frame number as a param ? */
func makeKeyStream() {
	// all registers contain 0s
	r1 = makeR1()
	r2 = makeR2()
	r3 = makeR3()
	r4 = makeR4()

	/* Initialiser Registers*/

	/* Initialize internal state with K_c and frame number */

	/* Force bits R1[15], R2[16], R3[18], R4[10] to be 1 */

	/* Run A5/2 for 99 clocks and ignore output */

	/* Run A5/2 for 228 clocks and use outputs as key-stream */

}

func main() {
	makeKeyStream()
}
