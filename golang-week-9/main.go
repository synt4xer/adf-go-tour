package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// A.32 - BUFFERED CHANNEL
	runtime.GOMAXPROCS(2)
	// messages := make(chan int, 3)

	// go func() {
	// 	for {
	// 		i := <-messages
	// 		fmt.Println("Receive Data: ", i)
	// 	}
	// }()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Send Data: ", i)
	// 	messages <- i
	// }

	// slices loop for string channel
	// strMsg := make(chan string, 3)
	// var items = []string{"John", "Wick", "Cenaa", "Edamame", "Brian", "Percival"}

	// go func() {
	// 	for {
	// 		i := <-strMsg
	// 		fmt.Println("Received Item: ", i)
	// 	}
	// }()

	// for _, item := range items {
	// 	fmt.Println("Send Item: ", item)
	// 	strMsg <- item
	// }
	// end slices channel

	// time.Sleep(1 * time.Second)

	// A.33 - CHANNEL - SELECT
	// var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	// fmt.Println("Numbers: ", numbers)

	// var maxNumbCh = make(chan int)
	// go getMax(numbers, maxNumbCh)

	// var avgCh = make(chan float64)
	// go getAverage(numbers, avgCh)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case avg := <-avgCh:
	// 		fmt.Println("Average Numbers: ", avg)
	// 	case max := <-maxNumbCh:
	// 		fmt.Println("Max Numbers: ", max)
	// 	}
	// }

	// fmt.Println("=============")
	// A.34 - CHANNEL - RANGE AND CLOSE
	// var sentMessages = make(chan string)
	// go sendMessage(sentMessages)
	// printMsg(sentMessages)

	// fmt.Println("=============")
	// A.35 - CHANNEL - TIMEOUT
	// var timeoutMsg = make(chan int)

	// go sendData(timeoutMsg)
	// retrieveData(timeoutMsg)

	fmt.Println("=============")
	// A.36 - Defer & Exit
	defer fmt.Println("halo, ini defer")

	// defer tidak bekerja jika aplikasi dihentikan secara paksa menggunakan perintah os.Exit
	// os.Exit(1)

	fmt.Println("Welcomeee")

	orderSomeFood("pizza")
	orderSomeFood("Ayam Geprek")

	// A.36 - Kombinasi Defer dan IIFE
	number := 3

	if number == 3 {
		fmt.Println("Halo 1")
		func() {
			defer fmt.Println("Halo 3")
		}()
	}

	fmt.Println("Halo 2")
}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}

	ch <- max
}

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data %d", i)
	}

	close(ch)
}

func printMsg(ch <-chan string) {
	for message := range ch {
		fmt.Println("Message: ", message)
	}
}

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		var totalTime = time.Duration(randomizer.Int()%10+1) * time.Second
		fmt.Println("Total Randomizer Time: ", totalTime)

		// fungsi time.Sleep ini akan melakukan sleep pada proses GOROUTINE yang terdapat pada Line :74
		time.Sleep(totalTime)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`Receive Data: "`, data, `"`, "\n")
		case <-time.After(time.Second * 5): // proses akan berhenti ketika tidak ada proses lain dalam waktu 5 detik setelah proses terakhir
			fmt.Println("Timeout, No Activities Under 5 Seconds")
			break loop
		}
	}
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, Silahkan Tunggu....")
	if menu == "pizza" {
		fmt.Print("Pilihan Tepat!", " ")
		fmt.Print("Pizza di Tempat Kami Paling Ena!", "\n")
		return
	}

	fmt.Println("Pesanan Anda: ", menu)
}
