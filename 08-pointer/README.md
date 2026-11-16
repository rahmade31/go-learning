# Bab 08 — Pointer

Pointer adalah konsep yang awalnya terasa menyeramkan, tapi sebenarnya
sederhana: **pointer menyimpan alamat memori** dari sebuah nilai, bukan nilai
itu sendiri.

## Tujuan

Setelah bab ini kamu paham:

- Apa itu alamat memori.
- Operator `&` (ambil alamat) dan `*` (baca nilai dari alamat).
- Kenapa pointer berguna saat meneruskan data ke fungsi.
- Konsep *pass by value* vs *pass by reference*.

## Konsep

### 1. Analogi

Bayangkan sebuah **rumah** (nilai) yang punya **alamat** (pointer).

- `&rumah` = alamat rumah (pointer).
- `*alamat` = masuk ke rumah yang dituju (membaca nilainya).

```go
x := 10
p := &x      // p menyimpan ALAMAT x (p bertipe *int)
fmt.Println(p)  // alamat, misal 0xc0000140a0
fmt.Println(*p) // 10 (nilai di alamat itu)

*p = 20      // ubah nilai lewat pointer
fmt.Println(x) // 20 (x ikut berubah)
```

### 2. Pass by value vs pass by reference

Secara default, saat kita meneruskan variabel ke fungsi, Go **menyalin** nilainya.
Perubahan di dalam fungsi tidak memengaruhi variabel asli. Ini disebut
*pass by value*.

```go
func ubah(n int) {
    n = 99 // hanya mengubah salinan
}

x := 10
ubah(x)
fmt.Println(x) // tetap 10
```

Agar perubahan memengaruhi nilai asli, teruskan **pointer** (*pass by
reference*):

```go
func ubah(n *int) {
    *n = 99 // ubah nilai yang ditunjuk
}

ubah(&x)
fmt.Println(x) // 99
```

### 3. Pointer ke struct

Sangat umum untuk meneruskan pointer struct agar hemat memori dan bisa
mengubah field.

```go
type Orang struct { Nama string }

func gantiNama(o *Orang, namaBaru string) {
    o.Nama = namaBaru // Go otomatis "dereference": sama dengan (*o).Nama
}
```

> 💡 Ini sebabnya method pada bab 07 memakai pointer receiver `(m *Mahasiswa)`.

### 4. Kapan pakai pointer?

- Ingin **mengubah nilai asli** di dalam fungsi.
- Data berukuran **besar** (misal struct besar) agar tidak menyalin seluruhnya.
- Perlu menyatakan "tidak ada nilai" dengan `nil` (dibahas di bab interface).

### 5. Kata kunci `new`

`new(T)` membuat pointer ke nilai zero-value bertipe T:

```go
p := new(int) // *int yang menunjuk ke 0
fmt.Println(*p) // 0
```

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./08-pointer
```

## Latihan

1. Buat variabel `a = 5`, buat pointer ke `a`, lalu ubah nilai via pointer dan
   cetak `a`.
2. Tulis fungsi `ganda(n *int)` yang menggandakan nilai, lalu uji.
3. Buat fungsi yang menerima pointer struct dan mengubah salah satu fieldnya.

> ⚠️ **Hati-hati:** membaca pointer `nil` (`*p` saat `p == nil`) menyebabkan
> program *panic* (crash). Selalu pastikan pointer tidak `nil` sebelum dipakai.

Selanjutnya: **[Bab 09 — Interface](../09-interface)**.
