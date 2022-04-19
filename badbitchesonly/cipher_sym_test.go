package main

import (
	_ "fmt"
	"reflect"
	"testing"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func TestBitEntry(t *testing.T) {
	reg := InitOneSymRegister()
	PrettySymPrintSlice(reg.ArrImposter)
	Bit_entry(reg)
	PrettySymPrintSlice(reg.ArrImposter)
	for i := 0; i < reg.Length; i++ {
		prints(reg.ArrImposter[i], "")
	}
}

func TestConstPrettyPrint(t *testing.T) {
	reg := InitOneSymRegister()
	Bit_entry(reg)
	PrettyPrint(reg)
}

func TestHowFrames(t *testing.T) {
	reg1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	// ree := make([][]int, 19)
	for i := 0; i < 19; i++ {
		reg1.ArrImposter[i] = make([]int, 22)
		// ree[i][i] = 1
	}
	// printmatrix(reg1.ArrImposter)

	for i := 0; i < 22; i++ {
		SymClock(reg1)
		reg1.ArrImposter[0][i] = 1
		// ree[i][i] = 1
	}

	printmatrix(reg1.ArrImposter)
	PrettySymPrintSlice(reg1.ArrImposter)

}

func TestSymRegistersSameAfterInitWithSameFrameNumber(t *testing.T) {
	current_frame_number = 22
	original_frame_number = 22
	makeSessionKey()
	makeRegisters()
	initializeRegisters()
	SymInitializeRegisters()

	reg1 := make([]int, 19)
	reg2 := make([]int, 22)
	reg3 := make([]int, 23)
	reg4 := make([]int, 17)
	copy(reg1, sr1.ArrImposter[0])
	copy(reg2, sr2.ArrImposter[0])
	copy(reg3, sr3.ArrImposter[0])
	copy(reg4, sr4.ArrImposter)

	sr1.ArrImposter[0][0] = 42
	sr2.ArrImposter[0][0] = 42
	sr3.ArrImposter[0][0] = 42
	sr4.ArrImposter[0] = 42

	makeRegisters()
	initializeRegisters()
	SymInitializeRegisters()

	if !reflect.DeepEqual(reg1, sr1.ArrImposter[0]) {
		t.Log("reg1 and r1 are different, but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg2, sr2.ArrImposter[0]) {
		t.Log("reg2 and r2 are different but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg3, sr3.ArrImposter[0]) {
		t.Log("reg3 and r3 are different but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg4, sr4.ArrImposter) {
		t.Log("reg4 and r4 are different but should be equal")
		t.Fail()
	}

}

func TestFinalXorSomething(t *testing.T) {
	current_frame_number = 22
	original_frame_number = 22
	makeSessionKey()
	makeRegisters()
	initializeRegisters()
	SymInitializeRegisters()
	// sr1.ArrImposter[sr1.Length-1][sr1.Length-1] = 1

	last_r1 := sr1.ArrImposter[sr1.Length-1]
	last_r2 := sr2.ArrImposter[sr2.Length-1]
	last_r3 := sr3.ArrImposter[sr3.Length-1]

	v1 := len(last_r1) - 1 //18
	v2 := len(last_r2) - 1
	v3 := len(last_r3) - 1

	maj_r1 := SymMajorityOutput(sr1)
	maj_r2 := SymMajorityOutput(sr2)
	maj_r3 := SymMajorityOutput(sr3)

	bit_entry1 := len(maj_r1) - 1
	// bit_entry2 := len(maj_r2) - 1
	// bit_entry3 := len(maj_r3) - 1

	res := SymMakeFinalXOR(sr1, sr2, sr3)

	prints(last_r1, "last sr1\n")
	prints(maj_r1[:v1], "maj_sr1\n")
	prints(res[:v1], "first r1 entries of finalxor\n")

	prints(last_r2, "last sr2\n")
	prints(maj_r2[:v2], "maj_sr2\n")
	prints(res[v1:v1+v2], "first r2 entries of finalxor\n")

	prints(last_r3, "last sr3\n")
	prints(maj_r3[:v3], "maj_sr3\n")
	prints(res[v1+v2:v1+v2+v3], "first r3 entries of finalxor\n")

	prints(maj_r1[v1:bit_entry1], "maj bits of sr1\n")
	print(len(maj_r1[v1:bit_entry1]))
	print("\n")
	print(len(res[v1+v2+v3 : v2+v3+bit_entry1]))
	prints(res[v1+v2+v3:v2+v3+bit_entry1], "maj bits sr1 finalxor\n")
	print(res[len(res)-1])
}
