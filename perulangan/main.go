package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
	
		fmt.Println()
	}

	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
    }
}
}