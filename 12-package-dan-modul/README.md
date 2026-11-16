# Bab 12 — Package & Modul

Saat program makin besar, kita perlu **mengorganisasi kode** ke dalam
paket-paket dan modul. Bab ini mengajarkan caranya.

## Tujuan

Setelah bab ini kamu paham:

- Apa itu paket dan kenapa penting.
- Aturan *exported* vs *unexported* (huruf besar/kecil).
- Membuat paket sendiri dan mengimpornya.
- Apa itu modul dan `go.mod`.

## Konsep

### 1. Paket

Paket adalah **folder berisi file `.go`** yang mengelompokkan kode terkait.
Setiap file di folder yang sama harus punya deklarasi `package` yang sama
(kecuali paket `main`).

Di bab ini kita membuat paket bernama `matematika` di folder
[`matematika/`](./matematika). Fungsinya menyimpan operasi hitung.

### 2. Exported vs unexported

Ini aturan penting dan unik di Go: **nama yang diawali huruf besar dapat
diakses dari paket lain** (disebut *exported*). Nama huruf kecil hanya bisa
dipakai di dalam paketnya sendiri (*unexported*).

```go
package matematika

// Tambah diawali huruf besar -> bisa diimpor paket lain.
func Tambah(a, b int) int { return a + b }

// helper diawali huruf kecil -> hanya untuk internal paket ini.
func helper() {}
```

### 3. Mengimpor paket sendiri

Di `main.go`, kita impor paket tadi menggunakan **jalur modul**:

```go
import "learning-go/12-package-dan-modul/matematika"

fmt.Println(matematika.Tambah(3, 4))
```

Jalur impor = nama modul (`learning-go`) + jalur folder dari akar modul.

### 4. Modul & `go.mod`

Modul adalah **unit versi kode** yang berisi satu atau lebih paket. File
`go.mod` di akar repo mendeklarasikan modul ini:

```
module learning-go

go 1.22
```

- Baris `module` = nama modul (dipakai sebagai awalan jalur impor).
- Baris `go` = versi Go minimum.

### 5. Mengimpor paket eksternal

Untuk memakai paket dari internet (misal dari GitHub), gunakan `go get`:

```bash
go get github.com/gin-gonic/gin
```

Ini akan menambahkannya ke `go.mod` dan mengunduh kodenya. (Kita tidak pakai
dependensi eksternal di tutorial ini agar tetap sederhana.)

## Kode Lengkap

- [`matematika/matematika.go`](./matematika/matematika.go) — paket custom.
- [`main.go`](./main.go) — pengguna paket.

Jalankan:

```bash
go run ./12-package-dan-modul
```

## Latihan

1. Tambah fungsi `Kali` dan `Kurang` ke paket `matematika`, lalu gunakan di `main.go`.
2. Buat paket baru `konversi` berisi fungsi `CelsiusKeFahrenheit`, lalu impor.
3. Coba ubah salah satu fungsi menjadi huruf kecil dan lihat error saat diimpor.

> 💡 **Mengapa huruf besar/kecil?**
> Ini membuat "API publik" sebuah paket terlihat jelas hanya dari nama. Tidak
> perlu kata kunci `public`/`private` seperti bahasa lain.

Selanjutnya: **[Bab 13 — Proyek Akhir](../13-proyek-akhir)** — saatnya merangkai
semua yang sudah dipelajari! 🚀
