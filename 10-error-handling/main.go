// Bab 10 — Error Handling
// Belajar error, defer, panic, dan recover.
//
// Jalankan: go run ./10-error-handling

package main

import (
	"errors"
	"fmt"
)

// --- Fungsi yang mengembalikan error ---
func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// --- Membungkus error dengan konteks ---
func proses(angka int) error {
	if angka < 0 {
		return fmt.Errorf("proses gagal: angka %d negatif", angka)
	}
	return nil
}

// --- defer: memastikan sesuatu selalu dijalankan ---
func contohDefer() {
	defer fmt.Println("  [defer] ini dijalankan TERAKHIR")
	defer fmt.Println("  [defer] ini dijalankan kedua (LIFO)")
	fmt.Println("  badan fungsi")
}

// --- panic + recover ---
func fungsiAman() {
	// recover harus ada di dalam defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  pulih dari panic:", r)
		}
	}()

	fmt.Println("  sebelum panic")
	// Perhatikan: baris apa pun yang diletakkan SETELAH panic() tidak akan
	// pernah dieksekusi, karena panic menghentikan alur normal fungsi.
	panic("terjadi kesalahan fatal")
}

func main() {
	// --- Error handling idiomatik ---
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10/2 =", hasil)
	}

	hasil, err = bagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10/0 =", hasil)
	}

	// --- Error dengan konteks ---
	if err := proses(-5); err != nil {
		fmt.Println("proses:", err)
	}

	// --- defer ---
	fmt.Println("contoh defer:")
	contohDefer()

	// --- panic + recover ---
	fmt.Println("contoh panic/recover:")
	fungsiAman()
	fmt.Println("program tetap berjalan setelah recover")
}
