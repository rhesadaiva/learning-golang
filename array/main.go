package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	var fruitsz = make([]string, 2)
	fruitsz[0] = "apple"
	fruitsz[1] = "manggo"

	fmt.Println(fruitsz) // [apple manggo]

}
