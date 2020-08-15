package main

import "fmt"

var (
	v = map[int][]int{
		92: []int{5, 17},
		5: []int{8, 1},
		17: []int{4, 35},
		8: []int{},
		1: []int{},
		4: []int{},
		35: []int{},
	}

	visited = make(map[int]bool, len(v))
	queue = make([]int, 0)
)

func main() {
	start := 92
	queue = append(queue, start)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		visited[cur] = true

		fmt.Println(cur)

		for _, i := range v[cur] {
			if visited[i] == false {
				queue = append(queue, i)
			}
		}
	}
}
