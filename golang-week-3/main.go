// A15 - A16

package main

import (
	"fmt"
)

func main() {

	// ARRAY

	var names [4]string

	names[0] = "James"
	names[1] = "Indro"
	names[2] = "Dono"
	names[3] = "Kasino"

	fmt.Println(names)

	var foods = [4]string{"pempek", "pecel lele", "bakso", "mie"}

	fmt.Println(foods)

	var numbers = [...]int{2, 3, 2, 4, 3, 2}

	fmt.Println(numbers)

	// array multidimensi
	var numbers1 = [3][3]int{{1, 2, 3}, {3, 2, 1}, {4, 2, 4}}

	fmt.Println(numbers1)

	// looping over array
	for i := 0; i < len(foods); i++ {
		fmt.Println(i, foods[i])
	}

	for _, food := range foods {
		fmt.Println(food)
	}

	// array allocation using make

	var fruits = make([]string, 4)

	fruits[0] = "orange"
	fruits[1] = "strawberry"
	fruits[2] = "mango"
	fruits[3] = "pear"

	fmt.Println(fruits)

	// SLICE

	var veggies = []string{"broccoli", "carrot", "spinach", "tomato", "potato"}

	fmt.Println(veggies)

	var sliceFoods = foods[0:3] // creating slice from array of foods

	fmt.Println(sliceFoods)

	sliceFoods[0] = "ketoprak"

	fmt.Println(foods)
	fmt.Println(sliceFoods)

	fmt.Println(&foods[0]) // address sebelum copy

	/*
		kalau buat copy dari array,
		dia akan beda memory address nya, ngga referencing ke array sebelumnya
		addres nya akan beda
	*/
	var food = &foods[0]

	fmt.Println(&food, *food) // address sesudah copy

	var alphabets = []string{"a", "b", "c", "d", "e"}
	var slicealphabets = alphabets[0:3]

	// fmt.Println(unsafe.Pointer(&alphabets))
	// fmt.Println(unsafe.Pointer(&slicealphabets))

	// var alphabet = &alphabets

	// fmt.Println(*alphabet, unsafe.Pointer(&alphabet), unsafe.Pointer(&alphabets))

	// // len

	// fmt.Println(len(alphabets), cap(alphabets))

	// var aAlphabet = alphabets[0:3]
	// var bAlphabet = alphabets[1:4]

	// fmt.Println(len(aAlphabet), cap(aAlphabet))
	// fmt.Println(len(bAlphabet), cap(bAlphabet))

	// var cAlphabet = append(alphabets, "f")

	// fmt.Println(cAlphabet, alphabets)

	/*
		APPEND FUNCTION
	*/

	var cAlphabet = append(slicealphabets, "A")

	fmt.Println(alphabets, cAlphabet)

	fmt.Printf("%p\n", alphabets)
	fmt.Printf("%p\n", cAlphabet)
	fmt.Printf("%p\n", slicealphabets)

	// COPYING A SLICE

	dst := make([]string, 3)
	src := []string{"watermelon", "melon", "apple"}

	copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)

	dst1 := []string{"watermelon", "melon", "apple", "kujang"}
	src1 := []string{"deer", "panda", "tiger", "lion"}

	copy(dst1, src1)
	fmt.Println(dst1, src1)

	var fruitsCap = []string{"apple", "grape", "banana"}
	var aFruits = fruitsCap[0:2]
	var bFruits = fruitsCap[0:2:2]

	fmt.Println(len(aFruits), cap(aFruits))
	fmt.Println(len(bFruits), cap(bFruits))

	/*
		ketika append ke variabel bfruits yang sudah di define kapasitas nya sejumlah x,
		kemudian di append, dia akan buat alamat baru,
		karena keterbatasan kapasitas yang sudah di define di awal
	*/
	var appendBfruits = append(bFruits, "durian")

	fmt.Println(fruitsCap)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(appendBfruits)
	fmt.Printf("%p\n", fruitsCap)
	fmt.Printf("%p\n", bFruits)
	fmt.Printf("%p\n", appendBfruits)

}
