package main

import (
	_ "fmt"
	_ "math/rand"
	_ "time"
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

func SymMakeRegister(length int, tabs []int, major_idx []int, compliment_idx int) SymRegister {
	reg := SymRegister{
		Length:      length,
		ArrImposter: make([][]int, length),
		Tabs:        tabs,
		Majs:        major_idx,
		Ært:         compliment_idx}
	return reg
}

func SymSetRegisters() {
	sr1 = SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14)
	sr2 = SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16)
	sr3 = SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13)
	sr4 = SymMakeRegister(17, []int{16, 11}, nil, -1)

}

//Calculate the new int slice by xor'ing the tab-slices together row-wise
func SymCalculateNewBit(r SymRegister) []int {
	slice_slice := r.ArrImposter

	newbit := make([]int, r.Length) //all 0 slice for first xor

	for i := range r.Tabs {
		tabslice := slice_slice[i] //get the slice for the tap

		for i := 0; i < r.Length; i++ { //loop through the slices and xor them index-wise
			newbit[i] = newbit[i] ^ tabslice[i]
		}
	}
	return newbit
}

func SymInitialiseRegisters() {
	// Reset registers
	SymSetRegisters()

	for i := 0; i < 64; i++ {
		//TODO: Clock symR1-R4

		//REVIEW: nomalt xor med sessionskey - skal dette stadig gøres?
	}

	// makes frame_number from int -> bits in array
	frame_bits := makeFrameNumberToBits(frame_number)

	for i := 0; i < 22; i++ {
		//TODO: clock

		//REVIEW: xor med framenumber
	}
}
