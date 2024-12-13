package main

import "fmt"

func main() {

	const username = "Admin"
	const password = "Admin"
	var ulang int

	for {
		//untuk membuat inputan saat program dijalankan
		var inputUsername, inputPassword string
		fmt.Print("Masukkan Username : ")
		fmt.Scanln(&inputUsername)
		fmt.Print("Masukan Password: ")
		fmt.Scanln(&inputPassword)

		if inputUsername == username && inputPassword == password {
			fmt.Printf("%d percobaan gagal login\n ", ulang)
			fmt.Println("Login Berhasil")
			break
		} else {
			fmt.Println("Username atau Password salah")
			ulang++
		}

	}

}
