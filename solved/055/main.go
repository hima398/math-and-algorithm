package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Matrix [2][2]int

func MulMatrix(a, b Matrix, p int) Matrix {
	var res [2][2]int
	res[0][0] = (a[0][0]*b[0][0] + a[0][1]*b[1][0]) % p
	res[0][1] = (a[0][0]*b[0][1] + a[0][1]*b[1][1]) % p
	res[1][0] = (a[1][0]*b[0][0] + a[1][1]*b[1][0]) % p
	res[1][1] = (a[1][0]*b[0][1] + a[1][1]*b[1][1]) % p
	return res
}

func PowMatrix(x Matrix, n, p int) Matrix {
	res := [2][2]int{{1, 0}, {0, 1}}
	for n > 0 {
		if n%2 == 1 {
			res = MulMatrix(res, x, p) // ret * x % p
		}
		n >>= 1
		x = MulMatrix(x, x, p) // x * x % p
	}
	return res
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7
	n := nextInt()
	e := [2][2]int{{2, 1}, {1, 0}}
	m := PowMatrix(e, n-2, p)
	//fmt.Println(m[0])
	//fmt.Println(m[1])
	ans := (m[0][0] + m[0][1]) % p
	fmt.Println(ans)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
