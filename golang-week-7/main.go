package main

/*
A.26 EXPORTED AND UNEXPORTED PROPERTY
*/

// func main() {
// 	var s1 = l.Student{"ethan", 21}
// 	fmt.Println("name", s1.Name)
// 	fmt.Println("grade", s1.Grade)
// }

// func main() {
// 	sayHello("ethan")
// }

// func main() {
// 	fmt.Printf("Name  : %s\n", library.Student.Name)
// 	fmt.Printf("Grade : %d\n", library.Student.Grade)
// }

/*
A.27 INTERFACE
*/
import (
	"fmt"
	"math"
	"strings"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

// EMBEDDED INTERFACE

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung2 interface {
	hitung2d
	hitung3d
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("====== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	var bangunRuang hitung2 = &kubus{4}

	fmt.Println("======== kubus")
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())

	fmt.Println("================================================================")

	/*
		A.28 ANY INTERFACE
	*/

	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "mango", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// interface{} punya alias bisa ditulis dengan any

	var data map[string]any

	data = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)
	// fmt.Println(secret.(int), "casting to int") //error

	var number = secret.(float64) * 10
	fmt.Println(secret, "multipled by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	fmt.Println("================================================================")

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

}
