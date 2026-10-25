# Codeforces Go Solutions

Kumpulan solusi problem [Codeforces](https://codeforces.com/) dalam **Go (Golang)**, ditulis dengan pendekatan eksplisit, prosedural, dan hemat dependensi.

Repository ini bukan sekadar tempat menyimpan jawaban ini catatan alur berpikir, cara membaca problem, dan cara menerjemahkannya jadi kode yang lolos limit.

## Filosofi

> Eksplisit, prosedural, dependensi minimal, penanganan batas ketat, alokasi memori terukur.

Semua solusi mengikuti prinsip yang sama:

- **Logika di atas library** utamakan pemahaman algoritma, bukan eksplorasi package.
- **Go dasar dan eksplisit** tidak ada trik sintaks yang menyembunyikan alur.
- **Dependensi eksternal minimal** `fmt`, `bufio`, `os`, `strconv`, `sort` secukupnya.
- **Penamaan semantik lokal** nama variabel berbahasa Indonesia kalau lebih jelas (`biaya`, `jalur`, `arah`, `terbaik`, `baca`, `tulis`).
- **Penanganan batas eksplisit** guard clause di awal, bukan `if` bersarang di tengah.
- **Alokasi terukur** slice dan buffer dialokasikan dengan ukuran yang jelas, bukan tumbuh tanpa kendali.

## Alur Menyelesaikan Problem

1. **Baca problem dua kali.**  
   Sekali untuk memahami cerita, sekali untuk menangkap batasan (`n`, `m`, `k`, limit waktu, limit memori).

2. **Identifikasi bentuk problem.**  
   DP, greedy, geometri, matematika, graf, atau simulasi? Bentuk problem menentukan struktur data.

3. **Cari jebakan sebelum menulis kode.**  
   Nilai `0`, `n = 1`, input kosong, overflow, dan nilai yang butuh perlakuan khusus.

4. **Tulis solusi paling sederhana dulu.**  
   `fmt` murni, logika langsung, tanpa abstraksi. Kalau lolos, selesai. Kalau TLE, baru optimasi.

5. **Optimasi hanya kalau perlu.**  
   Refactor diterima kalau tujuannya performa I/O lebih cepat, struktur data lebih ringan, alokasi lebih kecil.

6. **Uji dengan edge case manual.**  
   Nilai minimum, nilai maksimum, dan kasus khusus yang dicurigai sejak langkah 3.

## Ide yang Saya Pegang

### I/O dulu, bukan algoritma
TLE sering bukan karena algoritma salah, tapi karena I/O lambat. `fmt.Scan` pada ratusan ribu angka bisa jadi penyebab. Solusinya: `bufio.NewReader` + `bufio.NewWriter`, atau baca seluruh input sekaligus lalu parse manual.

### Matematika > brute-force
Kalau ada pola matematis yang bisa memangkas pencarian, pakai itu. Faktorisasi, rasio, locus geometri, eliminasi variabel sering mengubah solusi O(n²) jadi O(1).

### Kasus khusus butuh perlakuan eksplisit
Nilai `0`, input tunggal, dan kondisi degenerat harus ditangani di awal fungsi (guard clause), bukan di tengah loop.

### Flat lebih baik daripada bersarang
Array 2D jadi 1D kalau bisa. Rekursi jadi iterasi kalau bisa. Struktur datar lebih cepat diakses dan lebih mudah dilacak.

### Sedikit fungsi, sedikit abstraksi
Satu fungsi besar yang jelas lebih baik daripada lima helper yang saling memanggil.

## Gaya Kode Go

**Deklarasi**

```go
var biaya [][]int
biaya = make([][]int, n)
