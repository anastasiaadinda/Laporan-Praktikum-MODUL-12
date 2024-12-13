package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan Nilai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukan Nilai y : ")
	fmt.Scan(&y)

	hasil := 0
	for x >= y {
		x = x - y
		hasil++
	}
	fmt.Println(hasil)
}
