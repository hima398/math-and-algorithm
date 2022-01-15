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
	sp, sq := 0, 0
	ans := 0.0
	for i := 0; i < n; i++ {
		p, q := nextInt(), nextInt()
		sp += p
		sq += q
		ans += float64(q) / float64(p)
	}
	//ans := float64(sq) / float64(sp)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
