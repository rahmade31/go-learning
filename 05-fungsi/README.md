# Bab 05 — Fungsi

Fungsi adalah **blok kode bernama** yang bisa dipanggil berulang kali. Ini
fondasi untuk menulis kode yang rapi, tidak berulang, dan mudah dirawat.

## Tujuan

Setelah bab ini kamu paham:

- Cara mendefinisikan dan memanggil fungsi.
- Parameter dan nilai kembalian (*return value*).
- Nilai kembalian banyak (*multiple return*).
- Nilai kembalian bernama (*named return*).
- Fungsi *variadic* (jumlah parameter tak tentu).
- *Closure* dan fungsi sebagai nilai.

## Konsep

### 1. Definisi dasar

```go
func tambah(a int, b int) int {
    return a + b
}

hasil := tambah(3, 4) // 7
```

Struktur: `func nama(parameter) tipeReturn { ... return ... }`.

Jika dua parameter berurutan bertipe sama, bisa disingkat:

```go
func tambah(a, b int) int { ... }
```

### 2. Multiple return

Keunggulan khas Go: sebuah fungsi bisa mengembalikan **beberapa nilai**.
Sangat berguna untuk mengembalikan hasil + status error.

```go
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("pembagian dengan nol")
    }
    return a / b, nil
}

hasil, err := bagi(10, 0)
```

### 3. Named return

Nilai kembalian bisa diberi nama. Ini membuat `return` tanpa argumen
mengembalikan nilai yang sudah diisi.

```go
func luasPersegi(sisi float64) (luas float64) {
    luas = sisi * sisi
    return // otomatis mengembalikan "luas"
}
```

### 4. Fungsi variadic

Menerima jumlah parameter tak tentu, ditandai `...`.

```go
func jumlahkan(angka ...int) int {
    total := 0
    for _, n := range angka {
        total += n
    }
    return total
}

jumlahkan(1, 2, 3)     // 6
jumlahkan(10, 20)      // 30
```

### 5. Closure & fungsi sebagai nilai

Di Go, fungsi adalah *first-class citizen*: bisa disimpan dalam variabel,
dijadikan parameter, atau dikembalikan. *Closure* adalah fungsi yang
"mengingat" variabel di sekitarnya.

```go
tambah := func(a, b int) int { return a + b }

pembuat := func() func() int {
    hitung := 0
    return func() int {
        hitung++
        return hitung
    }
}
counter := pembuat()
counter() // 1
counter() // 2
```

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./05-fungsi
```

## Latihan

1. Buat fungsi `kali(a, b int) int` dan uji hasilnya.
2. Buat fungsi `infoNama` yang mengembalikan (nama, umur).
3. Buat fungsi variadic `rataRata` yang menghitung rata-rata banyak angka.
4. Buat closure `pembuatCounter` yang mengembalikan penghitung.

> 💡 **Kenapa fungsi penting?**
> Fungsi membuat kode bisa **dipakai ulang** (tidak copy-paste), **mudah diuji**,
> dan **mudah dibaca** karena memecah masalah besar jadi langkah-langkah kecil.

Selanjutnya: **[Bab 06 — Array, Slice, Map](../06-array-slice-map)**.
