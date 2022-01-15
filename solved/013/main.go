package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[int]struct{})
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			m[i] = struct{}{}
			m[n/i] = struct{}{}
		}
	}
	var ans []int
	for k := range m {
		ans = append(ans, k)
	}
	sort.Ints(ans)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
