# Chapter 1 Go-lang Basic (GLNG)
Go Programming Introduction

**Challenge 2**

## Biodata
> ID Regis = 1955617840-1126
> 
> Nama = Ilham Dwi Yanto
> 
> Kelas = Golang 2

<details><summary>Assigment 2</summary>
<p>

```ruby
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

```
![](<Output-Challenge 2.png?raw=true>)
</p>
</details>

