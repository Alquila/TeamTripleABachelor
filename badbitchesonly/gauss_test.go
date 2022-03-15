package main

import (
	"fmt"
	"log"
	"testing"
	"reflect"
)

func TestPrint42(t *testing.T) {
	plaintext := 42
	fmt.Printf("%d", plaintext)
}

func TestGauss(t *testing.T) {
	var tc = testCase{
		a: [][]int{
			{1, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 1},
			{0, 1, 0, 0, 1, 0},
			{0, 1, 1, 1, 0, 0},
			{0, 1, 1, 0, 1, 0}},
		b: []int{1, 0, 0, 0, 0, 0},
		x: []int{1, 1, 0, 1, 0, 1},
	}

	x, err := GaussPartial(tc.a, tc.b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
	for i, xi := range x {
		if Abs(tc.x[i]-xi) > ε {
			log.Println("out of tolerance")
			log.Fatal("expected", tc.x)
		}
	}
}

func TestWhatDoesGaussPartialOutput(t *testing.T) {		// i do not understand what partial output does
	res, err := GaussPartial(tc.a, tc.b)

	fmt.Printf("This is the input: \n %d \n", tc)
	fmt.Printf("This is the result: \n %d \n", res)
	fmt.Printf("This is the error: %d", err)

}

func TestBackSubstitution(t *testing.T) {
	augMa := make([][]int, 3)

	augMa[0] = []int{-3,  2, -1, -1}
	augMa[1] = []int{ 0, -2,  5, -9}
	augMa[2] = []int{ 0,  0, -2,  2}

	// fmt.Printf("augMa[0]: %d \n", augMa[0])
	// fmt.Printf("augMa[:][0]: %d \n", augMa[:][0])
	// fmt.Printf("augMa: %d \n", augMa)
	// fmt.Printf("length of augMa: %d \n", len(augMa))

	res := backSubstitution(augMa)

	fmt.Printf("This is the result: %d \n", res)  

	shouldBe := []int{2, 2, -1}

	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the back substitution is wrong")
		t.Fail()
	}

}

func TestGaussElimination(t *testing.T) {
	augMa := make([][]int, 3)

	augMa[0] = []int{-3,  2, -1, -1}
	augMa[1] = []int{ 6, -6,  7, -7}
	augMa[2] = []int{ 3, -4,  4, -6}

	res := gaussElimination(augMa) 
	fmt.Printf("This is the result of the Gauss elimination: %d \n", res) 

	shouldBe := make([][]int, 3) 
	shouldBe[0] = []int{-3,  2, -1, -1}
	shouldBe[1] = []int{ 0, -2,  5, -9}
	shouldBe[2] = []int{ 0,  0, -2,  2}

	// res2 := backSubstitution(res)
	// fmt.Printf("This is the result after back substitution: %d \n", res2) 

	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the Gauss elimination wrong")
		t.Fail()
	}

}

func TestMakeAugmentedMatrix(t *testing.T) {
	
	A := [][]int{{1,2,3,4}, {1,2,3,4}, {1,2,3,4}}
    
	b := []int{5,6,7}

	augMa := makeAugmentedMatrix(A, b)

	fmt.Printf("This is the augmented matrix: %d \n", augMa)


	shouldBe := [][]int{{1,2,3,4,5}, {1,2,3,4,6}, {1,2,3,4,7}}

	if !reflect.DeepEqual(augMa, shouldBe) {
		t.Log("The augmented matrix is wrong")
		t.Fail()
	}

}

func TestWholeGaussElimination(t *testing.T) {

	A := [][]int{{2, -1, 3}, {1, 4, -2}, {3, 1, 5}}

	b := []int{5, 1, 2}

	augMa := makeAugmentedMatrix(A, b) 
	fmt.Printf("This is augMa:			  %d \n", augMa) 

	afterGaussElim := gaussElimination(augMa) 
	fmt.Printf("This is after Gauss elim: %d \n", afterGaussElim)

	res := backSubstitution(afterGaussElim)
	fmt.Printf("This is the result after back substitution: %d \n", res) 


	shouldBe := []int{5, 3, -84}	// ifølge youtube video

	if !reflect.DeepEqual(res, shouldBe) {
		t.Log("The result of the back substitution is wrong")
		t.Fail()
	}



}

