package main

import "fmt"

func main() {
	var uts, uas int

	fmt.Print("Masukkan Nilai UTS : ")
	fmt.Scan(&uts)
	fmt.Print("Masukkan Nilai UAS : ")
	fmt.Scan(&uas)

	total := (uts * 30 / 100) + (uas * 70 / 100)

	if total <= 20 {
		fmt.Printf("Grade E")
	}
	if total > 20 && total <= 40 {
		fmt.Printf("Grade D")
	}
	if total > 41 && total <= 60 {
		fmt.Printf("Grade C")
	}
	if total > 61 && total <= 80 {
		fmt.Printf("Grade B")
	}
	if total > 81 && total <= 100 {
		fmt.Printf("Grade A")
	}
}
