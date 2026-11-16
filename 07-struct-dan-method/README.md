# Bab 07 — Struct & Method

Sejauh ini kita menyimpan data sederhana (angka, teks). Tapi dunia nyata butuh
data yang punya **beberapa bagian**. Misalnya seorang "mahasiswa" punya nama,
NIM, dan jurusan. Di sinilah **struct** berperan.

## Tujuan

Setelah bab ini kamu paham:

- Mendefinisikan struct sebagai tipe data custom.
- Membuat nilai struct dan membaca/mengubah field-nya.
- Menempelkan **method** pada struct.
- Struct bersarang dan *embedding* (pewarisan gaya Go).

## Konsep

### 1. Mendefinisikan struct

```go
type Mahasiswa struct {
    Nama    string
    NIM     string
    Jurusan string
    IPK     float64
}
```

`type` membuat tipe baru, `struct` mendefinisikan kumpulan *field*.

### 2. Membuat nilai struct

```go
// Cara 1: urutan field (rapuh, tidak disarankan)
m1 := Mahasiswa{"Budi", "101", "Informatika", 3.5}

// Cara 2: nama field (lebih jelas, disarankan)
m2 := Mahasiswa{
    Nama:    "Ani",
    NIM:     "102",
    Jurusan: "Sistem Informasi",
    IPK:     3.8,
}

fmt.Println(m2.Nama) // akses field dengan titik
m2.IPK = 3.9         // ubah field
```

### 3. Method

Method adalah **fungsi yang menempel pada struct**. Sintaksnya sama dengan
fungsi, tapi punya *receiver* (penerima) di antara `func` dan nama method.

```go
// Receiver ditulis dalam kurung: (m Mahasiswa)
func (m Mahasiswa) Perkenalan() string {
    return "Halo, saya " + m.Nama + " dari " + m.Jurusan
}

fmt.Println(m2.Perkenalan())
```

### 4. Pointer receiver

Method di atas menerima **salinan** struct, sehingga perubahan di dalam method
tidak mengubah nilai asli. Untuk mengubah nilai asli, gunakan **pointer
receiver** `(m *Mahasiswa)`:

```go
func (m *Mahasiswa) NaikkanIPK(x float64) {
    m.IPK += x // perubahan tersimpan ke nilai asli
}
```

> Detail pointer dibahas tuntas di bab 08. Untuk sekarang, cukup ingat: pakai
> `*` pada receiver jika method perlu mengubah data.

### 5. Embedding (komposisi)

Go tidak punya pewarisan (*inheritance*) seperti OOP klasik. Sebagai gantinya,
Go memakai **komposisi**: struct bisa "menempelkan" struct lain.

```go
type Orang struct {
    Nama string
}

type Karyawan struct {
    Orang       // embedding: Karyawan "mewarisi" field Orang
    Gaji int
}

k := Karyawan{Orang{"Budi"}, 5000}
fmt.Println(k.Nama) // akses langsung (promoted field)
```

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./07-struct-dan-method
```

## Latihan

1. Buat struct `Buku` (Judul, Penulis, Tahun) lalu buat 2 buku dan cetak.
2. Tambah method `Info()` yang mengembalikan ringkasan buku.
3. Buat struct `Mobil` dan `Engine`, lalu lakukan embedding.
4. Buat method pointer receiver yang mengubah salah satu field.

Selanjutnya: **[Bab 08 — Pointer](../08-pointer)**.
