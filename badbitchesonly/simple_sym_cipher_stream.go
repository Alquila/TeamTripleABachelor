package main

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
