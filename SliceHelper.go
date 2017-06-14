package main

import (
	"fmt"
	"math/rand"
	"time"
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

func checkSorted(size int, array []int, channel chan int) {
	var sorted = true

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if array[i] > array[j] {
				sorted = false
				break
			}
		}
		if !sorted {
			break
		}
	}

	if sorted {
		fmt.Printf("Array of size %d was successfully sorted\n", size)
	} else {
		fmt.Printf("Array of size %d was not sorted correctly\n", size)
	}

	channel <- size
}