// Bab 02 — Variabel & Tipe Data
// Belajar menyimpan data dengan variabel dan konstanta.
//
// Jalankan: go run ./02-variabel-dan-tipe-data

package main

import "fmt"

func main() {
	// -110
	// --- 1. Deklarasi dengan var ---
	var nama string // deklarasi tanpa nilai (zero value = "")
	nama = "Budi"   // pemberian nilai

	var umur int = 25 // deklarasi + nilai eksplisit

	var kota = "Jakarta" // tipe dihilangkan, Go menebak (string)

	// --- 2. Deklarasi singkat dengan := ---
	tinggi := 170.5 // float64
	menikah := false

	// --- 3. Deklarasi banyak sekaligus ---
	var a, b, c int = 1, 2, 3
	x, y := "dua", 2 // bisa beda tipe

	// --- 4. Konstanta ---
	const Pi = 3.14

	fmt.Println("Nama   :", nama)
	fmt.Println("Umur   :", umur)
	fmt.Println("Kota   :", kota)
	fmt.Println("Tinggi :", tinggi, "cm")
	fmt.Println("Menikah:", menikah)
	fmt.Println("a,b,c  :", a, b, c)
	fmt.Println("x,y    :", x, y)
	fmt.Println("Pi     :", Pi)

	// --- 5. Zero value ---
	var angka int
	var teks string
	var ok bool
	fmt.Println("Zero value int   :", angka) // 0
	fmt.Println("Zero value string:", teks)  // (kosong)
	fmt.Println("Zero value bool  :", ok)    // false

	// --- 6. Menampilkan tipe dengan %T ---
	fmt.Printf("Tipe variabel tinggi = %T\n", tinggi)
	fmt.Printf("Tipe variabel umur   = %T\n", umur)

}
