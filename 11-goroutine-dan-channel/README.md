# Bab 11 — Goroutine & Channel (Konkurensi)

Inilah keunggulan Go yang paling terkenal: **konkurensi**. Dengan kata kunci
`go`, kita bisa menjalankan banyak tugas secara bersamaan dengan sangat mudah.

## Tujuan

Setelah bab ini kamu paham:

- Apa itu goroutine dan cara memulainya (`go`).
- Apa itu channel dan cara menggunakannya untuk komunikasi antar goroutine.
- Sinkronisasi dengan `sync.WaitGroup`.
- Pola umum: buffered channel, select.

## Konsep

### 1. Goroutine

Goroutine adalah **eksekusi ringan** yang berjalan bersamaan. Memulainya cukup
menambahkan kata kunci `go` di depan pemanggilan fungsi:

```go
func cetak(n int) {
    fmt.Println(n)
}

go cetak(1) // jalan bersamaan dengan kode berikutnya
go cetak(2)
```

Goroutine jauh lebih ringan daripada thread sistem; satu program bisa menjalankan
ribuan goroutine.

### 2. Channel

Channel adalah **saluran komunikasi** antar goroutine. Satu goroutine mengirim
data, goroutine lain menerimanya.

```go
ch := make(chan int) // channel untuk int

go func() {
    ch <- 42 // kirim nilai ke channel
}()

nilai := <-ch // terima nilai dari channel (menunggu sampai ada data)
fmt.Println(nilai)
```

- `ch <- nilai` : mengirim.
- `nilai := <-ch` : menerima (memblokir sampai data tersedia).

### 3. Buffered channel

Channel default (unbuffered) mengharuskan pengirim dan penerima siap
bersamaan. Buffered channel punya antrean:

```go
ch := make(chan int, 3) // bisa menampung 3 nilai tanpa menunggu
ch <- 1
ch <- 2
ch <- 3
```

### 4. `select`

`select` menunggu beberapa operasi channel sekaligus, dan menjalankan yang
pertama siap:

```go
select {
case v := <-ch1:
    fmt.Println("dari ch1:", v)
case v := <-ch2:
    fmt.Println("dari ch2:", v)
default:
    fmt.Println("belum ada data")
}
```

### 5. `sync.WaitGroup`

Untuk menunggu semua goroutine selesai, gunakan `WaitGroup`:

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}

wg.Wait() // tunggu semua selesai
```

> ⚠️ **Tanpa `wg.Wait()`**, fungsi `main` bisa selesai lebih dulu sebelum
> goroutine selesai, sehingga output tidak muncul. Ini jebakan umum pemula.

### 6. Kenapa penting?

Konkurensi memungkinkan program menangani banyak tugas (misal banyak permintaan
HTTP) secara efisien, tanpa menunggu satu per satu.

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./11-goroutine-dan-channel
```

## Latihan

1. Jalankan 3 goroutine yang mencetak angka, gunakan `WaitGroup`.
2. Buat channel, kirim 2 nilai dari goroutine, terima dan cetak.
3. Coba buffered channel dan perhatikan bedanya.

Selanjutnya: **[Bab 12 — Package & Modul](../12-package-dan-modul)**.
