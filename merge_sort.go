// Это рекурсивный алгоритм, который постоянно разбивает список пополам.
// Если список пуст или состоит из одного элемента, то он отсортирован по определению (базовый случай).
// Если в списке больше, чем один элемент, мы разбиваем его и рекурсивно вызываем сортировку слиянием для каждой из половин.
// После того, как обе они уже отсортированы, выполняется основная операция, называемая слиянием.
// Слияние - это процесс комбинирования двух меньших сортированных списков в один новый, но тоже отсортированный.

package main

import "fmt"

func main() {
	var arr = []int{12, 44, 32, 54, 27, 2, 9, 10, 99}
	fmt.Print(mergeSort(arr))
}

func mergeSort(arr []int) []int {

	if len(arr) >= 2 {
		mid := len(arr) / 2
		arr = merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
	}

	return arr
}

func merge(listLeft []int, listRight []int) []int {
	lst := make([]int, 0, len(listLeft) + len(listRight))
	for len(listLeft) > 0 && len(listRight) > 0 {
		if listLeft[0] < listRight[0] {
			lst = append(lst, listLeft[0])
			listLeft = listLeft[1:]
		} else {
			lst = append(lst, listRight[0])
			listRight = listRight[1:]
		}
	}

	if len(listRight) > 0 {
		lst = append(lst, listRight...)
	}

	if len(listLeft) > 0 {
		lst = append(lst, listLeft...)
	}

	return lst
}
