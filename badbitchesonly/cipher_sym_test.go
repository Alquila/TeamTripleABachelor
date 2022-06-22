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
	PrettySymPrintSlice(reg.RegSlice)
	BitEntry(reg)
	PrettySymPrintSlice(reg.RegSlice)
	for i := 0; i < reg.Length; i++ {
		Prints(reg.RegSlice[i], "")
	}
}

func TestConstPrettyPrint(t *testing.T) {
	reg := InitOneSymRegister()
	BitEntry(reg)
	PrettyPrintSymRegister(reg)
}

func TestHowFrames(t *testing.T) {
	reg1 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	for i := 0; i < 17; i++ {
		reg1.RegSlice[i] = make([]int, 22)
	}

	for i := 0; i < 22; i++ {
		SymClock(reg1)
		reg1.RegSlice[0][i] = 1
	}

	PrintMatrix(reg1.RegSlice)
	PrettySymPrintSlice(reg1.RegSlice)

}

func TestDescribeRegistersFromFrame(t *testing.T) {
	sre1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sre2 := SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sre3 := SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)

	for i := 0; i < sre1.Length; i++ {
		sre1.RegSlice[i] = make([]int, 22)
	}

	for i := 0; i < sre2.Length; i++ {
		sre2.RegSlice[i] = make([]int, 22)
	}

	for i := 0; i < sre3.Length; i++ {
		sre3.RegSlice[i] = make([]int, 22)
	}

	reg4 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	for i := 0; i < 17; i++ {
		reg4.RegSlice[i] = make([]int, 22)
	}

	for i := 0; i < 22; i++ {
		SymClock(sre1)
		SymClock(sre2)
		SymClock(sre3)
		SymClock(reg4)
		sre1.RegSlice[0][i] = 1
		sre2.RegSlice[0][i] = 1
		sre3.RegSlice[0][i] = 1
		reg4.RegSlice[0][i] = 1
	}
	println("sr1")
	PrettySymPrintFrame(sre1.RegSlice)
	println("sr2")
	PrettySymPrintFrame(sre2.RegSlice)
	println("sr3")
	PrettySymPrintFrame(sre3.RegSlice)
	println("sr4")
	PrettySymPrintFrame(reg4.RegSlice)
}

func TestDescribeRegistersFromKey(t *testing.T) {
	sym := DescribeRegistersFromKey()
	PrettySymPrintFrame(sym)
	fmt.Printf("dims %d x %d of sym \n", len(sym), len(sym[0])) //81 x 64
	PrintMatrix(sym)
}

func TestRip(t *testing.T) {
	rr1 := MakeR1()
	rr2 := MakeR2()
	rr3 := MakeR3()
	rr4 := MakeR4()
	rr1.RegSlice = stringToIntArray("0 1 0 0 0 1 1 1 1 1 1 0 0 1 0 0 0 0 0")
	rr1.RegSlice = make([]int, rr1.Length)
	rr1.RegSlice[18] = 1
	rr2.RegSlice = stringToIntArray("0 1 1 0 0 1 0 0 0 0 1 1 0 1 1 1 1 1 1 0 0 1")
	rr3.RegSlice = stringToIntArray("1 0 1 0 1 0 1 0 0 0 0 1 0 1 0 0 1 1 1 1 0 1 1")
	rr4.RegSlice = stringToIntArray("1 0 0 1 0 1 1 1 0 0 1 1 0 0 0 1 0")

	frame := MakeFrameNumberToBits(6)
	Prints(rr1.RegSlice, "regs pre:")
	Prints(frame, "frame:   ")
	for i := 0; i < 22; i++ {
		Clock(rr1)
		Clock(rr2)
		Clock(rr3)
		Clock(rr4)
		rr1.RegSlice[0] = rr1.RegSlice[0] ^ frame[i]
		rr2.RegSlice[0] = rr2.RegSlice[0] ^ frame[i]
		rr3.RegSlice[0] = rr3.RegSlice[0] ^ frame[i]
		rr4.RegSlice[0] = rr4.RegSlice[0] ^ frame[i]
	}
	Prints(rr1.RegSlice, "regs:    ")

	sym := DescribeRegistersFromFrame()
	all_reg := append(rr1.RegSlice, rr2.RegSlice...)
	all_reg = append(all_reg, rr3.RegSlice...)
	all_reg = append(all_reg, rr4.RegSlice...)

	rriip6 := []int{frame[2] ^ frame[3] ^ frame[4] ^ frame[7] ^ frame[21],
		frame[1] ^ frame[2] ^ frame[3] ^ frame[6] ^ frame[20],
		frame[0] ^ frame[1] ^ frame[2] ^ frame[5] ^ frame[19],
		frame[0] ^ frame[1] ^ frame[4] ^ frame[18],
		frame[0] ^ frame[3] ^ frame[17],
		frame[2] ^ frame[16]}
	Prints(frameReg(sym, frame)[:19], "rii    ")
	// rip7 := frameReg(sym,frame)

	// x := SolveByGaussElimination(sym, all_reg)
	// print(x.ResType)

	// PrettySymPrintFrame(sym)
	Prints(rriip6, "equation")
}

func TestRip2(t *testing.T) {

	sym := DescribeRegistersFromKeyAndFrame()
	PrettySymPrintKeyFrame(sym)

}

func frameReg(sym [][]int, frame []int) []int {
	res := make([]int, 85) //change 85
	println(len(sym))
	println(len(sym[0]))
	for i := 0; i < len(sym); i++ {
		result := 0
		for j := 0; j < len(sym[0]); j++ {
			if sym[i][j] == 1 {
				result = result ^ frame[j]
			}
		}
		res[i] = result
	}
	return res
}

/*func TestRetrieveSessionKey(t *testing.T) {
	sessionKey = stringToIntArray("0 0 0 0 0 1 1 0 1 1 0 1 0 1 0 1 0 1 0 1 0 1 1 1 1 0 0 0 0 1 0 0 0 0 0 0 0 1 1 1 1 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0")
	real_ses := make([]int, 64)
	copy(real_ses, sessionKey)
	Prints(real_ses, "rree")

	MakeRegisters()
	InitializeRegisters()

	rr1 := stringToIntArray("0 1 0 0 0 1 1 1 1 1 1 0 0 1 0 0 0 0 0")
	rr2 := stringToIntArray("0 1 1 0 0 1 0 0 0 0 1 1 0 1 1 1 1 1 1 0 0 1")
	rr3 := stringToIntArray("1 0 1 0 1 0 1 0 0 0 0 1 0 1 0 0 1 1 1 1 0 1 1")
	rr4 := stringToIntArray("1 0 0 1 0 1 1 1 0 0 1 1 0 0 0 1 0")
	all_reg := append(rr1, rr2...)
	all_reg = append(all_reg, rr3...)
	all_reg = append(all_reg, rr4...)

	res := RetrieveSessionKey(all_reg)
	println(len(res))
	if !reflect.DeepEqual(res, real_ses) {
		Prints(res, "res was: ")
		t.Fail()
	}

}*/

func TestRetrieveSessionKey2(t *testing.T) {
	sessionKey = stringToIntArray("0 0 0 0 0 1 1 0 1 1 0 1 0 1 0 1 0 1 0 1 0 1 1 1 1 0 0 0 0 1 0 0 0 0 0 0 0 1 1 1 1 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0")
	MakeSessionKey()
	real_ses := make([]int, 64)
	copy(real_ses, sessionKey)
	Prints(real_ses, "rree")
	currentFrameNumber = 6

	MakeRegisters()
	InitializeRegisters()

	symkey := DescribeRegistersFromKey()
	symFrame := frameReg(DescribeRegistersFromFrame(), MakeFrameNumberToBits(currentFrameNumber))

	println(len(symkey))
	println(len(symkey[0]))
	println(len(symFrame))
	Prints(symFrame[:12], "fram")
	A := make([][]int, 81)
	for i := 0; i < len(symkey); i++ {
		A[i] = append(symkey[i], symFrame[i])
	}
	// PrintMatrix(A[:20])
	b := make([]int, 0)
	b = append(r1.RegSlice, r2.RegSlice...)
	b = append(b, r3.RegSlice...)
	b = append(b, r4.RegSlice...)

	x := SolveByGaussElimination(A, b)
	println(x.ResType)
	Prints(x.Multi[0], "reee")

}

func TestOldWay(t *testing.T) {
	MakeSessionKey()
	real_ses := make([]int, 64)
	copy(real_ses, sessionKey)
	Prints(real_ses, "rree")
	currentFrameNumber = 6

	MakeRegisters()
	r1.RegSlice = make([]int, r1.Length)
	r2.RegSlice = make([]int, r2.Length)
	r3.RegSlice = make([]int, r3.Length)
	r4.RegSlice = make([]int, r4.Length)

	// Clock all registers 64 times and XOR with the session key
	for i := 0; i < 64; i++ {
		Clock(r1)
		Clock(r2)
		Clock(r3)
		Clock(r4)

		r1.RegSlice[0] = r1.RegSlice[0] ^ sessionKey[i]
		r2.RegSlice[0] = r2.RegSlice[0] ^ sessionKey[i]
		r3.RegSlice[0] = r3.RegSlice[0] ^ sessionKey[i]
		r4.RegSlice[0] = r4.RegSlice[0] ^ sessionKey[i]
	}
	preFrame := make([]int, 0)
	preFrame = append(r1.RegSlice, r2.RegSlice...)
	preFrame = append(preFrame, r3.RegSlice...)
	preFrame = append(preFrame, r4.RegSlice...)
	// makes frame_number from int -> bits in slice
	frame_bits := MakeFrameNumberToBits(currentFrameNumber)

	// Clock all registers 22 times and XOR with frame number
	for i := 0; i < 22; i++ {
		Clock(r1)
		Clock(r2)
		Clock(r3)
		Clock(r4)

		r1.RegSlice[0] = r1.RegSlice[0] ^ frame_bits[i]
		r2.RegSlice[0] = r2.RegSlice[0] ^ frame_bits[i]
		r3.RegSlice[0] = r3.RegSlice[0] ^ frame_bits[i]
		r4.RegSlice[0] = r4.RegSlice[0] ^ frame_bits[i]
	}

	symFrame := frameReg(DescribeRegistersFromFrame(), MakeFrameNumberToBits(currentFrameNumber))
	ree := make([]int, 77)
	ree = remove(symFrame, 75)
	// copy(ree, symFrame[:15])                                                          //x_0 - x_14
	// x_15tox_16 := 16+3+16
	// x_17tox_17 := x_15tox_16 +
	// ree = append(ree, symFrame[16:x_15tox_16]...)                                        //x_16 - x_18, x_0 - x_15
	// ree = append(ree, symFrame[16+16+3+1:16+16+3+1+5+18]...)                          //x_17 - x_21, x_0-x_17
	// ree = append(ree, symFrame[16+16+3+1+5+18+1:16+16+3+1+5+18+1+4+10]...)            //x_19 - x_22, x_0-x_9
	// ree = append(ree, symFrame[16+16+3+1+5+18+1+4+10+1:16+16+3+1+5+18+1+4+10+1+6]...) //x_11-x_x_16
	fmt.Printf("ree: %d \n", len(ree))

	ee := helpSym()
	fmt.Printf("x_sym: %d x %d\n", (len(ee)), (len(ee[0])))
	A := make([][]int, 81)
	for i := 0; i < len(A); i++ {
		A[i] = append(ee[i], symFrame[i])
	}
	// PrintMatrix(A[:20])
	fmt.Printf("A: %d x %d \n", (len(A)), (len(A[0])))
	b := make([]int, 0)
	b = append(r1.RegSlice, r2.RegSlice...)
	b = append(b, r3.RegSlice...)
	b = append(b, r4.RegSlice...)

	// PrintMatrix(ee[:20])

	x := SolveByGaussElimination(A, b)
	println(x.ResType)
	Prints(x.Multi[0], "reee")
	Prints(preFrame, "pref")
	Prints(preFrame[:15], "15")
	Prints(preFrame[16:16+3], "16")

}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func helpSym() [][]int {
	sre1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sre2 := SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sre3 := SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)

	for i := 0; i < sre1.Length; i++ {
		sre1.RegSlice[i] = make([]int, sre1.Length)
		sre1.RegSlice[i][i] = 1
	}

	for i := 0; i < sre2.Length; i++ {
		sre2.RegSlice[i] = make([]int, sre2.Length)
		sre2.RegSlice[i][i] = 1
	}

	for i := 0; i < sre3.Length; i++ {
		sre3.RegSlice[i] = make([]int, sre3.Length)
		sre3.RegSlice[i][i] = 1
	}

	reg4 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	for i := 0; i < 17; i++ {
		reg4.RegSlice[i] = make([]int, reg4.Length)
		reg4.RegSlice[i][i] = 1
	}

	for i := 0; i < 22; i++ {
		SymClock(sre1)
		SymClock(sre2)
		SymClock(sre3)
		SymClock(reg4)
	}

	for i := 0; i < sre1.Length; i++ {
		rr := make([]int, 81)
		copy(rr, sre1.RegSlice[i])
		sre1.RegSlice[i] = rr
	}
	for i := 0; i < sre2.Length; i++ {
		rr := make([]int, 81)
		copy(rr[sre1.Length:], sre2.RegSlice[i])
		sre2.RegSlice[i] = rr
	}

	for i := 0; i < sre3.Length; i++ {
		rr := make([]int, 81)
		copy(rr[sre1.Length+sre2.Length:], sre3.RegSlice[i])
		sre3.RegSlice[i] = rr
	}

	for i := 0; i < 17; i++ {
		rr := make([]int, 81)
		copy(rr[sre1.Length+sre2.Length+sre3.Length:], reg4.RegSlice[i])
		reg4.RegSlice[i] = rr
	}

	symbolicDescription := make([][]int, 0)
	symbolicDescription = append(symbolicDescription, sre1.RegSlice...)
	symbolicDescription = append(symbolicDescription, sre2.RegSlice...)
	symbolicDescription = append(symbolicDescription, sre3.RegSlice...)
	symbolicDescription = append(symbolicDescription, reg4.RegSlice...)

	return symbolicDescription

}

func TestSymRegistersSameAfterInitWithSameFrameNumber(t *testing.T) {
	currentFrameNumber = 22
	originalFrameNumber = 22
	MakeSessionKey()
	MakeRegisters()
	InitializeRegisters()
	SymInitializeRegisters()

	reg1 := make([]int, 19)
	reg2 := make([]int, 22)
	reg3 := make([]int, 23)
	reg4 := make([]int, 17)
	copy(reg1, sr1.RegSlice[0])
	copy(reg2, sr2.RegSlice[0])
	copy(reg3, sr3.RegSlice[0])
	copy(reg4, sr4.RegSlice)

	sr1.RegSlice[0][0] = 42
	sr2.RegSlice[0][0] = 42
	sr3.RegSlice[0][0] = 42
	sr4.RegSlice[0] = 42

	MakeRegisters()
	InitializeRegisters()
	SymInitializeRegisters()

	if !reflect.DeepEqual(reg1, sr1.RegSlice[0]) {
		t.Log("reg1 and r1 are different, but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg2, sr2.RegSlice[0]) {
		t.Log("reg2 and r2 are different but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg3, sr3.RegSlice[0]) {
		t.Log("reg3 and r3 are different but should be equal")
		t.Fail()
	}
	if !reflect.DeepEqual(reg4, sr4.RegSlice) {
		t.Log("reg4 and r4 are different but should be equal")
		t.Fail()
	}

}

func TestFinalXorSomething(t *testing.T) {
	currentFrameNumber = 22
	originalFrameNumber = 22
	MakeSessionKey()
	MakeRegisters()
	InitializeRegisters()
	SymInitializeRegisters()

	last_r1 := sr1.RegSlice[sr1.Length-1]
	last_r2 := sr2.RegSlice[sr2.Length-1]
	last_r3 := sr3.RegSlice[sr3.Length-1]

	v1 := len(last_r1) - 1 //18
	v2 := len(last_r2) - 1
	v3 := len(last_r3) - 1

	maj_r1 := SymMajorityOutput(sr1)
	maj_r2 := SymMajorityOutput(sr2)
	maj_r3 := SymMajorityOutput(sr3)

	bit_entry1 := len(maj_r1) - 1

	res := SymMakeFinalXOR(sr1, sr2, sr3)

	Prints(last_r1, "last sr1\n")
	Prints(maj_r1[:v1], "maj_sr1\n")
	Prints(res[:v1], "first r1 entries of finalxor\n")

	Prints(last_r2, "last sr2\n")
	Prints(maj_r2[:v2], "maj_sr2\n")
	Prints(res[v1:v1+v2], "first r2 entries of finalxor\n")

	Prints(last_r3, "last sr3\n")
	Prints(maj_r3[:v3], "maj_sr3\n")
	Prints(res[v1+v2:v1+v2+v3], "first r3 entries of finalxor\n")

	Prints(maj_r1[v1:bit_entry1], "maj bits of sr1\n")
	print(len(maj_r1[v1:bit_entry1]))
	print("\n")
	print(len(res[v1+v2+v3 : v2+v3+bit_entry1]))
	Prints(res[v1+v2+v3:v2+v3+bit_entry1], "maj bits sr1 finalxor\n")
	print(res[len(res)-1])
}

func TestFinalXorLenght(t *testing.T) {
	currentFrameNumber = 22
	originalFrameNumber = 22
	sessionKey = make([]int, 64)
	MakeRegisters()
	InitializeRegisters()
	SymInitializeRegisters()

	res := SymMakeFinalXOR(sr1, sr2, sr3)
	println(len(res))

}

func TestSymClock(t *testing.T) {
	reg := InitOneSymRegister()
	BitEntry(reg)
	PrettySymPrintSliceBit(reg.RegSlice, reg.SetToOne)

	SymClock(reg)
	SymClock(reg)
	SymClock(reg)
	for i := 0; i < 16; i++ {
		SymClock(reg)
	}
	PrettySymPrintSliceBit(reg.RegSlice, reg.SetToOne)
	SymClock(reg)
	PrettySymPrintSliceBit(reg.RegSlice, reg.SetToOne)
}

func TestCompliance(t *testing.T) {
	symReg := InitOneSymRegister()
	reg := InitOneRegister()
	orgReg := make([]int, 19)
	copy(orgReg, reg.RegSlice)
	Prints(orgReg, "Original reg")
	BitEntry(symReg)

	// make output keystream in both
	reg1 := SimpleKeyStreamSym(symReg)
	reg2 := SimpleKeyStream(reg)

	Prints(reg1[0], "reg1[0")
	Prints(reg1[1], "reg1[1]")
	PrettySymPrintSliceBit(reg1[:20], symReg.SetToOne)
	Prints(reg2[:20], "res")

	res := SolveByGaussElimination(reg1, reg2)
	Prints(res.Solved, "gauss")
	Prints(orgReg, "Original reg")

}

func MakeLongIntSlice() []int {
	res := make([]int, 0)
	MakeRegisters()
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

	for i := 0; i < sr3.SetToOne; i++ {
		res = append(res, i)
	}
	for i := sr3.SetToOne + 1; i < r3.Length; i++ {
		res = append(res, i)
	}

	//lol the first 0-x products just becomes x
	prod1 := stringToIntArray("01 02 03 04 05 06 07 08 09 010 011 012 013 014 016 017 018 12 13 14 15 16 17 18 19 110 111 112 113 114 116 117 118 23 24 25 26 27 28 29 210 211 212 213 214 216 217 218 34 35 36 37 38 39 310 311 312 313 314 316 317 318 45 46 47 48 49 410 411 412 413 414 416 417 418 56 57 58 59 510 511 512 513 514 516 517 518 67 68 69 610 611 612 613 614 616 617 618 78 79 710 711 712 713 714 716 717 718 89 810 811 812 813 814 816 817 818 910 911 912 913 914 916 917 918 1011 1012 1013 1014 1016 1017 1018 1112 1113 1114 1116 1117 1118 1213 1214 1216 1217 1218 1314 1316 1317 1318 1416 1417 1418 1617 1618 1718")
	prod2 := stringToIntArray("01 02 03 04 05 06 07 08 09 010 011 012 013 014 015 017 018 019 020 021 12 13 14 15 16 17 18 19 110 111 112 113 114 115 117 118 119 120 121 23 24 25 26 27 28 29 210 211 212 213 214 215 217 218 219 220 221 34 35 36 37 38 39 310 311 312 313 314 315 317 318 319 320 321 45 46 47 48 49 410 411 412 413 414 415 417 418 419 420 421 56 57 58 59 510 511 512 513 514 515 517 518 519 520 521 67 68 69 610 611 612 613 614 615 617 618 619 620 621 78 79 710 711 712 713 714 715 717 718 719 720 721 89 810 811 812 813 814 815 817 818 819 820 821 910 911 912 913 914 915 917 918 919 920 921 1011 1012 1013 1014 1015 1017 1018 1019 1020 1021 1112 1113 1114 1115 1117 1118 1119 1120 1121 1213 1214 1215 1217 1218 1219 1220 1221 1314 1315 1317 1318 1319 1320 1321 1415 1417 1418 1419 1420 1421 1517 1518 1519 1520 1521 1718 1719 1720 1721 1819 1820 1821 1920 1921 2021")
	prod3 := stringToIntArray("01 02 03 04 05 06 07 08 09 010 011 012 013 014 015 016 017 019 020 021 022 12 13 14 15 16 17 18 19 110 111 112 113 114 115 116 117 119 120 121 122 23 24 25 26 27 28 29 210 211 212 213 214 215 216 217 219 220 221 222 34 35 36 37 38 39 310 311 312 313 314 315 316 317 319 320 321 322 45 46 47 48 49 410 411 412 413 414 415 416 417 419 420 421 422 56 57 58 59 510 511 512 513 514 515 516 517 519 520 521 522 67 68 69 610 611 612 613 614 615 616 617 619 620 621 622 78 79 710 711 712 713 714 715 716 717 719 720 721 722 89 810 811 812 813 814 815 816 817 819 820 821 822 910 911 912 913 914 915 916 917 919 920 921 922 1011 1012 1013 1014 1015 1016 1017 1019 1020 1021 1022 1112 1113 1114 1115 1116 1117 1119 1120 1121 1122 1213 1214 1215 1216 1217 1219 1220 1221 1222 1314 1315 1316 1317 1319 1320 1321 1322 1415 1416 1417 1419 1420 1421 1422 1516 1517 1519 1520 1521 1522 1617 1619 1620 1621 1622 1719 1720 1721 1722 1920 1921 1922 2021 2022 2122")

	res = append(res, prod1...)
	res = append(res, prod2...)
	res = append(res, prod3...)

	return res
}
