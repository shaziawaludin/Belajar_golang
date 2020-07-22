package main

import (
	"fmt"
	"strings"
)

func main() {
	var namaKakakStat string = "Merosa Pesona"
	var namaAdikStat string = "Anosep Asorem"

	var namaKakakDyn1 = "Merosa Pesona"
	var namaAdikDyn1 = "Anosep Asorem"

	namaKakakDyn2 := "Merosa Pesona"
	namaAdikdyn2 := "Anosep Asorem"

	var namaKakakStatMulti, namaAdikStatMulti string = "Merosa Pesona", "Anosep Asorem"

	namaKakakDynMulti, namaAdikDynMulti := "Merosa Pesona", "Anosep Asorem"

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

	var namaKakak string = "Merosa Pesona"
	temp := []rune(namaKakak)
	temp2 := []rune(namaKakak)

	for i := (len(namaKakak) - 1); i >= 0; i-- {
		temp2[len(namaKakak)-1-i] = temp[i]
	}

	namaAdik := strings.Title(strings.ToLower(string(temp2)))

	fmt.Println("Nama Kakak : ", namaKakak)
	fmt.Println("Nama Adik : ", namaAdik)
}
