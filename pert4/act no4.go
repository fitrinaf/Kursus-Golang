package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) BirthdayPointer() {
	p.Age++
	fmt.Println("dalam method BirthdayValue umurnya", p.age)
	fmt.Println()
}

func main() {
	var p Person
	p.Name = "putra"
	p
}
