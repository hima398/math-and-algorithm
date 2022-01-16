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
		i, d int
	}
	var q []Item
	q = append(q, Item{0, 0})
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, next := range e[p.i] {
			if dist[next] >= 0 {
				continue
			}
			q = append(q, Item{next, p.d + 1})
			dist[next] = p.d + 1
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range dist {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
