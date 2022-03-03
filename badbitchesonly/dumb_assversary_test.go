package main

import (
	"fmt"
	// "math/rand"
	// "reflect"
	// "strconv"
	// "strings"
	"testing"
	// "time"
	// //"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func TestPlaintext(t *testing.T) {
	plaintext := MakePlaintext()
	fmt.Printf("%d", plaintext)
}

func TestEncryptPlaintext(t *testing.T) {
	plaintext := MakePlaintext()
	fmt.Printf("This is the plaintext: %d \n", plaintext)
	cipher := EncryptSimplePlaintext(plaintext)
	fmt.Printf("%d \n", cipher)
}

func TestPrint2(t *testing.T) {
	print("hello worlds")
}
