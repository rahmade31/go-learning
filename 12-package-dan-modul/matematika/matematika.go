// Package matematika berisi operasi hitung sederhana.
//
// Catatan: nama fungsi diawali huruf besar (Tambah, Kali, dll.) agar bisa
// diimpor dari paket lain (exported). Nama huruf kecil hanya untuk internal.
package matematika

// Tambah menjumlahkan dua bilangan bulat.
func Tambah(a, b int) int {
	return a + b
}

// Kurang mengurangkan dua bilangan bulat.
func Kurang(a, b int) int {
	return a - b
}

// Kali mengalikan dua bilangan bulat.
func Kali(a, b int) int {
	return a * b
}

// Bagi membagi dua bilangan bulat. Mengembalikan 0 jika b == 0.
func Bagi(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
