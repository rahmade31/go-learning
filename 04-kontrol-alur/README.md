# Bab 04 — Kontrol Alur (if, switch, for)

Program yang berguna tidak selalu berjalan lurus ke bawah. Ia harus bisa
**mengambil keputusan** dan **mengulang**. Itulah kontrol alur.

## Tujuan

Setelah bab ini kamu paham:

- Percabangan dengan `if`, `else if`, dan `else`.
- Percabangan dengan `switch`.
- Perulangan dengan `for` (satu-satunya perulangan di Go).

## Konsep

### 1. `if` / `else if` / `else`

Digunakan untuk menjalankan kode hanya jika suatu kondisi bernilai `true`.

```go
umur := 20

if umur >= 17 {
    fmt.Println("Sudah boleh punya KTP")
} else {
    fmt.Println("Belum boleh punya KTP")
}
```

Bisa ditambah cabang tengah dengan `else if`:

```go
nilai := 75
if nilai >= 90 {
    fmt.Println("A")
} else if nilai >= 80 {
    fmt.Println("B")
} else if nilai >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("D")
}
```

> 💡 **Khas Go:** `if` bisa menyertakan pernyataan pendek sebelum kondisi,
> dipisah titik koma. Variabelnya hanya berlaku di dalam blok `if`:
>
> ```go
> if umur := 20; umur >= 17 {
>     fmt.Println("Umur", umur, "sudah boleh KTP")
> }
> ```

### 2. `switch`

`switch` adalah cara lebih rapi untuk menulis banyak `if else if`.

```go
hari := "Senin"
switch hari {
case "Senin":
    fmt.Println("Awal pekan")
case "Sabtu", "Minggu": // beberapa nilai sekaligus
    fmt.Println("Akhir pekan")
default:
    fmt.Println("Hari biasa")
}
```

> ⚠️ Berbeda dari bahasa C/Java: di Go, `case` **tidak otomatis lanjut** ke case
> berikutnya. Tidak perlu menulis `break`. Ini menghindari bug yang umum.

### 3. `for` (perulangan)

Go **hanya punya `for`** — tidak ada `while` atau `do while`. Tapi `for` cukup
fleksibel untuk meniru semuanya.

```go
// Bentuk 1: klasik (inisialisasi; kondisi; peningkatan)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Bentuk 2: hanya kondisi (seperti while)
n := 0
for n < 3 {
    fmt.Println(n)
    n++
}

// Bentuk 3: tanpa kondisi (perulangan selamanya)
for {
    fmt.Println("loop tak berujung...")
    break // hentikan
}
```

Kata kunci penting di dalam perulangan:

- `break` — berhenti total.
- `continue` — lompat ke iterasi berikutnya.

### 4. `range`

Untuk menjelajah koleksi (slice, map, string), Go punya `range`. Kita akan
pakai lebih banyak di bab 06.

```go
for index, huruf := range "Halo" {
    fmt.Println(index, string(huruf))
}
```

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./04-kontrol-alur
```

## Latihan

1. Cetak angka 1 sampai 10 memakai `for`.
2. Cetak hanya angka genap dari 1 sampai 10 (pakai `%` dan `continue`).
3. Buat program kecil: jika `nilai >= 75` cetak "Lulus", selain itu "Remedial".
4. Tulis `switch` untuk nama buah favoritmu.

Selanjutnya: **[Bab 05 — Fungsi](../05-fungsi)**.
