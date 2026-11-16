// Bab 06 — Array, Slice, dan Map
// Belajar struktur data koleksi.
//
// Jalankan: go run ./06-array-slice-map

package main

import "fmt"

func main() {
	// --- Array (ukuran tetap) ---
	var nilai [3]int
	nilai[0] = 90
	nilai[1] = 85
	nilai[2] = 78
	fmt.Println("array nilai :", nilai, "panjang:", len(nilai))

	// --- Slice (ukuran fleksibel) ---
	angka := []int{1, 2, 3}
	angka = append(angka, 4)
	fmt.Println("slice angka :", angka)

	// Slice kosong dengan make
	kota := make([]string, 0, 5)
	kota = append(kota, "Jakarta", "Bandung")
	fmt.Println("slice kota  :", kota)

	// --- Slicing (mengambil potongan) ---
	h := []int{10, 20, 30, 40, 50}
	fmt.Println("h[1:3]      :", h[1:3]) // [20 30]
	fmt.Println("h[:2]       :", h[:2])  // [10 20]
	fmt.Println("h[2:]       :", h[2:])  // [30 40 50]

	// --- range ---
	buah := []string{"apel", "jeruk", "mangga"}
	fmt.Println("daftar buah :")
	for i, nama := range buah {
		fmt.Printf("  indeks %d -> %s\n", i, nama)
	}

	// --- Map ---
	usia := map[string]int{
		"Budi": 25,
		"Ani":  30,
	}
	usia["Caca"] = 22   // tambah
	usia["Budi"] = 26   // ubah
	delete(usia, "Ani") // hapus

	fmt.Println("map usia    :", usia)

	// Cek keberadaan kunci
	umur, ada := usia["Budi"]
	fmt.Println("usia Budi   :", umur, "ada:", ada)

	_, ada = usia["Ani"]
	fmt.Println("usia Ani ada:", ada) // false (sudah dihapus)

	// Menjelajah map
	fmt.Println("isi map usia:")
	for kunci, nilai := range usia {
		fmt.Printf("  %s -> %d\n", kunci, nilai)

	}
}
