# Tugas Bab 11 — Goroutine & Channel

Kerjakan dari atas ke bawah. Coba sendiri dulu, baru cek jawaban.

## Tingkat Dasar

### Tugas 1 — Goroutine + WaitGroup
Jalankan 5 goroutine yang masing-masing mencetak angka 1–5, gunakan
`sync.WaitGroup` agar semua selesai sebelum program berakhir.

### Tugas 2 — Kirim & Terima di Channel
Buat channel `chan string`. Kirim satu pesan dari goroutine, terima di `main`,
lalu cetak.

### Tugas 3 — Buffered Channel
Buat buffered channel berkapasitas 3, kirim 3 nilai, lalu terima dan cetak.

## Tingkat Menengah

### Tugas 4 — Jumlahkan Lewat Channel
Jalankan 3 goroutine yang masing-masing mengirim sebuah angka ke channel yang
sama. Terima ketiganya di `main`, lalu cetak totalnya.

### Tugas 5 — select Dua Channel
Buat 2 channel. Kirim data ke salah satunya dari goroutine, lalu gunakan `select`
untuk menerima dari yang pertama siap.

### Tugas 6 — Ping-Pong
Buat 2 goroutine yang saling mengirim pesan bolak-balik lewat channel (misal 3
kali), lalu berhenti.

## Tingkat Mahir

### Tugas 7 — Worker Pool Sederhana
Buat channel `pekerjaan` berisi angka 1–5, dan 3 goroutine "worker" yang menerima
pekerjaan dan mencetak hasil kuadratnya. Gunakan `WaitGroup`.

### Tugas 8 — Batas Waktu (Timeout)
Gunakan `time.After` bersama `select` untuk membatasi menunggu channel maksimal
1 detik. Jika tidak ada data dalam 1 detik, cetak "timeout".

## Kriteria Selesai

- [ ] Semua program berhasil dijalankan.
- [ ] Kamu paham cara memulai goroutine dan pentingnya `WaitGroup`.
- [ ] Kamu paham kirim/terima di channel.
- [ ] Kamu paham buffered channel dan `select`.
