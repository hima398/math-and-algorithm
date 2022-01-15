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

	n := nextInt()
	a := nextIntSlice(n - 1)
	m := nextInt()
	b := nextIntSlice(m)

	// 0-indexed
	for i := 0; i < m; i++ {
		b[i]--
	}
	s := make([]int, n)
	for i := 1; i <= n-1; i++ {
		s[i] = s[i-1] + a[i-1]
	}
	// 0, 8, 14, 23
	// 1-0 -> 8
	//
	ans := 0
	l := b[0]
	for i := 1; i < m; i++ {
		r := b[i]
		if l > r {
			l, r = r, l
		}
		ans += s[r] - s[l]
		l = b[i]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
