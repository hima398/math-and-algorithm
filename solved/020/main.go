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

	//dp := make([][]map[int]int, n+1)
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 6)
		for j := 0; j <= 5; j++ {
			dp[i][j] = make([]int, 1001)
		}
	}
	dp[0][0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= 5; j++ {
			//for k := 0; k <= 1000; k++ {
			for k := 0; k <= 1000; k++ {
				dp[i][j][k] += dp[i-1][j][k]
				if j > 0 && k-a[i-1] >= 0 {
					dp[i][j][k] += dp[i-1][j-1][k-a[i-1]]
				}
			}
		}
	}
	//fmt.Println(dp[n][1])
	fmt.Println(dp[n][5][1000])
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
