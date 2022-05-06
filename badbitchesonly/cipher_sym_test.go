package main

import (
	"fmt"
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
	// reg1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	reg1 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	// ree := make([][]int, 19)
	for i := 0; i < 17; i++ {
		reg1.ArrImposter[i] = make([]int, 22)
		// ree[i][i] = 1
	}
	// printmatrix(reg1.ArrImposter)

	for i := 0; i < 22; i++ {
		SymClock(reg1)
		reg1.ArrImposter[0][i] = 1 //should this be xor?
		// ree[i][i] = 1
	}

	printmatrix(reg1.ArrImposter)
	PrettySymPrintSlice(reg1.ArrImposter)

}

func TestDescribeRegistersFromFrame(t *testing.T) {
	sre1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sre2 := SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sre3 := SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)

	for i := 0; i < sre1.Length; i++ {
		sre1.ArrImposter[i] = make([]int, 22)
	}

	for i := 0; i < sre2.Length; i++ {
		sre2.ArrImposter[i] = make([]int, 22)
	}

	for i := 0; i < sre3.Length; i++ {
		sre3.ArrImposter[i] = make([]int, 22)
	}

	reg4 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	for i := 0; i < 17; i++ {
		reg4.ArrImposter[i] = make([]int, 22)
	}

	for i := 0; i < 22; i++ {
		SymClock(sre1)
		SymClock(sre2)
		SymClock(sre3)
		SymClock(reg4)
		sre1.ArrImposter[0][i] = 1 //should this be xor? <- no den påvirkes kun af den i'te bit én gang
		sre2.ArrImposter[0][i] = 1
		sre3.ArrImposter[0][i] = 1
		reg4.ArrImposter[0][i] = 1
	}
	println("sr1")
	PrettySymPrintFrame(sre1.ArrImposter)
	println("sr2")
	PrettySymPrintFrame(sre2.ArrImposter)
	println("sr3")
	PrettySymPrintFrame(sre3.ArrImposter)
	println("sr4")
	PrettySymPrintFrame(reg4.ArrImposter)
}

func TestDescribeRegistersFromKey(t *testing.T) {
	sre1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sre2 := SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sre3 := SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)

	for i := 0; i < sre1.Length; i++ {
		sre1.ArrImposter[i] = make([]int, 64)
	}

	for i := 0; i < sre2.Length; i++ {
		sre2.ArrImposter[i] = make([]int, 64)
	}

	for i := 0; i < sre3.Length; i++ {
		sre3.ArrImposter[i] = make([]int, 64)
	}

	reg4 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	for i := 0; i < 17; i++ {
		reg4.ArrImposter[i] = make([]int, 64)
	}

	for i := 0; i < 64; i++ {
		SymClock(sre1)
		SymClock(sre2)
		SymClock(sre3)
		SymClock(reg4)
		sre1.ArrImposter[0][i] = 1 //should this be xor? <- no den påvirkes kun af den i'te bit én gang
		sre2.ArrImposter[0][i] = 1
		sre3.ArrImposter[0][i] = 1
		reg4.ArrImposter[0][i] = 1
	}
	println("sr1")
	PrettySymPrintFrame(sre1.ArrImposter)
	println("sr2")
	PrettySymPrintFrame(sre2.ArrImposter)
	println("sr3")
	PrettySymPrintFrame(sre3.ArrImposter)
	println("sr4")
	PrettySymPrintFrame(reg4.ArrImposter)

	sym := DescribeRegistersFromKey()
	PrettySymPrintFrame(sym)
	fmt.Printf("dims %d x %d of sym \n", len(sym), len(sym[0]))
	printmatrix(sym)

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

func TestSymClock(t *testing.T) {
	reg := InitOneSymRegister()
	// for i := 0; i < 19; i++ {
	// 	prints(reg.ArrImposter[i], "")
	// } <- her er der bare masser af nuller
	Bit_entry(reg)
	// for i := 0; i < 19; i++ {
	// 	prints(reg.ArrImposter[i], "")
	// } // <- her er der 1 taller diagonal
	PrettySymPrintSliceBit(reg.ArrImposter, reg.set1)
	SymClock(reg)
	SymClock(reg)
	SymClock(reg)
	for i := 0; i < 16; i++ {
		SymClock(reg)
	}
	PrettySymPrintSliceBit(reg.ArrImposter, reg.set1)
	SymClock(reg)
	// SymClock(reg)
	PrettySymPrintSliceBit(reg.ArrImposter, reg.set1)
}

func TestCompliance(t *testing.T) {
	symReg := InitOneSymRegister()
	reg := InitOneRegister()
	orgReg := make([]int, 19)
	copy(orgReg, reg.ArrImposter)
	prints(orgReg, "Original reg")
	Bit_entry(symReg)

	// make output keystream in both
	reg1 := SimpleKeyStreamSym(symReg)
	reg2 := SimpleKeyStream(reg)

	prints(reg1[0], "reg1[0")
	prints(reg1[1], "reg1[1]")
	PrettySymPrintSliceBit(reg1[:20], symReg.set1)
	prints(reg2[:20], "res")

	res := solveByGaussEliminationTryTwo(reg1, reg2)
	prints(res.Solved, "gauss")
	prints(orgReg, "Original reg")

}

func MakeLongIntSlice() []int {
	res := make([]int, 0)
	makeRegisters()
	SymSetRegisters()

	for i := 0; i < 15; i++ {
		res = append(res, i)
	}
	for i := 16; i < 19; i++ {
		res = append(res, i)
	}

	for i := 0; i < 16; i++ {
		res = append(res, i)
	}
	for i := 16 + 1; i < r2.Length; i++ {
		res = append(res, i)
	}

	for i := 0; i < sr3.set1; i++ {
		res = append(res, i)
	}
	for i := sr3.set1 + 1; i < r3.Length; i++ {
		res = append(res, i)
	}

	// prints(res, "res before products")

	//lol the first 0-x products just becomes x
	prod1 := stringToIntArray("01 02 03 04 05 06 07 08 09 010 011 012 013 014 016 017 018 12 13 14 15 16 17 18 19 110 111 112 113 114 116 117 118 23 24 25 26 27 28 29 210 211 212 213 214 216 217 218 34 35 36 37 38 39 310 311 312 313 314 316 317 318 45 46 47 48 49 410 411 412 413 414 416 417 418 56 57 58 59 510 511 512 513 514 516 517 518 67 68 69 610 611 612 613 614 616 617 618 78 79 710 711 712 713 714 716 717 718 89 810 811 812 813 814 816 817 818 910 911 912 913 914 916 917 918 1011 1012 1013 1014 1016 1017 1018 1112 1113 1114 1116 1117 1118 1213 1214 1216 1217 1218 1314 1316 1317 1318 1416 1417 1418 1617 1618 1718")
	prod2 := stringToIntArray("01 02 03 04 05 06 07 08 09 010 011 012 013 014 015 017 018 019 020 021 12 13 14 15 16 17 18 19 110 111 112 113 114 115 117 118 119 120 121 23 24 25 26 27 28 29 210 211 212 213 214 215 217 218 219 220 221 34 35 36 37 38 39 310 311 312 313 314 315 317 318 319 320 321 45 46 47 48 49 410 411 412 413 414 415 417 418 419 420 421 56 57 58 59 510 511 512 513 514 515 517 518 519 520 521 67 68 69 610 611 612 613 614 615 617 618 619 620 621 78 79 710 711 712 713 714 715 717 718 719 720 721 89 810 811 812 813 814 815 817 818 819 820 821 910 911 912 913 914 915 917 918 919 920 921 1011 1012 1013 1014 1015 1017 1018 1019 1020 1021 1112 1113 1114 1115 1117 1118 1119 1120 1121 1213 1214 1215 1217 1218 1219 1220 1221 1314 1315 1317 1318 1319 1320 1321 1415 1417 1418 1419 1420 1421 1517 1518 1519 1520 1521 1718 1719 1720 1721 1819 1820 1821 1920 1921 2021")
	prod3 := stringToIntArray("01 02 03 04 05 06 07 08 09 010 011 012 013 014 015 016 017 019 020 021 022 12 13 14 15 16 17 18 19 110 111 112 113 114 115 116 117 119 120 121 122 23 24 25 26 27 28 29 210 211 212 213 214 215 216 217 219 220 221 222 34 35 36 37 38 39 310 311 312 313 314 315 316 317 319 320 321 322 45 46 47 48 49 410 411 412 413 414 415 416 417 419 420 421 422 56 57 58 59 510 511 512 513 514 515 516 517 519 520 521 522 67 68 69 610 611 612 613 614 615 616 617 619 620 621 622 78 79 710 711 712 713 714 715 716 717 719 720 721 722 89 810 811 812 813 814 815 816 817 819 820 821 822 910 911 912 913 914 915 916 917 919 920 921 922 1011 1012 1013 1014 1015 1016 1017 1019 1020 1021 1022 1112 1113 1114 1115 1116 1117 1119 1120 1121 1122 1213 1214 1215 1216 1217 1219 1220 1221 1222 1314 1315 1316 1317 1319 1320 1321 1322 1415 1416 1417 1419 1420 1421 1422 1516 1517 1519 1520 1521 1522 1617 1619 1620 1621 1622 1719 1720 1721 1722 1920 1921 1922 2021 2022 2122")

	res = append(res, prod1...)
	// prints(res, "res after prod1")
	res = append(res, prod2...)
	res = append(res, prod3...)

	return res
}

// func TestTest(*testing.T) {
// 	p := []int{1, 2, 3, 4}

// 	w, err := os.Create("output.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer w.Close()

// 	wri := bufio.NewWriter(w)

// 	encoder := gob.NewEncoder(wri)
// 	err = encoder.Encode(p)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fi, err := os.Open("output.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// close fi on exit and check for its returned error
// 	defer func() {
// 		if err := fi.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// rea := bufio.NewReader(fi)

// 	// decoder := gob.NewDecoder(rea)
// 	// p = []int{}

// 	// err = decoder.Decode(&p)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// prints(p, "")
// }

// func TestTest2(t *testing.T) {
// 	file, _ := os.Create("user.gob")

// 	p := []int{1, 2, 3, 4}

// 	defer file.Close()

// 	encoder := gob.NewEncoder(file)

// 	encoder.Encode(p)

// 	decoder := gob.NewDecoder(file)
// 	q := []int{}

// 	err := decoder.Decode(&q)
// 	if err != nil {
// 		panic(err)
// 	}
// 	prints(p, "")
// }
