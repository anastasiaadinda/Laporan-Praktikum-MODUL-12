package main

import "fmt"

func main() {

	var digit int
	fmt.Print("Masukan digit: ")
	fmt.Scan(&digit)

	if digit <= 0 {
		fmt.Print("harap masukan bilangan bulat positif")
		return
	}

	for digit > 0 {
		angka := digit % 10
		fmt.Println(angka)
		digit /= 10
	}
}
