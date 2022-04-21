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

func Create_KG_matrix() [][]int {
	// Make first slice: 184 columns
	KG := make([][]int, 456)

	// Make 184 slices of length 456
	for i := 0; i < 456; i++ {
		col_slice := make([]int, 272)

		// Set diagonal to 1 after 184
		if i > 184 && i < 272 {
			col_slice[i] = 1
		}

		KG[i] = col_slice
	}

	return KG
}
