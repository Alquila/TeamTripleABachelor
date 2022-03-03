package main

import (
	. "fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func TestPrint(t *testing.T) {
	print("hello world!")
	print("1 ⨁ 0 ⨁ 1 ⨁ 0 = 0")
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
	r1 = Register{
		Length:      10,
		ArrImposter: make([]int, 10),
		Tabs:        []int{3, 5, 9}, // [0] = [3] ^ [5] ^ [9]
		Majs:        []int{4, 7},
		Ært:         6}

	return r1
}

func TestSimpleKeyStream(t *testing.T) {
	r0 := SymRegister{Length: 10,
		ArrImposter: make([][]int, 10),
		Tabs:        []int{3, 5, 9}, // [0] = [3] ^ [5] ^ [9]
		Majs:        []int{4, 7},
		Ært:         6}

	print("rip1")
	for i := 0; i < r0.Length; i++ {
		r0.ArrImposter[i] = make([]int, r0.Length)
		r0.ArrImposter[i][i] = 1
	}
	print("rip")
	key := SimpleKeyStream(r0)
	Println("rip3")
	for i := 0; i < 228; i++ {
		// PrettyPrint()
		for i := 0; i < r0.Length; i++ {
			accString := "["
			for j := 0; j < r0.Length; j++ {
				if key[i][j] == 1 {
					str := strconv.Itoa(j)
					accString += "x" + (str) + " ⨁ "
					// accString += " x" + (str) + " xor "
				}
			}
			// accString = strings.TrimRight(accString, " xor ")
			accString = strings.TrimRight(accString, "⨁ ")
			//Printf("xor)
			//Printf("")
			accString += " ]"
			print(accString)
		}
		Println()

	}

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

func TestSmallPrint(t *testing.T) {
	r0 := SymRegister{Length: 10,
		ArrImposter: make([][]int, 10),
		Tabs:        []int{3, 5, 9}, // [0] = [3] ^ [5] ^ [9]
		Majs:        []int{4, 7},
		Ært:         6}

	for i := 0; i < r0.Length; i++ {
		r0.ArrImposter[i] = make([]int, r0.Length)
	}

	r0.ArrImposter[8][8] = 1
	r0.ArrImposter[1][1] = 1
	r0.ArrImposter[5][5] = 1
	r0.ArrImposter[3][3] = 1
	PrettyPrint(r0)

	SymClock(r0)
	//Printf("%+v \n", r0.ArrImposter)
	//println(" 1st clock")
	//PrettyPrint(r0)
	SymClock(r0)
	//println(" 2nd clock")
	//PrettyPrint(r0)
	//Printf("%+v \n", r0.ArrImposter)
	SymClock(r0)
	//Printf("%+v \n", r0.ArrImposter)
	SymClock(r0)
	//Printf("%+v \n", r0.ArrImposter)
	SymClock(r0)
	//Printf("%+v \n", r0.ArrImposter)
	SymClock(r0)
	SymClock(r0)
	SymClock(r0)
	SymClock(r0)
	SymClock(r0)
	PrettyPrint(r0)
	SymClock(r0)
	SymClock(r0)
	SymClock(r0)
	SymClock(r0)
	SymClock(r0)
	Printf("%+v \n", r0.ArrImposter)
	PrettyPrint(r0)
	//Printf("%+v \n", r0.ArrImposter)

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

func TestMakeFrameNumber(t *testing.T) { // REVIEW: denne test, tester ikke metoden i vores kode, det bør den måske istedet
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
	if r1.ArrImposter[15] != 1 {
		t.Log("r1[15] should be 1 but was ", r1.ArrImposter[15])
		t.Fail()
	}
	if r2.ArrImposter[16] != 1 {
		t.Log("r2[16] should be 1 but was ", r2.ArrImposter[16])
		t.Fail()
	}
	if r3.ArrImposter[18] != 1 {
		t.Log("r3[18] should be 1 but was ", r3.ArrImposter[18])
		t.Fail()
	}
	if r4.ArrImposter[10] != 1 {
		t.Log("r4[10] should be 1 but was ", r4.ArrImposter[10])
		t.Fail()
	}
}

func TestRegistersAreSameAfterInitWithSameFrameNumber(t *testing.T) { // TODO: test at initreg er forskellig når framenumber er forskellig :-)
	makeRegisters()
	frame_number = 22
	makeSessionKey()
	initialiseRegisters()
	reg1 := make([]int, 19)
	reg2 := make([]int, 22)
	reg3 := make([]int, 23)
	reg4 := make([]int, 17)
	copy(reg1, r1.ArrImposter)
	copy(reg2, r2.ArrImposter)
	copy(reg3, r3.ArrImposter)
	copy(reg4, r4.ArrImposter)
	//TODO få amalie til at forklare den her test

	Printf("First initialisation: \n")
	printAll()

	r1.ArrImposter[6] = 42
	r2.ArrImposter[6] = 42
	r3.ArrImposter[6] = 42
	r4.ArrImposter[6] = 42

	//printAll()
	initialiseRegisters()
	Printf("Initialise registers again: \n")
	printAll()

	if !reflect.DeepEqual(reg1, r1.ArrImposter) {
		t.Log("reg1 and r1 are different, but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg2, r2.ArrImposter) {
		t.Log("reg2 and r2 are different but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg3, r3.ArrImposter) {
		t.Log("reg3 and r3 are different but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg4, r4.ArrImposter) {
		t.Log("reg4 and r4 are different but should be equal")
		t.Fail()
	}

}

func TestCalculateNewBit(t *testing.T) {
	makeRegisters()

	a1 := r1.ArrImposter
	a1[13] = 1
	a1[16] = 0
	a1[17] = 1
	a1[18] = 0 //set the tap indexes to concrete values 1 ⨁ 0 ⨁ 1 ⨁ 0 = 0
	res := calculateNewBit(r1)
	if res != 0 {
		t.Fail()
	}

	a1[13] = 1
	a1[16] = 0
	a1[17] = 1
	a1[18] = 1 //set the tap indexes to concrete values 1 ⨁ 0 ⨁ 1 ⨁ 1 = 1
	res = calculateNewBit(r1)
	if res != 1 {
		t.Fail()
	}

	a2 := r2.ArrImposter
	a2[20] = 0
	a2[21] = 0 //set the tap indexes to concrete values  0 ⨁ 0 = 0
	res = calculateNewBit(r2)
	if res != 0 {
		t.Fail()
	}

	a2[20] = 1
	a2[21] = 0 //set the tap indexes to concrete values 1 ⨁ 0 = 1
	res = calculateNewBit(r2)
	if res != 1 {
		t.Fail()
	}

	a3 := r3.ArrImposter
	a3[22] = 0
	a3[21] = 1
	a3[20] = 0
	a3[7] = 1 //set the tap indexes to concrete values 0 ⨁ 1 ⨁ 0 ⨁ 1 = 0
	res = calculateNewBit(r3)
	if res != 0 {
		t.Fail()
	}

	a3[22] = 1
	a3[21] = 1
	a3[20] = 0
	a3[7] = 1 //set the tap indexes to concrete values 1 ⨁ 1 ⨁ 0 ⨁ 1 = 1
	res = calculateNewBit(r3)
	if res != 1 {
		t.Fail()
	}

}

func TestMajorityOutput(t *testing.T) {
	makeRegisters()

	a := r1.ArrImposter
	a[12] = 1
	a[14] = 1
	a[15] = 0
	//set the tap indexes to concrete values maj(1,(1 ⨁ 1), 0)
	res := majorityOutput(r1)
	if res != 0 {
		t.Errorf(" x is not 0 but %d", res)
	}
	a[12] = 1
	a[14] = 0
	a[15] = 0
	//set the tap indexes to concrete values maj(1,(0 ⨁ 1), 0)
	res = majorityOutput(r1)
	if res != 1 {
		t.Errorf(" x is not 1 but %d", res)
	}

	a = r2.ArrImposter
	a[9] = 1
	a[13] = 1
	a[16] = 0
	//set the tap indexes to concrete values maj(1, 1, (0 ⨁ 1))
	res = majorityOutput(r2)
	if res != 1 {
		t.Errorf(" x is not 1 but %d", res)
	}
	a[9] = 1
	a[13] = 0
	a[16] = 1
	//set the tap indexes to concrete values maj(1, 0, (1 ⨁ 1))
	res = majorityOutput(r2)
	if res != 0 {
		t.Errorf(" x is not 0 but %d", res)
	}

	a = r3.ArrImposter
	a[13] = 1
	a[16] = 0
	a[18] = 0
	//set the tap indexes to concrete values maj((1 ⨁ 1), 0, 0)
	res = majorityOutput(r3)
	if res != 0 {
		t.Errorf(" x is not 0 but %d", res)
	}
	a[13] = 1
	a[16] = 1
	a[18] = 1
	//set the tap indexes to concrete values maj((1 ⨁ 1), 1, 1)
	res = majorityOutput(r3)
	if res != 1 {
		t.Errorf(" x is not 0 but %d", res)
	}
}

func TestClockingUnit(t *testing.T) {
	makeRegisters()

	a := r4.ArrImposter
	//clock R2 og R3
	a[3] = 1
	a[7] = 1
	a[10] = 0
	clockingUnit(r4) //will print those it clocks

	//clock all
	Clock(r4) //will have 0's in the indexes
	clockingUnit(r4)
}

func TestFinalXor(t *testing.T) {}

func TestKeyStreamSimple(t *testing.T) {
	makeSessionKey() // TODO snak om hvor vores loop skal være, som kalder makeKeyStream for nye frames
	frame_number = -1
	x := makeKeyStream()
	Printf("%+v \n", x)
	Printf("mean %f \n", mean(x))
}

func mean(a []int) float64 {
	sum := 0.0
	for _, v := range a {
		sum += (float64(v))
	}
	return (sum / float64(len(a)))
}
