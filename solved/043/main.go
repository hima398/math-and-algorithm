package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	visited := make([]bool, n)
	var Dfs func(i, par int)
	Dfs = func(i, par int) {
		visited[i] = true
		for _, next := range e[i] {
			if next == par || visited[next] {
				continue
			}
			Dfs(next, i)
		}
	}
	Dfs(0, -1)
	for _, ok := range visited {
		if !ok {
			fmt.Println("The graph is not connected.")
			return
		}
	}
	fmt.Println("The graph is connected.")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
