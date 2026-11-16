# Tugas Bab 09 — Interface

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Interface Suara
Buat interface `Hewan` dengan method `Suara() string`. Implementasikan pada
`Kucing` (mengembalikan `"Meong"`) dan `Anjing` (mengembalikan `"Guk"`).

### Tugas 2 — Fungsi Polimorfik
Buat fungsi `panggil(h Hewan)` yang mencetak suara hewan. Panggil untuk `Kucing`
dan `Anjing`.

### Tugas 3 — Interface Kosong
Buat variabel `var x any`, isi berturut-turut dengan `42`, `"teks"`, dan `3.14`,
cetak setiap kali.

## Tingkat Menengah

### Tugas 4 — Shape & Luas Total
Buat interface `Shape` dengan method `Luas() float64`. Implementasikan pada
`Lingkaran` dan `Persegi`. Buat fungsi `totalLuas(shapes []Shape) float64` yang
menjumlahkan luas semua shape.

### Tugas 5 — Type Assertion
Simpan nilai `"Halo"` ke `any`. Ambil kembali sebagai `string` memakai type
assertion dengan pola `nilai, ok := ...`. Cetak hasil dan `ok`.

### Tugas 6 — Type Switch
Buat `pilih(v any)` yang mencetak pesan berbeda untuk `int`, `string`, dan tipe
lain (gunakan type switch).

## Tingkat Mahir

### Tugas 7 — Banyak Tipe dalam Satu Slice
Karena `any` bisa menampung apa pun, buat slice `[]any` berisi `1`, `"dua"`,
`3.0`, lalu jelajahi dan cetak tipe tiap elemen memakai type switch.

### Tugas 8 — Desain Interface Sendiri
Rancang interface `Kendaraan` dengan method `JarakTempuh() float64` dan
`Deskripsi() string`. Implementasikan pada minimal 2 tipe berbeda (misal `Mobil`
dan `Sepeda`), lalu cetak info keduanya.

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham implementasi interface implisit.
- [ ] Kamu paham polimorfisme.
- [ ] Kamu paham type assertion dan type switch.
