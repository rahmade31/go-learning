// Bab 12 — Package & Modul
// Belajar membuat dan mengimpor paket sendiri.
//
// Jalankan: go run ./12-package-dan-modul

package main

import (
	"fmt"

	// Impor paket buatan sendiri menggunakan jalur modul:
	// "learning-go" (nama modul) + "/12-package-dan-modul/matematika" (folder).
	"learning-go/12-package-dan-modul/matematika"
)

func main() {
	a, b := 10, 3

	fmt.Println("Tambah :", matematika.Tambah(a, b))
	fmt.Println("Kurang :", matematika.Kurang(a, b))
	fmt.Println("Kali   :", matematika.Kali(a, b))
	fmt.Println("Bagi   :", matematika.Bagi(a, b))
	fmt.Println("Bagi 0 :", matematika.Bagi(a, 0)) // mengembalikan 0
}
