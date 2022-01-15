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

	n, x, y := nextInt(), nextInt(), nextInt()
	nx, ny, nxy := 0, 0, 0
	for i := 1; i <= n; i++ {
		if i%x == 0 {
			nx++
		}
		if i%y == 0 {
			ny++
		}
		if i%x == 0 && i%y == 0 {
			nxy++
		}
	}
	ans := nx + ny - nxy
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
