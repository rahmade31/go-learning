// Bab 11 — Goroutine & Channel (Konkurensi)
// Belajar menjalankan tugas secara bersamaan.
//
// Jalankan: go run ./11-goroutine-dan-channel

package main

import (
	"fmt"
	"sync"
	"time"
)

// --- Contoh goroutine ---
func cetak(angka int) {
	fmt.Println("  goroutine mencetak:", angka)
}

func main() {
	// --- 1. Goroutine + WaitGroup ---
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // tambah hitungan goroutine
		go func(n int) {
			defer wg.Done() // kurangi saat selesai
			cetak(n)
		}(i)
	}
	wg.Wait() // tunggu semua selesai
	fmt.Println("semua goroutine selesai")

	// --- 2. Channel (komunikasi antar goroutine) ---
	ch := make(chan int)

	go func() {
		ch <- 42 // kirim nilai
	}()

	nilai := <-ch // terima nilai
	fmt.Println("nilai dari channel:", nilai)

	// --- 3. Buffered channel ---
	chBuffered := make(chan string, 3)
	chBuffered <- "a"
	chBuffered <- "b"
	chBuffered <- "c"
	fmt.Println("buffered:", <-chBuffered, <-chBuffered, <-chBuffered)

	// --- 4. select ---
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "dari ch1"
	}()
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "dari ch2"
	}()

	// pilih yang pertama siap (dalam contoh ini ch2)
	select {
	case pesan := <-ch1:
		fmt.Println("menerima:", pesan)
	case pesan := <-ch2:
		fmt.Println("menerima:", pesan)
	}
}
