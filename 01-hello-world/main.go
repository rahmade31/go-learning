// Bab 01 — Hello World
// Program pertama: mencetak teks ke layar.
//
// Cara menjalankan dari folder utama (learning-go):
//   go run ./01-hello-world

// package main: program yang bisa dijalankan langsung.
package main

// import: mengambil fungsi dari paket standar "fmt".
import (
	"fmt"
)

// func main: titik awal program, selalu dijalankan pertama kali.
func main() {

	fmt.Println("Benar", true)
	fmt.Println("Salah", false)

	var (
		firstname  string = "Rahmad"
		middlename        = "Windi"
		lastname          = "Simanullang"
	)

	const (
		address   string = "Jalan Swadaya Raya	"
		age              = 25
		noaddress        = "Kembang Beji Village"
	)
	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
	fmt.Println(age)
	fmt.Println(address)

}
