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

// func TestGauss(t *testing.T) {
// 	var tc = testCase{
// 		a: [][]int{
// 			{1, 0, 0, 1, 0, 0},
// 			{1, 0, 0, 0, 0, 0},
// 			{1, 1, 1, 1, 0, 1},
// 			{0, 1, 0, 0, 1, 0},
// 			{0, 1, 1, 1, 0, 0},
// 			{0, 1, 1, 0, 1, 0}},
// 		b: []int{1, 0, 0, 0, 0, 0},
// 		x: []int{1, 1, 0, 1, 0, 1},
// 	}

// 	x, err := GaussPartial(tc.a, tc.b)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(x)
// 	for i, xi := range x {
// 		if Abs(tc.x[i]-xi) > ε {
// 			log.Println("out of tolerance")
// 			log.Fatal("expected", tc.x)
// 		}
// 	}
// }

func TestWhatDoesGaussPartialOutput(t *testing.T) { // i do not understand what partial output does
	res, err := GaussPartial(tc.a, tc.b)

	fmt.Printf("This is the input: \n %d \n", tc)
	fmt.Printf("This is the result: \n %d \n", res)
	fmt.Printf("This is the error: %d", err)

}

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

func TestGaussEliminationPart2(t *testing.T) {
	augMa := make([][]int, 4)

	augMa[0] = []int{1, 1, 1, 0, 0, 1}
	augMa[1] = []int{1, 1, 0, 1, 0, 1}
	augMa[2] = []int{1, 0, 1, 1, 0, 0}
	augMa[3] = []int{0, 1, 1, 1, 0, 1}

	res := gaussEliminationPart2(augMa)
	fmt.Printf("This is the result of the Gauss elimination: %d \n", res)

	shouldBe := make([][]int, 4)
	shouldBe[0] = []int{1, 1, 1, 0, 0, 1}
	shouldBe[1] = []int{0, 1, 0, 1, 0, 1}
	shouldBe[2] = []int{0, 0, 1, 1, 0, 0}
	shouldBe[3] = []int{0, 0, 0, 1, 0, 0}

	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the Gauss elimination wrong")
		t.Fail()
	}

	if reflect.DeepEqual(res, shouldBe) {
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

	res := backSubstitution(augMatrix)
	fmt.Printf("This is the result of the backSubstitution: %d \n", res)

	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the backSubstitution is wrong\n")
		t.Fail()
	}

	if reflect.DeepEqual(res, shouldBe) {
		fmt.Print("The result of the backSubstitution was correct\n")
	}

	augMatrix = make([][]int, 5)
	augMatrix[0] = []int{1, 1, 1, 0, 0, 1}
	augMatrix[1] = []int{0, 1, 0, 1, 0, 1}
	augMatrix[2] = []int{0, 0, 1, 1, 1, 0}
	augMatrix[3] = []int{0, 0, 0, 1, 0, 0}
	augMatrix[4] = []int{0, 0, 0, 0, 1, 1}

	//shouldBe := make([]int, 4)
	shouldBe = []int{1, 1, 1, 0}

	res = backSubstitution(augMatrix)
	fmt.Printf("This is the second result of the backSubstitution: %d \n", res)

	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The second result of the backSubstitution is wrong \n")
		t.Fail()
	}

	if reflect.DeepEqual(res, shouldBe) {
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
	shouldBe := []int{0, 1, 1, 0}

	res := backSubstitution(augMatrix)
	fmt.Printf("This is the result of the backSubstitution: %d \n", res)

	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the backSubstitution is wrong")
		t.Fail()
	}

	if reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the backSubstitution was correct")

	}
}

func TestGaussElimOnSecondExample(t *testing.T) {
	//FIXME this test is wrong as there is a 1=0 in shouldbe for bit entry
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

	fmt.Printf("This is res:	   %d \n", res)
	fmt.Printf("This is should be: %d \n", shouldBe)
	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the gauss elimination is wrong")
		t.Fail()
	}

	// TODO: check if backsub is working here also
	// res should be [0,1,0,1,1,0]
	//shouldBe := make([]int, 4)
	regShouldBe := []int{1, 0, 0, 1, 0}

	resBack := backSubstitution(res)
	fmt.Printf("This is the result of the backSubstitution: %d \n", resBack)

	// gets a row with [0,0,0,0,0,1,0] might be caught by error-gauss-stuff
	if !reflect.DeepEqual(resBack, regShouldBe) {
		t.Log("The result of the backSubstitution is wrong")
		t.Fail()
	}

}

// func TestGaussElimination(t *testing.T) {
// 	matrix := make([][]int, 4)

// 	matrix[0] = []int{1, 0, 1, 1}
// 	matrix[1] = []int{1, 0, 1, 0}
// 	matrix[2] = []int{0, 1, 0, 1}
// 	matrix[3] = []int{0, 1, 1, 1}

// 	res := gaussEliminationPart2(matrix)

// 	fmt.Printf("res is: \n %d", res)

// 	// This should fail, as the last row is [0, 0, 0, 1]
// 	// we need to make sure this says its wrong

// }
