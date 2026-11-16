# Belajar Go (Golang) — Zero to Hero

Selamat datang! Repositori ini adalah **template belajar Go dari nol sampai mahir**
("zero to hero") dengan penjelasan lengkap dan bahasa yang mudah dipahami.

Go (sering disebut **Golang**) adalah bahasa pemrograman modern yang dibuat oleh
Google. Go terkenal karena:

- **Sederhana** — sintaksnya bersih dan mudah dibaca.
- **Cepat** — dikompilasi langsung ke kode mesin (bukan dijalankan oleh interpreter).
- **Kuat untuk konkurensi** — punya goroutine dan channel bawaan.
- **Cocok untuk backend** — banyak dipakai untuk API, microservice, dan alat CLI.

## Peta Belajar (Roadmap)

Ikuti urutan berikut dari atas ke bawah. Setiap bab berisi tiga file:

- `README.md` — penjelasan konsep.
- `main.go` — kode contoh yang bisa langsung dijalankan.
- `TUGAS.md` — latihan wajib bertingkat (Dasar → Menengah → Mahir).

| No | Bab | Topik | Folder |
|----|-----|-------|--------|
| 00 | Persiapan | Instalasi & alat kerja | [`00-persiapan`](./00-persiapan) |
| 01 | Hello World | Struktur program pertama | [`01-hello-world`](./01-hello-world) |
| 02 | Variabel & Tipe Data | Deklarasi, konstanta, tipe dasar | [`02-variabel-dan-tipe-data`](./02-variabel-dan-tipe-data) |
| 03 | Operator | Aritmatika, perbandingan, logika | [`03-operator`](./03-operator) |
| 04 | Kontrol Alur | `if`, `switch`, perulangan `for` | [`04-kontrol-alur`](./04-kontrol-alur) |
| 05 | Fungsi | Fungsi, multiple return, variadic, closure | [`05-fungsi`](./05-fungsi) |
| 06 | Array, Slice, Map | Koleksi data | [`06-array-slice-map`](./06-array-slice-map) |
| 07 | Struct & Method | Tipe data custom | [`07-struct-dan-method`](./07-struct-dan-method) |
| 08 | Pointer | Alamat memori | [`08-pointer`](./08-pointer) |
| 09 | Interface | Polimorfisme | [`09-interface`](./09-interface) |
| 10 | Error Handling | `error`, `defer`, `panic`, `recover` | [`10-error-handling`](./10-error-handling) |
| 11 | Goroutine & Channel | Konkurensi | [`11-goroutine-dan-channel`](./11-goroutine-dan-channel) |
| 12 | Package & Modul | Organisasi kode | [`12-package-dan-modul`](./12-package-dan-modul) |
| 13 | Proyek Akhir | Aplikasi CLI TODO lengkap | [`13-proyek-akhir`](./13-proyek-akhir) |

## Cara Menjalankan Kode

Setiap bab adalah folder terpisah dengan fungsi `main`-nya sendiri. Untuk
menjalankan salah satu bab, buka terminal di folder ini lalu ketik:

```bash
# Menjalankan bab 01
go run ./01-hello-world

# Menjalankan proyek akhir
go run ./13-proyek-akhir
```

> **Catatan:** Pastikan Go sudah terinstal di komputermu. Lihat bab
> [`00-persiapan`](./00-persiapan) untuk panduan instalasi.

## Tips Belajar

1. **Jangan hanya dibaca** — ketik ulang kodenya sendiri, jangan copy-paste.
2. **Ubah-ubah kodenya** — coba modifikasi nilai, tambah logika, lihat hasilnya.
3. **Kerjakan latihan** di file `TUGAS.md` setiap bab sebelum lanjut ke bab
   berikutnya.
4. **Pahami konsep**, bukan sekadar hafal sintaks. Setiap bab menjelaskan
   "kenapa" di balik "bagaimana".
5. **Satu bab selesai dulu**, baru lanjut ke bab berikutnya.

Selamat belajar! 🚀
