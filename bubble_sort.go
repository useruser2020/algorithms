// Идея метода: шаг сортировки состоит в проходе снизу вверх по массиву.
// По пути просматриваются пары соседних элементов.
// Если элементы некоторой пары находятся в неправильном порядке, то меняем их местами.

package main

import "fmt"

func main() {
	var arr = []int{12, 44, 32, 54, 27, 2, 9, 10, 99}
	fmt.Print(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1]  > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	return arr
}


