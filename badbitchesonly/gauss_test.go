package main

import (
	"fmt"
	//"log"
	"reflect"
	"testing"
)

func TestPrint42(t *testing.T) {
	plaintext := 42
	fmt.Printf("%d", plaintext)
}

func TestWhatDoesGaussPartialOutput(t *testing.T) { // i do not understand what partial output does
	res, err := GaussPartial(tc.a, tc.b)

	fmt.Printf("This is the input: \n %d \n", tc)
	fmt.Printf("This is the result: \n %d \n", res)
	fmt.Printf("This is the error: %d", err)

}

func TestMakeAugmentedMatrix(t *testing.T) {

	A := [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	b := []int{5, 6, 7}

	augMa := makeAugmentedMatrix(A, b)

	fmt.Printf("This is the augmented matrix: %d \n", augMa)

	shouldBe := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 6}, {1, 2, 3, 4, 7}}

	if !reflect.DeepEqual(augMa, shouldBe) {
		t.Log("The augmented matrix is wrong")
		t.Fail()
	}

}

func TestGaussEliminationPart2(t *testing.T) {
	augMa := make([][]int, 4)

	augMa[0] = []int{1, 1, 1, 0, 0, 1}
	augMa[1] = []int{1, 1, 0, 1, 0, 1}
	augMa[2] = []int{1, 0, 1, 1, 0, 0}
	augMa[3] = []int{0, 1, 1, 1, 0, 1}

	res := gaussEliminationPart2(augMa)
	fmt.Printf("This is the result of the Gauss elimination: %d \n", res.TempRes)

	shouldBe := make([][]int, 4)
	shouldBe[0] = []int{1, 1, 1, 0, 0, 1}
	shouldBe[1] = []int{0, 1, 0, 1, 0, 1}
	shouldBe[2] = []int{0, 0, 1, 1, 0, 0}
	shouldBe[3] = []int{0, 0, 0, 1, 0, 0}

	if !reflect.DeepEqual(res.TempRes, shouldBe) {
		t.Log("The result of the Gauss elimination wrong")
		t.Fail()
	}

	if reflect.DeepEqual(res.TempRes, shouldBe) {
		fmt.Print("The result of the Gauss elimination was correct")
	}
}

func TestBackSubstitutionBinary(t *testing.T) {
	augMatrix := make([][]int, 5)
	augMatrix[0] = []int{1, 1, 1, 0, 0, 1}
	augMatrix[1] = []int{0, 1, 0, 1, 0, 1}
	augMatrix[2] = []int{0, 0, 1, 1, 0, 0}
	augMatrix[3] = []int{0, 0, 0, 1, 0, 0}
	augMatrix[4] = []int{0, 0, 0, 0, 1, 1}

	//shouldBe := make([]int, 4)
	shouldBe := []int{0, 1, 0, 0}

	res := GaussRes{
		ResType: "Valid",
		TempRes: augMatrix}

	res = backSubstitution(res)
	fmt.Printf("This is the result of the backSubstitution: %d \n", res.Solved)

	if !reflect.DeepEqual(res.Solved, shouldBe) {
		t.Log("The result of the backSubstitution is wrong\n")
		t.Fail()
	} else {
		fmt.Print("The result of the backSubstitution was correct\n")
	}

	augMatrix = make([][]int, 5)
	augMatrix[0] = []int{1, 1, 1, 0, 0, 1}
	augMatrix[1] = []int{0, 1, 0, 1, 0, 1}
	augMatrix[2] = []int{0, 0, 1, 1, 1, 0}
	augMatrix[3] = []int{0, 0, 0, 1, 0, 0}
	augMatrix[4] = []int{0, 0, 0, 0, 1, 1}

	shouldBe = []int{1, 1, 1, 0}

	res = GaussRes{
		ResType: "Valid",
		TempRes: augMatrix}

	res = backSubstitution(res)
	fmt.Printf("This is the second result of the backSubstitution: %d \n", res.Solved)
	if !reflect.DeepEqual(res.Solved, shouldBe) {
		t.Log("The second result of the backSubstitution is wrong \n")
		t.Fail()
	} else {
		fmt.Print("The second result of the backSubstitution was correct \n")
	}

}

func TestBackSubstitutionConstant(t *testing.T) {
	augMatrix := make([][]int, 4)
	augMatrix[0] = []int{1, 1, 1, 0, 0, 1}
	augMatrix[1] = []int{0, 1, 0, 1, 0, 1}
	augMatrix[2] = []int{0, 0, 1, 1, 1, 0}
	augMatrix[3] = []int{0, 0, 0, 1, 0, 0}

	//shouldBe := make([]int, 4)
	shouldBe := []int{1, 1, 1, 0}

	res := GaussRes{
		ResType: "Valid",
		TempRes: augMatrix}

	res = backSubstitution(res)
	fmt.Printf("This is the result of the backSubstitution: %d \n", res.Solved)

	if !reflect.DeepEqual(res.Solved, shouldBe) {
		t.Log("The result of the backSubstitution is wrong")
		t.Fail()
	}

	if reflect.DeepEqual(res.Solved, shouldBe) {
		t.Log("The result of the backSubstitution was correct")

	}
}

func TestBackSubstitution2(t *testing.T) {
	a := make([][]int, 5)
	a[0] = []int{1, 0, 1, 0, 1, 1}
	a[1] = []int{0, 1, 1, 1, 0, 0}
	a[2] = []int{0, 0, 1, 0, 1, 1}
	a[3] = []int{0, 0, 0, 1, 1, 0}
	a[4] = []int{0, 0, 0, 0, 1, 1}

	shouldBe := []int{0, 1, 0, 1}

	res := GaussRes{
		ResType: "Valid",
		TempRes: a}

	res = backSubstitution(res)
	fmt.Printf("res is: %d \n", res.Solved)

	if !reflect.DeepEqual(res.Solved, shouldBe) {
		t.Log("Not correct :(")
		t.Fail()
	}

}

func TestGaussEliminationPart2_2(t *testing.T) {
	a := make([][]int, 4)
	a[0] = []int{1, 0, 1, 1, 0}
	a[1] = []int{1, 1, 0, 1, 1}
	a[2] = []int{1, 0, 0, 1, 0}
	a[3] = []int{0, 0, 1, 1, 1}

	shouldBe1 := make([][]int, 4)
	shouldBe1[0] = []int{1, 0, 1, 1, 0}
	shouldBe1[1] = []int{0, 1, 1, 0, 1}
	shouldBe1[2] = []int{0, 0, 1, 0, 0}
	shouldBe1[3] = []int{0, 0, 0, 1, 1}

	res := gaussEliminationPart2(a)
	fmt.Printf("res1 is:      %d \n", res.TempRes)
	fmt.Printf("shouldBe1 is: %d \n", shouldBe1)
	fmt.Printf("Res Type is:  %v \n", res.ResType)

	if !reflect.DeepEqual(res.TempRes, shouldBe1) {
		t.Log("Not correct :(")
		t.Fail()
	}

	b := make([][]int, 4)
	b[0] = []int{1, 0, 1, 0, 0}
	b[1] = []int{1, 1, 0, 1, 1}
	b[2] = []int{1, 0, 0, 1, 0}
	b[3] = []int{0, 0, 1, 0, 1}

	shouldBe2 := make([][]int, 4)
	shouldBe2[0] = []int{1, 0, 1, 0, 0}
	shouldBe2[1] = []int{0, 1, 1, 1, 1}
	shouldBe2[2] = []int{0, 0, 1, 1, 0}
	shouldBe2[3] = []int{0, 0, 0, 1, 1}

	res2 := gaussEliminationPart2(b)
	fmt.Printf("res2 is:      %d \n", res2.TempRes)
	fmt.Printf("shouldBe2 is: %d \n", shouldBe2)
	fmt.Printf("Res Type is:  %v \n", res2.ResType)

	if !reflect.DeepEqual(res2.TempRes, shouldBe2) {
		t.Log("Not correct :(")
		t.Fail()
	}
}

func TestGaussEliminationReturnsError(t *testing.T) {
	matrix := make([][]int, 6)

	matrix[0] = []int{0, 1, 1, 1, 0, 1, 0}
	matrix[1] = []int{1, 0, 0, 0, 1, 0, 1}
	matrix[2] = []int{0, 1, 1, 1, 1, 0, 1}
	matrix[3] = []int{0, 0, 0, 0, 1, 0, 1}
	matrix[4] = []int{0, 1, 1, 0, 1, 0, 0}
	matrix[5] = []int{0, 1, 0, 1, 1, 0, 1}

	res := gaussEliminationPart2(matrix)

	shouldBe := make([][]int, 6)
	shouldBe[0] = []int{1, 0, 0, 0, 1, 0, 1}
	shouldBe[1] = []int{0, 1, 1, 1, 0, 1, 0}
	shouldBe[2] = []int{0, 0, 1, 0, 1, 1, 1}
	shouldBe[3] = []int{0, 0, 0, 1, 1, 1, 0}
	shouldBe[4] = []int{0, 0, 0, 0, 1, 0, 1}
	shouldBe[5] = []int{0, 0, 0, 0, 0, 1, 0}

	fmt.Printf("This is res:	   %d \n", res.TempRes)
	fmt.Printf("This is should be: %d \n", shouldBe)
	fmt.Printf("This is restype: %v \n", res.ResType)
	if !reflect.DeepEqual(res.ResType, Error) {
		t.Log("The result of the gauss elimination is wrong")
		t.Fail()
	}

}

func TestGaussEliminationDep(t *testing.T) {
	matrix := make([][]int, 4)

	matrix[0] = []int{1, 0, 1, 0, 0, 1}
	matrix[1] = []int{1, 1, 1, 0, 0, 1}
	matrix[2] = []int{1, 0, 1, 0, 0, 1}
	matrix[3] = []int{1, 0, 1, 1, 0, 1}

	res := gaussEliminationPart2(matrix)

	shouldBe := make([][]int, 4)
	shouldBe[0] = []int{1, 0, 1, 0, 0, 1}
	shouldBe[1] = []int{0, 1, 0, 0, 0, 0}
	shouldBe[2] = []int{0, 0, 0, 0, 0, 0}
	shouldBe[3] = []int{0, 0, 0, 1, 0, 0}

	fmt.Printf("This is res:	   %d \n", res.TempRes)
	fmt.Printf("This is should be: %d \n", shouldBe)
	fmt.Printf("This is restype: %v \n", res.ResType)
	if !reflect.DeepEqual(res.ResType, Error) {
		t.Log("The result of the gauss elimination is wrong")
		t.Fail()
	}
}

func TestGaussEliminationFreeVar(t *testing.T) {
	matrix := make([][]int, 5)

	matrix[0] = []int{1, 0, 1, 0, 0, 1}
	matrix[1] = []int{0, 1, 1, 0, 0, 1}
	matrix[2] = []int{0, 0, 1, 0, 0, 1}
	matrix[3] = []int{0, 0, 0, 0, 0, 0}
	matrix[4] = []int{0, 0, 0, 0, 0, 0}

	res := gaussEliminationPart2(matrix)

	fmt.Printf("This is res:	   %d \n", res.TempRes)
	fmt.Printf("This is restype: %v \n", res.ResType)
	if !reflect.DeepEqual(res.ResType, EmptyCol) {
		t.Log("The result of the gauss elimination is wrong 1")
		t.Fail()
	}

	shouldBeDepVar := []int{3}
	fmt.Printf("This is shouldBeFreeVar:	   %d \n", res.ColNo)
	if !reflect.DeepEqual(res.ColNo, shouldBeDepVar) {
		t.Log("The result of the gauss elimination is wrong")
		t.Fail()
	}
}

func TestGaussEliminationFreeVar_2(t *testing.T) {
	matrix := make([][]int, 5)

	matrix[0] = []int{1, 0, 0, 0, 0, 1}
	matrix[1] = []int{0, 1, 0, 0, 0, 1}
	matrix[2] = []int{0, 0, 0, 1, 0, 1}
	matrix[3] = []int{0, 0, 0, 1, 0, 1}
	matrix[4] = []int{0, 1, 0, 1, 0, 0}

	res := gaussEliminationPart2(matrix)

	fmt.Printf("This is res:	   %d \n", res.TempRes)
	fmt.Printf("This is restype: %v \n", res.ResType)
	if !reflect.DeepEqual(res.ResType, "EmptyCol") {
		t.Log("The ResType does not match")
		t.Fail()
	}

	shouldBeFreeVar := []int{2}
	fmt.Printf("This is shouldBeFreeVar:	   %d \n", res.ColNo)
	if !reflect.DeepEqual(res.ColNo, shouldBeFreeVar) {
		t.Log("The correct FreeVar is not sat")
		t.Fail()
	}

	shouldBe := make([][]int, 5)
	shouldBe[0] = []int{1, 0, 0, 0, 0, 1}
	shouldBe[1] = []int{0, 1, 0, 0, 0, 1}
	shouldBe[2] = []int{0, 0, 0, 0, 0, 0}
	shouldBe[3] = []int{0, 0, 0, 1, 0, 1}
	shouldBe[4] = []int{0, 0, 0, 0, 0, 0}

	if !reflect.DeepEqual(res.TempRes, shouldBe) {
		t.Log("Should be does not match the result")
		t.Fail()
	}
}

func TestGaussEliminationFreeVar_3(t *testing.T) {
	matrix := make([][]int, 5)

	matrix[0] = []int{1, 0, 0, 0, 0, 1}
	matrix[1] = []int{0, 1, 0, 0, 0, 1}
	matrix[2] = []int{0, 0, 0, 1, 0, 1}
	matrix[3] = []int{0, 0, 0, 1, 0, 0}
	matrix[4] = []int{0, 1, 0, 1, 0, 0}

	res := gaussEliminationPart2(matrix)

	if !reflect.DeepEqual(res.ResType, Error) {
		t.Log("The ResType does not match")
		t.Fail()
	}
}

func TestHandleEmptyCol(t *testing.T) {
	matrix := make([][]int, 6)

	matrix[0] = []int{1, 0, 0, 0, 0, 0, 0, 1}
	matrix[1] = []int{0, 1, 0, 0, 0, 0, 0, 0}
	matrix[2] = []int{0, 0, 1, 0, 0, 0, 0, 1}
	matrix[3] = []int{0, 0, 0, 0, 0, 0, 0, 0}
	matrix[4] = []int{0, 0, 0, 0, 1, 0, 0, 0}
	matrix[5] = []int{0, 0, 0, 0, 0, 1, 0, 1}

	res := gaussEliminationPart2(matrix)
	res = backSubstitution(res)

	fmt.Printf("This is restype: %v \n", res.ResType)
	if !reflect.DeepEqual(res.ResType, Multi) {
		t.Log("The ResType does not match")
		t.Fail()
	}

	shouldBe := make([][]int, 2)
	shouldBe[0] = []int{1, 0, 1, 0, 0, 1}
	shouldBe[1] = []int{1, 0, 1, 1, 0, 1}

	fmt.Printf("This is Multi[0]:	   %d \n", res.Multi[0])
	fmt.Printf("This is Multi[1]:	   %d \n", res.Multi[1])

	if !reflect.DeepEqual(res.Multi[0], shouldBe[0]) {
		t.Log("Should be [0] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[1], shouldBe[1]) {
		t.Log("Should be [1] does not match the result")
		t.Fail()
	}
}

func TestHandleEmptyCol_2(t *testing.T) {
	matrix := make([][]int, 6)

	matrix[0] = []int{1, 0, 0, 0, 0, 0, 0, 1}
	matrix[1] = []int{0, 0, 0, 0, 0, 0, 0, 0}
	matrix[2] = []int{0, 0, 1, 0, 0, 0, 0, 1}
	matrix[3] = []int{0, 0, 0, 0, 0, 0, 0, 0}
	matrix[4] = []int{0, 0, 0, 0, 1, 0, 0, 0}
	matrix[5] = []int{0, 0, 0, 0, 0, 1, 0, 1}

	res := gaussEliminationPart2(matrix)
	res = backSubstitution(res)

	fmt.Printf("This is restype: %v \n", res.ResType)
	if !reflect.DeepEqual(res.ResType, Multi) {
		t.Log("The ResType does not match")
		t.Fail()
	}

	shouldBe := make([][]int, 4)
	shouldBe[0] = []int{1, 0, 1, 0, 0, 1}
	shouldBe[1] = []int{1, 1, 1, 0, 0, 1}
	shouldBe[2] = []int{1, 0, 1, 1, 0, 1}
	shouldBe[3] = []int{1, 1, 1, 1, 0, 1}

	if !reflect.DeepEqual(res.Multi[0], shouldBe[0]) {
		t.Log("Should be [0] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[1], shouldBe[1]) {
		t.Log("Should be [1] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[2], shouldBe[2]) {
		t.Log("Should be [2] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[3], shouldBe[3]) {
		t.Log("Should be [3] does not match the result")
		t.Fail()
	}
}

func TestHandleEmptyCol_3(t *testing.T) {
	matrix := make([][]int, 6)

	matrix[0] = []int{1, 0, 0, 0, 0, 0, 0, 1}
	matrix[1] = []int{0, 0, 0, 0, 0, 0, 0, 0}
	matrix[2] = []int{0, 0, 1, 0, 0, 0, 0, 1}
	matrix[3] = []int{0, 0, 0, 0, 0, 0, 0, 0}
	matrix[4] = []int{0, 0, 0, 0, 0, 0, 0, 0}
	matrix[5] = []int{0, 0, 0, 0, 0, 1, 0, 1}

	res := gaussEliminationPart2(matrix)
	res = backSubstitution(res)

	fmt.Printf("This is restype: %v \n", res.ResType)
	if !reflect.DeepEqual(res.ResType, Multi) {
		t.Log("The ResType does not match")
		t.Fail()
	}

	shouldBe := make([][]int, 8)
	shouldBe[0] = []int{1, 0, 1, 0, 0, 1}
	shouldBe[1] = []int{1, 1, 1, 0, 0, 1}
	shouldBe[2] = []int{1, 0, 1, 1, 0, 1}
	shouldBe[3] = []int{1, 1, 1, 1, 0, 1}
	shouldBe[4] = []int{1, 0, 1, 0, 1, 1}
	shouldBe[5] = []int{1, 1, 1, 0, 1, 1}
	shouldBe[6] = []int{1, 0, 1, 1, 1, 1}
	shouldBe[7] = []int{1, 1, 1, 1, 1, 1}

	if !reflect.DeepEqual(res.Multi[0], shouldBe[0]) {
		t.Log("Should be [0] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[1], shouldBe[1]) {
		t.Log("Should be [1] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[2], shouldBe[2]) {
		t.Log("Should be [2] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[3], shouldBe[3]) {
		t.Log("Should be [3] does not match the result")
		t.Fail()
	}
	if !reflect.DeepEqual(res.Multi[4], shouldBe[4]) {
		t.Log("Should be [0] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[5], shouldBe[5]) {
		t.Log("Should be [1] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[6], shouldBe[6]) {
		t.Log("Should be [2] does not match the result")
		t.Fail()
	}

	if !reflect.DeepEqual(res.Multi[7], shouldBe[7]) {
		t.Log("Should be [3] does not match the result")
		t.Fail()
	}
}

func TestHandleFreeVar(t *testing.T) {
	bit := MakeBitSlice(3)
	for i, _ := range bit {
		fmt.Printf("Bit string: %d \n", bit[i])
	}
}

// OLD Tests

// func TestBackSubstitution(t *testing.T) {
// 	augMa := make([][]int, 3)

// 	augMa[0] = []int{-3, 2, -1, -1}
// 	augMa[1] = []int{0, -2, 5, -9}
// 	augMa[2] = []int{0, 0, -2, 2}

// 	// fmt.Printf("augMa[0]: %d \n", augMa[0])
// 	// fmt.Printf("augMa[:][0]: %d \n", augMa[:][0])
// 	// fmt.Printf("augMa: %d \n", augMa)
// 	// fmt.Printf("length of augMa: %d \n", len(augMa))

// 	res := backSubstitution(augMa)

// 	fmt.Printf("This is the result: %d \n", res)

// 	shouldBe := []int{2, 2, -1}

// 	if !reflect.DeepEqual(res, shouldBe) {
// 		t.Log("The result of the back substitution is wrong")
// 		t.Fail()
// 	}

// }

// func TestGaussElimination(t *testing.T) {
// 	augMa := make([][]int, 3)

// 	augMa[0] = []int{-3, 2, -1, -1}
// 	augMa[1] = []int{6, -6, 7, -7}
// 	augMa[2] = []int{3, -4, 4, -6}

// 	res := gaussElimination(augMa)
// 	fmt.Printf("This is the result of the Gauss elimination: %d \n", res)

// 	shouldBe := make([][]int, 3)
// 	shouldBe[0] = []int{-3, 2, -1, -1}
// 	shouldBe[1] = []int{0, -2, 5, -9}
// 	shouldBe[2] = []int{0, 0, -2, 2}

// 	// res2 := backSubstitution(res)
// 	// fmt.Printf("This is the result after back substitution: %d \n", res2)

// 	if !reflect.DeepEqual(res, shouldBe) {
// 		t.Log("The result of the Gauss elimination wrong")
// 		t.Fail()
// 	}
// }

// func TestWholeGaussElimination(t *testing.T) {

// 	A := [][]int{{-3, 2, -1}, {6, -6, 7}, {3, -4, 4}}

// 	b := []int{-1, -7, -6}

// 	res := solveByGaussElimination(A, b)
// 	// fmt.Printf("This is res: %d \n", res)

// 	shouldBe := []int{2, 2, -1} // ifølge youtube video

// 	if !reflect.DeepEqual(res, shouldBe) {
// 		t.Log("The result is wrong")
// 		t.Fail()
// 	}

// }
