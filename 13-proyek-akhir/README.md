# Bab 13 — Proyek Akhir: Aplikasi TODO CLI

Selamat! Ini bab terakhir. Kita merangkai **semua konsep** yang sudah dipelajari
menjadi satu aplikasi nyata: **aplikasi daftar tugas (TODO) berbasis terminal**.

## Fitur

- ➕ Tambah tugas baru.
- 📋 Lihat daftar tugas (beserta status selesai/belum).
- ✅ Tandai tugas selesai.
- ❌ Hapus tugas.
- 🚪 Keluar dari aplikasi.

## Konsep yang Dipakai

| Konsep | Penggunaan |
|--------|------------|
| Struct | `Task` (deskripsi + status) |
| Slice | Menyimpan daftar tugas |
| Fungsi & method | `tambahTugas`, `tampilkan`, method `TandaiSelesai` |
| `switch` | Menu utama |
| `for` | Perulangan menu & tampilan |
| `range` | Menjelajah daftar tugas |
| Error handling | Validasi input |
| `bufio` + `os.Stdin` | Membaca input pengguna |

## Cara Menjalankan

```bash
go run ./13-proyek-akhir
```

Lalu ikuti menu yang muncul. Contoh interaksi:

```
===== APLIKASI TODO =====
1. Tambah tugas
2. Lihat tugas
3. Tandai selesai
4. Hapus tugas
5. Keluar
Pilih menu (1-5): 1
Deskripsi tugas: Belajar Go
Tugas "Belajar Go" ditambahkan!
```

## Struktur Kode

Kode sengaja ditulis dalam **satu file** agar mudah diikuti, lalu dibagi menjadi
beberapa fungsi kecil agar rapi:

- `main` — perulangan menu utama.
- `tambahTugas` — menambah tugas.
- `tampilkanTugas` — mencetak semua tugas.
- `tandaiSelesai` — menandai tugas selesai.
- `hapusTugas` — menghapus tugas.
- `bacaInput` — pembaca input helper.

## Tantangan (Coba Kembangkan Sendiri)

1. Simpan daftar tugas ke **file** agar tidak hilang saat program ditutup
   (gunakan `os.WriteFile` / `os.ReadFile` — baca dokumentasinya!).
2. Tambah fitur **edit** deskripsi tugas.
3. Tambah **prioritas** pada setiap tugas (tinggi/sedang/rendah).
4. Ubah penyimpanan agar memakai **map** ber-key ID.

## Selamat! 🎉

Kamu sudah menempuh perjalanan dari "Hello World" hingga membangun aplikasi.
Dari sini, kamu bisa lanjut mempelajari:

- **HTTP & API** (paket `net/http`, atau framework `gin`/`echo`).
- **Database** (`database/sql` + driver seperti PostgreSQL/MySQL).
- **Testing** (`go test`, paket `testing`).
- **Baca dokumentasi resmi** di https://go.dev/doc/

Teruslah menulis kode — itu cara terbaik untuk mahir!
