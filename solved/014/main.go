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
	var ans []int
	for n%2 == 0 {
		ans = append(ans, 2)
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			ans = append(ans, i)
			n /= i
		}
	}
	if n > 1 {
		ans = append(ans, n)
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
