# Golang Dasar

## Pengenalan Golang
- Go adalah bahasa pemrograman open source yang dikembangkan oleh Google
- Dirilis pertama kali pada tahun 2009 oleh Robert Griesemer, Rob Pike, dan Ken Thompson
- Didesain untuk memudahkan pengembangan aplikasi sederhana, handal dan efisien
- Memiliki fitur garbage collection dan concurrent programming
- Sintaks yang sederhana dan mudah dipelajari
- Mendukung pemrograman berorientasi objek
- Kompilasi yang cepat dan menghasilkan file binary
- Memiliki standard library yang lengkap
- Cocok untuk pengembangan aplikasi backend, CLI, dan cloud computing

## Cara Kerja Golang
- Go adalah bahasa yang dikompilasi (compiled language)
- Source code Go akan dikompilasi menjadi binary file
- Binary file bisa langsung dijalankan tanpa perlu interpreter
- Proses development:
  1. Tulis kode program (.go)
  2. Kompilasi dengan command `go build`
  3. Hasil kompilasi berupa executable binary
  4. Jalankan binary file
- Contoh sederhana:
  ```go
  package main
  
  import "fmt"
  
  func main() {
      fmt.Println("Hello World")
  }
  ```
  Simpan dalam file hello.go, lalu:
  ```bash
  go build hello.go
  ./hello
  ```
- Jika ingin menjalankan langsung tanpa kompilasi manual, gunakan:
  ```bash
  go run hello.go
  ```

## Hello World
Berikut adalah contoh program Hello World dalam Go:

1. Buat file `hello-world.go`
2. Tulis kode berikut:
```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}
```

3. Jalankan perintah berikut di terminal:
```bash
go run hello-world.go
```

Penjelasan:
- `package main`: Menandakan bahwa ini adalah package utama yang akan dieksekusi
- `import "fmt"`: Mengimpor package fmt untuk format output
- `func main()`: Fungsi utama yang akan dieksekusi saat program dijalankan
- `fmt.Println("Hello, World!")`: Mencetak string "Hello, World!" ke konsol
- `go run hello-world.go`: Perintah untuk menjalankan program Go secara langsung tanpa perlu kompilasi manual
- Output yang diharapkan adalah:
```
Hello, World!
```

[hello-world.go](hello-world.go)

## Multiple Main
- Go tidak mendukung multiple main package dalam satu direktori
- Setiap file Go harus memiliki satu package main jika ingin dieksekusi
- Jika kita mencoba build multiple file dengan package main, akan terjadi error berikut:
```
# command-line-arguments
./file1.go:6:6: main redeclared in this block
    previous declaration at ./file2.go:6:6
```

[multiple-main.go](multiple-main.go)

## Number
- Tipe data numerik di Go terdiri dari:
  - Integer: `int`, `int8`, `int16`, `int32`, `int64`
  - Unsigned Integer: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
  - Floating Point: `float32`, `float64`
  - Complex: `complex64`, `complex128`
- Detail integer:
    - `int8`: 8-bit signed integer, rentang -128 hingga 127
    - `int16`: 16-bit signed integer, rentang -32,768 hingga 32,767
    - `int32`: 32-bit signed integer, rentang -2,147,483,648 hingga 2,147,483,647
    - `int64`: 64-bit signed integer, rentang -9,223,372,036,854,775,808 hingga 9,223,372,036,854,775,807
- Detail unsigned integer:
    - `uint8`: 8-bit unsigned integer, rentang 0 hingga 255
    - `uint16`: 16-bit unsigned integer, rentang 0 hingga 65,535
    - `uint32`: 32-bit unsigned integer, rentang 0 hingga 4,294,967,295
    - `uint64`: 64-bit unsigned integer, rentang 0 hingga 18,446,744,073,709,551,615
- Detail floating point:
    - `float32`: 32-bit floating point, presisi sekitar 7 digit desimal
    - `float64`: 64-bit floating point, presisi sekitar 15 digit desimal
- Alias tipe data:
  - `int`: Tipe integer dengan ukuran platform-dependent (32-bit atau 64-bit)
  - `uint`: Tipe unsigned integer dengan ukuran platform-dependent
  - `byte`: Alias untuk `uint8`, sering digunakan untuk merepresentasikan byte data
  - `rune`: Alias untuk `int32`, digunakan untuk merepresentasikan karakter Unicode

[number.go](number.go)

## Boolean
- Tipe data boolean di Go hanya memiliki dua nilai: `true` dan `false`
- Digunakan untuk menyimpan kondisi logika

[boolean.go](boolean.go)

## String
- Tipe data string di Go adalah urutan karakter yang tidak dapat diubah (immutable)
- Dideklarasikan dengan menggunakan tanda kutip ganda (`"`)
- String dapat berisi karakter ASCII atau Unicode
- Fungsi-fungsi umum untuk string:
  - `len(s)`: Mengembalikan panjang string `s`
  - `s[i]`: Mengakses karakter pada indeks `i` (tipe byte)
  - `s[start:end]`: Mengiris string dari indeks `start` hingga `end-1`
  - `strings.ToUpper(s)`: Mengubah string menjadi huruf kapital
  - `strings.ToLower(s)`: Mengubah string menjadi huruf kecil
  - `strings.Contains(s, substr)`: Mengecek apakah string `s` mengandung substring `substr`
  - `strings.Split(s, sep)`: Memecah string `s` berdasarkan separator `sep`

[string.go](string.go)

## Variable
- Variabel di Go dideklarasikan dengan kata kunci `var`
- Tipe data variabel dapat ditentukan secara eksplisit atau dibiarkan otomatis
- Contoh deklarasi variabel:
  ```go
  var x int
  x = 10 // deklarasi dan inisialisasi
  var y int = 20 // tipe data otomatis
  z := 30    // shorthand untuk deklarasi dan inisialisasi
  var (
    a int = 1
    b string = "Hello"
  ) // deklarasi multi variabel
    ```
[variable.go](variable.go)
  
## Constant
- Konstanta di Go dideklarasikan dengan kata kunci `const`
- Nilai konstanta tidak dapat diubah setelah dideklarasikan
- Cara deklarasi konstanta sama seperti variabel, tetapi menggunakan `const`

## Konversi Tipe Data
- Go mendukung konversi tipe data eksplisit
- Contoh konversi:
  ```go
  var x int = 10
  var y float64 = float64(x) // konversi dari int ke float64
  var z string = strconv.Itoa(x) // konversi dari int ke string
  ```
- Untuk tipe data numeric terdapat dua jenis konversi:
    - Widening conversion: dari tipe yang lebih kecil ke tipe yang lebih besar (misal `int8` ke `int`)
    - Narrowing conversion: dari tipe yang lebih besar ke tipe yang lebih kecil (misal `int` ke `int8`), harus dilakukan secara eksplisit
- Contoh konversi numeric:
```go
package main

import "fmt"

func main() {
    var a int8 = 100
    var b int = int(a) // widening conversion
    fmt.Println(b) // Output: 100

    var c int = 200
    var d int8 = int8(c) // narrowing conversion, harus hati-hati
    fmt.Println(d) // Output: 200, tapi jika c > 127 akan overflow atau kembali ke nilai minimum
}
```
- Untuk konversi string ke numeric, gunakan package `strconv`:
```go
package main
import (
    "fmt"
    "strconv"
)
func main() {
    var str string = "123"
    num, err := strconv.Atoi(str) // konversi string ke int
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted number:", num) // Output: Converted number: 123
    }
}
```
- Untuk konversi numeric ke string, gunakan `strconv.Itoa()` atau `fmt.Sprintf()`:
```go
package main
import (
    "fmt"
    "strconv"
)
func main() {
    var num int = 456
    str := strconv.Itoa(num) // konversi int ke string
    fmt.Println("Converted string:", str) // Output: Converted string: 456

    str2 := fmt.Sprintf("%d", num) // alternatif konversi
    fmt.Println("Converted string with Sprintf:", str2) // Output: Converted string with Sprintf: 456
}
```
[konversi-tipe-data.go](konversi-tipe-data.go)

# Type Declaration
- Merupakan cara untuk membuat tipe data baru berdasarkan tipe data yang sudah ada
- Digunakan untuk meningkatkan keterbacaan kode dan mengelompokkan data yang terkait
- Contoh deklarasi tipe data baru:
```go
package main
import "fmt"
type Age int // mendeklarasikan tipe data baru Age berdasarkan int
func main() {
    var myAge Age = 25 // menggunakan tipe data Age
    fmt.Println("My age is:", myAge)
}
```
[type-declaration.go](type-declaration.go)
  
# Operasi Aritmatika
- Go mendukung operasi aritmatika dasar seperti penjumlahan, pengurangan, perkalian, pembagian, dan modulus
- Macam operasi aritmatika:
  - Penjumlahan: `+`
  - Pengurangan: `-`
  - Perkalian: `*`
  - Pembagian: `/`
  - Modulus: `%`
  - Pangkat: `math.Pow(x, y)` untuk menghitung x pangkat y

# Operasi Logika
- Go mendukung operasi logika seperti AND, OR, dan NOT
- Macam operasi logika:
  - AND: `&&`
  - OR: `||`
  - NOT: `!`

# Operasi Perbandingan
- Go mendukung operasi perbandingan untuk membandingkan nilai
- Macam operasi perbandingan:
  - Sama dengan: `==`
  - Tidak sama dengan: `!=`
  - Lebih besar dari: `>`
  - Lebih kecil dari: `<`
  - Lebih besar atau sama dengan: `>=`
  - Lebih kecil atau sama dengan: `<=`

# Operasi Bitwise
- Go mendukung operasi bitwise untuk manipulasi bit
- Macam operasi bitwise:
  - AND: `&`
  - OR: `|`
  - XOR: `^`
  - NOT: `^` (unary)
  - Shift kiri: `<<`
  - Shift kanan: `>>`
  - Shift kanan aritmatika: `>>` (untuk signed integer)
  - Shift kanan logika: `>>` (untuk unsigned integer)
- Contoh operasi bitwise:
```go
package main
import "fmt"
func main() {
    var a int = 5  // 0101 dalam biner
    var b int = 3  // 0011 dalam biner

    fmt.Println("a & b =", a & b) // AND: 0001 (1)
    fmt.Println("a | b =", a | b) // OR: 0111 (7)
    fmt.Println("a ^ b =", a ^ b) // XOR: 0110 (6)
    fmt.Println("a << 1 =", a << 1) // Shift kiri: 1010 (10)
    fmt.Println("a >> 1 =", a >> 1) // Shift kanan: 0010 (2)
}
```
[operasi-bitwise.go](operasi-bitwise.go)

# Array
- Array adalah kumpulan elemen dengan tipe data yang sama
- Deklarasi array:
  - Dengan ukuran tetap: `var arr [5]int` (array dengan 5 elemen bertipe int)
  - Inisialisasi langsung: `arr := [3]string{"a", "b", "c"}`
  - Array tanpa ukuran: `arr := [...]int{1, 2, 3}` (ukuran otomatis berdasarkan jumlah elemen)
  - Array multidimensi: `var matrix [2][3]int` (array 2x3)
- Fungsi untuk mengakses elemen:
  - Menggunakan indeks: `arr[0]` untuk elemen pertama
  - Menggunakan loop: `for i, v := range arr {}` untuk iterasi
  - Menggunakan fungsi `len(arr)` untuk mendapatkan panjang array
- Contoh penggunaan array:
```go
package main
import "fmt"
func main() {
    // Deklarasi array
    var numbers [5]int
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    numbers[3] = 40
    numbers[4] = 50

    // Inisialisasi array langsung
    fruits := [3]string{"Apple", "Banana", "Cherry"}

    // Array multidimensi
    matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}

    // Mengakses elemen array
    fmt.Println("First number:", numbers[0])
    fmt.Println("Second fruit:", fruits[1])
    fmt.Println("Matrix element:", matrix[1][2]) // Output: 6

    // Menggunakan loop untuk iterasi
    for i, v := range numbers {
        fmt.Printf("Element at index %d: %d\n", i, v)
    }
}
```
[array.go](array.go)

# Slice
- Slice adalah tipe data dinamis yang merupakan bagian dari array
- Slice tidak memiliki ukuran tetap, bisa bertambah atau berkurang
- Slice adalah referensi ke array, sehingga perubahan pada slice akan mempengaruhi array asli
- Slice memiliki tiga atribut utama:
  - Pointer: Menunjuk ke elemen pertama dari array yang mendasarinya
  - Length: Jumlah elemen yang ada di slice
  - Capacity: Jumlah elemen yang bisa ditampung oleh slice sebelum perlu alokasi ulang (dihitung dari pointer hingga akhir array)
- Deklarasi slice:
  - Dengan inisialisasi langsung: `var s []int` (slice kosong)
  - Dengan inisialisasi dengan elemen: `s := []string{"a", "b", "c"}`
  - Menggunakan fungsi `make`: `s := make([]int, 5)` (slice dengan panjang 5, elemen default 0)
  - Slice dari array: `s := arr[1:4]` (slice dari indeks 1 hingga 3)
  - Slice multidimensi: `s := [][]int{{1, 2}, {3, 4}}` (slice dari slice)
- Fungsi pada slice:
  - `len(s)`: Mengembalikan panjang slice
  - `cap(s)`: Mengembalikan kapasitas slice (jumlah elemen yang bisa ditampung)
  - `append(s, elem)`: Menambahkan elemen ke akhir slice
  - `copy(dest, src)`: Menyalin elemen dari slice sumber ke slice tujuan
  - `delete(s, index)`: Menghapus elemen pada indeks tertentu (tidak ada fungsi delete bawaan, harus menggunakan slicing)
- Contoh penggunaan slice:
```go
package main
import "fmt"
func main() {
    // Deklarasi slice
    var numbers []int // slice kosong
    numbers = append(numbers, 10, 20, 30) // menambahkan elemen
    fmt.Println("Numbers:", numbers)
    // Inisialisasi slice dengan elemen
    fruits := []string{"Apple", "Banana", "Cherry"}
    fmt.Println("Fruits:", fruits)
    // Slice dari array
    arr := [5]int{1, 2, 3, 4, 5}
    s := arr[1:4] // slice dari indeks 1 hingga 3
    fmt.Println("Slice from array:", s)
    // Menggunakan fungsi make
    s2 := make([]int, 5) // slice dengan panjang 5
    fmt.Println("Slice with make:", s2)
    // Menambahkan elemen ke slice
    s2 = append(s2, 6, 7)
    fmt.Println("Slice after append:", s2)
    // Menggunakan copy
    s3 := make([]int, len(s2))
    copy(s3, s2) // menyalin elemen dari s2 ke s3
    fmt.Println("Copied slice:", s3)
    // Menghapus elemen (tidak ada fungsi delete, menggunakan slicing)
    s2 = append(s2[:1], s2[2:]...) // menghapus elemen kedua
    fmt.Println("Slice after delete:", s2)
    // Menggunakan slice multidimensi
    matrix := [][]int{{1, 2}, {3, 4}}
      fmt.Println("Matrix slice:", matrix)
}
```
- Ketika menggunakan slice, perlu diperhatikan bahwa:
  - Slice adalah referensi ke array, jadi perubahan pada slice akan mempengaruhi array asli
  - Kapasitas slice bisa bertambah secara otomatis saat menambahkan elemen, tetapi jika kapasitas terlampaui, Go akan membuat array baru dengan kapasitas yang lebih besar
  - Slice dapat digunakan untuk mengelola data dinamis dengan lebih efisien dibandingkan array tetap

[slice.go](slice.go)

# Map
- Map adalah tipe data koleksi yang menyimpan pasangan kunci-nilai (key-value)
- Kunci harus unik, sedangkan nilai bisa duplikat
- Map dideklarasikan dengan menggunakan `make` atau literal map
- Deklarasi map:
  - Dengan `make`: `m := make(map[string]int)` (map dengan kunci string dan nilai int)
  - Dengan literal: `m := map[string]int{"a": 1, "b": 2}` (map dengan inisialisasi langsung)
  - Map kosong: `var m map[string]int` (map kosong, harus diinisialisasi sebelum digunakan)
  - Map dengan tipe data kompleks:
    - Kunci bisa berupa tipe data apa pun yang dapat dibandingkan (comparable), seperti string, int, atau struct
    - Nilai bisa berupa tipe data apa pun, termasuk slice, struct, atau map lainnya
- Fungsi pada map:
  - `len(m)`: Mengembalikan jumlah pasangan kunci-nilai dalam map
  - `m[key]`: Mengakses nilai berdasarkan kunci, jika kunci tidak ada, akan mengembalikan nilai default
  - `m[key] = value`: Menambahkan atau memperbarui pasangan kunci-nilai
  - `delete(m, key)`: Menghapus pasangan kunci-nilai dari map
  - `make(map[K]V)`: Membuat map dengan tipe kunci K dan nilai V
- Contoh penggunaan map:
```go
package main
import "fmt"
func main() {
    // Deklarasi map dengan make
    m := make(map[string]int)
    m["a"] = 1 // menambahkan pasangan kunci-nilai
    m["b"] = 2
    fmt.Println("Map:", m)

    // Inisialisasi map dengan literal
    m2 := map[string]int{"x": 10, "y": 20}
    fmt.Println("Map with literal:", m2)

    // Mengakses nilai berdasarkan kunci
    value := m["a"]
    fmt.Println("Value for key 'a':", value)

    // Menggunakan len untuk mendapatkan jumlah elemen
    fmt.Println("Length of map:", len(m))

    // Menghapus pasangan kunci-nilai
    delete(m, "b")
    fmt.Println("Map after deletion:", m)

    // Map dengan tipe data kompleks
    complexMap := make(map[string][]int)
    complexMap["numbers"] = []int{1, 2, 3}
    fmt.Println("Complex map:", complexMap)
}
```

# If Expression
- Pernyataan `if` digunakan untuk mengontrol alur program berdasarkan kondisi
- Sintaks dasar:
```go
if kondisi {
    // blok kode jika kondisi benar
} else if kondisi_lain {
    // blok kode jika kondisi_lain benar
} else {
    // blok kode jika semua kondisi salah
}
```
- If dengan inisialisasi:
```go
if variabel := ekspresi; kondisi {
    // blok kode jika kondisi benar
}
```

# Switch Expression
- Pernyataan `switch` digunakan untuk memilih salah satu dari beberapa blok kode berdasarkan nilai ekspresi
- Sintaks dasar:
```go
switch ekspresi {
case nilai1:
    // blok kode jika ekspresi == nilai1
case nilai2:
    // blok kode jika ekspresi == nilai2
default:
    // blok kode jika tidak ada case yang cocok
}
```
- Switch tanpa ekspresi:
```go
switch {
case kondisi1:
    // blok kode jika kondisi1 benar
case kondisi2:
    // blok kode jika kondisi2 benar
default:
    // blok kode jika semua kondisi salah
}
```
- Switch dengan fallthrough:
```go
switch nilai {
case 1:
    fmt.Println("Satu")
    fallthrough // melanjutkan ke case berikutnya
case 2:
    fmt.Println("Dua")
default:
    fmt.Println("Lainnya")
}
```

# For Loop
- Pernyataan `for` digunakan untuk melakukan iterasi atau pengulangan
- Sintaks dasar:
```go
for inisialisasi; kondisi; increment {
    // blok kode yang akan diulang
}
```
- For dengan range:
```go
for indeks, nilai := range koleksi {
    // blok kode yang akan diulang untuk setiap elemen
}
```
- For tanpa kondisi (infinite loop):
```go
for {
    // blok kode yang akan diulang selamanya
    if kondisi_berhenti {
        break // keluar dari loop
    }
}
```
# Break and Continue
- `break` digunakan untuk keluar dari loop atau switch
- `continue` digunakan untuk melewati iterasi saat ini dan melanjutkan ke iterasi berikutnya
- Contoh penggunaan:
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // keluar dari loop jika i == 5
    }
    fmt.Println(i)
}
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // melewati iterasi jika i genap
    }
    fmt.Println(i) // hanya mencetak angka ganjil
}
```

# Function
- Fungsi adalah blok kode yang dapat dipanggil untuk melakukan tugas tertentu
- Deklarasi fungsi:
```go
func namaFungsi(parameter1 tipe1, parameter2 tipe2) returnType {
    // blok kode fungsi
    return nilai // jika ada return type
}
func main()  {
    namaFungsi(arg1, arg2) // memanggil fungsi
}
```

# Function Parameter
- Parameter fungsi adalah variabel yang digunakan untuk menerima input saat fungsi dipanggil
- Deklarasi parameter:
```go
func namaFungsi(param1 tipe1, param2 tipe2) {
    // blok kode
}
```
- Parameter dapat memiliki tipe data yang berbeda
- Parameter dapat dideklarasikan sebagai variadic (menerima jumlah argumen yang tidak terbatas):
```go
func variadicFunction(params ...int) {
    for _, v := range params {
        fmt.Println(v)
    }
}
```
- Memanggil fungsi dengan parameter:
```go
namaFungsi(arg1, arg2) // arg1 dan arg2 sesuai dengan tipe parameter
```

# Function Return
- Fungsi dapat mengembalikan nilai menggunakan kata kunci `return`
- Ketika kita menggunakan return, fungsi harus dideklarasikan dengan return type
- Deklarasi fungsi dengan return type:
```go
func namaFungsi() returnType {
    // blok kode
    return nilai // mengembalikan nilai sesuai dengan return type
}
```
- Fungsi dapat mengembalikan beberapa nilai:
```go
func multipleReturn() (int, string) {
    return 42, "Hello"
}
```
- Return value dapat dihiraukan dengan menggunakan underscore `_`:
```go
func example() (int, string) {
    return 1, "example"
}
func main() {
    _, str := example() // hanya mengambil string, integer diabaikan
    fmt.Println(str) // Output: example
}
```

# Named Return Value
- Named return value adalah fitur di Go yang memungkinkan kita mendeklarasikan nama variabel untuk nilai yang dikembalikan oleh fungsi
- Dengan named return value, kita dapat mengembalikan nilai tanpa perlu menyebutkan nama variabel pada pernyataan `return`
- Contoh penggunaan named return value:
```go
package main
import "fmt"
func divide(a, b int) (result int, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return // mengembalikan named return value
    }
    result = a / b
    return // mengembalikan named return value
}
func main() {
    res, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", res) // Output: Result: 5
    }

    res, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: division by zero
    } else {
        fmt.Println("Result:", res)
    }
}
```
[named-return-value.go](named-return-value.go)

# Variadic Function
- Variadic function adalah fungsi yang dapat menerima jumlah argumen yang tidak terbatas
- Dideklarasikan dengan menggunakan `...` sebelum tipe parameter
- Contoh deklarasi variadic function:
```go
func variadicFunction(params ...int) {
    for _, v := range params {
        fmt.Println(v)
    }
}
```
- Variadic function bisa diisi dengan slice.
- Contoh penggunaan variadic function:
```go
package main
import "fmt"
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println("Sum:", sum(1, 2, 3)) // Output: Sum: 6
    fmt.Println("Sum:", sum(4, 5, 6, 7, 8)) // Output: Sum: 30

    // Menggunakan slice sebagai argumen variadic function
    nums := []int{10, 20, 30}
    fmt.Println("Sum with slice:", sum(nums...)) // Output: Sum with slice: 60
}
```
[variadic-function.go](variadic-function.go)

# Function as Value
- Function as value adalah konsep di Go di mana fungsi dapat diperlakukan sebagai nilai
- Fungsi dapat disimpan dalam variabel, diteruskan sebagai argumen, atau dikembalikan dari fungsi lain
- Contoh penggunaan function value:
```go
package main
import "fmt"
func add(a, b int) int {
    return a + b
}
func main() {
    // Menyimpan fungsi dalam variabel
    sum := add
    result := sum(3, 4) // memanggil fungsi melalui variabel
    fmt.Println("Result:", result) // Output: Result: 7

    // Menggunakan fungsi sebagai argumen
    applyFunction := func(f func(int, int) int, x, y int) int {
        return f(x, y)
    }
    fmt.Println("Apply Function Result:", applyFunction(add, 5, 6)) // Output: Apply Function Result: 11
}
```
- Konsepnya mirip dengan higher-order functions di bahasa pemrograman lain.
- Jika fungsi ingin dijadikan sebagai nilai, cukup mendeklarasikan fungsi tersebut tanpa tanda kurung `()` di akhir karena jika menggunakan tanda kurung, maka fungsi tersebut akan dieksekusi dan bukan disimpan sebagai nilai.

[function-as-value.go](function-as-value.go)

# Function as Parameter
- Fungsi dapat diteruskan sebagai parameter ke fungsi lain
- Contoh penggunaan fungsi sebagai parameter:
```go
package main
import "fmt"
func applyFunction(f func(int, int) int, x, y int) int {
    return f(x, y)
}
func add(a, b int) int {
    return a + b
}
func main() {
    result := applyFunction(add, 3, 5) // meneruskan fungsi add sebagai parameter
    fmt.Println("Result:", result) // Output: Result: 8

    // Meneruskan fungsi anonim sebagai parameter
    multiply := func(a, b int) int {
        return a * b
    }
    result2 := applyFunction(multiply, 4, 6)
    fmt.Println("Multiply Result:", result2) // Output: Multiply Result: 24
}
```
- Apakah mirip dengan callback function di bahasa pemrograman lain? Ya, konsep ini mirip dengan callback function di bahasa pemrograman lain, di mana fungsi dapat diteruskan sebagai argumen untuk dipanggil kembali di dalam fungsi lain.

[function-as-parameter.go](function-as-parameter.go)

# Function Type Declaration
- Kita dapat mendeklarasikan tipe data baru dengan tipe berupa fungsi.
- Fungsinya adalah untuk membuat kode lebih terstruktur dan mudah dibaca karena kadang ada fungsi yang memiliki struktur yang banyak dan ingin kita sederhanakan dengan mendeklarasikan tipe fungsi.
- Contoh deklarasi tipe fungsi:
```go
package main
import "fmt"
type Operation func(int, int) int // mendeklarasikan tipe fungsi Operation
func add(a, b int) int {
    return a + b
}
func main() {
    var op Operation // mendeklarasikan variabel dengan tipe Operation
    op = add // mengassign fungsi add ke variabel op
    result := op(3, 4) // memanggil fungsi melalui variabel
    fmt.Println("Result:", result) // Output: Result: 7

    // Menggunakan tipe fungsi sebagai parameter
    applyOperation := func(f Operation, x, y int) int {
        return f(x, y)
    }
    fmt.Println("Apply Operation Result:", applyOperation(op, 5, 6)) // Output: Apply Operation Result: 11
}
```
[function-type-declaration.go](function-type-declaration.go)

# Anonymous Function
- Anonymous function adalah fungsi yang tidak memiliki nama
- Dapat digunakan untuk membuat fungsi sekali pakai atau fungsi yang hanya digunakan di dalam konteks tertentu
- Contoh penggunaan anonymous function:
```go
package main
import "fmt"
func main() {
    // Mendeklarasikan anonymous function
    greet := func(name string) {
        fmt.Println("Hello,", name)
    }
    greet("Alice") // Output: Hello, Alice

    // Menggunakan anonymous function sebagai parameter
    apply := func(f func(string), name string) {
        f(name)
    }
    apply(greet, "Bob") // Output: Hello, Bob

    // Anonymous function dengan return value
    square := func(x int) int {
        return x * x
    }
    fmt.Println("Square of 5:", square(5)) // Output: Square of 5: 25
}
```
- Anonymous function juga dapat digunakan untuk membuat closure, yaitu fungsi yang mengakses variabel di luar scope-nya:
```go
package main
import "fmt"
func main() {
    counter := 0
    increment := func() int {
        counter++
        return counter
    }
    fmt.Println("Increment:", increment()) // Output: Increment: 1
    fmt.Println("Increment:", increment()) // Output: Increment: 2
}
```

# Recursion
- Recursion adalah teknik di mana fungsi memanggil dirinya sendiri untuk menyelesaikan tugas
- Digunakan untuk menyelesaikan masalah yang dapat dipecah menjadi sub-masalah yang lebih kecil
- Contoh penggunaan recursion:
```go
package main
import "fmt"
func factorial(n int) int {
    if n == 0 {
        return 1 // basis kasus
    }
    return n * factorial(n-1) // memanggil dirinya sendiri
}
func main() {
    fmt.Println("Factorial of 5:", factorial(5)) // Output: Factorial of 5: 120
    fmt.Println("Factorial of 0:", factorial(0)) // Output: Factorial of 0: 1
}
```
- Penting untuk memiliki basis kasus (base case) untuk menghentikan rekursi, jika tidak akan terjadi infinite recursion.
- Recursion dapat menyebabkan stack overflow jika terlalu dalam, jadi perlu hati-hati dalam penggunaannya.

# Closure
- Closure adalah fungsi yang mengakses variabel di luar scope-nya, bahkan setelah scope tersebut selesai dieksekusi
- Closure memungkinkan kita untuk menyimpan state di dalam fungsi
- Contoh penggunaan closure:
```go
package main
import "fmt"
func main() {
    counter := 0
    increment := func() int {
        counter++ // mengakses variabel di luar scope
        return counter
    }
    fmt.Println("Increment:", increment()) // Output: Increment: 1
    fmt.Println("Increment:", increment()) // Output: Increment: 2
    fmt.Println("Increment:", increment()) // Output: Increment: 3

    // Closure dengan parameter
    multiplier := func(factor int) func(int) int {
        return func(x int) int {
            return x * factor
        }
    }
    double := multiplier(2)
    fmt.Println("Double of 5:", double(5)) // Output: Double of 5: 10
}
```

# Defer, Panic, dan Recover
- `defer` digunakan untuk menunda eksekusi fungsi hingga fungsi yang memanggilnya selesai. Meskipun fungsi `defer` dideklarasikan di awal, eksekusinya akan dilakukan setelah fungsi utama selesai.
- `panic` digunakan untuk menghentikan eksekusi program dan menghasilkan error
- `recover` digunakan untuk menangkap panic dan melanjutkan eksekusi program. Recover hanya dapat digunakan di dalam fungsi yang dipanggil oleh `defer`.
- Contoh penggunaan defer, panic, dan recover:
```go
package main
import (
    "fmt"
)
func main() {
    defer fmt.Println("This will be printed at the end") // akan dieksekusi setelah main selesai

    fmt.Println("Starting main function")
    
    // Contoh panic
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    panic("Something went wrong!") // akan menghentikan eksekusi program
    fmt.Println("This line will not be executed") // tidak akan dieksekusi
}
```
- Output dari contoh di atas akan menunjukkan bahwa pesan dari `defer` dieksekusi setelah panic, dan program tidak crash karena panic ditangkap oleh `recover`.
- Mirip dengan try-catch di bahasa pemrograman lain, `defer` digunakan untuk membersihkan sumber daya atau mengeksekusi kode tertentu sebelum keluar dari fungsi, sedangkan `panic` dan `recover` digunakan untuk menangani error secara terstruktur.

[defer-panic-recover.go](defer-panic-recover.go)

# Komentar
- Komentar di Go digunakan untuk memberikan penjelasan atau catatan dalam kode
- Ada dua jenis komentar:
  - Komentar satu baris: menggunakan `//` untuk komentar satu baris
  - Komentar multi-baris: menggunakan `/* ... */` untuk komentar yang mencakup beberapa baris
  - Komentar dokumentasi: menggunakan `//` sebelum deklarasi fungsi, variabel, atau tipe untuk memberikan dokumentasi yang akan digunakan oleh `godoc`

# Struct
- Struct adalah tipe data yang digunakan untuk mengelompokkan beberapa nilai dengan tipe data yang berbeda menjadi satu entitas
- Dideklarasikan dengan menggunakan kata kunci `type` dan `struct`
- Contoh deklarasi struct:
```go
package main
import "fmt"
type Person struct {
    Name string
    Age  int
    Address string
}
func main() {
    // Inisialisasi struct
    p := Person{Name: "Alice", Age: 30, Address: "123 Main St"}
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
    fmt.Println("Address:", p.Address)

    // Mengakses field struct
    p.Age = 31 // mengubah nilai field Age
    fmt.Println("Updated Age:", p.Age)
}
```
[struct.go](struct.go)

# Struct Method
- Struct dapat memiliki method yang terkait dengan tipe struct tersebut
- Method dideklarasikan dengan menggunakan kata kunci `func` dan memiliki tipe receiver yang merupakan tipe struct
- Contoh penggunaan struct method:
```go
package main
import "fmt"
type Rectangle struct {
    Width  float64
    Height float64
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height // method untuk menghitung luas
}
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height) // method untuk menghitung keliling
}
func main() {
    rect := Rectangle{Width: 5, Height: 10}
    fmt.Println("Area:", rect.Area())         // Output: Area: 50
    fmt.Println("Perimeter:", rect.Perimeter()) // Output: Perimeter: 30
}
```
- Sebelum nama method, terdapat receiver yang merupakan tipe struct. Receiver ini memungkinkan method untuk mengakses field dari struct tersebut. 
- Anggapannya kita menandai bahwa method ini milik tipe struct tersebut.

[struct-method.go](struct-method.go)

# Interface
- Interface adalah tipe data yang mendefinisikan sekumpulan method tanpa implementasi/body
- Digunakan untuk mendefinisikan perilaku yang harus dimiliki oleh tipe data lain
- Interface dideklarasikan dengan menggunakan kata kunci `type` dan `interface`
- Contoh deklarasi interface:
```go
package main
import "fmt"
type Shape interface {
    Area() float64 // mendefinisikan method Area
    Perimeter() float64 // mendefinisikan method Perimeter
}
type Rectangle struct {
    Width  float64
    Height float64
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height // implementasi method Area
}
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height) // implementasi method Perimeter
}
func printShapeInfo(s Shape) {
    fmt.Println("Area:", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}
func main() {
    rect := Rectangle{Width: 5, Height: 10}
    printShapeInfo(rect) // memanggil fungsi dengan tipe Shape
}
```
- Interface pada dasarnya adalah kontrak yang harus diikuti oleh tipe data yang mengimplementasikannya.
- Jika tipe data tersebut memiliki semua method yang didefinisikan dalam interface, maka tipe data tersebut dianggap mengimplementasikan interface tersebut.

[interface.go](interface.go)

# Interface Kosong
- Golang bukan merupakan object-oriented programming (OOP) seperti Java atau C#, tetapi memiliki konsep interface yang mirip dengan OOP.
- Biasanya dalam OOP, kita memiliki objek puncak yang merupakan superclass atau parent. Contohnya di Java ada `Java.Lang.Object` yang merupakan superclass dari semua kelas.
- Di Go, kita memiliki interface kosong `interface{}`.
- Interface kosong adalah interface yang tidak memiliki method sama sekali, sehingga dapat menyimpan nilai dari tipe data apa pun.
- Penggunaan interface kosong pada Golang:
  - `fmt.Println(...interface{})`: Fungsi ini dapat menerima argumen dari tipe data apa pun karena menggunakan interface kosong.
  - `panic(v interface{})`: Fungsi ini menerima argumen dari tipe data apa pun dan akan menghentikan eksekusi program.
  - `recover() interface{}`: Fungsi ini digunakan untuk menangkap panic dan mengembalikan nilai yang menyebabkan panic.
  - dll
- Mirip dengan any di TypeScript, interface kosong di Go memungkinkan kita untuk menangani berbagai tipe data tanpa perlu mendefinisikan tipe data secara eksplisit.
- Contoh penggunaan interface kosong:
```go
package main
import "fmt"
func printValue(v interface{}) {
    fmt.Println("Value:", v) // mencetak nilai dari interface kosong
}
func main() {
    printValue(42)          // Output: Value: 42
    printValue("Hello")     // Output: Value: Hello
    printValue(3.14)        // Output: Value: 3.14
    printValue([]int{1, 2, 3}) // Output: Value: [1 2 3]
}
```
[interface-kosong.go](interface-kosong.go)

# Nil
- `nil` adalah nilai khusus di Go yang menunjukkan bahwa suatu variabel tidak memiliki nilai atau belum diinisialisasi
- `nil` dapat digunakan dengan berbagai tipe data, termasuk pointer, slice, map, channel, dan interface
- Contoh penggunaan `nil`:
```go
package main
import "fmt"
func main() {
    var p *int // pointer yang belum diinisialisasi, nil secara default
    if p == nil {
        fmt.Println("Pointer is nil") // Output: Pointer is nil
    }

    var s []int // slice yang belum diinisialisasi, nil secara default
    if s == nil {
        fmt.Println("Slice is nil") // Output: Slice is nil
    }

    var m map[string]int // map yang belum diinisialisasi, nil secara default
    if m == nil {
        fmt.Println("Map is nil") // Output: Map is nil
    }

    var ch chan int // channel yang belum diinisialisasi, nil secara default
    if ch == nil {
        fmt.Println("Channel is nil") // Output: Channel is nil
    }

    var i interface{} // interface kosong yang belum diinisialisasi, nil secara default
    if i == nil {
        fmt.Println("Interface is nil") // Output: Interface is nil
    }
}
```

# Type Assertion
- Type assertion adalah cara untuk mengonversi tipe data dari interface kosong ke tipe data yang lebih spesifik
- Digunakan untuk mengambil nilai dari interface kosong dan mengonversinya ke tipe data yang diinginkan
- Sintaks type assertion:
```go
value, ok := interfaceValue.(T)
```
- Di mana `interfaceValue` adalah nilai dari interface kosong, `T` adalah tipe data yang diinginkan, `value` adalah nilai yang diambil, dan `ok` adalah boolean yang menunjukkan apakah type assertion berhasil
- Contoh penggunaan type assertion:
```go
package main
import "fmt"
func main() {
    var i interface{} = "Hello, World!" // interface kosong yang berisi string

    // Type assertion untuk mengonversi ke string
    str, ok := i.(string)
    if ok {
        fmt.Println("String value:", str) // Output: String value: Hello, World!
    } else {
        fmt.Println("Type assertion failed")
    }

    // Type assertion untuk mengonversi ke int (akan gagal)
    num, ok := i.(int)
    if ok {
        fmt.Println("Integer value:", num)
    } else {
        fmt.Println("Type assertion failed for int") // Output: Type assertion failed for int
    }
}
```
- Hati-hati saat menggunakan type assertion, karena jika tipe data yang diinginkan tidak cocok, program akan panic. Untuk menghindari panic, gunakan variabel boolean `ok` untuk memeriksa apakah type assertion berhasil.
- Contoh yang salah:
```go
package main
import "fmt"
func main() {
    var i interface{} = "Hello, World!" // interface kosong yang berisi string

    // Type assertion yang salah (akan panic jika tidak menggunakan ok)
    str := i.(string) // ini akan berhasil
    fmt.Println("String value:", str)

    // Type assertion yang salah (akan panic)
    num := i.(int) // ini akan panic karena i bukan int
    fmt.Println("Integer value:", num) // tidak akan dieksekusi
}
```
- Ada cara lain untuk melakukan type assertion yang aman, yaitu dengan menggunakan `switch`, Hal ini tidak akan menyebabkan panic jika tipe data tidak cocok, dan kita bisa menangani berbagai tipe data dengan lebih mudah.:
```go
package main
import "fmt"
func main() {
    var i interface{} = "Hello, World!" // interface kosong yang berisi string

    switch v := i.(type) {
    case string:
        fmt.Println("String value:", v) // Output: String value: Hello, World!
    case int:
        fmt.Println("Integer value:", v)
    default:
        fmt.Println("Unknown type")
    }
}
```
[type-assertion.go](type-assertion.go)

# Pointer
- Ada dua jenis passing data di Go: 
  - Passing by value: Mengirimkan salinan nilai variabel ke fungsi
  - Passing by reference: Mengirimkan alamat memori dari variabel ke fungsi (menggunakan pointer)
- Pointer adalah variabel yang menyimpan alamat memori dari variabel lain
- Digunakan untuk mengakses dan memanipulasi nilai variabel secara langsung
- Deklarasi pointer:
  - Dengan menggunakan tanda `*` sebelum tipe data: `var p *int` (pointer ke int)
  - Menggunakan operator `&` untuk mendapatkan alamat memori dari variabel: `p = &x` (mendapatkan alamat dari variabel x)
  - Menggunakan operator `*` untuk mengakses nilai yang ditunjuk oleh pointer: `*p` (mengambil nilai dari alamat yang ditunjuk oleh p)
- Contoh penggunaan pointer:
```go
package main
import "fmt"
func main() {
    var x int = 10
    var p *int = &x // mendeklarasikan pointer yang menyimpan alamat x

    fmt.Println("Value of x:", x) // Output: Value of x: 10
    fmt.Println("Address of x:", &x) // Output: Address of x: <alamat memori>
    fmt.Println("Value of p (pointer):", p) // Output: Value of p (pointer): <alamat memori>
    fmt.Println("Value pointed by p:", *p) // Output: Value pointed by p: 10

    *p = 20 // mengubah nilai yang ditunjuk oleh pointer
    fmt.Println("Updated value of x:", x) // Output: Updated value of x: 20
}
```
[pointer.go](pointer.go)

# Asterisk Operator
- Asterisk operator (`*`) digunakan untuk mendeklarasikan pointer dan mengakses nilai yang ditunjuk oleh pointer
- Digunakan untuk mendapatkan alamat memori dari variabel dan mengakses nilai yang disimpan di alamat tersebut
- Contoh penggunaan asterisk operator:
```go
package main
import "fmt"
func main() {
    var x int = 10
    var p *int = &x // mendeklarasikan pointer yang menyimpan alamat x

    fmt.Println("Value of x:", x) // Output: Value of x: 10
    fmt.Println("Address of x:", &x) // Output: Address of x: <alamat memori>
    fmt.Println("Value of p (pointer):", p) // Output: Value of p (pointer): <alamat memori>
    fmt.Println("Value pointed by p:", *p) // Output: Value pointed by p: 10

    *p = 20 // mengubah nilai yang ditunjuk oleh pointer
    fmt.Println("Updated value of x:", x) // Output: Updated value of x: 20
}
```
[pointer.go](pointer.go)

# Operator New
- Operator `new` digunakan untuk mengalokasikan memori untuk tipe data tertentu dan mengembalikan pointer ke memori tersebut. Nilai yang dialokasikan akan diinisialisasi dengan nilai default untuk tipe data tersebut (misalnya, 0 untuk int, false untuk bool, dan "" untuk string).
- Digunakan untuk membuat pointer ke tipe data yang belum diinisialisasi
- Contoh penggunaan operator `new`:
```go
package main
import "fmt"
func main() {
    p := new(int) // mengalokasikan memori untuk int dan mengembalikan pointer
    fmt.Println("Value of p (pointer):", p) // Output: Value of p (pointer): <alamat memori>
    fmt.Println("Value pointed by p:", *p) // Output: Value pointed by p: 0 (nilai default int)

    *p = 42 // mengubah nilai yang ditunjuk oleh pointer
    fmt.Println("Updated value pointed by p:", *p) // Output: Updated value pointed by p: 42
}
```
[operator-new.go](operator-new.go)

# Pointer di Function
- Pointer dapat digunakan sebagai parameter fungsi untuk mengubah nilai variabel yang diteruskan
- Dengan menggunakan pointer, kita dapat menghindari salinan nilai yang besar dan menghemat memori
- Contoh penggunaan pointer di fungsi:
```go
package main
import "fmt"
func increment(x *int) {
    *x++ // mengubah nilai yang ditunjuk oleh pointer
}
func main() {
    var num int = 10
    fmt.Println("Value before increment:", num) // Output: Value before increment: 10

    increment(&num) // meneruskan alamat memori dari num
    fmt.Println("Value after increment:", num) // Output: Value after increment: 11

    // Menggunakan pointer untuk mengalokasikan memori
    p := new(int)
    *p = 20
    fmt.Println("Value pointed by p:", *p) // Output: Value pointed by p: 20
}
```
[function-pointer.go](function-pointer.go)

# Pointer di Method/Struct
- Pointer juga dapat digunakan di dalam method pada struct untuk mengubah nilai field struct
- Dengan menggunakan pointer, kita dapat mengubah nilai field struct tanpa membuat salinan dari struct tersebut
- Best practice adalah menggunakan pointer pada method struct jika kita ingin mengubah nilai field struct tersebut
- Contoh penggunaan pointer di method struct:
```go
package main
import "fmt"
type Counter struct {
    count int
}
func (c *Counter) Increment() {
    c.count++ // mengubah nilai field count
}
func (c *Counter) GetCount() int {
    return c.count // mengembalikan nilai field count
}
func main() {
    c := &Counter{count: 0} // membuat pointer ke struct Counter
    fmt.Println("Initial count:", c.GetCount()) // Output: Initial count: 0

    c.Increment() // memanggil method Increment
    fmt.Println("Count after increment:", c.GetCount()) // Output: Count after increment: 1

    c.Increment() // memanggil method Increment lagi
    fmt.Println("Count after another increment:", c.GetCount()) // Output: Count after another increment: 2
}
```
[struct-pointer.go](struct-pointer.go)

# Package
- Package adalah kumpulan kode Go yang dapat digunakan kembali
- Digunakan untuk mengorganisir kode dan memisahkan fungsionalitas
- Setiap file Go harus berada dalam package, dan setiap package harus memiliki nama yang unik

# Access Modifier
- Go memiliki tiga tingkat akses modifier untuk mengontrol visibilitas variabel, fungsi, dan tipe data:
  - Public: Dapat diakses dari package lain, ditandai dengan huruf kapital pertama (contoh: `PublicFunction`)
  - Private: Hanya dapat diakses dalam package yang sama, ditandai dengan huruf kecil pertama (contoh: `privateFunction`)

# Package Initialization
- Package dapat memiliki fungsi `init()` yang akan dieksekusi secara otomatis saat package diimpor
- Fungsi `init()` digunakan untuk inisialisasi variabel atau konfigurasi yang diperlukan sebelum fungsi lain dalam package dapat dipanggil
- Dalam real case, fungsi `init()` sering digunakan untuk:
  - Menginisialisasi koneksi database
  - Membaca konfigurasi dari file atau environment variable
  - Mendaftarkan fungsi atau tipe data ke dalam sistem
```go
package main
import (
    "fmt"
)
func init() {
    fmt.Println("Package initialized") // akan dieksekusi saat package diimpor
}
func main() {
    fmt.Println("Main function executed") // akan dieksekusi setelah init()
}
```
[package-initialization/init.go](package-initialization/init.go)
[package-initialization.go](package-initialization.go)

# Error
- Error adalah tipe data built-in di Go yang digunakan untuk menangani kesalahan
- Tipe data error adalah interface yang memiliki satu method `Error() string`
```go
type Error interface {
    Error() string // mengembalikan pesan kesalahan
}
```

- Jika ingin membuat error sendiri, kita dapat membuat struct yang mengimplementasikan interface `Error`:
```go
package main
import (
    "fmt"
)
type MyError struct {
    Message string
}
func (e MyError) Error() string {
    return e.Message // mengembalikan pesan kesalahan
}

func main() {
    err := MyError{Message: "Something went wrong"}
    fmt.Println("Error:", err.Error()) // Output: Error: Something went wrong

    // Menggunakan error dalam fungsi
    if err := doSomething(); err != nil {
        fmt.Println("Error occurred:", err)
    }
}
```

# Membuat Custom Error
- Kita dapat membuat custom error dengan mendefinisikan struct yang mengimplementasikan interface `error`
- Contoh membuat custom error:
```go
package main
import (
    "fmt"
)
type validationError struct {
	Message string
}

type notFoundError struct {
    Message string
}

func (e *validationError) Error() string {
    return e.Message
}

func (e *notFoundError) Error() string {
    return e.Message
}

func saveData(data string) error {
    if data == "" {
		return &validationError{Message: "Data cannot be empty"} // mengembalikan custom error
    }
    // Simulasi penyimpanan data
	if data == "yoga" {
        return &notFoundError{Message: "Yoga data not found"} // mengembalikan custom error
    }
    fmt.Println("Data saved:", data)
    return nil
}

func main() {
    err := saveData("") // mencoba menyimpan data kosong
    if err != nil {
        if vErr, ok := err.(*validationError); ok {
            fmt.Println("Validation Error:", vErr.Message) // Output: Validation Error: Data cannot be empty
        } else if nErr, ok := err.(*notFoundError); ok {
            fmt.Println("Not Found Error:", nErr.Message) // Output: Not Found Error: Yoga data not found
        } else {
            fmt.Println("Error:", err)
        }
    }
}

```

# FINISH