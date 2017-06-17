//SortAlgos.go
//This file contains the sorting algorithms called from Sorter.go

package main

import (
	//"fmt" //debugging
)

//Insertion sort algorithm - modify array in place
func insertionSort(size int, array []int, channel chan int) {
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if array[i] > array[i-1] {
				break //Leave in place (insert at back)
			}
			if array[i] < array[j] {
				var toInsert = array[i]
				for k := i - 1; k >= j; k-- {
					array[k+1] = array[k]
				}
				array[j] = toInsert
			}
		}
	}
	channel <- size //Signal end of sort
}

//Merge sort algorithm - merge in place
func mergeSort(size int, array []int, channel chan int) {
	var preMergeChannel = make(chan int)
	go mergeRecurse(array, 0, len(array), preMergeChannel)
	for range preMergeChannel{} //Wait for sort to finish
	channel <- size //Signal end of sort
}

//Recursively merge and sort list until full list is sorted
func mergeRecurse(array []int, start, end int, mergeChannel chan int) {
	defer close(mergeChannel)

	if end-start < 2 { //Do nothing, vacuously sorted
	} else {
		//Get indices of left and right half of array
		var lStart, lEnd, rStart, rEnd = mergeSplit(start, end)

		//Sort left half
		var newMergeChannel = make(chan int)
		go mergeRecurse(array, lStart, lEnd, newMergeChannel)
		for range newMergeChannel{} //wait for left half sorted

		//Sort right half
		var otherNewChannel = make(chan int)
		go mergeRecurse(array, rStart, rEnd, otherNewChannel)
		for range otherNewChannel{} //wait for right half sorted
		
		//Merge left and right half
		merge(array, lStart, lEnd, rStart, rEnd)
	}
}

//Return indices represented array split in half
func mergeSplit(start, end int) (int, int, int, int) {
	return start, (end + start) / 2, (end + start) / 2, end
}

//Merge sorted subarrays into one sorted array
//Merging occurs in place - adds complexity to normal algorithm
//THIS IS NOT WORKING
//TODO - MAKE THIS NOT BAD
func merge(array []int, lStart, lEnd, rStart, rEnd int){
	var lCounter = 0
	var rCounter = 0
	var leftSize = (lEnd - lStart)
	var rightSize = (rEnd - rStart)
	var newSize = leftSize + rightSize

	//For debugging - not to be used in final version unless necessary
	var newArray = make([]int, newSize, newSize) 

	//Merge process - combine subarrays in place
	//Precondition: left and right sections are sorted 
	//This is broken
	for i:=0; i < newSize; i++{
		if lCounter == leftSize{
			newArray[i] = array[rCounter + rStart]
			rCounter++
		} else if rCounter == rightSize{
			newArray[i] = array[lCounter + lStart]
			lCounter++
		} else{
			if array[lCounter + lStart] < array[rCounter + rStart]{
				newArray[i] = array[lCounter + lStart]
				lCounter++;
			} else{
				newArray[i] = array[rCounter + rStart]
				rCounter++;
			}
		}
	}

	//For debugging purposes
	if (len(array) == 10){
		printArray(len(newArray), newArray)
	}
}

