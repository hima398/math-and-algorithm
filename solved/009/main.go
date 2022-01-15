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

	n, s := nextInt(), nextInt()
	a := nextIntSlice(n)
	dp := make([]bool, s+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		ndp := make([]bool, s+1)
		for j := 0; j <= s; j++ {
			if !dp[j] {
				continue
			}
			ndp[j] = dp[j]
			next := j + a[i]
			if next <= s {
				ndp[next] = true
			}
		}
		dp = ndp
		//fmt.Println(dp)
	}
	if dp[s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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
