package main

func RunA5_2() ([]int, [][]int) {
	keyStreamSym := make([][]int, 228)

	keyStream := make([]int, 228)

	/* Run A5/2 for 99 clocks and ignore output  */
	for i := 0; i < 99; i++ {
		// do the clock thingy and ignore
		ClockingUnit(r4)
		SymClockingUnit(r4)
		Clock(r4)
	}

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		// do the clock thingy and output
		ClockingUnit(r4)
		SymClockingUnit(r4)
		Clock(r4)
		keyStream[i] = MakeFinalXOR()
		keyStreamSym[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}

	return keyStream, keyStreamSym
}

//Makes a r1 symregister with all 0's
func InitOneSymRegister() SymRegister {
	reg := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15) // equivalent to reg1
	return reg
}

func SimpleKeyStreamSym(r SymRegister) [][]int {

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
		keyStream[i] = r.RegSlice[r.Length-1]
		if i == 200 {
		}
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
		OverwriteXorSlice(r.RegSlice[r.Length-1], aaa)
		keyStream[i] = aaa
	}
	return keyStream
}
