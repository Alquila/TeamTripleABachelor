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
	key := SimpleKeyStreamSym(r0)
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
		r0.ArrImposter[i] = make([]int, r0.Length+1)
	}

	r0.ArrImposter[8][8] = 1
	r0.ArrImposter[4][10] = 1
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
	current_frame_number = 22
	makeSessionKey()
	initializeRegisters()
	setIndicesToOne()
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
	initializeRegisters()
	setIndicesToOne()
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

func TestFinalXor(t *testing.T) {
	//hooow to teeeeest
	//right now we just show that it takes the three registers and calculates the majority and takes the last slice in each and xors it all together. returns a long array with all the stuff
	r1 := SymMakeRegister(4, []int{1, 3}, []int{0, 2}, 3, 0)
	for i := 0; i < 4; i++ {
		r1.ArrImposter[i][i] = 1 // each entry in the diagonal set to 1 as x_i is only dependent on x_i when initialized
	}
	r2 := SymMakeRegister(5, []int{1, 3}, []int{2, 4}, 0, 0)
	for i := 0; i < 5; i++ {
		r2.ArrImposter[i][i] = 1 // each entry in the diagonal set to 1 as x_i is only dependent on x_i when initialized
	}
	r3 := SymMakeRegister(8, []int{3, 4, 6, 5}, []int{1, 2}, 0, 0)
	for i := 0; i < 8; i++ {
		r3.ArrImposter[i][i] = 1 // each entry in the diagonal set to 1 as x_i is only dependent on x_i when initialized
	}

	res := SymMakeFinalXOR(r1, r2, r3)
	prints(res, "result")
	Printf("lenght %d \n", len(res))
	// 4+5+8 = 17
	// 4*3/2 + 5*4/2+ 8*7/2 = 44

	maj_r1 := SymMajorityOutput(r1)
	maj_r2 := SymMajorityOutput(r2)
	maj_r3 := SymMajorityOutput(r3)

	prints(maj_r1, "r1 majority")
	prints(maj_r2, "r2 majority")
	prints(maj_r3, "r3 majority")
	// r1 majority[1 0 1 0 0 1 1 0 0 1]
	// r2 majority[0 0 1 0 1 0 1 0 1 0 0 0 0 1 0]
	// r3 majority[0 1 1 0 0 0 0 0 1 1 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	// last_r1 := r1.ArrImposter[r1.Length-1]
	// last_r2 := r2.ArrImposter[r2.Length-1]
	// last_r3 := r3.ArrImposter[r3.Length-1]
	// prints(last_r1, "last r1")
	// prints(last_r2, "last r2")
	// prints(last_r3, "last r3")
	/*
			last r1[0 0 0 1]
			last r2[0 0 0 0 1]
			last r3[0 0 0 0 0 0 0 1]

		result[1 0 1 1 0 1 1 0 0 1 0 0 1 0 0 0 1 0 1 0 0 0 0 1 0 0 1 1 0 0 0 0 1 1 1 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		lenght 61

	*/

}

func TestKeyStreamSimple(t *testing.T) {
	makeSessionKey() // TODO snak om hvor vores loop skal være, som kalder makeKeyStream for nye frames
	current_frame_number = -1
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

func Symaaa(c []string, d []string) []string {
	lenc := len(c) - 1 //REVIEW -1 fordi vi ikke vil loop over den konkrete bit til sidst
	leng := lenc * (lenc - 1) / 2
	res := make([]string, leng+lenc+1) //REVIEW +1 fordi der bliver lagt bit ind til sidst
	acc := 0
	for i := 0; i < lenc; i++ {
		res[i] = c[i] + d[i] + c[lenc] + d[lenc] //REVIEW d[lenc] er bit plads
		for j := i + 1; j < lenc; j++ {
			res[lenc+acc] = c[i] + d[j] //+ c[j] + d[i]
			//Printf("res[%d] = %d*%d ^ %d*%d = %d \n", lenc+acc, c[i], d[j], c[j], d[i], res[lenc+acc])
			acc++
		}
	}
	res[len(res)-1] = c[lenc] + d[lenc]

	return res
}

func TestSymaa(t *testing.T) {
	c := []string{"0", "1", "2", "3", "a"}
	d := []string{"0", "1", "2", "3", "b"}

	res := Symaaa(c, d)
	// print(res)
	Printf("%+v \n", res)
	Printf("lenght: %d \n", len(res))
}

func TestSymMajorityMultiply(t *testing.T) {

	c := []int{0, 1, 0, 1, 1}
	d := []int{0, 1, 1, 0, 1}

	res := SymMajorityMultiply(c, d)

	Printf("%+v \n", res)
	Printf("lenght: %d \n", len(res))

}

// TODO
func TestSymMajorityOutput(t *testing.T) {
	c := []int{0, 1, 0, 1}
	d := []int{0, 1, 1, 0}
	e := []int{1, 0, 1, 0}

	r0 := SymRegister{Length: 4,
		ArrImposter: make([][]int, 4),
		Tabs:        []int{0, 0, 0},
		Majs:        []int{0, 1},
		Ært:         2}

	r0.ArrImposter[0] = c
	r0.ArrImposter[1] = d
	r0.ArrImposter[2] = e

	Printf("cd  %+v \n", SymMajorityMultiply(c, d))
	Printf("de  %+v \n", SymMajorityMultiply(d, e))
	Printf("ce  %+v \n", SymMajorityMultiply(c, e))

	shouldBe := []int{0, 1, 0, 1, 0, 1, 1, 1, 1, 0} //see notes
	res := SymMajorityOutput(r0)
	if !reflect.DeepEqual(res, shouldBe) {
		t.Logf("res is wrong %+v \n", res)
		t.Fail()
	}
}

func PrettySymPrint(symReg SymRegister) {
	for i := 0; i < symReg.Length; i++ {
		accString := "["
		for j := 0; j < symReg.Length; j++ {
			if symReg.ArrImposter[i][j] == 1 {
				str := strconv.Itoa(j)
				accString += "x" + (str) + " ⨁ "
			}
		}
		accString = strings.TrimRight(accString, "⨁ ")
		accString += "]"
		print(accString)
	}
	Println()

}

//works on slice_slice

func TestInitOneSymRegister(t *testing.T) {
	reg := InitOneSymRegister()
	PrettySymPrint(reg)
}

func TestSimpleKeyStreamSym(t *testing.T) {
	reg := InitOneSymRegister()

	keyStream := SimpleKeyStreamSym(reg)

	PrettySymPrintSlice(keyStream)
}

func TestOverwriteXorSlice(t *testing.T) {
	short := []int{1, 1, 0, 1, 1}
	long := []int{0, 1, 0, 1, 1, 1, 1, 1}
	OverwriteXorSlice(short, long)
	Printf("%+v \n", long)
}

func TestAppend(t *testing.T) {

	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 5}
	c := []int{7, 8, 9, 10, 11, 12, 13}

	start := make([]int, len(a))
	copy(start, a)              //start by res = [vars1 | prod1]
	start = append(start, b...) //now [vars1 | prod1 | vars2 | prod2 ]
	start = append(start, c...) //now [vars1 | prod1 | vars2 | prod2 | vars3 | prod3]

	Printf("a %+v \n", a)
	Printf("b %+v \n", b)
	Printf("c %+v \n", c)
	Printf("start %+v \n", start)

}

func TestSymInit(t *testing.T) {
	SymInitializeRegisters()
}
