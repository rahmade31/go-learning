# Tugas Bab 03 — Operator

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Kalkulator Dua Angka
Buat `a := 15` dan `b := 4`. Cetak hasil `+`, `-`, `*`, `/`, dan `%`:

```
15 + 4 = 19
15 - 4 = 11
15 * 4 = 60
15 / 4 = 3
15 % 4 = 3
```

### Tugas 2 — Genap atau Ganjil
Buat `angka := 7`. Dengan operator `%`, cetak apakah angka tersebut genap atau
ganjil (pakai `bool`).

### Tugas 3 — Operator Penugasan
Buat `nilai := 10`. Terapkan `+= 5`, `-= 2`, `*= 3`, `/= 4` secara berurutan,
cetak hasil setelah setiap langkah.

## Tingkat Menengah

### Tugas 4 — Cek Kelulusan (Logika)
Buat `nilai := 75` dan `kehadiran := 90`. Cetak `true`/`false` untuk:
- Lulus jika `nilai >= 60` **dan** `kehadiran >= 75`.
- Remedial jika `nilai < 60` **atau** `kehadiran < 75`.

### Tugas 5 — Luas & Keliling Lingkaran
Buat `r := 7.0`. Hitung luas (`3.14 * r * r`) dan keliling (`2 * 3.14 * r`),
cetak keduanya.

### Tugas 6 — Gabung Nama
Buat `depan := "Budi"` dan `belakang := "Santoso"`. Gabung menjadi nama lengkap
dengan spasi di tengah, lalu cetak.

## Tingkat Mahir

### Tugas 7 — Pembagian Desimal vs Bulat
Jelaskan (di komentar) kenapa `5 / 2` menghasilkan `2`, sedangkan `5.0 / 2.0`
menghasilkan `2.5`. Tulis program kecil yang membuktikannya.

### Tugas 8 — Ekspresi Kombinasi
Buat `a := 10`, `b := 3`, `c := 5`. Hitung `(a + b) * c / b` dan `a % b + c`,
lalu cetak hasil dan jelaskan urutan pengerjaannya.

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham operator `%` dan pembagian int vs float.
- [ ] Kamu paham beda `&&`, `||`, dan `!`.
