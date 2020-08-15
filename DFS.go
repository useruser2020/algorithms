package main

import "fmt"

func main() {
	start := 92

	dfs(start)
}

func dfs(s int) {
	if visited[s] == true{
		return
	}

	visited[s] = true

	fmt.Println(s)

	for _, i := range v[s] {
		dfs(i)
	}
}