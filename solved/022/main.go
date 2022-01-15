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

	const max = int(1e5)
	n := nextInt()
	a := nextIntSlice(n)
	s := make([]int, max+1)
	for _, v := range a {
		s[v]++
	}
	ans := 0
	for i := 1; i < max; i++ {
		if i == 50000 {
			continue
		}
		ans += s[i] * s[max-i]
	}
	ans /= 2
	ans += s[50000] * (s[50000] - 1) / 2
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
