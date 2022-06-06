package main

var originalFrameNumber int

/* global variable declaration */
var sr1 SymRegister
var sr2 SymRegister
var sr3 SymRegister
var sr4 Register

type SymRegister struct {
	Length   int
	RegSlice [][]int
	Taps     []int
	MajsTaps []int
	NegTap   int
	SetToOne int
}

func SymMakeRegister(length int, tabs []int, major_idx []int, compliment_idx int, bit_idx int) SymRegister {
	reg := SymRegister{
		Length:   length,
		RegSlice: make([][]int, length),
		Taps:     tabs,
		MajsTaps: major_idx,
		NegTap:   compliment_idx,
		SetToOne: bit_idx}

	for i := 0; i < reg.Length; i++ {
		reg.RegSlice[i] = make([]int, reg.Length)
	}

	return reg
}

// BitEntry
// initialises a Register's RegSlice
func BitEntry(reg SymRegister) {
	bit_entry := reg.SetToOne
	for i := 0; i < reg.Length; i++ {
		if i < bit_entry {
			reg.RegSlice[i][i] = 1
		}
		if i == bit_entry {
			reg.RegSlice[i][reg.Length-1] = 1
		}
		if i > bit_entry {
			reg.RegSlice[i][i-1] = 1
		}
	}
}

// SymSetRegisters
// calls SymMakeRegister on each register.
// Each register is initialised with array initialised for every entry, but no values inserted.
// Copies r4.RegSlice into sr4
func SymSetRegisters() {
	sr1 = SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sr2 = SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sr3 = SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)
	sr4 = MakeRegister(17, []int{16, 11}, nil, -1)
	copy(sr4.RegSlice, r4.RegSlice)
}

// SymClock
// clock a single SymRegister by calculating the new bit and shifting all bits once to the right.
func SymClock(r SymRegister) {
	arr := r.RegSlice

	//calculate the new bit before shifting all the numbers, using the feedback function
	newbit := SymCalculateNewBit(r)

	//shift all the numbers to the right, start at the end, copy from index before it
	for i := r.Length - 1; i > 0; i-- { //stops after arr[1] = arr[0]
		arr[i] = arr[i-1]
	}
	//set arr[0] to the new bit
	arr[0] = newbit

}

// SymClockingUnit
// Clock SR1, SR2, SR3 based on given state
func SymClockingUnit(rr4 Register) {
	arr := rr4.RegSlice
	maj := Majority(arr[3], arr[7], arr[10])

	if maj == arr[10] {
		SymClock(sr1)
	}
	if maj == arr[3] {
		SymClock(sr2)
	}
	if maj == arr[7] {
		SymClock(sr3)
	}
}

// SymCalculateNewBit
// calculate the new int slice by XOR'ing the tab-slices together row-wise
func SymCalculateNewBit(r SymRegister) []int {
	slice_slice := r.RegSlice

	newbit := make([]int, len(r.RegSlice[0])) // all 0 slice for first XOR

	for i := range r.Taps {
		tabslice := slice_slice[r.Taps[i]]        // get the slice for the tap
		for i := 0; i < len(r.RegSlice[0]); i++ { // loop through the slices and XOR them index-wise
			newbit[i] = newbit[i] ^ tabslice[i]
		}
	}
	return newbit
}

// SymInitializeRegisters
// calls SymSetRegisters and initialises them all to 0.
// Describes the registers with the current and original frame number.
//Sets the bit entries to 1. Copies r4 into sr4
func SymInitializeRegisters() {
	// Reset registers, all indexes are set to 0
	SymSetRegisters()

	sr1.RegSlice = DescribeNewFrameWithOldVariables(originalFrameNumber, currentFrameNumber, sr1)
	sr2.RegSlice = DescribeNewFrameWithOldVariables(originalFrameNumber, currentFrameNumber, sr2)
	sr3.RegSlice = DescribeNewFrameWithOldVariables(originalFrameNumber, currentFrameNumber, sr3)

	//Set bits to 1
	BitEntry(sr1)
	BitEntry(sr2)
	BitEntry(sr3)
}

/*
SymMakeFinalXOR
Makes the final xor of:

r1[-1] ⨁ maj(r1) ⨁ r2[-1] ⨁ maj(r2) ⨁ r3[-1] ⨁ maj(r3)

returns: [vars1 | vars2 | vars3 | prod1 | prod2 | prod3 | b ]

Calls SymMajorityOutput and OverwriteXorSlice
*/
func SymMakeFinalXOR(r1 SymRegister, r2 SymRegister, r3 SymRegister) []int {
	// save last entry in each register
	last_r1 := r1.RegSlice[r1.Length-1]
	last_r2 := r2.RegSlice[r2.Length-1]
	last_r3 := r3.RegSlice[r3.Length-1]

	maj_r1 := SymMajorityOutput(r1)
	maj_r2 := SymMajorityOutput(r2)
	maj_r3 := SymMajorityOutput(r3)

	// XOR them "locally" together first
	OverwriteXorSlice(last_r1, maj_r1) // [vars1 | prods1][b] = [vars1][b] ⨁ [vars1 | prods1][b]
	OverwriteXorSlice(last_r2, maj_r2)
	OverwriteXorSlice(last_r3, maj_r3)

	v1 := len(last_r1) - 1
	v2 := len(last_r2) - 1
	v3 := len(last_r3) - 1

	vars1 := maj_r1[0:v1] // vars1 points to the 19 [vars1] entries of maj_1 The 18 variables, maj_r1[v1] vil være x_01 som vi ikke vil have med

	start := make([]int, len(vars1))
	copy(start, vars1)                     // start by res = [vars1] (without the bit)
	start = append(start, maj_r2[0:v2]...) // now [vars1 | vars2 ]
	start = append(start, maj_r3[0:v3]...) // now [vars1 | vars2 | vars3]

	bitEntry1 := len(maj_r1) - 1
	bitEntry2 := len(maj_r2) - 1
	bitEntry3 := len(maj_r3) - 1

	start = append(start, maj_r1[v1:bitEntry1]...) // now [vars1 | vars2 | vars3 | prod1] without the bit entry
	start = append(start, maj_r2[v2:bitEntry2]...) // now [vars1 | vars2 | vars3 | prod1 | prod2 ]
	start = append(start, maj_r3[v3:bitEntry3]...) // now [vars1 | vars2 | vars3 | prod1 | prod2 | prod3]

	finalBit := maj_r1[bitEntry1] ^ maj_r2[bitEntry2] ^ maj_r3[bitEntry3]

	start = append(start, []int{finalBit}...)

	return start
}

/*
SymMajorityOutput
is the symbolic Majority function. Calls SymMajorityMultiply and XorSlice.
Performs xy ⨁ xz ⨁ yz ⨁ x ⨁ y on the Majority tabs of the register.
Returns slice of length
len(r)+ (len(r)*(len(r)-1))/2
with the original variables in the first len(r)-1 entries and products in the rest.
*/
func SymMajorityOutput(r SymRegister) []int {
	arr := r.RegSlice
	x := arr[r.MajsTaps[0]]
	y := arr[r.MajsTaps[1]]
	z := arr[r.NegTap]
	// xy ⨁ xz ⨁ yz ⨁ x ⨁ y
	xy := SymMajorityMultiply(x, y) // [vars | products ] [x1, x2, x3, ..., x12, x23]
	xz := SymMajorityMultiply(x, z)
	yz := SymMajorityMultiply(y, z)
	ee := XorSlice(xy, xz)
	long_slice := XorSlice(ee, yz) // This is a xor of normal and product variables //the last entry should be the bit entry
	short_slice := XorSlice(x, y)  // This is only an XOR of the normal variables
	// the long and short slice both have a bit in the last entry. these should be XOR'ed and added to the end of the res slice
	// XOR the "normal" variables in the start of the long slice [ vars | products ] ⨁ [ vars ]
	for i := 0; i < len(short_slice)-1; i++ { // -1 fordi så undgås bit entry
		long_slice[i] = long_slice[i] ^ short_slice[i]
	}
	long_slice[len(long_slice)-1] = long_slice[len(long_slice)-1] ^ short_slice[len(short_slice)-1] // XOR the bits together and place at the end
	return long_slice
}

// XorSlice
// takes two slices and XOR's them index wise together.
//Assumed to be of same length. Returns new slice of size len(a).
func XorSlice(a []int, b []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] ^ b[i]
	}
	return res
}

/*
OverwriteXorSlice
takes two slices with first shorter than the second.
Overwrites the first len(short) entries in long with
long[i] = short[i] ^ long[i]

long[len(long)-1] = long[len(long)-1] ^ short[len(short)-1].
*/
func OverwriteXorSlice(short []int, long []int) {
	for i := 0; i < len(short)-1; i++ {
		long[i] = short[i] ^ long[i]
	}
	long[len(long)-1] = long[len(long)-1] ^ short[len(short)-1]
}

/*
	SymMajorityMultiply
	multiplies two decision vectors with result being c[i]d[j] ^ c[j]d[i] for i /= j and result = c[i]d[j] for i=j.
	res slice has length len(c)*(len(c)-1)/2 + len(c).  c and d are assumed to be same length.
	The original len(c) variables will be in the first len(c) indexes of result
	The slice should be the full slice including the last bit index.
	The decision vectors corresponds to our symbolic representation, and express whether a specific unknown variable
	is part of the expression
*/
func SymMajorityMultiply(c []int, d []int) []int {
	lenc := len(c) - 1              // -1 fordi vi ikke vil loop over den konkrete bit til sidst
	leng := lenc * (lenc - 1) / 2   // for r1 : (18 * 17) / 2
	res := make([]int, leng+lenc+1) // +1 fordi der bliver lagt bit ind til sidst
	acc := 0
	for i := 0; i < lenc; i++ {
		res[i] = c[i]*d[i] ^ d[i]*c[lenc] ^ c[i]*d[lenc] // d[lenc] er bit plads
		for j := i + 1; j < lenc; j++ {
			res[lenc+acc] = c[i]*d[j] ^ c[j]*d[i]
			acc++
		}
	}
	res[len(res)-1] = c[lenc] * d[lenc] // sidste plads er bits ganget sammen

	return res
}

/*
	MakeSymKeyStream
	calls SymInitRegisters which clocks with key and frame.
	Calls ClockForKey with sr4 to get the 228-bit keystream.
*/
func MakeSymKeyStream() [][]int {
	SymInitializeRegisters()

	return ClockForKey(sr4)
}

// ClockForKey
// does the 99 + 228 clocking rounds based on the given register.
// Returns the 228 bit keystream.
func ClockForKey(r Register) [][]int {
	keyStream := make([][]int, 228)

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		SymClockingUnit(r)
		Clock(r)
	}

	// clock 228 times and save keystream
	for i := 0; i < 228; i++ {
		SymClockingUnit(r)
		Clock(r)
		keyStream[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}
	return keyStream

}
