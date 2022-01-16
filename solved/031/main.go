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
	a := nextIntSlice(n)
	dp := make([]int, n+1)
	dp[1] = a[0]
	for i := 2; i <= n; i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+a[i-1])
	}
	fmt.Println(dp[n])
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
