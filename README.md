<h1>Indonesian Numeral Speller</h1>
by Leonardo / 13517048 - Calon IRK 2019
<h2>Fungsi Program</h2>

Berikut adalah fungsi dan prosedur yang terdapat pada file source code
- <i>function</i> oneToEleven(angka : <i>integer</i>) -> string
- <i>function</i> TwelvetoBillion(angka : <i>integer</i>) -> string
- <i>function</i> IntToString(angka : <i>integer</i>) -> string<br>
{fungsi utama pengubahan dari integer menjadi string}
- <i>function</i> find(arr : <i>array of</i> string, val : string) -> <i>integer</i><br>
{mengembalikan kemunculan pertama elemen val pada arr}
- <i>function</i> baseStringToInt(kata : string) -> <i>integer</i>
- <i>procedure</i> normalize(arr : <i>array of</i> string)
- <i>function</i> getNum(arr_kata : <i>array</i> [0..12] <i>of</i> string) -> <i>integer</i>
- <i>function</i> StringToInt(kata : string) -> <i>integer</i><br>
{fungsi utama pengubahan dari string menjadi integer}
- <i>procedure</i> methodPOST(gc : gin.Context)<br>
{prosedur untuk fitur POST}
- <i>procedure</i> methodGET(gc : gin.Context)<br>
{prosedur untuk fitur GET}

<h2>Cara menjalankan</h2>
Requirements:<br>

- Pada komputer anda terinstalasi compiler bahasa pemrograman go (dapat dilihat di https://golang.org/doc/install).
- Pada komputer anda terinstalasi library bahasa GO: gin dan cors (install dengan command ```go get -u github.com/gin_gonic/gin``` dan ```go get -u github.com/gin_contrib/cors```).
- Memiliki aplikasi POSTMAN.
- Memiliki sebuah browser yang support HTML dan JavaScript.

Langkah penjalanan program:<br>

<h4>1. Pada POSTMAN</h4>

- Masuk ke direktori mainprog.go disimpan.
- Jalankan perintah ```go run mainprog.go```.
- Buka pada Aplikasi POSTMAN dan buka ```localhost:8080``` (tempat Rest API dijalankan).
- <b>CONVERTING</b>
    - Untuk pengubahan String ke Integer, pindahkan method ke ```POST``` dan Send ```localhost:8080/read?text=<isi string anda>```.
    - Untuk pengubahan Integer ke String, pindahkan method ke ```GET``` dan Send ```localhost:8080/spell?number=<isi angka anda>```.
- Lihat hasilnya pada Body hasil status pengerjaan.

<h4>2. Pada Web</h4>

- Masuk ke direktori mainprg.go disimpan.
- Jalankan perintah ```go run mainprog.go```.
- Buka file main.html pada browser anda.
- <b>CONVERTING</b>
    - Untuk pengubahan Integer ke String, isi textarea yang atas dan klik perintah 'Convert to String'.
    - Untuk pengubahan String ke Integer, isi textarea yang bawah dan klik perintah 'Convert to Integer'.

    *catatan: textarea pada web fleksibel ukuran
- Jawaban akan ditunjukkan pada sebuah alert.

<h2>Contoh request and response</h2>
Gambar 1: POST dan input valid:<br>

![postvalid](Test/post_valid.jpg)<br>

Gambar 2: POST dan input tidak valid:<br>

![postinvalid](Test/post_invalid.jpg)<br>

Gambar 3: GET dan input valid:<br>

![getvalid](Test/get_valid.jpg)<br>

Gambar 4: GET dan input tidak valid:<br>

![getinvalid](Test/get_invalid.jpg)<br>

Gambar 5: Penjalanan Number to Text/String valid pada aplikasi Web:<br>

![nttvalid](Test(JS)/numtotext_valid.jpg)<br>

Gambar 6: Penjalanan Number to Text/String invalid pada aplikasi Web:<br>

![nttinvalid](Test(JS)/numtotext_invalid.jpg)<br>

Gambar 7: Penjalanan Text/String to Number valid pada aplikasi Web:<br>

![ttnvalid](Test(JS)/texttonum_valid.jpg)<br>

Gambar 8: Penjalanan Text/String to Number invalid pada aplikasi Web:<br>

![ttnvalid](Test(JS)/texttonum_invalid.jpg)<br>

<h2>Arsitektur kode Go</h2>

```
$GOPATH
|---src
    |---bufio
    |   |---*lib
    |
    |---fmt
    |   |---*lib
    |
    |---github.com
    |   |---Indonesian-Numeral-Spellers
    |   |   |---src
    |   |   |   |---mainprog.go
    |   |   |
    |   |   |---Test
    |   |   |   |---*screenshot hasil
    |   |   |
    |   |   |---doc
    |   |       |---README.md
    |   |
    |   |---gin_contrib
    |   |   |---cors
    |   |       |---*lib
    |   |
    |   |---gin_gonic
    |   |   |---gin
    |   |       |---*lib
    |   
    |---os
    |   |---*lib
    |
    |---strconv
    |   |---*lib
    |
    |---strings
    |   |---*lib

.
.
.

JavaScript
|---Indonesian-Numeral-Spellers
    |---src
        |---main.html
        |---script.js
    |--Test(JS)
        |---*screenshot hasil
...
```
*beberapa directory/file yang tidak dirasa penting untuk dicantumkan tidak dicantumkan<br>
*beberapa directory/file yang tidak dirasa penting untuk dijelaskan tidak dijelaskan