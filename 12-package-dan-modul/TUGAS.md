# Tugas Bab 12 — Package & Modul

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Tambah Fungsi ke Paket
Di paket `matematika`, tambahkan fungsi `Modulo(a, b int) int` dan `Pangkat(a, b int) int`.
Gunakan keduanya di `main.go`.

### Tugas 2 — Pahami Exported/Unexported
Di paket `matematika`, tambahkan fungsi huruf kecil `internal()` yang tidak bisa
dipanggil dari `main`. Catat di komentar apa yang terjadi jika dipaksa memanggilnya.

### Tugas 3 — Impor dan Gunakan
Di `main.go`, impor dan gunakan semua fungsi yang sudah ada di paket `matematika`.

## Tingkat Menengah

### Tugas 4 — Paket Baru: konversi
Buat folder & paket baru `konversi` dengan fungsi:
- `CelsiusKeFahrenheit(c float64) float64`
- `FahrenheitKeCelsius(f float64) float64`

Impor dan uji dari `main.go`.

### Tugas 5 — Paket String Utilitas
Buat paket `teks` berisi fungsi:
- `Balik(s string) string` — membalik string.
- `Kapital(s string) string` — mengubah huruf pertama jadi kapital (gunakan `strings.ToUpper`).

Gunakan keduanya di `main.go`.

### Tugas 6 — Struktur Multi-file
Pisahkan paket `matematika` menjadi dua file (misal `tambah.go` dan `kali.go`)
di folder yang sama, pastikan masih berfungsi tanpa mengubah `main.go`.

## Tingkat Mahir

### Tugas 7 — Dokumentasi Paket
Tulis komentar dokumentasi (dimulai nama fungsi) pada semua fungsi ekspor di
paket `matematika`, lalu jalankan `go doc ./12-package-dan-modul/matematika`
untuk melihatnya.

### Tugas 8 — Analisis Import Path
Jelaskan (di komentar) kenapa jalur impor dimulai dengan `learning-go/...`.
Apa yang terjadi jika nama modul di `go.mod` diganti?

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan dengan `go run ./12-package-dan-modul`.
- [ ] Kamu paham aturan huruf besar/kecil (exported vs unexported).
- [ ] Kamu bisa membuat dan mengimpor paket sendiri.
- [ ] Kamu paham fungsi `go.mod` sebagai penentu jalur impor.
