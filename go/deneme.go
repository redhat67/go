package main

import "fmt"

func deneme() {

	var name = "Atamer"
	var age int16 = -256
	var İsMarried bool = true // veriyi atadığımız kısım
	var weight float32 = 72.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(İsMarried) // = üstte atama yaparken veriyi çalıştırdığımız kısım

	fmt.Println(weight)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", İsMarried) //= değişkenlerini sıralıyor
	fmt.Printf("%T\n", weight)

}


