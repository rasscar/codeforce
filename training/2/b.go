package main

import (
	"bufio"
	"fmt"
	"os"
)

func faktor(x, p int) int {
	if x == 0 {
		return -1
	}
	c := 0
	for x%p == 0 {
		c++
		x /= p
	}
	return c
}

func dpMin(f [][]int, n int) (int, [][]byte) {
	INF := 1 << 30
	cost := make([][]int, n)
	dir := make([][]byte, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		dir[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			cost[i][j] = INF
		}
	}
	if f[0][0] >= 0 {
		cost[0][0] = f[0][0]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || f[i][j] < 0 {
				continue
			}
			if i > 0 && cost[i-1][j] < INF {
				cost[i][j] = cost[i-1][j] + f[i][j]
				dir[i][j] = 'D'
			}
			if j > 0 && cost[i][j-1] < INF && cost[i][j-1]+f[i][j] < cost[i][j] {
				cost[i][j] = cost[i][j-1] + f[i][j]
				dir[i][j] = 'R'
			}
		}
	}
	return cost[n-1][n-1], dir
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([][]int, n)
	zi, zj := -1, -1
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
			if a[i][j] == 0 && zi == -1 {
				zi, zj = i, j
			}
		}
	}

	f2 := make([][]int, n)
	f5 := make([][]int, n)
	for i := 0; i < n; i++ {
		f2[i] = make([]int, n)
		f5[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f2[i][j] = faktor(a[i][j], 2)
			f5[i][j] = faktor(a[i][j], 5)
		}
	}

	c2, d2 := dpMin(f2, n)
	c5, d5 := dpMin(f5, n)

	cost, dir := c2, d2
	if c5 < c2 {
		cost, dir = c5, d5
	}

	if cost == 0 || zi == -1 {
		fmt.Fprintln(out, cost)
		path := make([]byte, 0, 2*n-2)
		i, j := n-1, n-1
		for i > 0 || j > 0 {
			path = append(path, dir[i][j])
			if dir[i][j] == 'D' {
				i--
			} else {
				j--
			}
		}
		for l, r := 0, len(path)-1; l < r; l, r = l+1, r-1 {
			path[l], path[r] = path[r], path[l]
		}
		out.Write(path)
		out.WriteByte('\n')
		return
	}

	fmt.Fprintln(out, 1)
	for i := 0; i < zi; i++ {
		out.WriteByte('D')
	}
	for j := 0; j < zj; j++ {
		out.WriteByte('R')
	}
	for i := zi; i < n-1; i++ {
		out.WriteByte('D')
	}
	for j := zj; j < n-1; j++ {
		out.WriteByte('R')
	}
	out.WriteByte('\n')
}
