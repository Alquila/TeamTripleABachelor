package main

import (
	. "fmt"
	"math/rand"
	"testing"
	"time"
	"reflect"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func TestPrint(t *testing.T) {
	print("hello world!")
}

func TestMajority(t *testing.T) {

	x := majority(0, 0, 0)
	if x != 0 {
		t.Errorf(" x is not 0 but %d", x)
	}

	x = majority(0, 0, 1)
	if x != 0 {
		t.Errorf(" x is not 0 but %d", x)
	}

	x = majority(0, 1, 1)
	if x != 1 {
		t.Errorf(" x is not 1 but %d", x)
	}

	x = majority(1, 1, 1)
	if x != 1 {
		t.Errorf(" x is not 1 but %d", x)
	}

	x = majority(1, 0, 1)
	if x != 1 {
		t.Errorf(" x is not 1 but %d", x)
	}
}

func makeSmallReg() Register { //[0 0 0 0 0 0 0 0 0 0]
	r1 = Register{Number: 0,
		Length:      10,
		ArrImposter: make([]int, 10),
		Tabs:        []int{3, 5, 9}, // [0] = [3] ^ [5] ^ [9]
		Majs:        []int{4, 7},
		Ært:         6}

	return r1
}

func TestClock(t *testing.T) {

	r0 := makeSmallReg()

	r0.ArrImposter[8] = 1

	Clock(r0)

	if r0.ArrImposter[9] != 1 {
		t.Errorf("x_9 should be 1 but was %d", r0.ArrImposter[9])
	}

	for i := 0; i < 10; i++ {
		prettyPrint(r0)
		Clock(r0)
	}

}

func TestMakeSessionKey(t *testing.T) {
	rand.Seed(time.Now().Unix())

	key := make([]int, 64)
	for i := 0; i < 64; i++ {
		key[i] = rand.Intn(2)
	}
	Printf("%+v \n", key)
}

//Generate a random array of length n
//Println(rand.Perm(64))

func TestMakeFrameNumber(t *testing.T) {
	f := 55

	frameBit := make([]int, 22)

	//LSB is at the 21'th index
	// for i := 0; i < 22; i++ {
	// 	frameBit[21-i] = (f >> i) & 1
	// }

	//opposite way - LSB is at the 0'th index
	for i := 0; i < 22; i++ {
		frameBit[i] = (f >> i) & 1
	}

	Printf("%+v \n", frameBit)
	Printf("0'th bit is %v \n", frameBit[0])
}

func TestSetIndiciesToOne(t *testing.T) {
	makeRegisters()
	setIndicesToOne()
	printAll()
}

func TestInitialiseRegisters(t *testing.T) {	// TODO: test at initreg er forskellig når framenumber er forskellig :-)
	makeRegisters()
	frame_number = 22
	makeSessionKey()
	initialiseRegisters()
	reg1 := r1
	reg2 := r2
	reg3 := r3 
	reg4 := r4

	r1.ArrImposter[6] = 42
	r2.ArrImposter[6] = 42
	r3.ArrImposter[6] = 42
	r4.ArrImposter[6] = 42
	printAll()
	Printf("Initialise registers again: \n")
	initialiseRegisters()
	printAll()
	
	if reflect.DeepEqual(reg1.ArrImposter, r1.ArrImposter) {
		t.Log("reg1 and r1 are different")
		t.Fail()
	}
	if reflect.DeepEqual(reg2.ArrImposter, r2.ArrImposter) {
		t.Log("reg2 and r2 are different")
		t.Fail()
	}
	if reflect.DeepEqual(reg3.ArrImposter, r3.ArrImposter) {
		t.Log("reg3 and r3 are different")
		t.Fail()
	}
	if reflect.DeepEqual(reg4.ArrImposter, r4.ArrImposter) {
		t.Log("reg4 and r4 are different")
		t.Fail()
	}

}

func TestCalculateNewBit(t *testing.T){
	makeRegisters()	

	a1 := r1.ArrImposter
	a1[13] = 1; a1[16] = 0; a1[17] = 1; a1[18] = 0; //set the tap indexes to concrete values 1 ⨁ 0 ⨁ 1 ⨁ 0 = 0
	res := calculateNewBit(r1)
	if( res != 0){t.Fail()}
	
	
	a1[13] = 1; a1[16] = 0; a1[17] = 1; a1[18] = 1; //set the tap indexes to concrete values 1 ⨁ 0 ⨁ 1 ⨁ 1 = 1
	res = calculateNewBit(r1)
	if( res != 1){t.Fail()}
}

func TestMajorityOutput(t *testing.T){}

func TestClockingUnit(t *testing.T){}

func TestFinalXor(t *testing.T){}

func TestKeyStream(t *testing.T){}
