package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeDistance2(x1, y1, x2, y2 int) int {
	tx := x2 - x1
	ty := y2 - y1
	return tx*tx + ty*ty
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	dist2 := 1 << 60
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			dist2 = Min(dist2, ComputeDistance2(x[i], y[i], x[j], y[j]))
		}
	}
	ans := math.Sqrt(float64(dist2))
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
