# Tugas Bab 13 — Proyek Akhir

Ini adalah **proyek pengembangan**. Kerjakan satu per satu; setiap tugas
memperkaya aplikasi TODO yang sudah ada di `main.go`.

## Tingkat Dasar

### Tugas 1 — Jalankan & Pahami
Jalankan aplikasi dengan `go run ./13-proyek-akhir`. Coba semua menu (tambah,
lihat, selesai, hapus, keluar). Tulis di catatanmu alur setiap menu.

### Tugas 2 — Tambah Menu "Hapus Semua"
Tambahkan opsi menu baru untuk menghapus **semua** tugas sekaligus. Konfirmasi
dulu ke pengguna sebelum menghapus.

### Tugas 3 — Jumlah Tugas
Saat menampilkan daftar, tampilkan juga jumlah total tugas dan berapa yang sudah
selesai.

## Tingkat Menengah

### Tugas 4 — Fitur Edit
Tambahkan menu "Edit tugas" yang mengubah deskripsi tugas berdasarkan nomornya.

### Tugas 5 — Prioritas
Tambahkan field `Prioritas` (misal "tinggi", "sedang", "rendah") pada struct
`Task`. Tampilkan prioritas di daftar, dan izinkan pengguna mengaturnya saat
menambah tugas.

### Tugas 6 — Urutkan Tugas
Tambahkan menu untuk menampilkan tugas **hanya yang belum selesai** dan
**hanya yang selesai**.

## Tingkat Mahir

### Tugas 7 — Simpan ke File
Simpan daftar tugas ke file (misal `tugas.txt`) saat keluar, dan muat kembali
saat program dibuka. Gunakan `os.ReadFile` dan `os.WriteFile`.

### Tugas 8 — Batas Waktu (Due Date)
Tambahkan field `Tenggat` (string tanggal). Tampilkan tenggat di daftar, dan
beri tanda khusus (misal `!`) untuk tugas yang tenggatnya sudah lewat (bandingkan
string tanggal sederhana).

### Tugas 9 — Pisahkan ke Paket
Refactor aplikasi: pindahkan logika TODO (struct `Task`, fungsi tambah/tampil/
hapus) ke paket terpisah (misal folder `todo/`), biarkan `main.go` hanya
menangani menu & input.

## Kriteria Selesai

- [ ] Aplikasi berjalan tanpa error.
- [ ] Fitur yang kamu tambahkan berfungsi dan sudah diuji manual.
- [ ] Kode tetap rapi dan terbagi dalam fungsi yang jelas.

## Ide Lanjutan

- Ganti penyimpanan ke format JSON (paket `encoding/json`).
- Tambahkan ID unik per tugas.
- Buat versi web sederhana dengan `net/http`.
