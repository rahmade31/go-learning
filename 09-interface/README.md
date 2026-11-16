# Bab 09 — Interface

Interface adalah salah satu konsep paling kuat di Go. Interface adalah
**kontrak perilaku**: "tipe apa pun yang punya method ini, dianggap bertipe ini."

## Tujuan

Setelah bab ini kamu paham:

- Mendefinisikan interface.
- Mengimplementasikan interface secara **implisit**.
- Polimorfisme: satu variabel bisa menampung banyak tipe.
- Interface kosong `interface{}` (atau `any`).

## Konsep

### 1. Mendefinisikan interface

```go
type Shape interface {
    Luas() float64
    Keliling() float64
}
```

Interface di atas menyatakan: "segala sesuatu yang punya method `Luas()` dan
`Keliling()` bertipe `Shape`."

### 2. Implementasi implisit

Di Go, kita **tidak perlu menulis** `implements Shape`. Cukup tipe itu punya
method yang diminta interface, maka otomatis dianggap mengimplementasikannya.

```go
type Lingkaran struct { JariJari float64 }

func (l Lingkaran) Luas() float64     { return 3.14 * l.JariJari * l.JariJari }
func (l Lingkaran) Keliling() float64 { return 2 * 3.14 * l.JariJari }

type Persegi struct { Sisi float64 }

func (p Persegi) Luas() float64     { return p.Sisi * p.Sisi }
func (p Persegi) Keliling() float64 { return 4 * p.Sisi }
```

### 3. Polimorfisme

Karena `Lingkaran` dan `Persegi` sama-sama `Shape`, kita bisa menulis fungsi
yang menerima `Shape` apa pun:

```go
func cetakInfo(s Shape) {
    fmt.Printf("Luas=%.2f, Keliling=%.2f\n", s.Luas(), s.Keliling())
}

cetakInfo(Lingkaran{JariJari: 7})
cetakInfo(Persegi{Sisi: 5})
```

### 4. Interface kosong (`any`)

Interface tanpa method (`interface{}`, alias `any` di Go 1.18+) bisa menampung
**nilai apa pun**:

```go
var apaSaja any
apaSaja = 42
apaSaja = "teks"
apaSaja = Lingkaran{JariJari: 3}
```

Untuk mengambil nilai aslinya kembali, gunakan **type assertion**:

```go
nilai, ok := apaSaja.(string) // ok = true jika memang string
```

### 5. Type switch

Untuk menangani banyak kemungkinan tipe sekaligus:

```go
switch v := apaSaja.(type) {
case int:
    fmt.Println("integer:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("tipe lain")
}
```

### 6. Mengapa interface penting?

- **Fleksibel** — fungsi bisa bekerja dengan banyak tipe berbeda.
- **Mudah diuji** — bisa mengganti implementasi asli dengan *mock* saat testing.
- **Decoupling** — kode bergantung pada "perilaku", bukan implementasi konkret.

## Kode Lengkap

Lihat [`main.go`](./main.go). Jalankan:

```bash
go run ./09-interface
```

## Latihan

1. Buat interface `Hewan` dengan method `Suara() string`, lalu implementasikan
   pada `Kucing` dan `Anjing`.
2. Tulis fungsi yang menerima `Hewan` dan mencetak suaranya.
3. Coba simpan nilai ke `any`, lalu ambil kembali dengan type assertion.

Selanjutnya: **[Bab 10 — Error Handling](../10-error-handling)**.
