// Bab 09 — Interface
// Belajar kontrak perilaku dan polimorfisme.
//
// Jalankan: go run ./09-interface

package main

import "fmt"

// --- Definisi interface ---
type Shape interface {
	Luas() float64
	Keliling() float64
}

// --- Lingkaran (implementasi implisit) ---
type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return 3.14 * l.JariJari * l.JariJari
}

func (l Lingkaran) Keliling() float64 {
	return 2 * 3.14 * l.JariJari
}

// --- Persegi (implementasi implisit) ---
type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

// --- Fungsi polimorfik: menerima Shape apa pun ---
func cetakInfo(s Shape) {
	fmt.Printf("  Luas=%.2f, Keliling=%.2f\n", s.Luas(), s.Keliling())
}

func main() {
	fmt.Println("Info Lingkaran:")
	cetakInfo(Lingkaran{JariJari: 7})

	fmt.Println("Info Persegi:")
	cetakInfo(Persegi{Sisi: 5})

	// --- Interface kosong (any) ---
	var apaSaja any
	apaSaja = 42
	fmt.Println("apaSaja =", apaSaja)

	apaSaja = "sebuah teks"
	fmt.Println("apaSaja =", apaSaja)

	// --- Type assertion ---
	if teks, ok := apaSaja.(string); ok {
		fmt.Println("apaSaja adalah string:", teks)
	}

	// --- Type switch ---
	nilai := any(3.14)
	switch v := nilai.(type) {
	case int:
		fmt.Println("integer:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Printf("tipe lain: %T = %v\n", v, v)
	}
}
