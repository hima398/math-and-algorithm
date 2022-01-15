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
	e := make([]bool, n+1)
	e[0] = true
	e[1] = true
	for i := 2; i <= n; i++ {
		for j := i + i; j <= n; j += i {
			e[j] = true
		}
	}
	var ans []int
	for i := 2; i <= n; i++ {
		if !e[i] {
			ans = append(ans, i)
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "%d", ans[0])
	for i := 1; i < len(ans); i++ {
		fmt.Fprintf(out, " %d", ans[i])
	}
	fmt.Fprintln(out)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
