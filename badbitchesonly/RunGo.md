# GO kommandoer
go test -run TestName (kør specific test)
go test -v (kør alle tests)
go test -timeout 0 -run TestTryAllReg4 (no timeout)

# When running the first time run: 
Run this inside the project:
 * go mod init module_name
 * go mod tidy


This is a wonderful monster: 
```
// 	for i := 0; i < 2; i++ {
// 		r4_guess[0] = i
// 		for i := 0; i < 2; i++ {
// 			r4_guess[1] = i
// 			for i := 0; i < 2; i++ {
// 				r4_guess[2] = i
// 				for i := 0; i < 2; i++ {
// 					r4_guess[3] = i
// 					for i := 0; i < 2; i++ {
// 						r4_guess[4] = i
// 						for i := 0; i < 2; i++ {
// 							r4_guess[5] = i
// 							for i := 0; i < 2; i++ {
// 								r4_guess[6] = i
// 								for i := 0; i < 2; i++ {
// 									r4_guess[7] = i
// 									for i := 0; i < 2; i++ {
// 										r4_guess[8] = i
// 										for i := 0; i < 2; i++ {
// 											r4_guess[9] = i
// 											for i := 0; i < 2; i++ {
// 												r4_guess[11] = i
// 												for i := 0; i < 2; i++ {
// 													r4_guess[12] = i
// 													for i := 0; i < 2; i++ {
// 														r4_guess[13] = i
// 														for i := 0; i < 2; i++ {
// 															r4_guess[14] = i
// 															for i := 0; i < 2; i++ {
// 																r4_guess[15] = i
// 																for i := 0; i < 2; i++ {
// 																	r4_guess[16] = i
// 																	//do the gauss or whatever
// 																}
// 															}
// 														}
// 													}
// 												}
// 											}
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
```