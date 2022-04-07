package main

import (
	. "fmt"
	"strconv"
)

/* global variable declaration */
var sr1 SymRegister
var sr2 SymRegister
var sr3 SymRegister
var sr4 Register

type SymRegister struct {
	Length      int
	ArrImposter [][]int
	Tabs        []int
	Majs        []int
	Ært         int
	set1        int
}

// var sym_session_key [][]int

func SymMakeRegister(length int, tabs []int, major_idx []int, compliment_idx int, bit_idx int) SymRegister {
	reg := SymRegister{
		Length:      length,
		ArrImposter: make([][]int, length),
		Tabs:        tabs,
		Majs:        major_idx,
		Ært:         compliment_idx,
		set1:        bit_idx} 

	for i := 0; i < reg.Length; i++ {
		reg.ArrImposter[i] = make([]int, reg.Length)
	}

	return reg
}

// This does what
func Bit_entry(reg SymRegister) {
	// reg.ArrImposter[reg.bit_entry] = make([]int, reg.Length)
	// reg.ArrImposter[reg.bit_entry][reg.Length-1] = 1
	bit_entry := reg.set1
	for i := 0; i < reg.Length; i++ {
		reg.ArrImposter[i] = make([]int, reg.Length)
		if i < bit_entry {
			reg.ArrImposter[i][i] = 1
		}
		if i == bit_entry {
			reg.ArrImposter[i][reg.Length-1] = 1
		}
		if i > bit_entry {
			reg.ArrImposter[i][i-1] = 1
		}
	}
}

//Calls SymMakeRegister on each register. Each register is initialised and the symbolic slices are +1 of r.Lenght to make space for extra bit
func SymSetRegisters() {
	sr1 = SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sr2 = SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sr3 = SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)
	sr4 = makeR4()
}

//Clock R1, R2, R3 based on R4 state
func SymClockingUnit(r4 Register) {
	arr := r4.ArrImposter
	maj := majority(arr[3], arr[7], arr[10])

	if maj == arr[10] {
		SymClock(sr1)
		print("clock R1\n")
	}
	if maj == arr[3] {
		//clock R2
		SymClock(sr2)
		print("clock R2\n")
	}
	if maj == arr[7] {
		//clock R3
		SymClock(sr3)
		print("clock R3\n")
	}
}

func SymClock(r SymRegister) {
	arr := r.ArrImposter

	//calculate the new bit before shifting all the numbers, using the feedback function
	newbit := SymCalculateNewBit(r)
	//print(newbit)

	//shift all the numbers to the right, start at the end, copy from index before it
	for i := r.Length - 1; i > 0; i-- { //stops after arr[1] = arr[0]
		arr[i] = arr[i-1]
	}
	//set arr[0] to the new bit
	arr[0] = newbit

}

//Calculate the new int slice by xor'ing the tab-slices together row-wise
func SymCalculateNewBit(r SymRegister) []int {
	slice_slice := r.ArrImposter

	newbit := make([]int, len(r.ArrImposter[0])) //all 0 slice for first xor

	for i := range r.Tabs {
		tabslice := slice_slice[r.Tabs[i]] //get the slice for the tap
		//Printf("slice %d is %+v \n",r.Tabs[i], tabslice)
		for i := 0; i < len(r.ArrImposter[0]); i++ { //loop through the slices and xor them index-wise
			newbit[i] = newbit[i] ^ tabslice[i]
		}
	}
	return newbit
}

func SymInitializeRegisters() {
	// Reset registers, all indexes are set to 0
	SymSetRegisters()

	/* for i := 0; i < 64; i++ {
	 	SymClock(sr1)
		SymClock(sr2)
	 	SymClock(sr3)
	 	SymClock(sr4)
	}
	*/

	/* for i := 0; i < 22; i++ {
		SymClock(sr1)
	 	SymClock(sr2)
	 	SymClock(sr3)
	 	SymClock(sr4)
	 	// : xor med framenumber
	 	// : we pretend that the framenumber is 0
	 	// frame_bits[i] skal XORs her
	}*/
	//REVIEW noget med noget framenumber f' her eller et eller andet sted

	//Set bits to 1 
	Bit_entry(sr1)
	Bit_entry(sr2)
	Bit_entry(sr3)
	sr4.ArrImposter[10] = 1
}

func SymInitializeRegistersFrame(old_frame int, new_frame int) {

	SymSetRegisters()

	sr1.ArrImposter = DescribeNewFrameWithOldVariables(old_frame, new_frame, sr1)
	sr2.ArrImposter = DescribeNewFrameWithOldVariables(old_frame, new_frame, sr2)
	sr3.ArrImposter = DescribeNewFrameWithOldVariables(old_frame, new_frame, sr3)

	//FIXME R4?

	Bit_entry(sr1)
	Bit_entry(sr2)
	Bit_entry(sr3)
	sr4.ArrImposter[10] = 1
}


/*
Makes the final xor of r1[-1] ⨁ maj(r1) ⨁ r2[-1] ⨁ maj(r2) ⨁ r3[-1] ⨁ maj(r3)
returns [vars1 | prod1 | vars2 | prod2 | vars3 | prod3 | b]
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
	OverwriteXorSlice(last_r1, maj_r1) //[vars1 | prods1][b] = [vars1][b] ⨁ [vars1 | prods1][b]
	OverwriteXorSlice(last_r2, maj_r2)
	OverwriteXorSlice(last_r3, maj_r3)

	v1 := len(last_r1) - 1 //19
	v2 := len(last_r2) - 1
	v3 := len(last_r3) - 1
	print("lenght of v1")
	print(v1)
	vars1 := maj_r1[0:v1] //vars1 points to the 19 [vars1] entries of maj_1 //REVIEW the 18 variables, maj_r1[v1] vil være x_01

	start := make([]int, len(vars1))
	copy(start, vars1)                     //start by res = [vars1] (without the bit)
	start = append(start, maj_r2[0:v2]...) //now [vars1 | vars2 ]
	start = append(start, maj_r3[0:v3]...) //now [vars1 | vars2 | vars3]

	bit_entry1 := len(maj_r1) - 1
	bit_entry2 := len(maj_r2) - 1
	bit_entry3 := len(maj_r3) - 1

	start = append(start, maj_r1[v1:bit_entry1]...) //now [vars1 | vars2 | vars3 | prod1] without the bit entry
	start = append(start, maj_r2[v2:bit_entry2]...) //now [vars1 | prod1 | vars2 | prod2 ]
	start = append(start, maj_r3[v2:bit_entry3]...) //now [vars1 | prod1 | vars2 | prod2 | vars3 | prod3]

	final_bit := maj_r1[bit_entry1] ^ maj_r2[bit_entry2] ^ maj_r3[bit_entry3] 
	
	start = append(start, []int{final_bit}...)
	//now [vars1 | prod1 | vars2 | prod2 | vars3 | prod3 | b ]

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
	long_slice := XorSlice(ee, yz) //This is a xor of normal and product variables //the last entry should be the bit entry
	short_slice := XorSlice(x, y)  //This is only a xor of the normal variables
	//REVIEW the long and short slice both have a bit in the last entry. these should be xored and added to the end of the res slice
	// xor the "normal" variables in the start of the long slice [ vars | products ] ⨁ [ vars ]
	for i := 0; i < len(short_slice)-1; i++ { //REVIEW -1 fordi så undgå bit indgang
		long_slice[i] = long_slice[i] ^ short_slice[i]
	}
	long_slice[len(long_slice)-1] = long_slice[len(long_slice)-1] ^ short_slice[len(short_slice)-1] //REVIEW xor the bits together and place at the end
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
//long[len(long)-1] = long[len(long)-1] ^ short[len(short)-1]
func OverwriteXorSlice(short []int, long []int) {
	for i := 0; i < len(short)-1; i++ {
		long[i] = short[i] ^ long[i]
	}
	long[len(long)-1] = long[len(long)-1] ^ short[len(short)-1]
}

/*
multiplies two decision vectors with result being c[i]d[j] ^ c[j]d[i] for i /= j and result = c[i]d[j] for i=j.
res slice has lenght len(c)*(len(c)-1)/2 + len(c).  c and d are assumed to be same lenght.
The original len(c) variables will be in the first len(c) indexes of result
The slice should be the full slice including the last bit index
*/
func SymMajorityMultiply(c []int, d []int) []int {
	lenc := len(c) - 1 //REVIEW -1 fordi vi ikke vil loop over den konkrete bit til sidst
	leng := lenc * (lenc - 1) / 2
	res := make([]int, leng+lenc+1) //REVIEW +1 fordi der bliver lagt bit ind til sidst
	acc := 0
	for i := 0; i < lenc; i++ {
		res[i] = c[i]*d[i] ^ c[lenc] ^ d[lenc] //REVIEW d[lenc] er bit plads
		for j := i + 1; j < lenc; j++ {
			res[lenc+acc] = c[i]*d[j] ^ c[j]*d[i]
			//Printf("res[%d] = %d*%d ^ %d*%d = %d \n", lenc+acc, c[i], d[j], c[j], d[i], res[lenc+acc])
			acc++
		}
	}
	res[len(res)-1] = c[lenc] * d[lenc] //REVIEW sidste plads er bits ganget sammen

	return res
}

func makeSymKeyStream() [][]int {
	SymInitializeRegisters()

	keyStream := make([][]int, 228)

	// Clock the register 99 times
	for i := 0; i < 99; i++ {
		
		SymClockingUnit(r4)
		Clock(sr4)
	}

	// clock 228 times and save keystream
	for i := 0; i < 228; i++ {
		SymClockingUnit(r4)
		Clock(sr4)
		//OverwriteXorSlice(r.ArrImposter[r.Length-1], aaa)
		keyStream[i] = SymMakeFinalXOR(sr1, sr2, sr3)
		// overw(SymMajorityOutput(r), r.ArrImposter[r.Length-1])
		//Printf("Length of output from symMajorFunc %d\n", len(SymMajorityOutput(r)))
	}
	return keyStream
}

/*
###########################################################
#### THIS IS WHERE THE SIMPLE CIPHER SYM STREAM EXISTS ####
###########################################################
*/

func InitOneSymRegister() SymRegister {
	reg := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15) // equvalent to reg1
	for i := 0; i < 19; i++ {
		// reg.ArrImposter[i] = make([]int, 19)
		reg.ArrImposter[i][i] = 1 // each entry in the diagonal set to 1 as x_i is only dependent on x_i when initialized
	}
	reg.ArrImposter[15][15] = 0
	reg.ArrImposter[15][len(reg.ArrImposter[15])-1] = 1
	return reg
}

func SimpleKeyStreamSym(r SymRegister) [][]int {

	// Init key stream array
	keyStream := make([][]int, 228)
	for i := 0; i < 228; i++ {
		keyStream[i] = make([]int, r.Length+1)
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
		aaa := SymMajorityOutput(r)
		OverwriteXorSlice(r.ArrImposter[r.Length-1], aaa)
		keyStream[i] = aaa
		// overw(SymMajorityOutput(r), r.ArrImposter[r.Length-1])
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
	rMatrix := r.ArrImposter
	rBit := r.set1

	PrettySymPrintSliceBit(rMatrix, rBit)
}

func PrettySymPrintSlice(slice [][]int) {
	for i := 0; i < len(slice); i++ { //19
		accString := "["
		for j := 0; j < len(slice[0])-1; j++ { //19
			if slice[i][j] == 1 {
				str := strconv.Itoa(j)
				accString += "x" + (str) + " ⨁ "
			}
		}
		accString += strconv.Itoa(slice[i][len(slice[0])-1]) + "]  \n"
		print(accString)
	}
	Println()

}

func PrettySymPrintSliceBit(rMatrix [][]int, bit_entry int) {
	rLength := len(rMatrix)
	for i := 0; i < rLength; i++ {
		accString := "r" + strconv.Itoa(i) + " = "
		for j := 0; j < len(rMatrix[0])-1; j++ {
			if j >= bit_entry {
				if rMatrix[i][j] == 1 {
					str := strconv.Itoa(j + 1)
					accString += " x" + (str) + " ⨁ "
				}
			} else if rMatrix[i][j] == 1 {
				str := strconv.Itoa(j)
				accString += " x" + (str) + " ⨁ "
			}
		}
		// accString = strings.TrimRight(accString, "⨁ ")
		accString = accString + strconv.Itoa(rMatrix[i][rLength-1])
		Printf("")
		println(accString)
	}

}


func prints(res []int, text string) {
	Printf(text+"%+v \n", res)
}

func printmatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		prints(matrix[i], "")
	}
}
