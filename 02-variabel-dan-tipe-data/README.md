# Bab 02 — Variabel & Tipe Data

Di bab ini kita belajar **menyimpan data** di dalam program. Tempat menyimpan
data itu disebut **variabel**.

## Tujuan

Setelah bab ini kamu paham:

- Cara mendeklarasikan variabel dengan beberapa gaya (`var`, `:=`).
- Tipe-tipe data dasar Go: angka, teks, boolean, dan lain-lain.
- Apa itu **konstanta** (`const`).
- Konsep *zero value*.

## Konsep

### 1. Apa itu variabel?

Variabel adalah **kotak bernama** yang menyimpan nilai. Misalnya `umur` adalah
kotak yang isinya `25`. Nanti kita bisa ambil atau ubah isinya.

### 2. Tipe data dasar

| Tipe | Contoh nilai | Keterangan |
|------|--------------|------------|
| `int` | `25`, `-3` | Bilangan bulat |
| `float64` | `3.14`, `-0.5` | Bilangan desimal |
| `string` | `"Halo"` | Teks (selalu pakai tanda kutip ganda) |
| `bool` | `true`, `false` | Benar / salah |
| `rune` | `'a'` | Satu karakter (kutip tunggal) |
| `byte` | `byte(65)` | Satu byte data |

Go juga punya varian ukuran angka seperti `int8`, `int64`, `float32`, dan
sebagainya, tapi untuk pemula `int` dan `float64` sudah cukup.

### 3. Cara mendeklarasikan variabel

Ada beberapa gaya, hasilnya sama:

```go
// Gaya 1: pakai kata kunci var (tipe ditulis eksplisit)
var nama string
nama = "Budi"

// Gaya 2: var + inisialisasi sekaligus (tipe bisa dihilangkan)
var umur = 25

// Gaya 3: deklarasi singkat dengan := (paling umum)
kota := "Jakarta"
```

Aturan praktis:

- Gunakan `:=` untuk deklarasi singkat **di dalam fungsi**.
- Gunakan `var` jika ingin deklarasi tanpa langsung memberi nilai, atau untuk
  variabel tingkat paket.

> `:=` artinya "deklarasikan DAN beri nilai". Variabel yang sama tidak boleh
> dideklarasikan dua kali dengan `:=`.

### 4. Zero value

Jika variabel dideklarasikan tanpa nilai awal, Go memberinya **nilai default**
yang disebut *zero value*:

| Tipe | Zero value |
|------|-----------|
| `int`, `float64` | `0` |
| `string` | `""` (teks kosong) |
| `bool` | `false` |

Ini mencegah "nilai sampah" seperti di bahasa lain.

### 5. Konstanta

Konstanta adalah nilai yang **tidak bisa diubah** setelah dideklarasikan.
Dipakai untuk nilai yang memang tetap, misalnya angka Pi.

```go
const Pi = 3.14
```

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan dengan:

```bash
go run ./02-variabel-dan-tipe-data
```

## Latihan

1. Buat variabel `nama`, `umur`, dan `tinggi` (dalam meter) lalu cetak semuanya.
2. Buat konstanta `Phi` bernilai `3.14` dan hitung luas lingkaran
   `luas = Phi * r * r` (r = jari-jari).
3. Coba deklarasikan variabel tanpa nilai, lalu cetak untuk melihat *zero value*-nya.

> 💡 **Kenapa Go peduli sekali dengan tipe data?**
> Karena Go adalah bahasa *statically typed*: tipe diperiksa saat kompilasi.
> Ini menangkap banyak bug sejak awal — misalnya mencoba menjumlahkan angka
> dengan teks akan langsung error, bukan error aneh saat program berjalan.

Selanjutnya: **[Bab 03 — Operator](../03-operator)**.
