package main

func Create_G_matrix() [][]int {
	// Make first slice: 184 columns
	G := make([][]int, 184)

	// Make 184 slices of length 456
	for i := 0; i < 184; i++ {
		col_slice := make([]int, 456)

		// Set diagonal to 1
		col_slice[i] = 1

		G[i] = col_slice
	}

	return G
}

func new_Gauss() {
	
}
