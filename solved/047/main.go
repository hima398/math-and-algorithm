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
	type Item struct {
		i, c int
	}
	var q []Item
	colors := make([]int, n)
	for i := range colors {
		colors[i] = -1
	}
	for i := 0; i < n; i++ {
		if colors[i] < 0 {
			q = append(q, Item{i, 0})
			colors[i] = 0
		}

		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, next := range e[p.i] {
				nc := p.c ^ 1
				if colors[next] >= 0 {
					if colors[next] == nc {
						continue
					} else {
						fmt.Println("No")
						return
					}
				}
				q = append(q, Item{next, nc})
				colors[next] = nc
			}
		}
	}
	fmt.Println("Yes")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
