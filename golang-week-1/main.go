package main

import (
	"fmt"
)

func main() {
	// * manifest typing
	// * declare + initialize
	var firstName string = "john"

	// * declare
	var lastName string

	// * assign value
	lastName = "wick"

	// * multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// * ===================

	// * type inference
	anotherName := "Bambang"
	var anotherLastName = "Udin"

	anotherName = "Ilham"
	anotherLastName = "Parker"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// * dengan tipe data berbeda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	_, _, _, _, _, _, _, _, _ = first, second, third, fourth, fifth, sixth, seventh, eight, ninth
	_, _, _, _ = one, isFriday, twoPointTwo, say

	fmt.Println("Hello world!")
	fmt.Printf("Hello %s %s, and %s %s !\n", firstName, lastName, anotherName, anotherLastName)

	// * keyword new()
	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)

	*name = "sarah"

	fmt.Println(*name)
	fmt.Println(&name)

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("%T, %T\n", negativeNumber, positiveNumber)

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var message = `Nama saya "John Wick".	
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message)

	const namaOrang string = "idham"
	fmt.Printf("Halo %s !\n", namaOrang)

	const (
		square        = "kotak"
		phi           = 3.14
		numeric uint8 = 8
		isToday bool  = false
	)

	_, _, _, _ = square, phi, numeric, isToday

	var point uint8 = 1
	var anotherPoint int16 = 2561

	fmt.Printf("%T %t \n", point, (point == uint8(anotherPoint)))

	fmt.Printf("%T %d \n", uint8(anotherPoint), uint8(anotherPoint))

	// * jika cakupan bilangan pada variable decimal lebih kecil daripada nilai yg di assign, value akan menjadi 1
}
