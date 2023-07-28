package main

import "fmt"

func main() {
	var i int

	for i := i; i < 10; {

		fmt.Print("Input Nilai I :")
		fmt.Scan(&i)

		if i%2 == 0 {
			fmt.Println(i, " Adalah Angka Genap\n")
		} else {
			fmt.Println(i, " Adalah Angka Ganjil\n")
		}
	}
	fmt.Print("Pertanyaan Selesai, Karena Angka I yang di input sudah melebihi dari 10. Terimakasih")

}
