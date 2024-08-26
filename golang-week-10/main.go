package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	/*
		A.37. Error, Panic and Recover
	*/

	// var input string
	// fmt.Println("Type some number: ")

	// fmt.Scanln(&input)

	// var number int
	// var err error

	// number, err = strconv.Atoi(input)

	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(input, "is not number")
	// 	fmt.Println(err.Error())
	// }

	// custom error handling

	// defer catch()

	// var name string
	// fmt.Println("Type your name: ")
	// fmt.Scanln(&name)

	// if valid, err := validate(name); valid {
	// 	fmt.Println("halo", name)
	// } else {
	// 	panic(err.Error())
	// 	fmt.Println("end")
	// }

	// recover pada IIFE

	// data := []string{"superman", "batman", "ironman", "loki"}

	// for _, each := range data {
	// 	func() {
	// 		defer func() {
	// 			if r := recover(); r != nil {
	// 				fmt.Println("Panic occured on looping", each, "| message:", r)
	// 			} else {
	// 				fmt.Println("Application running perfectly")
	// 			}
	// 		}()

	// 		panic("some error happened")
	// 	}()
	// }

	/*
		A.38 Layout Format String
	*/

	type student struct {
		name        string
		height      float64
		age         int32
		isGraduated bool
		hobbies     []string
	}

	var data = student{
		name:        "wick",
		age:         32,
		height:      182.5,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}

	// Digunakan untuk memformat data numerik yang merupakan kode unicode,
	// menjadi bentuk string karakter unicode-nya.

	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)

	// Digunakan untuk memformat data numerik,
	// menjadi bentuk string numerik berbasis 10 (basis bilangan yang kita gunakan).

	fmt.Printf("%d\n", data.age)

	// Digunakan untuk memformat data numerik desimal
	// ke dalam bentuk notasi numerik standar Scientific notation.

	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)

	// Berfungsi untuk memformat data numerik desimal
	// dengan lebar desimal bisa ditentukan.
	// Secara default lebar digit desimal adalah 6 digit.

	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)

	// Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan.
	// Lebar kapasitasnya sangat besar, pas digunakan untuk data yang jumlah digit desimalnya cukup banyak.

	fmt.Printf("%g\n", 0.12131312232312132343545546) // ada maksimal nya, ga bisa sampai 26 character
	fmt.Printf("%.26f\n", 0.12131312232312132343545546)

	// Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 8 (oktal).

	fmt.Printf("%o\n", data.age)

	// Digunakan untuk memformat data pointer, mengembalikan alamat pointer referensi variabel-nya.

	fmt.Printf("%p\n", &data.name)

	/*
		A.39 RANDOMIZER
	*/

	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println(time.Now().UTC().UnixNano())
	fmt.Println("random ke-1:", randomizer.Int()) // 5221277731205826435
	fmt.Println("random ke-2:", randomizer.Int()) // 3852159813000522384
	fmt.Println("random ke-3:", randomizer.Int()) // 8532807521486154107

	fmt.Println("random int:", randomizer.Int())
	fmt.Println("random float32:", randomizer.Float32())
	fmt.Println("random uint:", randomizer.Uint32())

	fmt.Println("random index tertentu:", randomizer.Intn(99))

	fmt.Println("Random string 5 karakter:", randomString(5))

}

func randomString(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured: ", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}
