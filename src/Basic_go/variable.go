package main

import (
	"fmt"
	"strings"
)

func main() {
	//deklarasi dengan var dan tipe data
	var namaKakakStat string = "Merosa Pesona"
	var namaAdikStat string = "Anosep Asorem"

	//deklarasi dengan var saja
	//tipe data menyesuaikan nilai yang diberikan
	var namaKakakDyn1 = "Merosa Pesona"
	var namaAdikDyn1 = "Anosep Asorem"

	//deklarasi tanpa var dan tipe data
	//menyesuaikan nilai yang diberikan
	namaKakakDyn2 := "Merosa Pesona"
	namaAdikdyn2 := "Anosep Asorem"

	//multiple declaration dengan var dan tipe data
	var namaKakakStatMulti, namaAdikStatMulti string = "Merosa Pesona", "Anosep Asorem"

	//multiple declaration tanpa var dan tipe data
	namaKakakDynMulti, namaAdikDynMulti := "Merosa Pesona", "Anosep Asorem"

	//output dari semua variable dari berbagai cara deklarasi
	fmt.Printf("Nama Kakak : %s\n", namaKakakStat)
	fmt.Printf("Nama Adik : %s\n\n", namaAdikStat)

	fmt.Printf("Nama Kakak : %s\n", namaKakakDyn1)
	fmt.Printf("Nama Adik : %s\n\n", namaAdikDyn1)

	fmt.Printf("Nama Kakak : %s\n", namaKakakDyn2)
	fmt.Printf("Nama Adik : %s\n\n", namaAdikdyn2)

	fmt.Println("Nama Kakak : ", namaKakakStatMulti)
	fmt.Println("Nama Adik : ", namaAdikStatMulti+"\n")

	fmt.Println("Nama Kakak : ", namaKakakDynMulti)
	fmt.Println("Nama Adik : ", namaAdikDynMulti+"\n")

	//block fungsi untuk mendapatkan nilai string
	//var ke pertama secara terbalik
	//yang akan dijadikan nilai var ke dua

	var namaKakak string = "Merosa Pesona"
	temp := []rune(namaKakak)  //rune adalah tipe data selayaknya char pada c++
	temp2 := []rune(namaKakak) //rune adalah tipe data selayaknya char pada c++
	//temp2 dideklarasi ulang untuk mengalokasikan memori.

	for i := (len(namaKakak) - 1); i >= 0; i-- {
		temp2[len(namaKakak)-1-i] = temp[i]
	}

	//ToLower untuk membuat seluruh char dalam string menjadi lowercase
	//Title untuk membuat seluruh huruf pertama pada kata menjadi Uppercase(kapital)
	namaAdik := strings.Title(strings.ToLower(string(temp2)))

	fmt.Println("Nama Kakak : ", namaKakak)
	fmt.Println("Nama Adik : ", namaAdik)
}
