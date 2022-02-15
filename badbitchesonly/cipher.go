package badbitchesonly

import (
	_ "fmt"
)

type Register struct {
	Number      int
	Length      int
	ArrImposter []int
	Tabs        []int
}

func makeR1() Register {
	r1 := Register{Number: 1,
		Length:      19,
		ArrImposter: make([]int, 19),
		Tabs:        []int{18, 17, 16, 13}}
	return r1
}

func makeR2() Register {
	r2 := Register{Number: 2,
		Length:      22,
		ArrImposter: make([]int, 22),
		Tabs:        []int{21, 20}}
	return r2
}

func makeR3() Register {
	r3 := Register{Number: 3,
		Length:      23,
		ArrImposter: make([]int, 23),
		Tabs:        []int{22, 21, 20, 7}}
	return r3
}

func makeR4() Register {
	r4 := Register{Number: 4,
		Length:      17,
		ArrImposter: make([]int, 17),
		Tabs:        []int{16, 11}}
	return r4
}

// Returns the majority bit of input x, y, z
func majority(x int, y int, z int) int {
	if x+y+z >= 2 {
		return 1
	} else {
		return 0
	}
}

func clockingUnit(r4 Register) int {
	arr := r4.ArrImposter
	maj := majority(arr[3], arr[7], arr[10])
	return maj
}

func initializeRegisters(session_key []int, frame_number []int) {
	/* do A5/2 */
	r1 := makeR1()
	r2 := makeR2()
	r3 := makeR3()
	r4 := makeR4()

	for i := 0; i < 63; i++ {
		/* Clock all registers */
		r1.ArrImposter[0] = ^session_key[i]
		r2.ArrImposter[0] = ^session_key[i]
		r3.ArrImposter[0] = ^session_key[i]
		r4.ArrImposter[0] = ^session_key[i]
	}

	for i := 0; i < 21; i++ {
		/* Clock all registers */
		r1.ArrImposter[0] = ^frame_number[i]
		r2.ArrImposter[0] = ^frame_number[i]
		r3.ArrImposter[0] = ^frame_number[i]
		r4.ArrImposter[0] = ^frame_number[i]
	}
}

/* Should we give frame number as a param ? */
func makeKeyStream() {
	/* Initialiser Registers*/

	/* Initialize internal state with K_c and frame number */

	/* Force bits R1[15], R2[16], R3[18], R4[10] to be 1 */

	/* Run A5/2 for 99 clocks and ignore output */

	/* Run A5/2 for 228 clocks and use outputs as key-stream */

}
