// Bab 04 — Kontrol Alur (if, switch, for)
// Belajar pengambilan keputusan dan perulangan.
//
// Jalankan: go run ./04-kontrol-alur

package main

import "fmt"

func main() {
	// --- if / else ---
	umur := 20
	if umur >= 17 {
		fmt.Println("Sudah boleh punya KTP")
	} else {
		fmt.Println("Belum boleh punya KTP")
	}

	// --- if dengan pernyataan pendek ---
	if nilai := 85; nilai >= 75 {
		fmt.Println("Lulus dengan nilai", nilai)
	}

	// --- if / else if / else ---
	nilai := 72
	if nilai >= 90 {
		fmt.Println("Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 {
		fmt.Println("Nilai C")
	} else {
		fmt.Println("Nilai D")
	}

	// --- switch ---
	hari := "Minggu"
	switch hari {
	case "Senin":
		fmt.Println("Awal pekan")
	case "Sabtu", "Minggu":
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Hari biasa")
	}

	// --- for bentuk klasik ---
	fmt.Println("for klasik:")
	for i := 0; i < 3; i++ {
		fmt.Println("  i =", i)
	}

	// --- for seperti while ---
	fmt.Println("for seperti while:")
	n := 0
	for n < 3 {
		fmt.Println("  n =", n)
		n++
	}

	// --- for + break + continue ---
	fmt.Println("angka genap 1..10:")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // lewati angka ganjil
		}
		if i > 8 {
			break // berhenti di 8
		}
		fmt.Println("  ", i)
	}
}
