# Tugas Bab 08 — Pointer

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Ambil Alamat
Buat `x := 10`. Buat pointer `p` yang menunjuk ke `x`. Cetak `x`, `p`, dan `*p`.

### Tugas 2 — Ubah Lewat Pointer
Dari `x` di atas, ubah nilainya menjadi `99` melalui `p`, lalu cetak `x` untuk
membuktikan `x` ikut berubah.

### Tugas 3 — `new`
Gunakan `new(int)` untuk membuat pointer ke int, lalu cetak nilai yang ditunjuk.

## Tingkat Menengah

### Tugas 4 — Tukar Nilai dengan Pointer
Buat fungsi `tukar(a, b *int)` yang menukar nilai dua variabel. Uji dengan
`a := 1, b := 2` dan pastikan setelah dipanggil `a == 2` dan `b == 1`.

### Tugas 5 — Ganda dengan Pointer
Buat fungsi `gandakan(n *int)` yang mengalikan nilai yang ditunjuk dengan 2.
Panggil dua kali dan cetak hasilnya.

### Tugas 6 — Pointer ke Struct
Buat struct `Orang` dengan field `Nama`. Buat fungsi `gantiNama(o *Orang, baru string)`
yang mengubah nama, lalu uji.

## Tingkat Mahir

### Tugas 7 — Pass by Value vs Reference
Tulis dua fungsi:
- `ubahNilai(n int)` (by value).
- `ubahNilaiPtr(n *int)` (by reference).

Buktikan dengan mencetak nilai sebelum dan sesudah pemanggilan bahwa hanya versi
pointer yang mengubah nilai asli.

### Tugas 8 — Nilai Nil
Buat `var p *int` (pointer `nil`). Jelaskan di komentar apa yang terjadi jika
kita mencoba membaca `*p`, dan bagaimana cara aman memeriksanya (`if p != nil`).

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham `&` dan `*`.
- [ ] Kamu paham pass by value vs pass by reference.
- [ ] Kamu tahu bahaya membaca pointer `nil`.
