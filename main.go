package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)


func QuickSort(arr []int, wg *sync.WaitGroup, depth int) {
	if wg != nil {
		defer wg.Done()
	}

	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr)

	if depth > 0 {
		var innerWg sync.WaitGroup
		innerWg.Add(2)

		go QuickSort(arr[:pivotIndex], &innerWg, depth-1)

		go QuickSort(arr[pivotIndex+1:], &innerWg, depth-1)

		innerWg.Wait()
	} else {
		// sort in the current goroutine
		QuickSort(arr[:pivotIndex], nil, 0)
		QuickSort(arr[pivotIndex+1:], nil, 0)
	}
}


func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}


func getUserInput(prompt string) string { 
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	arraySizeInput := getUserInput("Enter the size of the array: ")
	arrSize, err := strconv.Atoi(arraySizeInput)
	if err != nil || arrSize <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer for the array size.")
		return
	}

	showSortedInput := getUserInput("Do you want to display the sorted array? (yes/no): ")
	showSorted := strings.ToLower(showSortedInput) == "yes" || strings.ToLower(showSortedInput) == "y"

	rand.Seed(time.Now().UnixNano())

	arr := make([]int, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(1000000) // up to 1 million
	}

	//fmt.Printf("Sorting array of size %d...\n", arrSize)

	numCPU := runtime.NumCPU()
	fmt.Println("Number of CPU cores:", numCPU)

	var wg sync.WaitGroup
	wg.Add(1)

	start := time.Now()

	go QuickSort(arr, &wg, numCPU)

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Sorted array of size %d in %s\n", arrSize, elapsed)

	if showSorted {
		fmt.Println("Sorted array:", arr)
	} else {
		fmt.Println("Sorted array not displayed.")
	}
}
