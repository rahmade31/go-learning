# Tugas Bab 05 — Fungsi

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Tambah & Kali
Buat fungsi `tambah(a, b int) int` dan `kali(a, b int) int`. Panggil keduanya
dari `main` dan cetak hasilnya.

### Tugas 2 — Salam
Buat fungsi `salam(nama string) string` yang mengembalikan `"Halo, <nama>!"`.
Cetak hasilnya.

### Tugas 3 — Luas Persegi Panjang
Buat fungsi `luas(panjang, lebar int) int` yang mengembalikan luasnya, lalu
cetak hasil untuk `panjang=10, lebar=5`.

## Tingkat Menengah

### Tugas 4 — Nilai Maksimum
Buat fungsi `maks(a, b, c int) int` yang mengembalikan angka terbesar dari tiga
angka.

### Tugas 5 — Multiple Return
Buat fungsi `bagi(a, b int) (int, error)` yang mengembalikan `(0, error)` jika
`b == 0`. Tampilkan error dengan rapi.

### Tugas 6 — Variadic: Rata-rata
Buat fungsi `rataRata(angka ...float64) float64` yang menghitung rata-rata.
Uji dengan 3 angka dan dengan 5 angka.

## Tingkat Mahir

### Tugas 7 — Faktorial (Rekursif)
Buat fungsi rekursif `faktorial(n int) int`. Ingat `faktorial(0) = 1` dan
`faktorial(n) = n * faktorial(n-1)`.

### Tugas 8 — Closure: Counter
Buat fungsi `pembuatCounter() func() int` yang mengembalikan closure penghitung.
Panggil 3 kali dan pastikan hasilnya 1, 2, 3.

### Tugas 9 — Fungsi sebagai Parameter
Buat fungsi `terapkan(f func(int) int, n int) int` yang memanggil `f(n)`.
Gunakan untuk menggandakan angka dan untuk menambah 10.

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham parameter, return value, dan multiple return.
- [ ] Kamu paham variadic dan closure.
- [ ] Kamu bisa menulis fungsi rekursif sederhana.
