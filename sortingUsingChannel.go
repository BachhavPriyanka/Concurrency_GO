package main

import (
	"fmt"
	"sort"
)

func part1(slice []float64, ch chan []float64) {
	part := slice[:len(slice)/4]
	sort.Float64s(part)
	fmt.Println(part)
	ch <- part
}

func part2(slice []float64, ch chan []float64) {
	part := slice[len(slice)/4 : 2*len(slice)/4]
	sort.Float64s(part)
	ch <- part
}

func part3(slice []float64, ch chan []float64) {
	part := slice[2*len(slice)/4 : 3*len(slice)/4]
	sort.Float64s(part)
	ch <- part
}

func part4(slice []float64, ch chan []float64) {
	part := slice[3*len(slice)/4:]
	sort.Float64s(part)
	ch <- part
}

func main() {
	inputSlice := []float64{99, 55, 1, 66, 11, 77}
	sort.Float64s(inputSlice)

	channel := make(chan []float64)

	go part1(inputSlice, channel)
	sortedPart1 := <-channel

	go part2(inputSlice, channel)
	sortedPart2 := <-channel

	go part3(inputSlice, channel)
	sortedPart3 := <-channel

	go part4(inputSlice, channel)
	sortedPart4 := <-channel

	var sortedArray []float64
	sortedArray = append(sortedArray, sortedPart1...)
	sortedArray = append(sortedArray, sortedPart2...)
	sortedArray = append(sortedArray, sortedPart3...)
	sortedArray = append(sortedArray, sortedPart4...)

	fmt.Println(sortedArray)
}
