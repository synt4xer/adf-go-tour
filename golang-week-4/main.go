// * A.17 - A.19
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	// * declare
	// * manifest typing
	var chicken map[string]int

	// * type inference
	// chicken := map[string]int{}

	// fmt.Println("ayam =", chicken)

	// * init
	chicken = map[string]int{}

	// * assign value
	chicken["january"] = 1
	chicken["february"] = 2

	// * overwrite
	chicken["february"] = 3

	// fmt.Println("ayam =", chicken)
	fmt.Println("bulan januari = ", chicken["january"])
	fmt.Println("bulan februari = ", chicken["february"])

	// * using make
	// var ayam = make(map[string]int)

	// ? NEW using pointer, further discussion on A.23
	var ayam = *new(map[string]int)

	ayam = map[string]int{
		"jan": 1,
		"feb": 2,
		"mar": 3,
		"apr": 4,
	}

	// * LOOPING USING FOR RANGE
	for key, val := range ayam {
		fmt.Println(key, " \t:", val)
	}

	//* before delete
	// fmt.Println(len(ayam))

	// * delete
	// delete(ayam, "apr")

	// * after delete
	// fmt.Println(len(ayam))
	// fmt.Println((ayam))

	// * SECOND RETURN IS OPTIONAL AND ITS RETURN A BOOLEAN
	var value, isExist = ayam["jan"]

	// * how to detect a key
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exists")
	}

	// * ==========================================
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chick := range chickens {
		fmt.Println(chick["name"], chick["gender"])
	}

	// * ==========================================
	// * FUNCTION SECTION

	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	// * call the function with multiple return

	var diameter float64 = 15
	var area, circum = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circum)
}

// * FUNCTION WITH RETURN
func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

// * FUNCTION WITH RETURN VOID / USING RETURN FOR STOP THE BLOCK CODE
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider\n")
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

// * FUNCTION WITH MULTIPLE RETURN
/* func calculate(d float64) (float64, float64) {
	// luas
	var area = math.Pi * math.Pow(d/2, 2)

	// keliling
	var circumference = math.Pi * d

	return area, circumference
} */

// * FUNCTION WITH PREDEFINED RETURN
func calculate(d float64) (area float64, circumference float64) {
	// luas
	area = math.Pi * math.Pow(d/2, 2)

	// keliling
	circumference = math.Pi * d

	return
}
