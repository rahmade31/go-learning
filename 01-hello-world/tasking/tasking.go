// Tugas Bab 01 — Hello World
// Program latihan ini dipisah dari main.go, jadi ia punya package dan
// func main sendiri.
//
// Cara menjalankan dari folder utama (go-learning):
//   go run ./01-hello-world/tasking
//
// Kenapa harus di folder sendiri?
// Dalam satu folder hanya boleh ada satu package dan satu func main.
// Kalau file ini ditaruh langsung di 01-hello-world/, ia akan bentrok
// dengan main.go dan muncul error "main redeclared in this block".

// package main: program yang bisa dijalankan langsung.
package main

// import: mengambil paket standard "fmt" untuk mencetak ke layar.
import "fmt"

// func main: titik awal program.
func main() {
	// Tugas 1 — Sapaan Pribadi
	fmt.Println("Halo, nama saya Rahmad")
	fmt.Println("Saya sedang belajar Go")
	fmt.Println("Semoga lancar!")

	fmt.Println()

	// Tugas 2 — Print vs Println
	// Print tidak menambah baris baru, Println menambah baris baru.
	fmt.Print("A ")
	fmt.Print("B")
	fmt.Println()
	fmt.Println("C")

	fmt.Println()

	// Tugas 3 — Format dengan Printf
	nama := "Rahmad"
	umur := 25
	kota := "<kotamu>" // ganti dengan kotamu

	fmt.Printf("Saya %s, umur %d tahun, tinggal di %s.\n", nama, umur, kota)
}
