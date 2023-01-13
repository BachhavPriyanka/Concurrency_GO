package main

import "fmt"

func findOdd(numbers []int, odd chan int) {
	for _, number := range numbers {
		if number%2 != 0 {
			odd <- number
		}
	}
	close(odd)
}

func findEven(numbers []int, even chan int) {
	for _, number := range numbers {
		if number%2 == 0 {
			even <- number
		}
	}
	close(even)
}

func main() {
	numbers := []int{10, 15, 20, 25, 30, 35, 40}
	odd := make(chan int)
	even := make(chan int)

	go findOdd(numbers, odd)
	go findEven(numbers, even)

	fmt.Print("Odd numbers:")
	for o := range odd {
		fmt.Print(o, " ")
	}

	fmt.Print("\nEven numbers:")
	for e := range even {
		fmt.Print(e, " ")
	}
	fmt.Println("")
}
