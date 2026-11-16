# Tugas Bab 06 — Array, Slice, dan Map

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Jumlah & Rata-rata
Buat slice `angka := []int{10, 20, 30, 40, 50}`. Hitung jumlah dan rata-rata,
lalu cetak keduanya.

### Tugas 2 — Tambah Elemen
Buat slice kosong `[]string{}`, lalu gunakan `append` untuk menambah 3 nama
kota. Cetak sebelum dan sesudah penambahan.

### Tugas 3 — Akses Map
Buat map `harga := map[string]int{"apel": 5000, "jeruk": 8000}`. Cetak harga
apel, lalu tambah kunci `"mangga"` dengan nilai `10000`.

## Tingkat Menengah

### Tugas 4 — Balik Slice
Buat fungsi (atau kode di `main`) yang membalik urutan elemen slice
`[1 2 3 4 5]` menjadi `[5 4 3 2 1]`.

### Tugas 5 — Nilai Min & Maks
Dari slice angka, temukan nilai terkecil dan terbesar, lalu cetak keduanya.

### Tugas 6 — Hitung Frekuensi Kata (Map)
Dari slice `["apel", "jeruk", "apel", "mangga", "apel"]`, hitung berapa kali
setiap kata muncul memakai map.

## Tingkat Mahir

### Tugas 7 — Hapus Duplikat
Buat fungsi yang menerima `[]int` dan mengembalikan slice tanpa duplikat
(misal `[1,2,2,3,3,3]` → `[1,2,3]`).

### Tugas 8 — Slicing
Dari `data := []int{1, 2, 3, 4, 5, 6}`, cetak:
- Tiga elemen pertama (`data[:3]`)
- Tiga elemen terakhir (`data[3:]`)
- Elemen tengah (`data[2:4]`)

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham beda array dan slice.
- [ ] Kamu paham `append`, slicing, dan `range`.
- [ ] Kamu paham cara cek keberadaan kunci di map.
