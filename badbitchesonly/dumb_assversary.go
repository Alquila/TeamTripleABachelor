package main

import (
	"fmt"
	"reflect"
)

//"fmt"

func idk() int {
	return 42
}

func doTheSimpleHack() {
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

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res, orgReg) {
		fmt.Printf("This is fucking wrong\n")
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("reg er: %d\n", reg.ArrImposter)
	}
}

// func Solve(A, k) {
// 	// should solve Ax = k
// 	// hopefully som go library has this func
// 	42
// 	// brug gauss
// }
