# Bab 03 — Operator

Operator adalah simbol yang dipakai untuk **melakukan operasi** pada data:
menjumlahkan, membandingkan, menggabungkan, dan lain-lain.

## Tujuan

Setelah bab ini kamu paham:

- Operator aritmatika (`+`, `-`, `*`, `/`, `%`).
- Operator perbandingan (`==`, `!=`, `<`, `>`, `<=`, `>=`).
- Operator logika (`&&`, `||`, `!`).
- Operator penugasan (`=`, `+=`, `-=`, dll.).

## Konsep

### 1. Operator Aritmatika (untuk angka)

| Operator | Arti | Contoh | Hasil |
|----------|------|--------|-------|
| `+` | Tambah | `5 + 2` | `7` |
| `-` | Kurang | `5 - 2` | `3` |
| `*` | Kali | `5 * 2` | `10` |
| `/` | Bagi | `5 / 2` | `2` (karena int, dibulatkan ke bawah) |
| `%` | Sisa bagi (modulo) | `5 % 2` | `1` |

> ⚠️ Pembagian dua bilangan `int` menghasilkan `int`. Untuk hasil desimal,
> gunakan `float64`: `5.0 / 2.0` → `2.5`.

### 2. Operator Perbandingan (menghasilkan `bool`)

| Operator | Arti | Contoh |
|----------|------|--------|
| `==` | Sama dengan | `5 == 5` → `true` |
| `!=` | Tidak sama | `5 != 5` → `false` |
| `<` | Kurang dari | `3 < 5` → `true` |
| `>` | Lebih dari | `3 > 5` → `false` |
| `<=` | Kurang dari/sama | `5 <= 5` → `true` |
| `>=` | Lebih dari/sama | `4 >= 5` → `false` |

### 3. Operator Logika (menggabungkan kondisi `bool`)

| Operator | Arti | Contoh |
|----------|------|--------|
| `&&` | DAN (semua harus `true`) | `true && false` → `false` |
| `\|\|` | ATAU (salah satu `true`) | `true \|\| false` → `true` |
| `!` | BUKAN (dibalik) | `!true` → `false` |

### 4. Operator Penugasan (mempersingkat)

| Bentuk | Sama dengan |
|--------|-------------|
| `x += 3` | `x = x + 3` |
| `x -= 3` | `x = x - 3` |
| `x *= 3` | `x = x * 3` |
| `x /= 3` | `x = x / 3` |

### 5. Operator String

Teks bisa "dijumlahkan" (disebut *concatenation*):

```go
nama := "Budi" + " " + "Santoso" // "Budi Santoso"
```

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./03-operator
```

## Latihan

1. Hitung keliling dan luas persegi panjang (`p=10`, `l=5`) dan cetak hasilnya.
2. Buat variabel `umur = 20`, lalu cek (cetak `true`/`false`) apakah:
   - `umur` lebih dari 17 **dan** kurang dari 30.
   - `umur` kurang dari 17 **atau** lebih dari 60.
3. Buat variabel `nilai := 10` lalu naikkan nilainya 2 kali dengan `+=`.

> 💡 **Apa bedanya `=` dan `==`?**
> `=` memberi nilai (penugasan), sedangkan `==` membandingkan nilai. Ini sumber
> bug klasik pemula, jadi perhatikan baik-baik.

Selanjutnya: **[Bab 04 — Kontrol Alur](../04-kontrol-alur)**.
