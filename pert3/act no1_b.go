package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan input: ")
	fmt.Scan(&input)

	for i := 1; i <= input; i += 2 {
		fmt.Println("Angka ", i)
	}
}
