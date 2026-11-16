# Tugas Bab 10 — Error Handling

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Bagi dengan Error
Tulis fungsi `bagi(a, b float64) (float64, error)` yang mengembalikan error saat
`b == 0`. Panggil untuk kasus valid dan kasus error, lalu tampilkan dengan
`if err != nil`.

### Tugas 2 — Akar Kuadrat
Tulis fungsi `akar(n float64) (float64, error)` yang mengembalikan error jika
`n < 0`. Gunakan `math.Sqrt` untuk menghitung akar nilai positif.

### Tugas 3 — defer Sederhana
Tulis fungsi yang memakai `defer fmt.Println("selesai")` dan buktikan pesan itu
muncul di akhir.

## Tingkat Menengah

### Tugas 4 — Membungkus Error
Tulis fungsi `baca(namaFile string) error` yang mengembalikan
`fmt.Errorf("gagal membaca %s: %w", namaFile, errAsli)`. Panggil dan cetak
pesannya.

### Tugas 5 — Banyak defer (LIFO)
Tulis 3 `defer` yang mencetak `"pertama"`, `"kedua"`, `"ketiga"`. Amati urutan
cetaknya dan tuliskan kesimpulanmu.

### Tugas 6 — panic & recover
Tulis fungsi yang memanggil `panic`, lalu tangkap dengan `recover` di dalam
`defer`. Pastikan program tetap berjalan setelahnya.

## Tingkat Mahir

### Tugas 7 — Validasi Input
Tulis fungsi `validasiUsia(umur int) error` yang mengembalikan error jika umur
`< 0` atau `> 150`. Tampilkan pesan error dengan jelas.

### Tugas 8 — Gabungkan Semua
Tulis program kecil yang:
1. Meminta input angka (bisa pakai nilai tetap untuk menyederhanakan).
2. Memanggil fungsi yang bisa error.
3. Menangani error dengan `if err != nil`.
4. Menggunakan `defer` untuk membersihkan (misal mencetak "program selesai").

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham pola `if err != nil`.
- [ ] Kamu paham `defer` dan urutan LIFO.
- [ ] Kamu tahu kapan memakai `panic`/`recover` (dan kapan tidak).
