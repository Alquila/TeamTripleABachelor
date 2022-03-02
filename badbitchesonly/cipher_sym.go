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
	
	for i:=0; i< reg.Length; i++ {
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
	symClock(sr1)
	print("clock R1\n")
	//}
	//if maj == arr[3] {
	//clock R2
	symClock(sr2)
	print("clock R2\n")
	//}
	//if maj == arr[7] {
	//clock R3
	symClock(sr3)
	print("clock R3\n")
	//}
}

func symClock(r SymRegister) {
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

func symMakeSessionKey() {
	rand.Seed(time.Now().Unix())

	key := make([]int, 64)
	for i := 0; i < 64; i++ {
		key[i] = rand.Intn(2)
	}
	sym_session_key = make([][]int, 4)
}

func SymInitialiseRegisters() {
	// Reset registers
	SymSetRegisters()

	for i := 0; i < 64; i++ {
		symClock(sr1)
		symClock(sr2)
		symClock(sr3)
		symClock(sr4)

		// REVIEW: nomalt xor med sessions key - skal dette stadig gøres?
		// REVIEW: we pretend that the session key is 0 #verySafe sorry Ivan
		// session_key[i] skal XORs her
	}

	// makes frame_number from int -> bits in array
	//frame_bits := makeFrameNumberToBits(frame_number)

	for i := 0; i < 22; i++ {
		symClock(sr1)
		symClock(sr2)
		symClock(sr3)
		symClock(sr4)

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
