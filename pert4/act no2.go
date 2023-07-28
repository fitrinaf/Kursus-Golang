package main

import "fmt"

// FITRI NUR AFIFAH 10121491

func main() {
	var kursus = []string{"Go", "Python", "Ruby", "Java", "C++", "C#"}
	var kursus_saya = kursus[1:5]
	kursus_saya = append(kursus_saya, "Manajemen Informatika")

	fmt.Println("Isi dari kursus awal adalah : ", kursus)
	fmt.Printf("panjang kursus adalah %d dan kapasitas %d\n", len(kursus), cap(kursus))

	fmt.Println("Isi dari kursus awal adalah : ", kursus_saya)
	fmt.Printf("panjang kursus_saya adalah %d dan kapasitas %d\n", len(kursus_saya), cap(kursus_saya))

}
