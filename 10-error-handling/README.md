# Bab 10 — Error Handling

Di dunia nyata, banyak hal bisa gagal: file tidak ada, jaringan putus, input
salah. Go menangani kesalahan secara **eksplisit** menggunakan nilai `error`.

## Tujuan

Setelah bab ini kamu paham:

- Tipe `error` dan cara membuat/mengembalikannya.
- Pola idiomatik Go: periksa `if err != nil`.
- `defer` untuk memastikan sesuatu selalu dijalankan.
- `panic` dan `recover` untuk situasi darurat.

## Konsep

### 1. Tipe `error`

`error` adalah interface bawaan dengan satu method: `Error() string`.

```go
import "errors"

err := errors.New("terjadi kesalahan")
fmt.Println(err.Error()) // "terjadi kesalahan"
```

### 2. Pola idiomatik: kembalikan error

Fungsi yang bisa gagal sebaiknya mengembalikan `(hasil, error)`:

```go
func bagi(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("tidak bisa membagi dengan nol")
    }
    return a / b, nil
}

hasil, err := bagi(10, 0)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Hasil:", hasil)
}
```

> 💡 Di Go, kita **tidak** memakai try/catch. Sebaliknya, error adalah nilai
> biasa yang diperiksa eksplisit dengan `if err != nil`. Ini membuat alur error
> terlihat jelas.

### 3. Membungkus error (`fmt.Errorf`)

Untuk menambah konteks pada error, gunakan `fmt.Errorf` dengan `%w`:

```go
return fmt.Errorf("gagal membuka file %s: %w", namaFile, err)
```

### 4. `defer`

`defer` menunda eksekusi sebuah pernyataan sampai fungsi selesai. Sangat
berguna untuk membersihkan sumber daya (menutup file, koneksi).

```go
func bacaFile() {
    f, err := os.Open("data.txt")
    if err != nil {
        return
    }
    defer f.Close() // dijamin dijalankan saat fungsi selesai

    // ... proses file ...
}
```

Aturan `defer`: dieksekusi **LIFO** (Last In, First Out) — `defer` terakhir
dijalankan pertama.

### 5. `panic` & `recover`

- `panic` menghentikan program secara tiba-tiba (untuk error fatal yang tak
  bisa dipulihkan).
- `recover` menangkap panic dan melanjutkan program (biasanya dipasangkan
  dengan `defer`).

```go
func aman() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Pulih dari panic:", r)
        }
    }()
    panic("sesuatu buruk terjadi")
}
```

> ⚠️ `panic`/`recover` **bukan** pengganti error handling normal. Gunakan untuk
> situasi benar-benar darurat, bukan untuk alur bisnis biasa.

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./10-error-handling
```

## Latihan

1. Tulis fungsi `akar(n float64) (float64, error)` yang mengembalikan error
   jika `n` negatif.
2. Tulis fungsi dengan `defer` yang mencetak "selesai" saat fungsi berakhir.
3. Buat contoh `panic` yang ditangkap `recover`.

Selanjutnya: **[Bab 11 — Goroutine & Channel](../11-goroutine-dan-channel)**.
