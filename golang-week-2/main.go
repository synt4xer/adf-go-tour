// A13 - A14

package main

import "fmt"

func main() {
	/* condition */
	var point = 4

	if point > 8 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if (point >= 5) && (point < 8) {
		fmt.Println("lulus")
	} else {
		fmt.Println("remedial")
	}

	fmt.Println("====================")

	// variable temporary
	var pointTemp float32 = 8840.0
	if percent := pointTemp / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	fmt.Println("====================")

	var pointSwitch int8 = 3
	switch pointSwitch {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4: // multiple condition
		fmt.Println("good")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	fmt.Println("====================")

	var pointSwitchIfElse int8 = 8
	switch {
	case pointSwitchIfElse == 8:
		fmt.Println("perfect")
	case (pointSwitchIfElse > 3) && (pointSwitchIfElse < 8):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println("====================s")

	var pointSwitchFallthrough int8 = 6
	switch {
	case pointSwitchFallthrough == 8:
		fmt.Println("perfect")
	case (pointSwitchFallthrough > 3) && (pointSwitchFallthrough < 8):
		fmt.Println("awesome")
		fallthrough
	case pointSwitchFallthrough < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println("====================")

	var pointNestedIf = 0

	if pointNestedIf > 7 {
		switch pointNestedIf {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointNestedIf == 5 {
			fmt.Println("not bad")
		} else if pointNestedIf == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if pointNestedIf == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	fmt.Println("====================")
	fmt.Println("====================")

	/* looping */
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("====================")

	var i = 1
	for i <= 5 { // while style
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("====================")

	var iNoArgument = 0
	for {
		fmt.Println("Angka", iNoArgument)
		iNoArgument++

		if iNoArgument == 5 {
			break
		}
	}

	fmt.Println("====================")

	var xs string = "123"
	for _, v := range xs { // string
		fmt.Printf("Value=%c\n", v)
	}

	fmt.Println("====================")

	var ys = [5]int{10, 20, 30, 40, 50}
	for _, v := range ys { // array
		fmt.Println("Value=", v)
	}

	fmt.Println("====================")

	var zs = ys[0:2]
	for _, v := range zs { // slice
		fmt.Println("Value=", v)
	}

	fmt.Println("====================")

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Printf("Key=%c\n", k)
		fmt.Println("Value=", v)
	}

	fmt.Println("====================")

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	fmt.Println("====================")

	for i := range 6 {
		fmt.Println(i)
	}

	fmt.Println("====================")

	for i := 1; i <= 10; i++ { // loop with break and continue
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Println("====================")

	for i := 5; i >= 0; i-- { // nested loop
		for j := i; j <= 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	fmt.Println("====================")

outerLoop: // loop label
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
