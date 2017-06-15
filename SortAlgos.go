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
	var preMergeChannel = make(chan int)
	mergeRecurse(array, 0, len(array), preMergeChannel)
	for range preMergeChannel{
		//Wait for sort to finish
	}
	channel <- size //Signal end of sort
}

func mergeRecurse(array []int, start, end int, mergeChannel chan int) {
	if (len(array) == 10) {
		printArray(len(array), array)
	}
	defer close(mergeChannel)

	if end-start < 2 {
		//Do nothing
	} else {
		var lStart, lEnd, rStart, rEnd = mergeSplit(start, end)

		var newMergeChannel = make(chan int)
		go mergeRecurse(array, lStart, lEnd, newMergeChannel)
		for range newMergeChannel{
			//wait for first call to finish
		}

		var otherNewChannel = make(chan int)
		go mergeRecurse(array, rStart, rEnd, otherNewChannel)
		for range otherNewChannel{
			//wait for second call to finish
		}
		
		merge(array, lStart, lEnd, rStart, rEnd)
	}
}

func mergeSplit(start, end int) (int, int, int, int) {
	return start, (end + start) / 2, (end + start) / 2, end
}

func merge(array []int, lStart, lEnd, rStart, rEnd int){
	var lCounter = 0
	var rCounter = 0
	var leftSize = (lEnd - lStart)
	var rightSize = (rEnd - rStart)
	var newSize = leftSize + rightSize

	for i:=0; i < newSize; i++{
		if lCounter == leftSize{
			array[lStart + i] = array[rCounter + rStart]
			rCounter++
		} else if rCounter == rightSize{
			array[lStart + i] = array[lCounter + lStart]
			lCounter++
		} else{
			if array[lCounter + lStart] < array[rCounter + rStart]{
				array[lStart + i] = array[lCounter + lStart]
				lCounter++;
			} else{
				array[lStart + i] = array[rCounter + rStart]
				rCounter++;
			}
		}
	}
}

