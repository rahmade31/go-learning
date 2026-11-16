// Bab 13 — Proyek Akhir: Aplikasi TODO CLI
//
// Aplikasi daftar tugas sederhana yang merangkai semua konsep Go:
// struct, slice, fungsi, method, switch, for, range, dan error handling.
//
// Jalankan: go run ./13-proyek-akhir

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task merepresentasikan satu tugas.
type Task struct {
	Deskripsi string
	Selesai   bool
}

// TandaiSelesai mengubah status tugas menjadi selesai.
func (t *Task) TandaiSelesai() {
	t.Selesai = true
}

// daftarTugas menyimpan semua tugas (slice, karena jumlahnya dinamis).
var daftarTugas []Task

// reader membaca input dari terminal (dipakai ulang di seluruh program).
var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		tampilkanMenu()

		pilihan := bacaInput("Pilih menu (1-5): ")

		switch pilihan {
		case "1":
			tambahTugas()
		case "2":
			tampilkanTugas()
		case "3":
			tandaiSelesai()
		case "4":
			hapusTugas()
		case "5":
			fmt.Println("Sampai jumpa! 👋")
			return // keluar dari main -> program berhenti
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}

		fmt.Println() // baris kosong agar tampilan rapi
	}
}

// tampilkanMenu mencetak menu utama.
func tampilkanMenu() {
	fmt.Println("===== APLIKASI TODO =====")
	fmt.Println("1. Tambah tugas")
	fmt.Println("2. Lihat tugas")
	fmt.Println("3. Tandai selesai")
	fmt.Println("4. Hapus tugas")
	fmt.Println("5. Keluar")
}

// tambahTugas menambah tugas baru ke daftar.
func tambahTugas() {
	deskripsi := bacaInput("Deskripsi tugas: ")

	// Validasi: jangan izinkan deskripsi kosong.
	if strings.TrimSpace(deskripsi) == "" {
		fmt.Println("Deskripsi tidak boleh kosong.")
		return
	}

	daftarTugas = append(daftarTugas, Task{Deskripsi: deskripsi})
	fmt.Printf("Tugas %q ditambahkan!\n", deskripsi)
}

// tampilkanTugas mencetak semua tugas beserta statusnya.
func tampilkanTugas() {
	if len(daftarTugas) == 0 {
		fmt.Println("Belum ada tugas.")
		return
	}

	fmt.Println("Daftar tugas:")
	for i, tugas := range daftarTugas {
		status := "[ ]"
		if tugas.Selesai {
			status = "[x]"
		}
		fmt.Printf("  %d. %s %s\n", i+1, status, tugas.Deskripsi)
	}
}

// tandaiSelesai menandai tugas berdasarkan nomornya.
func tandaiSelesai() {
	tampilkanTugas()
	if len(daftarTugas) == 0 {
		return
	}

	idx, ok := bacaNomor("Nomor tugas yang selesai: ")
	if !ok {
		return
	}

	daftarTugas[idx].TandaiSelesai()
	fmt.Printf("Tugas %q ditandai selesai!\n", daftarTugas[idx].Deskripsi)
}

// hapusTugas menghapus tugas berdasarkan nomornya.
func hapusTugas() {
	tampilkanTugas()
	if len(daftarTugas) == 0 {
		return
	}

	idx, ok := bacaNomor("Nomor tugas yang dihapus: ")
	if !ok {
		return
	}

	dihapus := daftarTugas[idx].Deskripsi
	// Hapus elemen: gabungkan sebelum & sesudah indeks.
	daftarTugas = append(daftarTugas[:idx], daftarTugas[idx+1:]...)
	fmt.Printf("Tugas %q dihapus!\n", dihapus)
}

// bacaNomor membaca input nomor (1..n) dan mengembalikan indeks 0-based.
// Nilai balik kedua (ok) menyatakan apakah input valid.
func bacaNomor(prompt string) (int, bool) {
	raw := bacaInput(prompt)
	nomor, err := strconv.Atoi(raw)
	if err != nil || nomor < 1 || nomor > len(daftarTugas) {
		fmt.Println("Nomor tidak valid.")
		return 0, false
	}
	return nomor - 1, true // konversi ke indeks 0-based
}

// bacaInput membaca satu baris teks dari terminal (menghapus baris baru).
func bacaInput(prompt string) string {
	fmt.Print(prompt)
	teks, _ := reader.ReadString('\n')
	return strings.TrimSpace(teks)
}
