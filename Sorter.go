package main

import (
	"fmt"
)

const (
	NUM_ARRAYS = 5

	//Array Dimensions
	TINY   = 10
	SMALL  = 100
	MEDIUM = 1000
	LARGE  = 10000
	HUGE   = 100000
)

var (
	tinyArray   []int
	smallArray  []int
	mediumArray []int
	largeArray  []int
	hugeArray   []int
)

func main() {
	var populating = make(chan int)
	initialize(populating)
	for range populating {
		//Wait until arrays have populated to print
	}
	printArray(TINY, tinyArray)

	var sorting = make(chan int)
	mergeSortAll(sorting)
	for range sorting {
		//Wait until arrays done sorting to print
	}
	printArray(TINY, tinyArray)

	var checkingSorted = make(chan int)
	checkArraysSorted(checkingSorted)
	for range checkingSorted {
		//Wait until done checking arrays properly sorted
	}
}

func initialize(channel chan int) {
	defer close(channel)

	tinyArray = make([]int, TINY, TINY)
	smallArray = make([]int, SMALL, SMALL)
	mediumArray = make([]int, MEDIUM, MEDIUM)
	largeArray = make([]int, LARGE, LARGE)
	hugeArray = make([]int, HUGE, HUGE)

	go populateArrayRand(TINY, tinyArray, channel)
	go populateArrayRand(SMALL, smallArray, channel)
	go populateArrayRand(MEDIUM, mediumArray, channel)
	go populateArrayRand(LARGE, largeArray, channel)
	go populateArrayRand(HUGE, hugeArray, channel)

	for i := 0; i < NUM_ARRAYS; i++ {
		/*flag := */ <-channel
		//fmt.Printf("Done with array of size: %d\n" , flag)
	}
}

func insertionSortAll(channel chan int) {
	defer close(channel)

	go insertionSort(TINY, tinyArray, channel)
	go insertionSort(SMALL, smallArray, channel)
	go insertionSort(MEDIUM, mediumArray, channel)
	go insertionSort(LARGE, largeArray, channel)
	go insertionSort(HUGE, hugeArray, channel)

	for i := 0; i < NUM_ARRAYS; i++ {
		/*flag := */ <-channel
		//fmt.Printf("Done with array of size: %d\n" , flag)
	}
}

func mergeSortAll(channel chan int) {
	defer close(channel)

	go mergeSort(TINY, tinyArray, channel)
	go mergeSort(SMALL, smallArray, channel)
	go mergeSort(MEDIUM, mediumArray, channel)
	go mergeSort(LARGE, largeArray, channel)
	go mergeSort(HUGE, hugeArray, channel)

	for i := 0; i < NUM_ARRAYS; i++ {
		/*flag :=*/  <-channel
		//fmt.Printf("Done with array of size: %d\n" , flag)
	}
}

func checkArraysSorted(channel chan int) {
	defer close(channel)

	go checkSorted(TINY, tinyArray, channel)
	go checkSorted(SMALL, smallArray, channel)
	go checkSorted(MEDIUM, mediumArray, channel)
	go checkSorted(LARGE, largeArray, channel)
	go checkSorted(HUGE, hugeArray, channel)

	for i := 0; i < NUM_ARRAYS; i++ {
		<-channel
	}
}

func printArray(size int, array []int) {
	for i := 0; i < size; i++ {
		fmt.Print(array[i])
		fmt.Print(" ")
	}
	fmt.Printf("\n")
}



