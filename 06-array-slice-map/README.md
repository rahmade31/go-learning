# Bab 06 — Array, Slice, dan Map

Kita sering perlu menyimpan **banyak data sekaligus**. Go menyediakan tiga
struktur utama: **array**, **slice**, dan **map**.

## Tujuan

Setelah bab ini kamu paham:

- Perbedaan array dan slice.
- Cara membuat, membaca, dan mengubah slice.
- Operasi umum slice (`append`, `copy`, slicing).
- Cara memakai map untuk pasangan kunci-nilai.

## Konsep

### 1. Array

Array adalah kumpulan elemen dengan **ukuran tetap** yang ditentukan saat
deklarasi.

```go
var nilai [3]int      // array berisi 3 int, semuanya 0
nilai[0] = 90
nilai[1] = 85
nilai[2] = 78

nama := [2]string{"Budi", "Ani"}
fmt.Println(len(nama)) // 2 (panjang array)
```

Array jarang dipakai langsung karena ukurannya kaku. Yang lebih sering dipakai
adalah **slice**.

### 2. Slice

Slice adalah **tampilan fleksibel** atas array. Ukurannya bisa bertambah
dengan `append`. Ini struktur koleksi paling umum di Go.

```go
angka := []int{1, 2, 3} // slice literal
angka = append(angka, 4) // menambah elemen -> [1 2 3 4]

// Membuat slice kosong dengan make
kota := make([]string, 0, 5) // panjang 0, kapasitas 5
kota = append(kota, "Jakarta", "Bandung")
```

### 3. Slicing

Kita bisa mengambil "potongan" slice dengan notasi `[awal:akhir]`.

```go
h := []int{10, 20, 30, 40, 50}
h[1:3] // [20 30] (indeks 1 sampai sebelum 3)
h[:2]  // [10 20]
h[2:]  // [30 40 50]
```

### 4. Menjelajah dengan `range`

```go
buah := []string{"apel", "jeruk", "mangga"}
for i, nama := range buah {
    fmt.Println(i, nama)
}
```

- `i` = indeks, `nama` = nilai.
- Jika indeks tidak dipakai, gunakan `_` (blank identifier): `for _, nama := range buah`.

### 5. Map

Map menyimpan data sebagai **pasangan kunci-nilai** (seperti kamus).

```go
usia := map[string]int{
    "Budi": 25,
    "Ani":  30,
}
usia["Caca"] = 22           // tambah
usia["Budi"] = 26           // ubah
delete(usia, "Ani")         // hapus

umur, ada := usia["Budi"]   // cek keberadaan kunci
fmt.Println(umur, ada)      // 26 true
```

- `ada` bernilai `false` jika kunci tidak ada (ini pola umum di Go).

### 6. Ringkasan perbedaan

| Struktur | Ukuran | Ubah ukuran? |
|----------|--------|--------------|
| Array | Tetap | Tidak |
| Slice | Fleksibel | Ya (`append`) |
| Map | Fleksibel | Ya (tambah/hapus kunci) |

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./06-array-slice-map
```

## Latihan

1. Buat slice berisi 5 angka, lalu cetak total dan rata-ratanya.
2. Gunakan `append` untuk menambah angka ke slice, lalu cetak sebelum & sesudah.
3. Buat map `harga` untuk 3 barang, lalu tampilkan semua kunci-nilai dengan `range`.
4. Coba hapus satu kunci dari map dan pastikan ia hilang.

> 💡 **Kenapa slice lebih disukai daripada array?**
> Karena slice bisa tumbuh dan sangat murah untuk diteruskan ke fungsi
> (hanya menyalin "penunjuk" ke data, bukan seluruh data). Kita bahas lebih
> dalam di bab pointer.

Selanjutnya: **[Bab 07 — Struct & Method](../07-struct-dan-method)**.
