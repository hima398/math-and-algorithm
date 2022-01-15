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

	n, q := nextInt(), nextInt()
	var l, r, x []int
	for i := 0; i < q; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
		x = append(x, nextInt())
	}
	s := make([]int, n+2)
	for i := 0; i < q; i++ {
		s[l[i]] += x[i]
		s[r[i]+1] -= x[i]
	}
	for i := 0; i <= n; i++ {
		s[i+1] += s[i]
	}
	var ans []rune
	for i := 1; i < n; i++ {
		if s[i] < s[i+1] {
			ans = append(ans, '<')
		} else if s[i] > s[i+1] {
			ans = append(ans, '>')
		} else {
			ans = append(ans, '=')
		}
	}
	fmt.Println(string(ans))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
