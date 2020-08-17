package main

import "fmt"

func main() {
	var arr = []int{12, 44, 32, 54, 27, 2, 9, 10, 99}
	fmt.Print(heapSort(arr))
}

func buildMaxHeap(arr []int, l int, i int) {
	mxIdx := i
	leftChildIdx := 2 * i + 1
	rightChildIdx := 2 * i + 2

	if leftChildIdx < l && arr[leftChildIdx] > arr[mxIdx] {
		mxIdx = leftChildIdx
	}

	if rightChildIdx< l && arr[rightChildIdx] > arr[mxIdx] {
		mxIdx = rightChildIdx
	}

	if arr[mxIdx] > arr[i] {
		arr[mxIdx], arr[i] = arr[i], arr[mxIdx]
		buildMaxHeap(arr, l, mxIdx)
	}
}

func heapSort(arr []int) []int {
	l := len(arr)
	for i := l / 2 - 1; i >= 0; i-- {
		buildMaxHeap(arr, l, i)
	}

	for i := l - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		buildMaxHeap(arr, i, 0)
	}

	return arr
}