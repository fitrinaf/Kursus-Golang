package main

import "fmt"

// FITRI NUR AFIFAH 10121491
type mahasiswa struct {
	nama  string
	kelas string
}

func main() {
	var data = map[string]mahasiswa{
		"10121491": {
			"Fitri Nur Afifah",
			"2KA23",
		},
		"10121401": {
			"Balya Najla Nafisah",
			"2KA23",
		},
	}
	var search string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari : ")
	fmt.Scanln(&search)
	var result, ok = data[search]
	if ok {
		fmt.Println("Nama : ", result.nama)
		fmt.Println("Kelas : ", result.kelas)
	} else {
		fmt.Println("NIM tidak ditemukan")
	}
}
