# Bab 01 — Hello World (Program Pertamamu)

Pada bab ini kamu akan menulis program pertama: mencetak teks "Hello, World!"
ke layar. Ini tradisi wajib setiap programmer baru. 🎉

## Tujuan

Setelah bab ini kamu paham:

- Struktur dasar file Go.
- Apa itu `package` dan fungsi `main`.
- Cara memakai fungsi `fmt.Println` untuk mencetak teks.
- Cara menjalankan program dengan `go run`.

## Konsep

Setiap program Go terdiri dari **paket** (*package*). Paket adalah cara Go
mengelompokkan kode. Program yang bisa *dijalankan langsung* harus memakai
`package main` dan punya fungsi bernama `main`.

```go
package main   // (1) nama paket

import "fmt"   // (2) mengimpor paket bawaan untuk I/O

func main() {  // (3) titik awal program
    fmt.Println("Hello, World!")
}
```

Mari bedah baris per baris:

1. **`package main`** — menyatakan file ini bagian dari paket `main`. Paket
   `main` istimewa karena menjadi titik masuk program.
2. **`import "fmt"`** — mengambil fungsi-fungsi dari paket standar bernama
   `fmt` (singkatan *format*). Tanpa ini, kita tidak bisa pakai `Println`.
3. **`func main() { ... }`** — fungsi utama yang **selalu** dijalankan pertama
   kali saat program mulai. Tanpa fungsi ini, program tidak bisa dieksekusi.

### Apa itu `fmt.Println`?

- `fmt` = nama paket yang tadi kita impor.
- `Println` = nama fungsi di dalam paket `fmt`.
- `Println` mencetak teks lalu menambah **baris baru** di akhir.

Ada juga `fmt.Print` (tanpa baris baru) dan `fmt.Printf` (dengan format, kita
pelajari nanti).

## Kode Lengkap

Lihat file [`main.go`](./main.go) di folder ini.

## Cara Menjalankan

Dari folder utama (`learning-go`), ketik:

```bash
go run ./01-hello-world
```

Keluaran yang diharapkan:

```
Hello, World!
```

## Coba Ubah Sendiri

1. Ganti teks `"Hello, World!"` menjadi namamu, misal `"Halo, Budi!"`.
2. Tambahkan baris `fmt.Println` kedua untuk mencetak kalimat lain.
3. Jalankan lagi dan lihat hasilnya.

## Latihan

1. Cetak 3 baris berbeda menggunakan `Println`.
2. Coba `fmt.Print` dan perhatikan bedanya (apakah ada baris baru?).

> 💡 **Kenapa kurung kurawal `{` harus sebaris dengan `func main()`?**
> Di Go, aturan peletakan `{` sudah baku (dipaksakan oleh tool `gofmt`).
> Menaruh `{` di baris baru akan menyebabkan error. Ini salah satu cara Go
> menjaga agar semua kode terlihat seragam.

Selanjutnya: **[Bab 02 — Variabel & Tipe Data](../02-variabel-dan-tipe-data)**.
