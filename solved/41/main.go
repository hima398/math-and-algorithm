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

	t, n := nextInt(), nextInt()
	var l, r []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
	}
	ans := make([]int, t+1)
	for i := 0; i < n; i++ {
		ans[l[i]]++
		ans[r[i]]--
	}
	//fmt.Println(ans)
	for i := 0; i < t; i++ {
		ans[i+1] += ans[i]
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < t; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
