# Golang Modules

## Sebelum Ada Go Modules
- Saat membuat aplikasi, biasanya kita akan menggunakan library atau dependensi dari projek lain.
- Manajemen dependensi ini sangat sulit awalnya, karena kita harus mengunduh dependensi secara manual dan mengatur versinya sendiri.
- Biasanya kita akan meng-copy source code dari dependensi tersebut ke dalam projek kita, atau mengunduhnya secara manual dan menaruhnya di dalam folder `vendor`.
- Akhirnya muncullah Go Modules pada versi Go 1.11 dan 1.12.
- Dengan Go Modules, kita bisa mengelola dependensi dengan lebih mudah dan efisien.

## Membuat Go Modules
- Untuk membuat Go Modules, kita harus berada di dalam folder projek kita dan menjalankan perintah:
```bash
go mod init nama-modul
```
- Perintah ini akan membuat file `go.mod` yang berisi informasi tentang modul kita, seperti nama modul, versi Go yang digunakan, dan dependensi yang dibutuhkan.
- Setelah itu, kita bisa menambahkan dependensi yang kita butuhkan dengan perintah:
```bash
go get nama-dependensi
```
- Untuk merilis modul kita pada Git, kita harus membuat tag pada commit terakhir dengan format `vX.Y.Z`, di mana `X`, `Y`, dan `Z` adalah angka versi yang kita inginkan.


## Menambahkan Dependensi
- Untuk menambahkan dependensi, kita bisa menggunakan perintah `go get` diikuti dengan nama paket yang ingin kita tambahkan.
- Contoh:
```bash
go get github.com/gin-gonic/gin
```

## Upgrade Module
- Untuk mengupgrade modul ke versi terbaru, kita cukup membuat tag baru pada commit terakhir dengan format `vX.Y.Z`.

## Upgrade Dependensi
- Ada 2 cara untuk mengupgrade dependensi:
  - Menggunakan perintah `go get -u` untuk mengupgrade semua dependensi ke versi terbaru.
  - Mengubah isi dari file `go.mod` secara manual dengan menambahkan versi yang diinginkan pada nama paket.

## Major Upgrade
- Major upgrade adalah ketika kita mengupgrade versi dari dependensi yang memiliki perubahan besar, biasanya ditandai dengan perubahan pada angka pertama dari versi (X pada vX.Y.Z).
- Tapi, major upgrade ini bisa menyebabkan perubahan yang tidak kompatibel dengan versi sebelumnya.
- Oleh karena itu, kita harus berhati-hati dalam melakukan major upgrade dan memastikan bahwa kode kita masih berjalan dengan baik setelah upgrade.

# FINISH
