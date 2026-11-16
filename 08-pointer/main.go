// Bab 08 — Pointer
// Belajar alamat memori dan pass by reference.
//
// Jalankan: go run ./08-pointer

package main

import "fmt"

// --- Fungsi pass by value (mengubah salinan saja) ---
func ubah(n int) {
	n = 99
}

// --- Fungsi pass by reference (mengubah nilai asli) ---
func ubahViaPointer(n *int) {
	*n = 99
}

// --- Pointer ke struct ---
type Orang struct {
	Nama string
}

func gantiNama(o *Orang, namaBaru string) {
	o.Nama = namaBaru // Go otomatis dereference: (*o).Nama
}

func main() {
	// --- Dasar pointer ---
	x := 10
	p := &x // p menyimpan alamat x

	fmt.Println("x      :", x)  // 10
	fmt.Println("p (&x) :", p)  // alamat memori
	fmt.Println("*p     :", *p) // 10 (nilai yang ditunjuk)

	*p = 20                             // ubah nilai lewat pointer
	fmt.Println("x setelah *p=20 :", x) // 20

	// --- Pass by value ---
	y := 10
	ubah(y)
	fmt.Println("y setelah ubah()      :", y) // tetap 10

	// --- Pass by reference ---
	ubahViaPointer(&y)
	fmt.Println("y setelah ubahPointer():", y) // 99

	// --- Pointer ke struct ---
	org := Orang{Nama: "Budi"}
	gantiNama(&org, "Ani")
	fmt.Println("nama setelah ganti :", org.Nama)

	// --- new() ---
	ptr := new(int) // pointer ke int zero-value (0)
	fmt.Println("nilai *ptr (new)   :", *ptr)
}
