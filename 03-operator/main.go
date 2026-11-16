// Bab 03 — Operator
// Belajar operator aritmatika, perbandingan, logika, dan penugasan.
//
// Jalankan: go run ./03-operator

package main

import "fmt"

func main() {
	// --- Operator aritmatika ---
	a, b := 10, 3
	fmt.Println("a + b =", a+b) // 13
	fmt.Println("a - b =", a-b) // 7
	fmt.Println("a * b =", a*b) // 30
	fmt.Println("a / b =", a/b) // 3 (int dibagi int -> int)
	fmt.Println("a % b =", a%b) // 1 (sisa bagi)

	// Pembagian desimal
	fmt.Println("10.0 / 3.0 =", 10.0/3.0) // 3.333...

	// --- Operator perbandingan ---
	fmt.Println("a == b :", a == b) // false
	fmt.Println("a != b :", a != b) // true
	fmt.Println("a > b  :", a > b)  // true
	fmt.Println("a <= b :", a <= b) // false

	// --- Operator logika ---
	sukaBuku := true
	sukaFilm := false
	fmt.Println("sukaBuku && sukaFilm :", sukaBuku && sukaFilm) // false
	fmt.Println("sukaBuku || sukaFilm :", sukaBuku || sukaFilm) // true
	fmt.Println("!sukaBuku            :", !sukaBuku)            // false

	// --- Operator penugasan ---
	nilai := 10
	nilai += 5 // nilai = 15
	fmt.Println("nilai setelah +=5 :", nilai)
	nilai *= 2 // nilai = 30
	fmt.Println("nilai setelah *=2 :", nilai)

	// --- Operator string (concatenation) ---
	depan, belakang := "Budi", "Santoso"
	namaLengkap := depan + " " + belakang
	fmt.Println("Nama lengkap :", namaLengkap)
}
