package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"
)

func TestCiphertextOnlyAttackWithNoPrints(t *testing.T) {
	start := time.Now()

	r4Found := make([][]int, 0)
	r4Guess := make([]int, 17)

	MakeSessionKey()
	originalFrameNumber, currentFrameNumber = 42, 42
	r4Bin, binKey, keyForTest := MakeRealKeyStreamSixFrames(originalFrameNumber)

	/* calculate ciphertext */
	c := CalculateXFrameCiphertext(binKey, 6)
	cForTest := CalculateXFrameCiphertext(keyForTest, 2)
	println(c[0] + cForTest[0])

	/* Calculate KG*C */
	KG := CreateKgMatrix()
	KG_C := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[:456])))
	KG_C2 := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[456:912])))
	KG_C3 := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[912:])))
	full_KGC := append(KG_C, KG_C2...)
	full_KGC = append(full_KGC, KG_C3...)

	guesses := int(math.Pow(2, 16))

	for i := 0; i < guesses; i++ {
		originalFrameNumber = 42 //reset the framenumber for the symbolic version
		currentFrameNumber = 42

		r4Guess = MakeR4Guess(i)
		r4Guess = PutConstantBackInRes(r4Guess, 10)

		symKeyStream := make([][]int, 0)

		for i := 0; i < 6; i++ { //Make six frame long sym-keystream for the guess
			frame_influenced_bits := SimulateClockingR4WithFrameDifference(originalFrameNumber, currentFrameNumber)
			r4.RegSlice = XorSlice(frame_influenced_bits, r4Guess)
			r4.RegSlice[10] = 1
			key1 := MakeSymKeyStream() //this clocks sr4 which has r4_guess as its array
			symKeyStream = append(symKeyStream, key1...)
			currentFrameNumber++
		}

		/* Multiply KG with the SymbolicKeyStream to make KGK */
		KGk := CalculateKgTimesSymKeyStream(KG, symKeyStream[:456])
		KGk2 := CalculateKgTimesSymKeyStream(KG, symKeyStream[456:912])
		KGk3 := CalculateKgTimesSymKeyStream(KG, symKeyStream[912:])
		full_KGk := append(KGk, KGk2...)
		full_KGk = append(full_KGk, KGk3...)

		x := SolveByGaussElimination(full_KGk, full_KGC)

		if x.ResType == Multi {
			for i := 0; i < len(x.Multi); i++ {
				if VerifyKeyStream(x.Multi[i]) {
					r4Found = append(r4Found, r4Guess)
				}
			}
		}
		if x.ResType == Error {
			continue
		}

	}

	fmt.Printf("This is the value of the found R4 Register: %v\n", r4Found[0])
	fmt.Printf("This is r4_bin: %v\n", r4Bin)

	fmt.Println("Have we found the right r4?")
	if reflect.DeepEqual(r4Bin, r4Found[0]) {
		fmt.Println("Fuck yes we found it gutterne")
	} else {
		fmt.Println("RIP we dit not manage to find the correct R4")
	}

	executionTime := time.Since(start)
	fmt.Printf("Ciphertext-only Attack took: %s", executionTime)
}

func TestKnownPlaintextAttackWithNoPrints(t *testing.T) {
	start := time.Now()

	r4Found := make([][]int, 0) // append results to this boi
	r4Guess := make([]int, 17)

	sessionKey = make([]int, 64) //all zero session key
	originalFrameNumber = 42
	r4Real, realKey, r4ForTest := MakeRealKeyStreamFourFrames(originalFrameNumber)

	guesses := int(math.Pow(2, 16))
	println(guesses)
	for i := 0; i < guesses; i++ {
		originalFrameNumber = 42 // reset the frame number for the symbolic version
		currentFrameNumber = 42

		r4Guess = MakeR4Guess(i) // for all possible value of r4 we need three frames
		r4Guess = PutConstantBackInRes(r4Guess, 10)

		// do this such that r4 guess can be copied into sr4 in SymSetRegisters()
		r4 = MakeR4()
		copy(r4.RegSlice, r4Guess)
		key1 := MakeSymKeyStream() // this clocks sr4 which has r4Guess as its array

		currentFrameNumber++

		// update r4Guess with new frame value
		r4 = MakeR4()
		copy(r4.RegSlice, r4Guess)

		frameInfluencedBits := SimulateClockingR4WithFrameDifference(originalFrameNumber, currentFrameNumber)
		r4.RegSlice = XorSlice(frameInfluencedBits, r4Guess)

		r4.RegSlice[10] = 1

		key2 := MakeSymKeyStream() //this will now copy the updated r4_regSlice into sr4

		currentFrameNumber++
		r4 = MakeR4()
		fakeR4 := MakeR4()
		copy(r4.RegSlice, r4Guess)
		diff := FindDifferenceOfFrameNumbers(originalFrameNumber, currentFrameNumber)

		//fakeR4.RegSlice clockes således at det er [...1..] de steder hvor diff påvirker indgangene
		for i := 0; i < 22; i++ {
			Clock(fakeR4)
			fakeR4.RegSlice[0] = fakeR4.RegSlice[0] ^ diff[i]
		}

		r4.RegSlice = XorSlice(fakeR4.RegSlice, r4.RegSlice)
		r4.RegSlice[10] = 1
		key3 := MakeSymKeyStream()
		currentFrameNumber++

		key := append(key1, key2...)
		key = append(key, key3...)

		// this returns a gauss struct
		gauss := SolveByGaussElimination(key, realKey)

		if gauss.ResType == Error {
			continue
		} else if gauss.ResType == Multi {
			fmt.Printf("found multi in %d of lenght %d \n", i, len(gauss.Multi))
			for i := 0; i < len(gauss.Multi); i++ {
				if VerifyKeyStream(gauss.Multi[i]) {
					// If a solution is verified it is added to a list of verified guesses
					r4Found = append(r4Found, r4Guess)
				}
			}
		}
	}

	// 'trial encryptions'
	correctR4 := make([]int, len(r4Guess))
	numberOfValidR4 := len(r4Found)
	if numberOfValidR4 <= 0 {
		fmt.Printf("No solutions were found \n")
	} else if numberOfValidR4 > 1 {
		// we have multiple plausible solutions to r4
		for i := 0; i > numberOfValidR4; i++ {
			r4.RegSlice = r4Found[i]
			// We make an extra keystream with frame number: base frame + 4
			currentFrameNumber = originalFrameNumber + 4
			ks := MakeKeyStream()
			if reflect.DeepEqual(ks, r4ForTest) {
				fmt.Printf("This is the right one: %d\n", r4ForTest)
				correctR4 = r4Found[i]
				break
			}
		}
	} else {
		correctR4 = r4Found[0]
	}

	fmt.Printf("This is original r4:       %d\n", r4Real)
	for i := range r4Found {
		fmt.Printf("This is %d'th found r4:    %d\n", i, r4Found[i])
	}
	fmt.Println("Have we found the right r4?")
	if reflect.DeepEqual(correctR4, r4Real) {
		fmt.Println("Fuck yes we found it gutterne")
	} else {
		fmt.Println("RIP we dit not")
	}

	executionTime := time.Since(start)
	fmt.Printf("Known Plaintext Attack took: %s", executionTime)
}
