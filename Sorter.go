//Sorter.go
//This file handles command line arguments and calls sort functions
//Can be seen as "driver"

package main

import (
	"fmt" 
	"os" //for command line arguments
)

//Global constant declaration
const (
	NUM_ARRAYS = 5

	//Array Dimensions
	TINY   = 10
	SMALL  = 100
	MEDIUM = 1000
	LARGE  = 10000
	HUGE   = 100000
)

//Global variable declartion
var (
	//Arrays for sorting
	tinyArray   []int
	smallArray  []int
	mediumArray []int
	largeArray  []int
	hugeArray   []int

	//Arrays to veryify successful sort
	tinyArrayCopy   []int
	smallArrayCopy  []int
	mediumArrayCopy []int
	largeArrayCopy  []int
	hugeArrayCopy   []int
)

func main() {
   	var args = os.Args[1:] //Get command line arguments, ignoring first

   	if (validArgs(args)) {
		var populating = make(chan int)
		initializeArrays(populating)
		for range populating {} //Wait till done populating arrays

		if len(args) == 2 { //User specified to verify success of sort
			var copying = make(chan int)
			copyArrays(copying) 
			for range copying{} //Wait till done copying
		}

		var sorting = make(chan int)
		doSpecifiedSort(args[0], sorting)
		for range sorting {} //Wait until arrays done being sorted

		if len(args) == 2 { //Verify success of sort algorithm 
			var checkingSorted = make(chan int)
			checkArraysSorted(checkingSorted)
			for range checkingSorted {} //Wait till done checking
		}	
	}
}

//Sort all arrays according to command line arguments
//Extend switch statement as more sort algorithms are written
func doSpecifiedSort(parameter string, channel chan int) {
	switch parameter {
		case "insertion":
			insertionSortAll(channel);
		case "merge":
			mergeSortAll(channel);
	}
}

//Initialize all arrays using goroutines
func initializeArrays(channel chan int) {
	defer close(channel)

	tinyArray = make([]int, TINY, TINY)
	smallArray = make([]int, SMALL, SMALL)
	mediumArray = make([]int, MEDIUM, MEDIUM)
	largeArray = make([]int, LARGE, LARGE)
	hugeArray = make([]int, HUGE, HUGE)

	//Choose Rand, Increasing, Decreasing to test algorithm cases
	//For example, insertion sort should be very fast on increasing array
	//TODO - make this choice a command line argument option
	go populateArrayRand(TINY, tinyArray, channel)
	go populateArrayRand(SMALL, smallArray, channel)
	go populateArrayRand(MEDIUM, mediumArray, channel)
	go populateArrayRand(LARGE, largeArray, channel)
	go populateArrayRand(HUGE, hugeArray, channel)

	for i := 0; i < NUM_ARRAYS; i++ { //Wait for goroutines to finish
		<-channel
	}
}

//Copy arrays for future comparison using goroutines
func copyArrays(channel chan int) {
	defer close(channel)

	tinyArrayCopy = make([]int, TINY, TINY)
	smallArrayCopy = make([]int, SMALL, SMALL)
	mediumArrayCopy = make([]int, MEDIUM, MEDIUM)
	largeArrayCopy = make([]int, LARGE, LARGE)
	hugeArrayCopy = make([]int, HUGE, HUGE)

	go copyArray(TINY, tinyArray, tinyArrayCopy, channel)
	go copyArray(SMALL, smallArray, smallArrayCopy, channel)
	go copyArray(MEDIUM, mediumArray, mediumArrayCopy, channel)
	go copyArray(LARGE, largeArray, largeArrayCopy, channel)
	go copyArray(HUGE, hugeArray, hugeArrayCopy, channel)

	for i := 0; i < NUM_ARRAYS; i++ { //Wait for goroutines to finish
		<-channel
	}
}

//Insertion sort all arrays using goroutines
func insertionSortAll(channel chan int) {
	defer close(channel)

	go insertionSort(TINY, tinyArray, channel)
	go insertionSort(SMALL, smallArray, channel)
	go insertionSort(MEDIUM, mediumArray, channel)
	go insertionSort(LARGE, largeArray, channel)
	go insertionSort(HUGE, hugeArray, channel)

	for i := 0; i < NUM_ARRAYS; i++ { //Wait for goroutines to finish
		<-channel
	}
}

//Merge sort all arrays using goroutines
func mergeSortAll(channel chan int) {
	defer close(channel)

	go mergeSort(TINY, tinyArray, channel)
	go mergeSort(SMALL, smallArray, channel)
	go mergeSort(MEDIUM, mediumArray, channel)
	go mergeSort(LARGE, largeArray, channel)
	go mergeSort(HUGE, hugeArray, channel)

	for i := 0; i < NUM_ARRAYS; i++ { //Wait for goroutines to finish
		<-channel
	}
}

//Check all arrays are properly sorted using goroutines
func checkArraysSorted(channel chan int) {
	defer close(channel)

	go checkSorted(TINY, tinyArray, tinyArrayCopy, channel)
	go checkSorted(SMALL, smallArray, smallArrayCopy, channel)
	go checkSorted(MEDIUM, mediumArray, mediumArrayCopy, channel)
	go checkSorted(LARGE, largeArray, largeArrayCopy, channel)
	go checkSorted(HUGE, hugeArray, hugeArrayCopy, channel)

	for i := 0; i < NUM_ARRAYS; i++ { //Wait for goroutines to finish
		<-channel
	}
}

//Verify arugments passed from command line are valid instructions
//Extend switch statement as more sort algos are added
func validArgs(args []string) (bool){
	if len(args) > 2 || len(args) == 0 { //1 or 2 args is valid
		displayUsage()
		return false;
	}

	switch args[0] {
		case "insertion":
			return true;
		case "merge":
			return true;
		default:
			displayUsage();
			return false;
	}
}

//Display valid options for command line input
func displayUsage() {
	fmt.Println("Usage: mode [verify]")
	fmt.Println("Valid modes: 'insertion' 'merge' (more to be added)")
}



