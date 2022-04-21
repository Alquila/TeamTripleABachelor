package main

import (
	"fmt"
	_ "fmt"
	_ "reflect"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

func TestCreate_G_matrix(t *testing.T) {
	G := Create_G_matrix()

	// fmt.Printf("G Matrix: %d \n", G)
	for i := 0; i < 184; i++ {
		if G[i][i] != 1 {
			t.Fail()
			fmt.Printf("G[%d][%d] er 0 men burde være 1\n", i, i)
		}
	}
}
