# Bab 00 — Persiapan (Instalasi & Alat Kerja)

Sebelum menulis kode, kita siapkan dulu "dapur" kerjanya. Bab ini tidak berisi
kode Go, hanya panduan instalasi.

## 1. Instal Go

Unduh installer resmi dari situs resmi Go:

👉 https://go.dev/dl/

Pilih sesuai sistem operasimu (Windows, macOS, Linux).

### Cek instalasi

Setelah selesai instal, buka **terminal** (di Windows bisa pakai *PowerShell*
atau *Git Bash*, di macOS/Linux pakai *Terminal*) lalu ketik:

```bash
go version
```

Jika muncul tulisan seperti `go version go1.22.x ...`, berarti Go sudah siap.

## 2. Pilih Editor / IDE

Kamu bisa pakai editor apa pun, tapi yang paling populer untuk Go adalah:

| Editor | Alasan |
|--------|--------|
| **VS Code** | Gratis, ringan, dukungan Go sangat baik lewat ekstensi |
| **GoLand** (JetBrains) | Fitur lengkap, berbayar |
| **Zed** | Editor modern yang kamu pakai sekarang 😄 |

### Rekomendasi: VS Code + Ekstensi Go

1. Instal [VS Code](https://code.visualstudio.com/).
2. Buka menu *Extensions* (ikon kotak-kotak di kiri).
3. Cari **"Go"** (penerbit: Go Team at Google) lalu klik *Install*.
4. Saat VS Code menawarkan instal alat bantu (gopls, dlv, dll.), klik *Install All*.

Ekstensi ini memberi *auto-complete*, *format otomatis*, dan deteksi error.

## 3. Struktur Folder

Repo ini sudah diatur agar setiap bab punya folder sendiri. Satu bab = satu
paket kecil yang bisa dijalankan sendiri.

```
learning-go/
├── go.mod                ← deklarasi modul & versi Go
├── README.md             ← kamu sedang membacanya
├── 00-persiapan/         ← bab ini
├── 01-hello-world/       ← bab selanjutnya
│   ├── README.md
│   └── main.go
└── ...                   ← dan seterusnya
```

File `go.mod` di atas menyatakan bahwa semua folder adalah bagian dari **satu
modul** bernama `learning-go`.

## 4. Perintah Dasar yang Akan Sering Dipakai

| Perintah | Fungsi |
|----------|--------|
| `go run ./nama-folder` | Menjalankan program tanpa membuat file biner |
| `go build ./nama-folder` | Mengompilasi menjadi file biner (`.exe` di Windows) |
| `go vet ./...` | Memeriksa kesalahan umum dalam kode |
| `go fmt ./...` | Merapikan format kode otomatis |
| `go test ./...` | Menjalankan unit test |

> Istilah **kompilasi** artinya menerjemahkan kode sumber menjadi file yang bisa
> langsung dieksekusi komputer. `go run` sebenarnya juga mengompilasi, tapi
> hasilnya langsung dijalankan tanpa disimpan.

## 5. Cheat Sheet Singkat

- Go adalah bahasa **statically typed** (tipe data ditentukan saat kompilasi).
- Program Go selalu mulai dari fungsi bernama `main` di dalam paket `main`.
- Go memakai **kurung kurawal `{}`** untuk blok, dan tidak memakai titik koma.

---

Selanjutnya, mari tulis program pertamamu di **[Bab 01 — Hello World](../01-hello-world)**. 🎉
