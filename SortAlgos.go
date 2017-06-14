package main

import (
	//"fmt" //debugging
)

func insertionSort(size int, array []int, channel chan int) {
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if array[i] > array[i-1] {
				break
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

func mergeSort(size int, array []int, channel chan int) {
	//var newArray = mergeRecurse(array)
	array = mergeRecurse(array)
	//copySlice(array, newArray)
	channel <- size //Signal end of sort
}

func copySlice(target []int, source []int) {
	for i:= 0; i < len(source); i++{
		target[i] = source[i]
	}
}

func mergeRecurse(array []int) ([]int) {
	if (len(array) < 2){
		return array;
	}else{
		var leftSub, rightSub = mergeSplit(array)
		var sortedLeft = mergeRecurse(leftSub)
		var sortedRight = mergeRecurse(rightSub)
		return (merge(sortedLeft, sortedRight))
	}
	
}

func mergeSplit(toSplit []int) ([]int, []int) {	
	var leftSize int = len(toSplit) / 2
	var rightSize int = len(toSplit) - leftSize
	//fmt.Printf("leftSize = %d, rightSize = %d, total = %d\n", leftSize, rightSize, len(toSplit))

	var leftSub = make([]int, leftSize, leftSize)
	var rightSub = make([]int, rightSize, rightSize)

	for i:=0; i < leftSize; i++{
		leftSub[i] = toSplit[i];
	}
	for i:=0; i < rightSize; i++{
		rightSub[i] = toSplit[i + leftSize]
	}

	return leftSub, rightSub
}

func merge(leftSub []int, rightSub []int) ([]int) {
	var newSize = len(leftSub) + len(rightSub)
	var leftSize =len(leftSub)
	var rightSize = len(rightSub)
	var leftCounter = 0
	var rightCounter = 0
	var mergedList = make([]int, newSize, newSize)

	for i:=0; i < newSize; i++{
		if leftCounter == leftSize{
			mergedList[i] = rightSub[rightCounter]
			rightCounter++
		} else if rightCounter == rightSize{
			mergedList[i] = leftSub[leftCounter]
			leftCounter++
		} else{
			if leftSub[leftCounter] < rightSub[rightCounter]{
				mergedList[i] = leftSub[leftCounter]
				leftCounter++;
			} else{
				mergedList[i] = rightSub[rightCounter]
				rightCounter++;
			}
		}
	}
	
	return mergedList
}