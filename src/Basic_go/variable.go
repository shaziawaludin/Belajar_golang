package main

import "fmt"

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
	fmt.Printf("Nama Adik : %s\n", namaAdikStat)

	fmt.Printf("Nama Kakak : %s\n", namaKakakDyn1)
	fmt.Printf("Nama Adik : %s\n", namaAdikDyn1)

	fmt.Printf("Nama Kakak : %s\n", namaKakakDyn2)
	fmt.Printf("Nama Adik : %s\n", namaAdikdyn2)

	fmt.Println("Nama Kakak : ", namaKakakStatMulti+"\n")
	fmt.Println("Nama Adik : ", namaAdikStatMulti+"\n")

	fmt.Println("Nama Kakak : ", namaKakakDynMulti+"\n")
	fmt.Println("Nama Adik : ", namaAdikDynMulti+"\n")
}
