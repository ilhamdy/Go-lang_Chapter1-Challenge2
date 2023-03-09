package main

import "fmt"

func main() {
	// Looping nilai i
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	// Looping nilai j
	for j := 0; j < 11; j++ {
		if j == 5 {
			// String yang akan diiterate
			s := "САШАРВО"
			// Looping karakter pada string
			for i, r := range s {
				fmt.Printf("character %U '%c' starts at byte position %d\n", r, r, i)
			}
			// Melompati nilai j = 5
			continue
		}
		fmt.Println("Nilai j =", j)
	}
}
