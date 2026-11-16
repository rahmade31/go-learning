# Tugas Bab 02 — Variabel & Tipe Data

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Kartu Identitas
Buat variabel `nama` (string), `umur` (int), `tinggi` (float64), dan `menikah`
(bool), lalu cetak semuanya dalam satu baris memakai `Printf`.

### Tugas 2 — Zero Value
Deklarasikan (tanpa memberi nilai) variabel bertipe `int`, `float64`, `string`,
dan `bool`. Cetak nilai masing-masing dan tuliskan komentar hasilnya.

### Tugas 3 — Luas Persegi Panjang
Buat `panjang := 10` dan `lebar := 5`, hitung luas dan keliling, lalu cetak:

```
Luas      = 50
Keliling  = 30
```

## Tingkat Menengah

### Tugas 4 — Tukar Nilai (Swap)
Buat `a := 5` dan `b := 10`. Tukar isinya tanpa variabel ketiga, lalu cetak.
(Hint: `a, b = b, a`.)

### Tugas 5 — Konversi Suhu
Simpan suhu `celsius := 30.0` (float64). Hitung fahrenheit dengan rumus
`F = C*9/5 + 32`, lalu cetak hasilnya.

### Tugas 6 — Konstanta
Buat konstanta `Phi = 3.14` dan `r = 7.0`. Hitung luas lingkaran
`Phi * r * r` dan cetak hasilnya.

## Tingkat Mahir

### Tugas 7 — Konversi Tipe Data
Buat `angkaInt := 42`. Ubah ke `float64` dan ke `string` (pakai `strconv.Itoa`),
lalu cetak ketiganya beserta tipenya (pakai `%T`).

### Tugas 8 — Identifikasi Error
Tulis (jangan dijalankan) kode berikut, lalu tuliskan di komentar apa yang salah
dan bagaimana membenarkannya:

```go
var x = 10
x := 20
```

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu memahami beda `var` dan `:=`.
- [ ] Kamu bisa menyebutkan zero value dari int, string, dan bool.
