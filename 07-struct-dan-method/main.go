// Bab 07 — Struct & Method
// Belajar tipe data custom dan method.
//
// Jalankan: go run ./07-struct-dan-method

package main

import "fmt"

// --- Definisi struct ---
type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	IPK     float64
}

// --- Method dengan value receiver (menerima salinan) ---
func (m Mahasiswa) Perkenalan() string {
	return "Halo, saya " + m.Nama + " dari " + m.Jurusan
}

// --- Method dengan pointer receiver (bisa mengubah nilai asli) ---
func (m *Mahasiswa) NaikkanIPK(tambah float64) {
	m.IPK += tambah
}

// --- Embedding (komposisi) ---
type Orang struct {
	Nama string
}

type Karyawan struct {
	Orang // embedding: Karyawan punya field Nama
	Gaji  int
}

func main() {
	// --- Membuat nilai struct ---
	// Cara 1: urutan field (rapuh, tidak disarankan)
	m1 := Mahasiswa{"Budi", "101", "Informatika", 3.5}

	// Cara 2: nama field (disarankan)
	m2 := Mahasiswa{
		Nama:    "Ani",
		NIM:     "102",
		Jurusan: "Sistem Informasi",
		IPK:     3.8,
	}

	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)

	// Akses & ubah field
	fmt.Println("Nama m2:", m2.Nama)
	m2.IPK = 3.9
	fmt.Println("IPK m2 setelah diubah:", m2.IPK)

	// --- Method ---
	fmt.Println(m2.Perkenalan())

	// Pointer receiver mengubah nilai asli
	m2.NaikkanIPK(0.1)
	fmt.Println("IPK m2 setelah naik:", m2.IPK)

	// --- Embedding ---
	k := Karyawan{
		Orang: Orang{Nama: "Caca"},
		Gaji:  5000,
	}
	fmt.Println("Karyawan:", k.Nama, "gaji", k.Gaji) // Nama di-promote
}
