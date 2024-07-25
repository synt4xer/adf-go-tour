// A.23 - A.24.8

package main

import (
	"fmt"
	"strings"
)

// A.24.1 Deklarasi struct
type Person struct {
	name string
	age  int
}

type Student struct {
	Person // embedded in struct
	age    int
	grade  int
}

type student struct {
	Person struct { // anonymous struct
		name string
		age  int
	}
	grade   int
	hobbies []string
}

type People = Person // type alias

type People2 = struct {
	name string
	age  int
}

type Students1 struct {
	name string
	age  int
}

func (s Students1) SayHello() {
	fmt.Println("Hello", s.name)
}

func (s Students1) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// to effect change name, need using pointer
func (s *Students1) changeName(name string) {
	s.name = name
}

func main() {
	// A.23 Pointer
	// A.23.1 Penerapan Pointer
	var numberA int = 4         // variable biasa
	var numberB *int = &numberA // mengambil nilai pointer dari variable numberA (metode referencing)

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // alamat pointer

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // alamat pointer

	// A.23.2 Efek Perubahan Nilai Pointer
	fmt.Println("")
	numberA = 5

	fmt.Println("numberA (value)   :", numberA)  // 5
	fmt.Println("numberA (address) :", &numberA) // alamat pointer

	fmt.Println("numberB (value)   :", *numberB) // 5
	fmt.Println("numberB (address) :", numberB)  // alamat pointer

	// A.23.3 Parameter Pointer
	fmt.Println("")
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10

	change(&number, 4)
	fmt.Println("after  :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 4

	// The end of A.23 Pointer

	// A.24 Struct
	fmt.Println("")

	// A.24.2 Penerapan Struct Untuk Membuat Object
	// var s1 Student
	// s1.name = "John Wick"
	// s1.grade = 2
	// fmt.Println("name  :", s1.name)
	// fmt.Println("grade :", s1.grade)

	// A.24.3 Inisialisasi object struct
	// var s2 = Student{}
	// s2.name = "wick"
	// s2.grade = 2

	// var s3 = Student{"Ethan", 2}

	// var s4 = Student{name: "Jason"}

	// fmt.Println("student 1 :", s2.name)
	// fmt.Println("student 2 :", s3.name)
	// fmt.Println("student 3 :", s4.name)

	// another way
	/* var s5 = student{name: "wayne", grade: 2}
	var s6 = student{grade: 2, name: "bruce"} */

	// A.24.4 Variable objek pointer
	fmt.Println("")
	// var s1Point = Student{name: "Dado", grade: 2}
	// var s2Point *Student = &s1Point

	// fmt.Println("student 1, name :", s1Point.name)
	// fmt.Println("student 4, name :", s2Point.name)
	// s2Point.name = "Feri"
	// fmt.Println("student 1, name :", s1Point.name)
	// fmt.Println("student 4, name :", s2Point.name)

	// A.24.5 - A.24.6 Embedded struct & Nama Property Yang Sama
	var s1Embed = Student{}
	s1Embed.name = "wick"
	s1Embed.age = 21        // age of student
	s1Embed.Person.age = 22 // age of person
	s1Embed.grade = 2

	fmt.Println("name  :", s1Embed.Person.name)
	fmt.Println("age   :", s1Embed.age)
	fmt.Println("Person age   :", s1Embed.Person.age)
	fmt.Println("grade :", s1Embed.grade)

	// A.24.7 Pengisian nilai substruct
	fmt.Println("")

	var p1 = Person{name: "Niko", age: 34}
	var s1 = Student{Person: p1, grade: 2}
	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("age of person   :", s1.Person.age)
	fmt.Println("grade :", s1.grade)

	// A.24.8 Anonymous struct
	// initialized with zero value
	fmt.Println("")
	var s1Anon = struct {
		Person
		grade int
	}{}

	s1Anon.Person = Person{"andar", 37}
	s1Anon.grade = 2

	// initialized with value
	var _ = struct {
		Person
		grade int
	}{
		Person: Person{"wick", 21},
		grade:  2,
	}

	fmt.Println("name  :", s1Anon.Person.name)
	fmt.Println("age   :", s1Anon.Person.age)
	fmt.Println("grade :", s1Anon.grade)

	// The end of A.24 Struct
	var students = []Person{
		{name: "diah", age: 2},
		{name: "ilham", age: 4},
		{name: "iki", age: 6},
	}

	for _, value := range students {
		fmt.Println(value.name, "age is", value.age)
	}

	var teachers = []struct {
		Person
		grade int
	}{
		{Person: Person{"udin", 2}, grade: 8},
		{Person: Person{"adil", 4}, grade: 10},
		{Person: Person{"jarwo", 6}, grade: 20},
	}

	for _, teacher := range teachers {
		fmt.Println("teacher name is", teacher.name)
		fmt.Println("teacher age is", teacher.age)
		fmt.Println("teacher grade is", teacher.grade)
	}

	var anotherStudent = student{}
	anotherStudent.Person = struct {
		name string
		age  int
	}{"adi", 2}
	anotherStudent.grade = 2
	anotherStudent.hobbies = []string{"mancing", "maling"}

	var orang = People{name: "orang", age: 20}
	fmt.Println(orang)

	// A.25 Method
	var dika = Students1{"dika rungin", 22}
	dika.SayHello()
	var dikaName = dika.getNameAt(2)
	fmt.Println(dikaName)

	// ??? gada bedanya
	var andar = &Students1{"andar niki", 30}
	andar.SayHello()
	andar.changeName("adi lego")
	andar.SayHello()
}

func change(original *int, value int) {
	*original = value
}
