// Bab 05 — Fungsi
// Belajar mendefinisikan dan memakai fungsi.
//
// Jalankan: go run ./05-fungsi

package main

import (
	"errors"
	"fmt"
)

// --- Fungsi sederhana ---
// Parameter bertipe sama bisa disingkat (a, b int).
func tambah(a, b int) int {
	return a + b
}

// --- Multiple return ---
// Mengembalikan hasil pembagian dan error (jika b == 0).
func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// --- Named return ---
// Variabel "luas" sudah dideklarasikan di tanda kurung return.
func luasPersegi(sisi float64) (luas float64) {
	luas = sisi * sisi
	return // return kosong -> mengembalikan "luas"
}

// --- Variadic ---
// Menerima banyak angka (angka ...int).
func jumlahkan(angka ...int) int {
	total := 0
	for _, n := range angka {
		total += n
	}
	return total
}

func main() {
	fmt.Println("tambah(3, 4)      =", tambah(3, 4))

	// Multiple return
	hasil, err := bagi(10, 2)
	fmt.Println("bagi(10, 2)       =", hasil, "err =", err)

	hasil, err = bagi(10, 0)
	fmt.Println("bagi(10, 0)       =", hasil, "err =", err)

	// Named return
	fmt.Println("luasPersegi(5)    =", luasPersegi(5))

	// Variadic
	fmt.Println("jumlahkan(1..3)   =", jumlahkan(1, 2, 3))
	fmt.Println("jumlahkan(10,20)  =", jumlahkan(10, 20))
	fmt.Println("jumlahkan()       =", jumlahkan()) // 0

	// --- Fungsi sebagai nilai (function literal) ---
	kali := func(a, b int) int { return a * b }
	fmt.Println("kali(3, 4)        =", kali(3, 4))

	// --- Closure ---
	// "counter" mengingat variabel "hitung" milik fungsi pembuatnya.
	pembuatCounter := func() func() int {
		hitung := 0
		return func() int {
			hitung++
			return hitung
		}
	}

	counter := pembuatCounter()
	fmt.Println("counter()         =", counter()) // 1
	fmt.Println("counter()         =", counter()) // 2
	fmt.Println("counter()         =", counter()) // 3
}
