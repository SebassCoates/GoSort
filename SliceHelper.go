package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func populateArrayRand(size int, array []int, channel chan int) {
	var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		array[i] = generator.Intn(size)
	}

	channel <- size //Signal end of sort
}

func populateArrayIncreasing(size int, array []int, channel chan int) {
	for i := 0; i < size; i++ {
		array[i] = i
	}

	channel <- size //Signal end of sort
}

func checkSorted(size int, toCheck, initial []int, channel chan int) {
	var sorted = true

	//Sort original contents using known functional algorithm to compare
	sort.Ints(initial);

	for i := 0; i < size; i++ {
		if toCheck[i] != initial[i]{
			sorted = false;
		}
	}

	if sorted {
		fmt.Printf("Array of size %d was successfully sorted\n", size)
	} else {
		fmt.Printf("Array of size %d was not sorted correctly\n", size)
	}

	channel <- size
}