package main

import (
	// "math/rand"
	// "reflect"
	// "strconv"
	// "strings"
	"fmt"
	"reflect"
	"testing"
	// "time"
	// //"golang.org/x/tools/go/analysis/passes/nilfunc"
)

// func TestPlaintext(t *testing.T) {
// 	plaintext := MakePlaintext()
// 	fmt.Printf("%d", plaintext)
// }

// func TestEncryptPlaintext(t *testing.T) {
// 	plaintext := MakePlaintext()
// 	fmt.Printf("This is the plaintext: %d \n", plaintext)
// 	cipher := EncryptSimplePlaintext(plaintext)
// 	fmt.Printf("%d \n", cipher)
// }

// func TestSymPlaintext(t *testing.T) {
// 	plaintext := MakeSymPlaintext()
// 	fmt.Printf("This is the plaintext: %d \n", plaintext)
// }

// func TestSymEncryptPlaintext(t *testing.T) {
// 	plaintext := MakeSymPlaintext()
// 	fmt.Printf("This is the plaintext: %d \n", plaintext)
// 	cipher := EncryptSimpleSymPlaintext()
// }

func NotAllowedBigBangTest(t *testing.T) {
	//plaintext := MakePlaintext()

}

// func TestDoTheSimpleHack(t *testing.T) {
// 	doTheSimpleHack()
// }

func TestPrint2(t *testing.T) {
	print("hello worlds")
}

func TestDoTheSimpleHack(t *testing.T) {
	// TODO: make this a function in dumb_assversary.go plz

	// init one register, in both OG and sym version
	symReg := InitOneSymRegister()
	reg := InitOneRegister()
	orgReg := make([]int, 19)
	copy(orgReg, reg.ArrImposter)

	// make output keystream in both
	symKeyStream := SimpleKeyStreamSym(symReg)
	keyStream := SimpleKeyStream(reg)

	// make sym version into [][]int if not allready

	// use gauss to solve equations
	//res := solveByGaussElimination(symKeyStream, keyStream)
	res := solveByGaussEliminationTryTwo(symKeyStream, keyStream)

	fmt.Printf("Res er: %d\n", res)
	fmt.Printf("reg er: %d\n", orgReg)

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res, orgReg) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("reg er: %d\n", orgReg)
	}
}

func TestDoTheSimpleHackSecondVersion(t *testing.T) {
	// TODO: make this a function in dumb_assversary.go plz

	// init one register, in both OG and sym version
	symReg := InitOneSymRegister()
	reg := InitOneRegister()
	orgReg := make([]int, 19)
	copy(orgReg, reg.ArrImposter)

	// make output keystream in both
	symKeyStream := SimpleKeyStreamSymSecondVersion(symReg)
	fmt.Printf("length of symKeyStream[0]: %d\n", len(symKeyStream[0]))
	keyStream := SimpleKeyStreamSecondVersion(reg)
	fmt.Printf("length of KeyStream: %d\n", len(keyStream))

	// make sym version into [][]int if not allready

	// use gauss to solve equations
	//res := solveByGaussElimination(symKeyStream, keyStream)
	res := solveByGaussEliminationTryTwo(symKeyStream, keyStream)

	fmt.Printf("længden af res er: %d\n", len(res))

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res[0:19], orgReg) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("reg er: %d\n", orgReg)
	}
	fmt.Printf("reg er: %d\n", orgReg)
	fmt.Printf("Res er: %d\n", res[0:19])
}

func TestPlaintextAttack(t *testing.T) {
	r4.ArrImposter = []int{0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1}
	orgReg := make([]int, 17)
	copy(orgReg, r4.ArrImposter)

	session_key = make([]int, 64)

	res := DoTheKnownPlainTextHack()

	fmt.Printf("len af res er: %d\n", len(res))

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res[0:16], orgReg) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("reg er: %d\n", orgReg)
	}
	fmt.Printf("reg er: %d\n", orgReg)
	fmt.Printf("Res er: %d\n", res[0:19])

}

func TestFindDiffOfFrameNumbers(t *testing.T) {

	res := FindDifferenceOfFrameNumbers(1, 2)

	// LEAST significant bit is at index 0, so the bit is kinda 'reversed'
	shouldBe := make([]int, 22)
	shouldBe[0] = 1
	shouldBe[1] = 1

	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("shouldBe er: %d\n", shouldBe)
	}
}

func TestDescribeNewFrameNumberWithOldVar(t *testing.T) {
	firstSymReg := InitOneSymRegister()
	prints(firstSymReg.ArrImposter[15], "række 15")

	res := DescribeNewFrameWithOldVariables(0, 1, firstSymReg.ArrImposter)

	fmt.Printf("res er: \n%d \n", res)
	fmt.Printf("res er %d \n", len(res))
	fmt.Printf("res[0] er %d \n", len(res[0]))
	PrettySymPrintSlice(res)
}

func TestSymInit(t *testing.T) {
	SymInitializeRegisters()
}
