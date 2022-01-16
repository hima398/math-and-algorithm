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
	h := nextIntSlice(n)
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = Abs(h[1] - h[0])
	for i := 2; i < n; i++ {
		dp[i] = Min(dp[i-2]+Abs(h[i]-h[i-2]), dp[i-1]+Abs(h[i]-h[i-1]))
	}
	fmt.Println(dp[n-1])
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
