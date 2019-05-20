package main

import (
	"bufio"   //untuk input string
	"fmt"     //untuk output
	"os"      //untuk input string
	"strconv" //untuk convert string berisi angka to integer
	"strings" //library strings

	"github.com/gin-contrib/cors" //untuk konfigurasi
	"github.com/gin-gonic/gin"    //untuk konfigurasi
)

//INT TO STRING
func oneToEleven(angka int) string {
	//fungsi dasar pengubah angka menjadi kata
	switch angka {
	case 0:
		return ""
	case 1:
		return "satu"
	case 2:
		return "dua"
	case 3:
		return "tiga"
	case 4:
		return "empat"
	case 5:
		return "lima"
	case 6:
		return "enam"
	case 7:
		return "tujuh"
	case 8:
		return "delapan"
	case 9:
		return "sembilan"
	case 10:
		return "sepuluh"
	case 11:
		return "sebelas"
	default:
		return "Input tidak Valid"
	}
}

func TwelveToBillion(angka int) string {
	//fungsi rekursif converter int to string
	if angka < 12 { //--BASIS 1
		return oneToEleven(angka)
	} else if angka/10 == 1 { //--BASIS 2
		return oneToEleven(angka%10) + " belas"
	} else { //angka 20 keatas --REKURENS
		if angka%100 == angka {
			//puluhan
			return oneToEleven(angka/10) + " puluh " + TwelveToBillion(angka%10)
		} else if angka%1000 == angka {
			//ratusan
			if angka/100 == 1 {
				return "seratus " + TwelveToBillion(angka%100)
			} else {
				return oneToEleven(angka/100) + " ratus " + TwelveToBillion(angka%100)
			}
		} else if angka%1000000 == angka {
			//ribuan
			if angka/1000 == 1 {
				return "seribu " + TwelveToBillion(angka%1000)
			} else {
				return TwelveToBillion(angka/1000) + " ribu " + TwelveToBillion(angka%1000)
			}
		} else if angka%1000000000 == angka {
			//jutaan
			return TwelveToBillion(angka/1000000) + " juta " + TwelveToBillion(angka%1000000)
		} else if angka%10000000000 == angka {
			//milyaran (spek hanya sampai satu digit di depan milyar)
			return oneToEleven(angka/1000000000) + " milyar " + TwelveToBillion(angka%1000000000)
		} else {
			return "Input tidak Valid"
		}
	}
}

func IntToString(angka int) string {
	//fungsi pemanggil utama untuk pengubahan int ke string
	if angka == 0 {
		return "nol"
	} else if angka > 0 {
		return TwelveToBillion(angka)
	} else {
		return "Input tidak Valid"
	}
}

//STRING TO INT
func find(arr []string, val string) int {
	//fungsi pencarian berdasarkan value untuk array of string
	//return value: index
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1 //not found
}

func baseStringToInt(kata string) int {
	switch kata {
	case "nol":
		return 0
	case "satu":
		return 1
	case "dua":
		return 2
	case "tiga":
		return 3
	case "empat":
		return 4
	case "lima":
		return 5
	case "enam":
		return 6
	case "tujuh":
		return 7
	case "delapan":
		return 8
	case "sembilan":
		return 9
	default:
		return -10000000000 //kata not found
	}
}

func normalize(arr []string) {
	//mengubah seribu menjadi satu ribu
	if find(arr, "seribu") != -1 {
		arr[find(arr, "seribu")] = "satu ribu"
	}
	//mengubah seratus menjadi satu ratus
	if find(arr, "seratus") != -1 {
		arr[find(arr, "seratus")] = "satu ratus"
	}
	//mengubah sepuluh menjadi satu puluh
	if find(arr, "sepuluh") != -1 {
		arr[find(arr, "sepuluh")] = "satu puluh"
	}
	//mengubah sebelas menjadi satu belas
	if find(arr, "sebelas") != -1 {
		arr[find(arr, "sebelas")] = "satu belas"
	}
}

func getNum(arr_kata [13]string) int {
	//prekondisi: arr_kata sudah dalam bentuk ratusan (paling maksimal)
	var res int = 0
	for i := 0; i < cap(arr_kata); i++ {
		if arr_kata[i] == "ratus" || arr_kata[i] == "puluh" || arr_kata[i] == "belas" || arr_kata[i] == "" {
			continue
		} else {
			if arr_kata[i+1] == "ratus" {
				res += baseStringToInt(arr_kata[i]) * 100
			} else if arr_kata[i+1] == "puluh" {
				res += baseStringToInt(arr_kata[i]) * 10
			} else if arr_kata[i+1] == "belas" {
				res += baseStringToInt(arr_kata[i]) + 10
			} else if arr_kata[i+1] == "" {
				res += baseStringToInt(arr_kata[i])
			}
		}
	}
	return res
}
func StringToInt(kata string) int {
	//ASUMSI: input pasti benar dan lowercase
	temp := strings.Split(kata, "\n")
	arr_kata := strings.Split(temp[0], " ")
	normalize(arr_kata)
	temp_str := strings.Join(arr_kata, " ")
	arr_kata = strings.Split(temp_str, " ")
	//1,999,999,999 will be seperated to 1 999 999 999
	var mily [13]string //13 angka asal
	var juta [13]string
	var ribu [13]string
	var pulu [13]string
	var i int = 0
	var j int = 0
	var k int = 0
	if find(arr_kata, "milyar") != -1 { //milyaran
		i = find(arr_kata, "milyar")
		j = find(arr_kata, "milyar")
		k = find(arr_kata, "milyar")
		for iter := 0; iter < i; iter++ {
			mily[iter] = arr_kata[iter]
		}
		i = i + 1
		j = j + 1
		k = k + 1
	}
	if find(arr_kata, "juta") != -1 { //jutaan
		j = find(arr_kata, "juta")
		k = find(arr_kata, "juta")
		for iter := 0; iter < j-i; iter++ {
			juta[iter] = arr_kata[iter+i]
		}
		j = j + 1
		k = k + 1
	}
	if find(arr_kata, "ribu") != -1 { //ribuan
		k = find(arr_kata, "ribu")
		for iter := 0; iter < k-j; iter++ {
			ribu[iter] = arr_kata[iter+j]
		}
		k = k + 1
	}
	for iter := 0; iter < len(arr_kata)-k; iter++ { //ratusan-puluhan-satuan
		pulu[iter] = arr_kata[iter+k]
	}
	return getNum(mily)*1000000000 + getNum(juta)*1000000 + getNum(ribu)*1000 + getNum(pulu)
}

func procAngkakeString() {
	var inp int
	fmt.Scan(&inp)
	fmt.Print("angka ", inp, " merupakan angka -> ")
	fmt.Println(IntToString(inp))
}

func procStringkeAngka() {
	reader := bufio.NewReader(os.Stdin)
	temp, _, _ := reader.ReadLine()
	var inp string = string(temp)
	fmt.Print("angka ", inp, " merupakan angka -> ")
	fmt.Println(StringToInt(inp))
}

func methodPOST(gc *gin.Context) {
	//untuk string ke angka
	inp := gc.Query("text")
	res := StringToInt(inp)
	if res < 0 { //string invalid (ada kata di string yang tidak ada)
		gc.JSON(200, gin.H{
			"status": "OK",
			"number": "Input tidak Valid",
		})
	} else {
		gc.JSON(200, gin.H{
			"status": "OK",
			"number": res,
		})
	}
}

func methodGET(gc *gin.Context) {
	//untuk angka ke string
	inp := gc.Query("number")
	num, err := strconv.Atoi(inp)
	if num < 0 || err != nil { //angka invalid
		gc.JSON(200, gin.H{
			"status": "OK",
			"text":   "Input tidak Valid",
		})
	} else {
		res := IntToString(num)
		gc.JSON(200, gin.H{
			"status": "OK",
			"text":   res,
		})
	}
}

func testDriver() { //test driver
	fmt.Print("Masukkan POST untuk string ke angka\nMasukkan GET untuk angka ke string\nInput: ")
	reader := bufio.NewReader(os.Stdin)
	temp, _, _ := reader.ReadLine()
	var inp string = string(temp)
	if inp == "POST" {
		procStringkeAngka()
	} else if inp == "GET" {
		procAngkakeString()
	}
}

func main() { //program utama -> inisialisasi dan run gin dan cors
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{"POST", "GET"} //fitur
	router := gin.Default()
	router.Use(cors.New(config))
	router.POST("/read", methodPOST) //POST dapat melakukan read
	router.GET("/spell", methodGET)  //GET dapat melakukan spell
	router.Run()
}
