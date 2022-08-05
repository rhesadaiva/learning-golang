package main

import (
	"fmt"
	"strings"
)

func main() {

	// 	Data array (yang didapat dari parameter pertama) akan di-looping.
	// Di tiap perulangannya, closure callback dipanggil, dengan disisipkan data tiap elemen perulangan sebagai parameter.
	// Closure callback berisikan kondisi filtering, dengan hasil bertipe bool yang kemudian dijadikan nilai balik dikembalikan.
	// Di dalam fungsi filter() sendiri, ada proses seleksi kondisi (yang nilainya didapat dari hasil eksekusi closure callback).
	// Ketika kondisinya bernilai true, maka data elemen yang sedang diulang dinyatakan lolos proses filtering.
	// Data yang lolos ditampung variabel result. Variabel tersebut dijadikan sebagai nilai balik fungsi filter().

	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
