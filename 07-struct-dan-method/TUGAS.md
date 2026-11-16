# Tugas Bab 07 — Struct & Method

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Struct Mahasiswa
Buat struct `Mahasiswa` dengan field `Nama`, `NIM`, dan `IPK`. Buat 2 nilai
struct dan cetak keduanya.

### Tugas 2 — Method Perkenalan
Tambahkan method `Perkenalan() string` pada `Mahasiswa` yang mengembalikan
`"Halo, saya <Nama> (NIM: <NIM>)"`. Cetak hasilnya.

### Tugas 3 — Akses & Ubah Field
Ubah nilai `IPK` salah satu mahasiswa, lalu cetak sebelum dan sesudah.

## Tingkat Menengah

### Tugas 4 — Rekening Bank
Buat struct `Rekening` dengan field `Nama` dan `Saldo`. Tambahkan method:
- `Setor(jumlah int)` — menambah saldo.
- `Tarik(jumlah int) bool` — mengurangi saldo jika cukup, mengembalikan `false` jika tidak.

Uji dengan beberapa transaksi.

### Tugas 5 — Pointer Receiver
Buat method `NaikkanIPK(tambah float64)` dengan **pointer receiver** pada
`Mahasiswa`. Buktikan bahwa perubahan tersimpan (cetak IPK sebelum & sesudah).

### Tugas 6 — Embedding
Buat struct `Orang` (field `Nama`) dan `Pegawai` yang meng-embed `Orang` plus
field `Gaji`. Cetak `Nama` pegawai langsung (tanpa `Orang.Nama`).

## Tingkat Mahir

### Tugas 7 — Method Rantai (Chaining)
Buat struct `Kotak` dengan field `Panjang`, `Lebar`, `Tinggi`. Tambahkan method
`Volume() int`. Lalu tulis fungsi yang menerima banyak `Kotak` dan mengembalikan
total volumenya.

### Tugas 8 — Struct Bersarang
Buat struct `Alamat` (Kota, Provinsi) dan struct `Mahasiswa` yang memiliki field
`Alamat`. Cetak kota tempat tinggal mahasiswa.

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham cara membuat struct dan akses field.
- [ ] Kamu paham value receiver vs pointer receiver.
- [ ] Kamu paham embedding (komposisi).
