// берем опорную точку (pointer),
// проходим массив, чтобы элементы, которые меньше опорной точки оказались слева от нее, а которые больше - справа.
// Дальше берем часть массива до опорной точки и вторую часть после опорной точки, повторяем на них сортировку.
// Продолжаем до того момента, как сортируемая часть массива будет пустой или состоять из одного элемента.

package main

import "fmt"

func main() {
	var arr = []int{12, 44, 32, 54, 27, 2, 9, 10, 99}
	fmt.Print(Quicksort(arr))
}

func Quicksort(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}

	sp := partition(ar)

	Quicksort(ar[:sp])
	Quicksort(ar[sp:])

	return ar
}

func partition(ar []int) int {
	p := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ar[left] < p {
			left++
		}

		for ar[right] > p {
			right--
		}

		if left >= right {
			return right
		}

		ar[left], ar[right] = ar[right], ar[left]
	}
}