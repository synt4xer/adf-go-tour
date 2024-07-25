package main

import (
	"fmt"
	"strings"
)

func main() {
	// 20.1 - Variadic Function
	// Variadic Function can only use in final params
	// var avg = calculate(0, 2, 8, 9, 10, 25, 5)
	// var msg = fmt.Sprintf("Rata-rata: %.2f", avg)
	// fmt.Println(msg)

	// 20.3 - Variadic Function with Data Slice
	var numbers = []int{2, 4, 5, 3, 1, 7}
	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Rata-rata: %.2f", avg)
	fmt.Println(msg)

	// 20.4 - Function with Simple Param and Variadic Param
	var hobbiesSlice = []string{"Sleeping", "Sleeping", "Sleeping"}
	yourHobbies("Iki", "Sleeping", "Eating", "Gaming", "Sleding")
	yourHobbies("Iki Pake Slice", hobbiesSlice...)

	// batas
	fmt.Println("=======================")

	// 21.1 - Closure
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var closureNumbers = []int{2, 3, 4, 5, 4, 2, 3}
	var min, max = getMinMax(closureNumbers)
	fmt.Printf("Data: %v \n Min: %v \n Max: %v \n", numbers, min, max)

	// 21.2 - Closure IIFE
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range closureNumbers {
			if e < min {
				continue
			}

			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("Original Number: ", numbers)
	fmt.Println("New Number: ", newNumbers)

	// batas
	fmt.Println("=======================")

	// 21.3 - Closure as return value
	var maxClosure = 3
	var howMany, getNumbers = findMax(closureNumbers, maxClosure)

	fmt.Println("Numbers \t: ", closureNumbers)
	fmt.Printf("Find Max \t: %d \n \n", maxClosure)

	var theNumbers = getNumbers()

	fmt.Println("found \t: ", howMany)
	fmt.Println("value \t: ", theNumbers)

	// batas
	fmt.Println("=======================")

	// 22 Function as Params
	var paramsData = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(paramsData, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(paramsData, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data List \t\t: ", paramsData)
	fmt.Println("Filter Kata yg ada O nya \t\t: ", dataContainsO)
	fmt.Println("Filter Kata yg panjangnya 5 karakter \t\t: ", dataLength5)
}

// Function Calculate for get Average from Numbers Submit
// variable numbers will treat as Slice, so u can use numbers[0]
func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, My name is: %s \n", name)
	fmt.Printf("My Hobbies are: %s \n", hobbiesAsString)
}

// 21.3 - Closure as return value example
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		fmt.Printf("Ini closure debug \n")
		return res
	}
}

// 22 - Function as Params
// 22.2 Function as Params using Alias
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
